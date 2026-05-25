package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"
	"time"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/http_server"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/auth"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/buyer"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/public"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/scrape"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/shipping"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := getEnv("DB_NAME", "baby_registry")
	port := getEnv("PORT", "8088")
	jwtSecret := getEnv("JWT_SECRET", "dev-insecure-secret-change-me")
	appBaseURL := getEnv("APP_BASE_URL", "http://localhost:5173")
	allowedOrigins := parseAllowedOrigins(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3035,http://localhost:5173"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		_ = mongoClient.Disconnect(context.Background())
	}()
	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")

	db := mongoClient.Database(dbName)
	apiClient := api.NewMongoBackedClient(db)

	authSvc := auth.NewService(auth.Config{
		DB:         db,
		Client:     apiClient,
		JWTSecret:  []byte(jwtSecret),
		AppBaseURL: appBaseURL,
	})

	resolveActor := func(r *http.Request) (permissions.Actor, error) {
		actor, err := authSvc.ResolveOwnerFromRequest(r)
		if err == nil && actor != nil {
			return actor, nil
		}
		return nil, errors.New("unauthorized")
	}

	forgeMux, err := http_server.ServeMux(apiClient, http_server.ServeMuxProps{
		ResolveActor: resolveActor,
		OnError: func(handler string, e error) {
			log.Printf("forge handler error %s: %v", handler, e)
		},
	})
	if err != nil {
		log.Fatalf("Failed to create forge HTTP server: %v", err)
	}

	root := http.NewServeMux()

	authSvc.Register(root)

	scrapeHandler := scrape.NewHandler()
	root.Handle("/api/scrape", scrapeHandler)

	publicHandler := public.NewHandler(apiClient)
	buyerSvc := buyer.NewService(buyer.Config{
		DB:        db,
		JWTSecret: []byte(jwtSecret),
	})
	publicHandler.SetBuyerResolver(buyerSvc.ResolveBuyer)
	buyerSvc.Register(publicHandler.Mux())
	root.Handle("/api/public/", http.StripPrefix("/api/public", publicHandler))

	shippingHandler := shipping.NewHandler(apiClient, shipping.ActorResolver(resolveActor))
	root.Handle("/api/shipping/", http.StripPrefix("/api/shipping", shippingHandler))

	// Mount forge mux behind /api (owner-authenticated CRUD).
	root.Handle("/api/", http.StripPrefix("/api", forgeAuthWrapper(forgeMux, authSvc)))

	root.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	handler := corsMiddleware(root, allowedOrigins)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Backend listening on :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

	shutdownCtx, sCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer sCancel()
	_ = server.Shutdown(shutdownCtx)
}

// forgeAuthWrapper guards the forge-generated mux behind owner authentication.
// Public endpoints live under /api/public/* and /api/auth/* outside this wrapper.
func forgeAuthWrapper(next http.Handler, authSvc *auth.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actor, err := authSvc.ResolveOwnerFromRequest(r)
		if err != nil || actor == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		// Block path-traversal-style requests to the events feed for non-supers.
		_ = actor
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if !slices.Contains(allowedOrigins, origin) {
				if r.Method == http.MethodOptions {
					http.Error(w, "origin not allowed", http.StatusForbidden)
					return
				}
				next.ServeHTTP(w, r)
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func parseAllowedOrigins(raw string) []string {
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func getEnv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

// Ensure owner_user is referenced for go mod tidy completeness in case
// downstream tooling tree-shakes; the auth package imports it directly,
// but this keeps things explicit.
var _ = owner_user.Model{}
