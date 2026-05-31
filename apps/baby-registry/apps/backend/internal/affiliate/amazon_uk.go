package affiliate

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// AmazonUK is the Amazon Associates UK provider.
type AmazonUK struct {
	// Tag is the Amazon Associates tracking id (e.g. "butterfeetlab-21").
	Tag string
	// HTTPClient is used to resolve short links (amzn.eu/d/...). If nil, a
	// safe default with a short timeout is used.
	HTTPClient *http.Client
}

// NewAmazonUK returns a configured Amazon UK provider.
func NewAmazonUK(tag string) *AmazonUK {
	return &AmazonUK{
		Tag: tag,
		HTTPClient: &http.Client{
			Timeout: 8 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 10 {
					return errors.New("too many redirects")
				}
				return nil
			},
		},
	}
}

func (a *AmazonUK) Name() string { return "amazon_uk" }

var amazonHostPattern = regexp.MustCompile(`(^|\.)(amazon\.co\.uk|amzn\.eu|amzn\.to)$`)

// Match returns true for amazon.co.uk and the amzn.eu / amzn.to short
// domains. Other Amazon TLDs (.com, .de, etc.) are intentionally excluded —
// the associate tag is region-specific.
func (a *AmazonUK) Match(rawURL string) bool {
	u, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil || u == nil || u.Host == "" {
		return false
	}
	host := strings.ToLower(u.Hostname())
	return amazonHostPattern.MatchString(host)
}

var asinPattern = regexp.MustCompile(`(?i)/(?:dp|gp/product|gp/aw/d|product)/([A-Z0-9]{10})`)

// Normalize collapses an Amazon URL down to https://www.amazon.co.uk/dp/<ASIN>,
// resolving amzn.eu/amzn.to short links via a single redirect-following GET.
func (a *AmazonUK) Normalize(ctx context.Context, rawURL string) (string, error) {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return "", errors.New("empty url")
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	host := strings.ToLower(u.Hostname())

	// Resolve short links by following redirects.
	if strings.HasSuffix(host, "amzn.eu") || strings.HasSuffix(host, "amzn.to") {
		resolved, err := a.resolveShortLink(ctx, rawURL)
		if err != nil || resolved == "" {
			// Caller (Registry.Resolve) will fall back to the original URL.
			return "", fmt.Errorf("resolve short link: %w", err)
		}
		u, err = url.Parse(resolved)
		if err != nil {
			return "", err
		}
		host = strings.ToLower(u.Hostname())
	}

	if asin := extractASIN(u.Path); asin != "" {
		return "https://www.amazon.co.uk/dp/" + asin, nil
	}
	// Couldn't find an ASIN — return the URL with tracking params stripped
	// so we still produce a tidy canonical form.
	if !strings.HasPrefix(host, "www.") && (host == "amazon.co.uk") {
		u.Host = "www.amazon.co.uk"
	}
	u.RawQuery = ""
	u.Fragment = ""
	return u.String(), nil
}

// GenerateAffiliateURL appends our associate tag, replacing any existing tag.
func (a *AmazonUK) GenerateAffiliateURL(normalizedURL string) (string, error) {
	if a.Tag == "" {
		return normalizedURL, nil
	}
	u, err := url.Parse(normalizedURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("tag", a.Tag)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func extractASIN(path string) string {
	m := asinPattern.FindStringSubmatch(path)
	if len(m) < 2 {
		return ""
	}
	return strings.ToUpper(m[1])
}

func (a *AmazonUK) resolveShortLink(ctx context.Context, rawURL string) (string, error) {
	client := a.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; StorkNestBot/1.0)")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.Request == nil || resp.Request.URL == nil {
		return "", errors.New("no final url")
	}
	return resp.Request.URL.String(), nil
}
