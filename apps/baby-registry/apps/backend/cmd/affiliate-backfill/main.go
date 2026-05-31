// Backfills affiliateUrl/canonicalUrl/originalUrl/retailer on existing
// RegistryItem records by running each productUrl through the affiliate
// provider registry. Safe to re-run: items already enriched are skipped
// unless --force is passed.
//
// Usage (inside the backend container, where MONGO_URI is set):
//
//	go run ./cmd/affiliate-backfill                 # dry run
//	go run ./cmd/affiliate-backfill -apply          # write changes
//	go run ./cmd/affiliate-backfill -apply -force   # re-enrich everything
//	go run ./cmd/affiliate-backfill -registry <id>  # limit to one registry
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	registryitemapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/affiliate"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	apply := flag.Bool("apply", false, "actually write changes (default: dry run)")
	force := flag.Bool("force", false, "re-enrich items even if affiliateUrl is already set")
	registryID := flag.String("registry", "", "optional registry id to limit to")
	flag.Parse()

	mongoURI := envOr("MONGO_URI", "mongodb://localhost:27017")
	dbName := envOr("DB_NAME", "baby_registry")
	amazonTag := envOr("AMAZON_UK_ASSOCIATE_TAG", "butterfeetlab-21")

	affiliate.RegisterProvider(affiliate.NewAmazonUK(amazonTag))

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

	where := registry_item.WhereClause{}
	if *registryID != "" {
		where.RegistryIdEq = registryID
	}

	bgCtx := context.Background()
	result, _, err := apiClient.RegistryItem().Search(bgCtx, super, where, registryitemapi.QueryOptions{Limit: 1000})
	if err != nil {
		log.Fatal().Err(err).Msg("search items")
	}

	var scanned, updated, skipped, failed int
	for _, m := range result.Data {
		scanned++
		item := m.Model
		productURL := strings.TrimSpace(item.ProductUrl)
		if productURL == "" {
			skipped++
			continue
		}
		if !*force && strings.TrimSpace(item.AffiliateUrl) != "" {
			skipped++
			continue
		}

		res := affiliate.Default.Resolve(bgCtx, productURL)
		fmt.Printf("[%s] retailer=%s\n  product=%s\n  canonical=%s\n  affiliate=%s\n",
			item.Id, res.Retailer, productURL, res.CanonicalURL, res.AffiliateURL)

		if !*apply {
			continue
		}

		item.OriginalUrl = res.OriginalURL
		item.CanonicalUrl = res.CanonicalURL
		item.AffiliateUrl = res.AffiliateURL
		item.Retailer = res.Retailer

		if _, _, err := apiClient.RegistryItem().Update(bgCtx, super, item, registry_item.NewProjection(true)); err != nil {
			failed++
			log.Error().Err(err).Str("itemId", item.Id).Msg("update failed")
			continue
		}
		updated++
	}

	mode := "dry-run"
	if *apply {
		mode = "applied"
	}
	log.Info().
		Str("mode", mode).
		Int("scanned", scanned).
		Int("updated", updated).
		Int("skipped", skipped).
		Int("failed", failed).
		Msg("backfill complete")
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}
