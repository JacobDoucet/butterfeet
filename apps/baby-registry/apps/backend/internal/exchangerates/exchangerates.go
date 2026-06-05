// Package exchangerates maintains a cached table of foreign-exchange rates so
// the public registry page can display item prices in a viewer-chosen
// currency. Rates are sourced from frankfurter.dev (ECB daily reference rates,
// free, no API key) on a slow poll and persisted to a single Mongo document so
// every request is served from cache rather than a live upstream call.
package exchangerates

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Base is the currency all stored rates are expressed against. A rate of N for
// currency X means 1 Base = N X.
const Base = "USD"

// collectionName is the Mongo collection holding the single cached rates doc.
const collectionName = "exchange_rates"

// docID is the fixed _id of the singleton rates document.
const docID = "latest"

// providerURL fetches the latest rates with USD as the base currency.
const providerURL = "https://api.frankfurter.dev/v1/latest?base=" + Base

// Snapshot is the cached set of rates returned to callers.
type Snapshot struct {
	Base      string             `bson:"base" json:"base"`
	Rates     map[string]float64 `bson:"rates" json:"rates"`
	FetchedAt time.Time          `bson:"fetchedAt" json:"fetchedAt"`
	Date      string             `bson:"date" json:"date"`
}

// Config controls the poller.
type Config struct {
	DB *mongo.Database
	// Interval between upstream refreshes. Defaults to 12h.
	Interval time.Duration
}

func (c *Config) defaults() {
	if c.Interval <= 0 {
		c.Interval = 12 * time.Hour
	}
}

// Store reads the cached rates snapshot. Callers (e.g. the public HTTP handler)
// use this to serve rates without hitting the upstream provider.
type Store struct {
	db *mongo.Database
}

// NewStore returns a read accessor for the cached rates.
func NewStore(db *mongo.Database) *Store { return &Store{db: db} }

// Get returns the cached rates snapshot, or false if none has been stored yet.
func (s *Store) Get(ctx context.Context) (Snapshot, bool) {
	var snap Snapshot
	err := s.db.Collection(collectionName).FindOne(ctx, bson.M{"_id": docID}).Decode(&snap)
	if err != nil {
		return Snapshot{}, false
	}
	return snap, true
}

// Run blocks until ctx is cancelled, refreshing the cached rates on startup and
// then on each interval tick.
func Run(ctx context.Context, cfg Config) {
	cfg.defaults()
	if cfg.DB == nil {
		log.Warn().Msg("exchangerates.Run skipped: db is nil")
		return
	}

	tick := time.NewTicker(cfg.Interval)
	defer tick.Stop()

	refresh(ctx, cfg)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			refresh(ctx, cfg)
		}
	}
}

type providerResponse struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func refresh(ctx context.Context, cfg Config) {
	fetchCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	snap, err := fetch(fetchCtx)
	if err != nil {
		log.Error().Err(err).Msg("exchangerates: refresh failed")
		return
	}

	opts := options.Update().SetUpsert(true)
	_, err = cfg.DB.Collection(collectionName).UpdateOne(
		fetchCtx,
		bson.M{"_id": docID},
		bson.M{"$set": bson.M{
			"base":      snap.Base,
			"rates":     snap.Rates,
			"date":      snap.Date,
			"fetchedAt": snap.FetchedAt,
		}},
		opts,
	)
	if err != nil {
		log.Error().Err(err).Msg("exchangerates: persist failed")
		return
	}
	log.Info().Int("currencies", len(snap.Rates)).Str("date", snap.Date).Msg("exchangerates: refreshed")
}

func fetch(ctx context.Context) (Snapshot, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, providerURL, nil)
	if err != nil {
		return Snapshot{}, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Snapshot{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return Snapshot{}, fmt.Errorf("exchangerates: provider returned status %d", resp.StatusCode)
	}

	var pr providerResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return Snapshot{}, err
	}
	if len(pr.Rates) == 0 {
		return Snapshot{}, fmt.Errorf("exchangerates: provider returned no rates")
	}

	// Ensure the base currency itself is present with a 1.0 rate so callers can
	// convert to/from the base without special-casing.
	rates := make(map[string]float64, len(pr.Rates)+1)
	for k, v := range pr.Rates {
		rates[k] = v
	}
	rates[Base] = 1

	return Snapshot{
		Base:      Base,
		Rates:     rates,
		Date:      pr.Date,
		FetchedAt: time.Now().UTC(),
	}, nil
}
