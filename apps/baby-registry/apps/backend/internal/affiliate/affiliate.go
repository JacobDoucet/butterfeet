// Package affiliate provides a retailer-agnostic interface for detecting,
// normalizing, and rewriting product URLs into affiliate-tagged links.
package affiliate

import (
	"context"
	"sync"
)

// Provider is the contract every retailer adapter implements.
type Provider interface {
	// Name returns a stable identifier for the retailer (e.g. "amazon_uk").
	Name() string
	// Match reports whether the provider can handle the given raw URL.
	Match(rawURL string) bool
	// Normalize resolves and rewrites the URL into a canonical, tag-free form.
	Normalize(ctx context.Context, rawURL string) (string, error)
	// GenerateAffiliateURL appends/sets the provider's affiliate tag.
	GenerateAffiliateURL(normalizedURL string) (string, error)
}

// Resolution is the full result of running a URL through the pipeline.
type Resolution struct {
	OriginalURL  string
	CanonicalURL string
	AffiliateURL string
	Retailer     string
}

// Registry is a thread-safe collection of providers.
type Registry struct {
	mu        sync.RWMutex
	providers []Provider
}

// NewRegistry returns an empty Registry.
func NewRegistry() *Registry { return &Registry{} }

// Register adds a provider to the registry. First-registered wins on Match.
func (r *Registry) Register(p Provider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.providers = append(r.providers, p)
}

// Find returns the first provider that matches the URL, or nil.
func (r *Registry) Find(rawURL string) Provider {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, p := range r.providers {
		if p.Match(rawURL) {
			return p
		}
	}
	return nil
}

// Resolve runs the full pipeline against the URL.
// If no provider matches, the original URL is returned in all three slots with
// retailer set to "unknown". Errors during normalization fall back to using the
// raw URL so item creation never fails because of affiliate rewriting.
func (r *Registry) Resolve(ctx context.Context, rawURL string) Resolution {
	if rawURL == "" {
		return Resolution{}
	}
	p := r.Find(rawURL)
	if p == nil {
		return Resolution{
			OriginalURL:  rawURL,
			CanonicalURL: rawURL,
			AffiliateURL: rawURL,
			Retailer:     "unknown",
		}
	}
	canonical, err := p.Normalize(ctx, rawURL)
	if err != nil || canonical == "" {
		canonical = rawURL
	}
	affiliate, err := p.GenerateAffiliateURL(canonical)
	if err != nil || affiliate == "" {
		affiliate = canonical
	}
	return Resolution{
		OriginalURL:  rawURL,
		CanonicalURL: canonical,
		AffiliateURL: affiliate,
		Retailer:     p.Name(),
	}
}

// Default is the process-wide registry consumed by HTTP middleware.
var Default = NewRegistry()

// RegisterProvider registers a provider on the default registry.
func RegisterProvider(p Provider) { Default.Register(p) }
