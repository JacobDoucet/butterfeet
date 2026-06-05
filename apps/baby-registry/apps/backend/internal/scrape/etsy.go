package scrape

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

// etsyListingIDRE pulls the numeric listing id out of an Etsy product URL,
// e.g. https://www.etsy.com/uk/listing/4440554569/some-slug -> 4440554569.
var etsyListingIDRE = regexp.MustCompile(`/listing/(\d+)`)

func isEtsyURL(host string) bool {
	return strings.Contains(strings.ToLower(host), "etsy.")
}

func etsyListingID(path string) string {
	if m := etsyListingIDRE.FindStringSubmatch(path); len(m) == 2 {
		return m[1]
	}
	return ""
}

type etsyListingResponse struct {
	ListingID int64  `json:"listing_id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Price     struct {
		Amount       int64  `json:"amount"`
		Divisor      int64  `json:"divisor"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price"`
	Images []struct {
		URLFullxfull string `json:"url_fullxfull"`
	} `json:"images"`
}

// fetchEtsyAPI resolves an Etsy listing via the official Open API v3. Unlike
// the listing HTML pages (which sit behind DataDome bot protection and return
// a captcha challenge to server IPs), the Open API serves clean JSON. It only
// needs an app keystring in the ETSY_API_KEY env var for public listing data
// (no per-user OAuth). Returns an error so callers fall through to HTML
// scraping when the key is missing or the request fails.
func (h *Handler) fetchEtsyAPI(ctx context.Context, listingID string) (Result, error) {
	key := strings.TrimSpace(os.Getenv("ETSY_API_KEY"))
	if key == "" {
		return Result{}, fmt.Errorf("etsy: ETSY_API_KEY not set")
	}

	endpoint := fmt.Sprintf("https://openapi.etsy.com/v3/application/listings/%s?includes=Images", listingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return Result{}, err
	}
	req.Header.Set("x-api-key", key)
	req.Header.Set("Accept", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1*1024*1024))
	if err != nil {
		return Result{}, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Warn().Str("listingId", listingID).Int("status", resp.StatusCode).Str("body", string(body)).Msg("scrape: etsy api error")
		return Result{}, fmt.Errorf("etsy api: status=%d", resp.StatusCode)
	}

	var lr etsyListingResponse
	if err := json.Unmarshal(body, &lr); err != nil {
		return Result{}, err
	}

	res := Result{
		Title:    strings.TrimSpace(lr.Title),
		Source:   "Etsy",
		Currency: strings.ToUpper(strings.TrimSpace(lr.Price.CurrencyCode)),
	}
	if lr.Price.Divisor > 0 {
		res.Price = float64(lr.Price.Amount) / float64(lr.Price.Divisor)
	}
	if len(lr.Images) > 0 {
		res.ImageUrl = strings.TrimSpace(lr.Images[0].URLFullxfull)
	}
	return res, nil
}
