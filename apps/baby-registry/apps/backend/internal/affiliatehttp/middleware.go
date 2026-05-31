// Package affiliatehttp wires the affiliate package into the HTTP layer.
package affiliatehttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/affiliate"
)

// EnrichMiddleware wraps an inner handler. For POST /registry-items/create and
// PATCH /registry-items/update it parses the request body, runs the affiliate
// pipeline against `productUrl`, and forwards a body augmented with
// originalUrl/canonicalUrl/affiliateUrl/retailer fields.
//
// Any failure to parse falls through to the inner handler unchanged so a
// malformed payload still produces the normal 400 from the forge handler.
func EnrichMiddleware(reg *affiliate.Registry) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !isItemMutation(r) {
				next.ServeHTTP(w, r)
				return
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			_ = r.Body.Close()

			enriched := enrichBody(r.Context(), reg, body)
			r.Body = io.NopCloser(bytes.NewReader(enriched))
			r.ContentLength = int64(len(enriched))
			next.ServeHTTP(w, r)
		})
	}
}

func isItemMutation(r *http.Request) bool {
	p := strings.TrimSuffix(r.URL.Path, "/")
	switch {
	case r.Method == http.MethodPost && strings.HasSuffix(p, "/registry-items/create"):
		return true
	case r.Method == http.MethodPatch && strings.HasSuffix(p, "/registry-items/update"):
		return true
	}
	return false
}

// enrichBody mutates the `data` object in-place. The body shape is
// { "data": { ... } } for both create and update; we touch only the affiliate
// fields when `productUrl` is present, leaving everything else untouched.
func enrichBody(ctx context.Context, reg *affiliate.Registry, body []byte) []byte {
	var envelope map[string]any
	if err := json.Unmarshal(body, &envelope); err != nil {
		return body
	}
	data, ok := envelope["data"].(map[string]any)
	if !ok {
		return body
	}
	productURL, _ := data["productUrl"].(string)
	productURL = strings.TrimSpace(productURL)
	if productURL == "" {
		return body
	}
	res := reg.Resolve(ctx, productURL)
	data["originalUrl"] = res.OriginalURL
	data["canonicalUrl"] = res.CanonicalURL
	data["affiliateUrl"] = res.AffiliateURL
	data["retailer"] = res.Retailer

	out, err := json.Marshal(envelope)
	if err != nil {
		log.Warn().Err(err).Msg("affiliate enrich: re-marshal failed; passing original body")
		return body
	}
	return out
}
