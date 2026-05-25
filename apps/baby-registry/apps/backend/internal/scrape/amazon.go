package scrape

import (
	"encoding/json"
	"math"
	"net/url"
	"regexp"
	"strings"
)

type amazonParser struct{}

func (amazonParser) Match(host string) bool {
	h := strings.ToLower(host)
	return strings.Contains(h, "amazon.")
}

func (amazonParser) Apply(s pageSignals, res *Result) {
	if v := firstNonEmpty(s.TextByID["productTitle"], s.Meta["title"]); v != "" {
		res.Title = v
	}
	amazonImage := firstNonEmpty(
		s.ImageByID["landingImage"],
		s.ImageByID["imgBlkFront"],
		firstNonEmpty(s.AmazonDynamicImages...),
		s.Meta["image"],
	)
	if amazonImage != "" && (res.ImageUrl == "" || isGenericAmazonImage(res.ImageUrl)) {
		res.ImageUrl = amazonImage
	}
	if res.Price == 0 {
		priceCandidates := []string{
			s.Meta["product:price:amount"],
			s.Meta["og:price:amount"],
			s.Meta["twitter:data1"],
			s.Meta["price"],
			s.Meta["price:amount"],
			s.Meta["amount"],
			s.TextByID["priceblock_ourprice"],
			s.TextByID["priceblock_dealprice"],
			s.TextByID["priceblock_saleprice"],
			s.TextByID["price_inside_buybox"],
			s.TextByID["corePriceDisplay_desktop_feature_div"],
			s.TextByID["corePrice_feature_div"],
			s.TextByID["newBuyBoxPrice"],
			s.TextByID["tp_price_block_total_price_ww"],
		}
		priceCandidates = append(priceCandidates, s.PriceTexts...)
		priceCandidates = append(priceCandidates, scriptPriceCandidates(s.ScriptTexts)...)
		res.Price = firstPrice(priceCandidates...)
	}
	if res.Currency == "" {
		res.Currency = firstNonEmpty(
			s.Meta["product:price:currency"],
			s.Meta["og:price:currency"],
			s.Meta["pricecurrency"],
			s.Meta["price:currency"],
			currencyFromValue(s.Meta["twitter:data1"]),
			currencyFromValues(s.PriceTexts...),
			currencyFromValues(scriptPriceCandidates(s.ScriptTexts)...),
		)
	}
}

var scriptPricePatterns = []*regexp.Regexp{
	regexp.MustCompile(`"priceAmount"\s*:\s*"?([0-9]+(?:[\.,][0-9]{1,2})?)"?`),
	regexp.MustCompile(`"displayPrice"\s*:\s*"([^"]+)"`),
	regexp.MustCompile(`"price"\s*:\s*"?([0-9]+(?:[\.,][0-9]{1,2})?)"?`),
	regexp.MustCompile(`"amount"\s*:\s*"?([0-9]+(?:[\.,][0-9]{1,2})?)"?`),
}

func scriptPriceCandidates(scripts []string) []string {
	out := make([]string, 0, 8)
	for _, s := range scripts {
		for _, re := range scriptPricePatterns {
			matches := re.FindAllStringSubmatch(s, 3)
			for _, m := range matches {
				if len(m) > 1 {
					out = append(out, m[1])
				}
			}
		}
	}
	return out
}

func currencyFromValues(vs ...string) string {
	for _, v := range vs {
		if cur := currencyFromValue(v); cur != "" {
			return cur
		}
	}
	for _, v := range vs {
		s := strings.ToLower(strings.TrimSpace(v))
		switch {
		case strings.Contains(s, "usd"):
			return "USD"
		case strings.Contains(s, "gbp"):
			return "GBP"
		case strings.Contains(s, "eur"):
			return "EUR"
		}
	}
	return ""
}

func isGenericAmazonImage(raw string) bool {
	raw = strings.TrimSpace(strings.ToLower(raw))
	if raw == "" {
		return false
	}
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	host := strings.ToLower(u.Host)
	path := strings.ToLower(u.Path)
	if !strings.Contains(host, "amazon.") {
		return false
	}
	if strings.Contains(path, "/images/g/01/social") {
		return true
	}
	if strings.Contains(path, "logo") || strings.Contains(path, "favicon") || strings.Contains(path, "icon") {
		return true
	}
	if strings.Contains(path, "nav") || strings.Contains(path, "sprite") {
		return true
	}
	return false
}

func currencyFromValue(v string) string {
	v = strings.TrimSpace(v)
	switch {
	case strings.Contains(v, "$"):
		return "USD"
	case strings.Contains(v, "£"):
		return "GBP"
	case strings.Contains(v, "€"):
		return "EUR"
	default:
		return ""
	}
}

func parseLargestImageFromDynamicAttr(raw string) string {
	type dims []float64
	var m map[string]dims
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		return ""
	}
	bestURL := ""
	bestArea := float64(0)
	for u, d := range m {
		if len(d) < 2 {
			if bestURL == "" {
				bestURL = u
			}
			continue
		}
		area := math.Abs(d[0] * d[1])
		if area > bestArea {
			bestArea = area
			bestURL = u
		}
	}
	return bestURL
}
