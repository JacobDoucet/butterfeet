// Refetches prices for the Amazon items in a single registry and saves
// them back to each RegistryItem (priceCents + currency). Intended to be
// run from inside the backend container on the prod droplet, where
// MONGO_URI is already set, so there's no build/deploy step:
//
//	go run ./cmd/refetch-prices -slug <registry-slug>            # dry run
//	go run ./cmd/refetch-prices -slug <registry-slug> -apply     # write changes
//	go run ./cmd/refetch-prices -slug <registry-slug> -apply -force  # include items that already have a price
//
// Only items whose productUrl is an Amazon marketplace are processed; the
// scraper's per-domain currency rules already make Amazon prices reliable.
package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	registryitemapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/scrape"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	slug := flag.String("slug", "", "registry slug to refetch prices for (required)")
	apply := flag.Bool("apply", false, "actually write changes (default: dry run)")
	force := flag.Bool("force", false, "refetch even if the item already has a price")
	delay := flag.Duration("delay", 2*time.Second, "delay between requests to avoid rate limiting")
	flag.Parse()

	if strings.TrimSpace(*slug) == "" {
		fmt.Fprintln(os.Stderr, "error: -slug is required")
		flag.Usage()
		os.Exit(2)
	}

	mongoURI := envOr("MONGO_URI", "mongodb://localhost:27017")
	dbName := envOr("DB_NAME", "baby_registry")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal().Err(err).Msg("mongo connect")
	}
	defer client.Disconnect(context.Background())
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal().Err(err).Msg("mongo ping")
	}

	db := client.Database(dbName)
	apiClient := api.NewMongoBackedClient(db)
	super := permissions.NewSuperActor()
	scraper := scrape.NewHandler()

	bgCtx := context.Background()

	regResult, _, err := apiClient.Registry().Search(bgCtx, super, registry.WhereClause{
		SlugEq: slug,
	}, registryapi.QueryOptions{Limit: 1})
	if err != nil {
		log.Fatal().Err(err).Msg("search registry")
	}
	if len(regResult.Data) == 0 {
		log.Fatal().Str("slug", *slug).Msg("registry not found")
	}
	reg := regResult.Data[0]
	registryID := reg.Id

	itemResult, _, err := apiClient.RegistryItem().Search(bgCtx, super, registry_item.WhereClause{
		RegistryIdEq: &registryID,
	}, registryitemapi.QueryOptions{Limit: 1000})
	if err != nil {
		log.Fatal().Err(err).Msg("search items")
	}

	var scanned, fetched, updated, skipped, failed int
	for _, m := range itemResult.Data {
		item := m.Model
		productURL := strings.TrimSpace(item.ProductUrl)
		if productURL == "" {
			continue
		}
		if !isAmazonURL(productURL) {
			continue
		}
		scanned++

		if !*force && item.PriceCents > 0 {
			skipped++
			fmt.Printf("[skip] %s already has price %d %s\n  %s\n", item.Id, item.PriceCents, item.Currency, productURL)
			continue
		}

		res, err := scraper.Fetch(bgCtx, productURL)
		if err != nil {
			failed++
			log.Warn().Err(err).Str("itemId", item.Id).Str("url", productURL).Msg("fetch failed")
			time.Sleep(*delay)
			continue
		}
		fetched++

		newCents := priceToCents(res.Price)
		if newCents <= 0 {
			failed++
			log.Warn().Str("itemId", item.Id).Str("url", productURL).Float64("price", res.Price).Msg("no usable price scraped")
			time.Sleep(*delay)
			continue
		}
		newCurrency := strings.ToUpper(strings.TrimSpace(res.Currency))
		if newCurrency == "" {
			newCurrency = item.Currency
		}

		fmt.Printf("[%s] %s\n  old=%d %s -> new=%d %s\n",
			item.Id, productURL, item.PriceCents, item.Currency, newCents, newCurrency)

		if *apply {
			item.PriceCents = newCents
			item.Currency = newCurrency
			if _, _, err := apiClient.RegistryItem().Update(bgCtx, super, item, registry_item.NewProjection(true)); err != nil {
				failed++
				log.Error().Err(err).Str("itemId", item.Id).Msg("update failed")
				time.Sleep(*delay)
				continue
			}
			updated++
		}

		time.Sleep(*delay)
	}

	mode := "dry-run"
	if *apply {
		mode = "applied"
	}
	log.Info().
		Str("mode", mode).
		Str("slug", *slug).
		Int("amazonItems", scanned).
		Int("fetched", fetched).
		Int("updated", updated).
		Int("skipped", skipped).
		Int("failed", failed).
		Msg("price refetch complete")
}

// isAmazonURL reports whether the product URL points at an Amazon
// marketplace. Only Amazon items are processed because the scraper's
// per-domain currency handling makes those prices reliable.
func isAmazonURL(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(u.Hostname()), "amazon.")
}

// priceToCents converts a scraped price (e.g. 23.95) to integer cents,
// rounding to the nearest cent. Returns 0 for non-positive values.
func priceToCents(price float64) int {
	if price <= 0 {
		return 0
	}
	return int(math.Round(price * 100))
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}
