package buyer

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	otpCollection = "buyer_otp_codes"
	otpTTL        = 15 * time.Minute
	sessionTTL    = 30 * 24 * time.Hour
	cookiePrefix  = "br_buyer_"
	// maxOtpAttempts caps how many wrong codes a (slug,email) pair can submit
	// before having to request a new one. Keeps brute-force costs visible.
	maxOtpAttempts = 6
)

type Config struct {
	DB        *mongo.Database
	JWTSecret []byte
}

type Service struct {
	cfg Config
}

func NewService(cfg Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("/buyer/verify/request", s.handleRequest)
	mux.HandleFunc("/buyer/verify/confirm", s.handleConfirm)
	mux.HandleFunc("/buyer/me", s.handleMe)
	mux.HandleFunc("/buyer/logout", s.handleLogout)
}

type requestBody struct {
	Email string `json:"email"`
	Slug  string `json:"slug"`
}

type confirmBody struct {
	Email string `json:"email"`
	Slug  string `json:"slug"`
	Code  string `json:"code"`
}

func (s *Service) handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body requestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	email := strings.ToLower(strings.TrimSpace(body.Email))
	slug := strings.TrimSpace(body.Slug)
	if email == "" || !strings.Contains(email, "@") || slug == "" {
		http.Error(w, "invalid email or slug", http.StatusBadRequest)
		return
	}

	code, err := randomDigits(6)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	codeHash := sha256Hex(code)
	now := time.Now()
	_, err = s.cfg.DB.Collection(otpCollection).UpdateOne(r.Context(),
		bson.M{"email": email, "slug": slug},
		bson.M{
			"$set": bson.M{
				"email":     email,
				"slug":      slug,
				"codeHash":  codeHash,
				"expiresAt": now.Add(otpTTL),
				"createdAt": now,
				"attempts":  0,
			},
		},
		mongoUpsert(),
	)
	if err != nil {
		log.Printf("buyer otp upsert error: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Dev: log the code. In prod with SMTP wired up, this would email the buyer.
	log.Printf("BUYER OTP for %s (slug=%s) -> %s", email, slug, code)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func (s *Service) handleConfirm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body confirmBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	email := strings.ToLower(strings.TrimSpace(body.Email))
	slug := strings.TrimSpace(body.Slug)
	code := strings.TrimSpace(body.Code)
	if email == "" || slug == "" || code == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	coll := s.cfg.DB.Collection(otpCollection)
	var doc struct {
		CodeHash  string    `bson:"codeHash"`
		ExpiresAt time.Time `bson:"expiresAt"`
		Attempts  int       `bson:"attempts"`
	}
	err := coll.FindOne(r.Context(), bson.M{"email": email, "slug": slug}).Decode(&doc)
	if err != nil {
		http.Error(w, "invalid code", http.StatusUnauthorized)
		return
	}
	if doc.Attempts >= maxOtpAttempts {
		http.Error(w, "too many attempts; request a new code", http.StatusTooManyRequests)
		return
	}
	if time.Now().After(doc.ExpiresAt) {
		http.Error(w, "code expired", http.StatusUnauthorized)
		return
	}
	if sha256Hex(code) != doc.CodeHash {
		_, _ = coll.UpdateOne(r.Context(),
			bson.M{"email": email, "slug": slug},
			bson.M{"$inc": bson.M{"attempts": 1}},
		)
		http.Error(w, "invalid code", http.StatusUnauthorized)
		return
	}

	_, _ = coll.DeleteOne(r.Context(), bson.M{"email": email, "slug": slug})

	token, err := s.mintJWT(email, slug)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, buyerCookie(slug, token, sessionTTL))

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "email": email})
}

func (s *Service) handleMe(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimSpace(r.URL.Query().Get("slug"))
	if slug == "" {
		http.Error(w, "missing slug", http.StatusBadRequest)
		return
	}
	email, err := s.ResolveBuyer(r, slug)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"email": email})
}

func (s *Service) handleLogout(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimSpace(r.URL.Query().Get("slug"))
	if slug == "" {
		http.Error(w, "missing slug", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName(slug),
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	w.WriteHeader(http.StatusOK)
}

// ResolveBuyer returns the verified email associated with a buyer cookie for
// the given registry slug, or an error if no valid cookie is present.
func (s *Service) ResolveBuyer(r *http.Request, slug string) (string, error) {
	c, err := r.Cookie(cookieName(slug))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(c.Value, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.cfg.JWTSecret, nil
	})
	if err != nil {
		return "", err
	}
	aud, _ := claims["aud"].(string)
	if aud != "buyer:"+slug {
		return "", errors.New("audience mismatch")
	}
	email, _ := claims["email"].(string)
	if email == "" {
		return "", errors.New("missing email")
	}
	return email, nil
}

func (s *Service) mintJWT(email, slug string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"aud":   "buyer:" + slug,
		"exp":   time.Now().Add(sessionTTL).Unix(),
		"iat":   time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(s.cfg.JWTSecret)
}

func cookieName(slug string) string {
	// Cookie names may not contain certain chars; slugs are already URL-safe.
	return cookiePrefix + slug
}

func buyerCookie(slug, value string, ttl time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName(slug),
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(ttl),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func randomDigits(n int) (string, error) {
	digits := make([]byte, n)
	for i := 0; i < n; i++ {
		v, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		digits[i] = byte('0' + v.Int64())
	}
	return string(digits), nil
}

func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// mongoUpsert builds the upsert option helper.
func mongoUpsert() *options.UpdateOptions {
	return options.Update().SetUpsert(true)
}

var _ = mongo.ErrNoDocuments
