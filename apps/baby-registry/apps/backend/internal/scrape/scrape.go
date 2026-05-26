package scrape

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

var privateRanges []*net.IPNet

func init() {
	for _, cidr := range []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"100.64.0.0/10",
		"::1/128",
		"fc00::/7",
		"fe80::/10",
		"0.0.0.0/8",
	} {
		_, network, _ := net.ParseCIDR(cidr)
		privateRanges = append(privateRanges, network)
	}
}

func isPrivateIP(ip net.IP) bool {
	for _, network := range privateRanges {
		if network.Contains(ip) {
			return true
		}
	}
	return false
}

func newSafeClient() *http.Client {
	safeDialer := &net.Dialer{Timeout: 10 * time.Second}
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			addrs, err := net.DefaultResolver.LookupHost(ctx, host)
			if err != nil {
				return nil, err
			}
			for _, a := range addrs {
				ip := net.ParseIP(a)
				if ip == nil || isPrivateIP(ip) {
					return nil, fmt.Errorf("requests to private or reserved addresses are not allowed")
				}
			}
			return safeDialer.DialContext(ctx, network, net.JoinHostPort(addrs[0], port))
		},
	}
	return &http.Client{
		Timeout:   15 * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return errors.New("too many redirects")
			}
			host := req.URL.Hostname()
			addrs, err := net.LookupHost(host)
			if err != nil {
				return err
			}
			for _, a := range addrs {
				ip := net.ParseIP(a)
				if ip == nil || isPrivateIP(ip) {
					return errors.New("redirect to private or reserved address not allowed")
				}
			}
			return nil
		},
	}
}

type Result struct {
	Title      string  `json:"title"`
	ImageUrl   string  `json:"imageUrl"`
	ProductUrl string  `json:"productUrl"`
	Price      float64 `json:"price"`
	Currency   string  `json:"currency"`
	Source     string  `json:"source"`
}

type Handler struct {
	client *http.Client
}

