package scrape

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

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
		client: &http.Client{Timeout: 15 * time.Second},
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

	res, err := parseHTML(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	res.ProductUrl = target
	res.Source = detectSource(parsed.Host)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func parseHTML(body string) (Result, error) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return Result{}, err
	}
	meta := map[string]string{}
	var pageTitle string

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil && pageTitle == "" {
					pageTitle = strings.TrimSpace(n.FirstChild.Data)
				}
			case "meta":
				var prop, name, content string
				for _, a := range n.Attr {
					switch strings.ToLower(a.Key) {
					case "property":
						prop = strings.ToLower(a.Val)
					case "name":
						name = strings.ToLower(a.Val)
					case "content":
						content = a.Val
					}
				}
				key := prop
				if key == "" {
					key = name
				}
				if key != "" && content != "" {
					if _, exists := meta[key]; !exists {
						meta[key] = content
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	res := Result{
		Title:    firstNonEmpty(meta["og:title"], meta["twitter:title"], pageTitle),
		ImageUrl: firstNonEmpty(meta["og:image"], meta["twitter:image"], meta["twitter:image:src"]),
		Currency: firstNonEmpty(meta["og:price:currency"], meta["product:price:currency"]),
	}
	if priceStr := firstNonEmpty(meta["og:price:amount"], meta["product:price:amount"]); priceStr != "" {
		if v, err := strconv.ParseFloat(priceStr, 64); err == nil {
			res.Price = v
		}
	}
	if res.Title == "" {
		return Result{}, errors.New("could not extract title")
	}
	return res, nil
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
