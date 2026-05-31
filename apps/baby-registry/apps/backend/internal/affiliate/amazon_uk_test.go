package affiliate

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAmazonUK_Match(t *testing.T) {
	p := NewAmazonUK("butterfeetlab-21")
	cases := []struct {
		url  string
		want bool
	}{
		{"https://www.amazon.co.uk/dp/B0ABC12345", true},
		{"https://amazon.co.uk/gp/product/B0ABC12345", true},
		{"https://amzn.eu/d/abc123", true},
		{"http://www.amazon.co.uk/anything", true},
		{"https://www.amazon.com/dp/B0ABC12345", false},
		{"https://amazon.de/dp/B0ABC12345", false},
		{"https://www.johnlewis.com/foo", false},
		{"", false},
		{"not a url", false},
	}
	for _, tc := range cases {
		t.Run(tc.url, func(t *testing.T) {
			if got := p.Match(tc.url); got != tc.want {
				t.Fatalf("Match(%q)=%v want %v", tc.url, got, tc.want)
			}
		})
	}
}

func TestAmazonUK_Normalize(t *testing.T) {
	p := NewAmazonUK("butterfeetlab-21")
	ctx := context.Background()
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"dp basic", "https://www.amazon.co.uk/dp/B0ABC12345", "https://www.amazon.co.uk/dp/B0ABC12345"},
		{"dp with query", "https://www.amazon.co.uk/dp/B0ABC12345?th=1&psc=1", "https://www.amazon.co.uk/dp/B0ABC12345"},
		{"gp product", "https://www.amazon.co.uk/gp/product/B0ABC12345/ref=something?tag=other-21", "https://www.amazon.co.uk/dp/B0ABC12345"},
		{"dp with seo slug", "https://www.amazon.co.uk/Some-Product-Name/dp/B0ABC12345/ref=sr_1_1", "https://www.amazon.co.uk/dp/B0ABC12345"},
		{"lowercase asin uppercased", "https://www.amazon.co.uk/dp/b0abc12345", "https://www.amazon.co.uk/dp/B0ABC12345"},
		{"bare domain no asin", "https://www.amazon.co.uk/gp/cart?ref=x", "https://www.amazon.co.uk/gp/cart"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p.Normalize(ctx, tc.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("Normalize(%q)\n got: %s\nwant: %s", tc.in, got, tc.want)
			}
		})
	}
}

func TestAmazonUK_NormalizeShortLink(t *testing.T) {
	// Simulate amzn.eu redirecting to the canonical product page.
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer target.Close()

	redirector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://www.amazon.co.uk/dp/B0ABC12345?ref=short", http.StatusFound)
	}))
	defer redirector.Close()

	p := NewAmazonUK("butterfeetlab-21")
	// Override the host check by feeding through the public API. We hit the
	// real Normalize path that calls resolveShortLink against our fake URL;
	// since our fake host doesn't match the short-link host suffix, we
	// instead test resolveShortLink directly with the redirector and then
	// the Amazon path normalization.
	final, err := p.resolveShortLink(context.Background(), redirector.URL)
	if err != nil {
		t.Fatalf("resolveShortLink: %v", err)
	}
	if !strings.Contains(final, "/dp/B0ABC12345") {
		t.Fatalf("expected redirect followed, got %s", final)
	}
}

func TestAmazonUK_GenerateAffiliateURL(t *testing.T) {
	p := NewAmazonUK("butterfeetlab-21")
	cases := []struct {
		in   string
		want string
	}{
		{"https://www.amazon.co.uk/dp/B0ABC12345", "https://www.amazon.co.uk/dp/B0ABC12345?tag=butterfeetlab-21"},
		{"https://www.amazon.co.uk/dp/B0ABC12345?th=1", "https://www.amazon.co.uk/dp/B0ABC12345?tag=butterfeetlab-21&th=1"},
		{"https://www.amazon.co.uk/dp/B0ABC12345?tag=stale-21", "https://www.amazon.co.uk/dp/B0ABC12345?tag=butterfeetlab-21"},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got, err := p.GenerateAffiliateURL(tc.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("GenerateAffiliateURL(%q)\n got: %s\nwant: %s", tc.in, got, tc.want)
			}
		})
	}
}

func TestAmazonUK_GenerateAffiliateURL_NoTag(t *testing.T) {
	p := NewAmazonUK("")
	got, err := p.GenerateAffiliateURL("https://www.amazon.co.uk/dp/B0ABC12345")
	if err != nil {
		t.Fatal(err)
	}
	if got != "https://www.amazon.co.uk/dp/B0ABC12345" {
		t.Fatalf("expected no-op when tag empty, got %s", got)
	}
}

func TestRegistry_Resolve_AmazonUK(t *testing.T) {
	r := NewRegistry()
	r.Register(NewAmazonUK("butterfeetlab-21"))

	res := r.Resolve(context.Background(), "https://www.amazon.co.uk/gp/product/B0ABC12345/ref=sr?tag=other-21")
	if res.Retailer != "amazon_uk" {
		t.Fatalf("retailer = %q want amazon_uk", res.Retailer)
	}
	if res.CanonicalURL != "https://www.amazon.co.uk/dp/B0ABC12345" {
		t.Fatalf("canonical = %s", res.CanonicalURL)
	}
	if res.AffiliateURL != "https://www.amazon.co.uk/dp/B0ABC12345?tag=butterfeetlab-21" {
		t.Fatalf("affiliate = %s", res.AffiliateURL)
	}
	if res.OriginalURL == res.CanonicalURL {
		t.Fatalf("original should preserve input")
	}
}

func TestRegistry_Resolve_UnknownRetailer(t *testing.T) {
	r := NewRegistry()
	r.Register(NewAmazonUK("butterfeetlab-21"))

	res := r.Resolve(context.Background(), "https://www.johnlewis.com/product/123")
	if res.Retailer != "unknown" {
		t.Fatalf("retailer = %s want unknown", res.Retailer)
	}
	if res.AffiliateURL != "https://www.johnlewis.com/product/123" {
		t.Fatalf("affiliate should pass through, got %s", res.AffiliateURL)
	}
}

func TestRegistry_Resolve_Empty(t *testing.T) {
	r := NewRegistry()
	r.Register(NewAmazonUK("butterfeetlab-21"))
	res := r.Resolve(context.Background(), "")
	if res != (Resolution{}) {
		t.Fatalf("expected zero Resolution for empty input, got %+v", res)
	}
}

func TestRegistry_Resolve_InvalidURL(t *testing.T) {
	r := NewRegistry()
	r.Register(NewAmazonUK("butterfeetlab-21"))
	res := r.Resolve(context.Background(), "not a url")
	if res.Retailer != "unknown" {
		t.Fatalf("retailer = %s want unknown", res.Retailer)
	}
	if res.AffiliateURL != "not a url" {
		t.Fatalf("affiliate passthrough expected")
	}
}