func NewHandler() *Handler {
	return &Handler{
		client: newSafeClient(),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	target := strings.TrimSpace(r.URL.Query().Get("url"))
	if target == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}
	parsed, err := url.Parse(target)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, target, nil)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; BabyRegistryBot/1.0)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	resp, err := h.client.Do(req)
	if err != nil {
		http.Error(w, "fetch error: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 4*1024*1024))
	if err != nil {
		http.Error(w, "read error", http.StatusBadGateway)
		return
	}

	res, err := parseHTML(string(body), parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	res.ProductUrl = target
	res.Source = detectSource(parsed.Host)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func parseHTML(body string, pageURL *url.URL) (Result, error) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return Result{}, err
	}
	signals := extractPageSignals(doc)

	res := baseResult(signals)
	for _, p := range parsers {
		if p.Match(pageURL.Hostname()) {
			p.Apply(signals, &res)
		}
	}

	res.ImageUrl = makeAbsoluteURL(pageURL, res.ImageUrl)
	if res.Title == "" {
		return Result{}, errors.New("could not extract title")
	}
	return res, nil
}

type pageSignals struct {
	Meta                map[string]string
	PageTitle           string
	FirstImgSrc         string
	TextByID            map[string]string
	PriceTexts          []string
	ScriptTexts         []string
	ImageByID           map[string]string
	AmazonDynamicImages []string
	JSONLDProducts      []jsonLDProduct
}

type jsonLDProduct struct {
	Name     string
	Image    string
	Price    float64
	Currency string
}

func extractPageSignals(doc *html.Node) pageSignals {
	signals := pageSignals{
		Meta:      map[string]string{},
		TextByID:  map[string]string{},
		ImageByID: map[string]string{},
	}
	var scripts []string

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil && signals.PageTitle == "" {
					signals.PageTitle = strings.TrimSpace(n.FirstChild.Data)
				}
			case "meta":
				var prop, name, itemprop, content string
				for _, a := range n.Attr {
					switch strings.ToLower(a.Key) {
					case "property":
						prop = strings.ToLower(a.Val)
					case "name":
						name = strings.ToLower(a.Val)
					case "itemprop":
						itemprop = strings.ToLower(a.Val)
					case "content":
						content = a.Val
					}
				}
				key := prop
				if key == "" {
					key = name
				}
				if key == "" {
					key = itemprop
				}
				if key != "" && content != "" {
					if _, exists := signals.Meta[key]; !exists {
						signals.Meta[key] = content
					}
				}
			case "link":
				var rel, href string
				for _, a := range n.Attr {
					switch strings.ToLower(a.Key) {
					case "rel":
						rel = strings.ToLower(a.Val)
					case "href":
						href = a.Val
					}
				}
				if strings.Contains(rel, "image_src") && href != "" {
					if _, exists := signals.Meta["link:image_src"]; !exists {
						signals.Meta["link:image_src"] = href
					}
				}
			case "script":
				typ := strings.ToLower(attrVal(n, "type"))
				if txt := nodeText(n); txt != "" {
					if typ == "application/ld+json" {
						scripts = append(scripts, txt)
					}
					if strings.Contains(txt, "price") || strings.Contains(txt, "Price") || strings.Contains(txt, "$") || strings.Contains(txt, "£") || strings.Contains(txt, "€") {
						signals.ScriptTexts = append(signals.ScriptTexts, txt)
					}
				}
			case "img":
				src := firstNonEmpty(attrVal(n, "data-old-hires"), attrVal(n, "src"), attrVal(n, "data-src"))
				if signals.FirstImgSrc == "" && src != "" {
					signals.FirstImgSrc = src
				}
				if id := strings.TrimSpace(attrVal(n, "id")); id != "" && src != "" {
					signals.ImageByID[id] = src
				}
				if dyn := strings.TrimSpace(attrVal(n, "data-a-dynamic-image")); dyn != "" {
					signals.AmazonDynamicImages = append(signals.AmazonDynamicImages, parseLargestImageFromDynamicAttr(dyn))
				}
			}

			if classContains(attrVal(n, "class"), "a-offscreen") {
				if txt := nodeText(n); txt != "" {
					signals.PriceTexts = append(signals.PriceTexts, txt)
				}
			}

			if id := strings.TrimSpace(attrVal(n, "id")); id == "productTitle" || looksLikePriceID(id) {
				txt := nodeText(n)
				if txt != "" {
					signals.TextByID[id] = txt
					if looksLikePriceID(id) {
						signals.PriceTexts = append(signals.PriceTexts, txt)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	signals.JSONLDProducts = parseJSONLDProducts(scripts)

	for i := range signals.AmazonDynamicImages {
		signals.AmazonDynamicImages[i] = strings.TrimSpace(signals.AmazonDynamicImages[i])
	}

	return signals
}

func baseResult(s pageSignals) Result {
	res := Result{
		Title: firstNonEmpty(
			s.Meta["og:title"],
			s.Meta["twitter:title"],
			s.PageTitle,
		),
		ImageUrl: firstNonEmpty(
			s.Meta["og:image"],
			s.Meta["twitter:image"],
			s.Meta["twitter:image:src"],
			s.Meta["link:image_src"],
		),
		Currency: firstNonEmpty(s.Meta["og:price:currency"], s.Meta["product:price:currency"]),
	}
	if priceStr := firstNonEmpty(s.Meta["og:price:amount"], s.Meta["product:price:amount"]); priceStr != "" {
		if v, err := strconv.ParseFloat(priceStr, 64); err == nil {
			res.Price = v
		}
	}
	if (res.ImageUrl == "" || res.Price == 0 || res.Currency == "") && len(s.JSONLDProducts) > 0 {
		for _, p := range s.JSONLDProducts {
			if res.Title == "" {
				res.Title = p.Name
			}
			if res.ImageUrl == "" {
				res.ImageUrl = p.Image
			}
			if res.Price == 0 && p.Price > 0 {
				res.Price = p.Price
			}
			if res.Currency == "" {
				res.Currency = p.Currency
			}
			if res.ImageUrl != "" && res.Price > 0 && res.Currency != "" {
				break
			}
		}
	}
	if res.ImageUrl == "" {
		res.ImageUrl = s.FirstImgSrc
	}
	return res
}

type siteParser interface {
	Match(host string) bool
	Apply(s pageSignals, res *Result)
}

var parsers = []siteParser{
	amazonParser{},
}

func detectSource(host string) string {
	host = strings.ToLower(host)
	switch {
	case strings.Contains(host, "amazon."):
		return "Amazon"
	case strings.Contains(host, "mamasandpapas"):
		return "MamasAndPapas"
	case strings.Contains(host, "etsy."):
		return "Etsy"
	case strings.Contains(host, "johnlewis"):
		return "JohnLewis"
	case strings.Contains(host, "ikea."):
		return "IKEA"
	default:
		return "Other"
	}
}

func firstNonEmpty(vs ...string) string {
	for _, v := range vs {
		v = strings.TrimSpace(v)
		if v != "" {
			return v
		}
	}
	return ""
}

func firstPrice(vs ...string) float64 {
	for _, v := range vs {
		if p, ok := parsePrice(v); ok {
			return p
		}
	}
	return 0
}

var priceRE = regexp.MustCompile(`[-+]?[0-9]*[\.,]?[0-9]+`)

func parsePrice(v string) (float64, bool) {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0, false
	}
	m := priceRE.FindString(v)
	if m == "" {
		return 0, false
	}
	m = strings.ReplaceAll(m, ",", "")
	f, err := strconv.ParseFloat(m, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func parseJSONLDProducts(scripts []string) []jsonLDProduct {
	out := make([]jsonLDProduct, 0, 1)
	for _, raw := range scripts {
		for _, root := range decodeJSONLDRoots(raw) {
			walkJSONLD(root, &out)
		}
	}
	return out
}

func decodeJSONLDRoots(raw string) []any {
	var single any
	if err := json.Unmarshal([]byte(raw), &single); err == nil {
		return []any{single}
	}

	var arr []any
	if err := json.Unmarshal([]byte(raw), &arr); err == nil {
		return arr
	}
	return nil
}

func walkJSONLD(v any, out *[]jsonLDProduct) {
	switch t := v.(type) {
	case map[string]any:
		if isJSONLDProduct(t) {
			*out = append(*out, parseJSONLDProduct(t))
		}
		if g, ok := t["@graph"]; ok {
			walkJSONLD(g, out)
		}
		for _, val := range t {
			switch val.(type) {
			case map[string]any, []any:
				walkJSONLD(val, out)
			}
		}
	case []any:
		for _, e := range t {
			walkJSONLD(e, out)
		}
	}
}

func isJSONLDProduct(m map[string]any) bool {
	t := strings.ToLower(strings.TrimSpace(toString(m["@type"])))
	if t == "product" {
		return true
	}
	if arr, ok := m["@type"].([]any); ok {
		for _, v := range arr {
			if strings.ToLower(strings.TrimSpace(toString(v))) == "product" {
				return true
			}
		}
	}
	return false
}

func parseJSONLDProduct(m map[string]any) jsonLDProduct {
	p := jsonLDProduct{
		Name:  toString(m["name"]),
		Image: parseJSONLDImage(m["image"]),
	}
	if offers, ok := m["offers"]; ok {
		o := firstJSONLDOffer(offers)
		if o != nil {
			p.Currency = toString(o["priceCurrency"])
			p.Price = parseJSONLDPrice(o["price"])
		}
	}
	return p
}

func parseJSONLDImage(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case []any:
		for _, it := range t {
			if s := toString(it); s != "" {
				return s
			}
			if obj, ok := it.(map[string]any); ok {
				if s := firstNonEmpty(toString(obj["url"]), toString(obj["contentUrl"])); s != "" {
					return s
				}
			}
		}
	case map[string]any:
		return firstNonEmpty(toString(t["url"]), toString(t["contentUrl"]))
	}
	return ""
}

func firstJSONLDOffer(v any) map[string]any {
	switch t := v.(type) {
	case map[string]any:
		return t
	case []any:
		for _, e := range t {
			if m, ok := e.(map[string]any); ok {
				return m
			}
		}
	}
	return nil
}

func parseJSONLDPrice(v any) float64 {
	switch t := v.(type) {
	case float64:
		return t
	case string:
		if p, ok := parsePrice(t); ok {
			return p
		}
	}
	return 0
}

func toString(v any) string {
	switch t := v.(type) {
	case string:
		return strings.TrimSpace(t)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	default:
		return ""
	}
}

func nodeText(n *html.Node) string {
	if n == nil {
		return ""
	}
	var b strings.Builder
	var walk func(*html.Node)
	walk = func(cur *html.Node) {
		if cur == nil {
			return
		}
		if cur.Type == html.TextNode {
			t := strings.TrimSpace(cur.Data)
			if t != "" {
				if b.Len() > 0 {
					b.WriteByte(' ')
				}
				b.WriteString(t)
			}
		}
		for c := cur.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return strings.TrimSpace(b.String())
}

func attrVal(n *html.Node, key string) string {
	if n == nil {
		return ""
	}
	key = strings.ToLower(key)
	for _, a := range n.Attr {
		if strings.ToLower(a.Key) == key {
			return strings.TrimSpace(a.Val)
		}
	}
	return ""
}

func makeAbsoluteURL(base *url.URL, raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" || base == nil {
		return raw
	}
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	return base.ResolveReference(u).String()
}

func looksLikePriceID(id string) bool {
	id = strings.TrimSpace(strings.ToLower(id))
	if id == "" {
		return false
	}
	return strings.Contains(id, "price") || strings.HasPrefix(id, "corepr") || strings.Contains(id, "deal")
}

func classContains(classList string, className string) bool {
	className = strings.TrimSpace(strings.ToLower(className))
	if className == "" {
		return false
	}
	for _, c := range strings.Fields(strings.ToLower(classList)) {
		if c == className {
			return true
		}
	}
	return false
}
