package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_role"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_role"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	owneruserapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	cookieName      = "br_session"
	tokenTTL        = 15 * time.Minute
	sessionTTL      = 30 * 24 * time.Hour
	magicCollection = "magic_link_tokens"
)

type Config struct {
	DB         *mongo.Database
	Client     api.Client
	JWTSecret  []byte
	AppBaseURL string
}

type Service struct {
	cfg Config
}

func NewService(cfg Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/auth/magic/request", s.handleRequest)
	mux.HandleFunc("/api/auth/magic/verify", s.handleVerify)
	mux.HandleFunc("/api/auth/me", s.handleMe)
	mux.HandleFunc("/api/auth/logout", s.handleLogout)
}

type requestBody struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type verifyBody struct {
	Token string `json:"token"`
}

type meResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
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
	if email == "" || !strings.Contains(email, "@") {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}

	token, err := randomToken(32)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = s.cfg.DB.Collection(magicCollection).InsertOne(r.Context(), bson.M{
		"email":     email,
		"name":      strings.TrimSpace(body.Name),
		"token":     token,
		"expiresAt": time.Now().Add(tokenTTL),
		"createdAt": time.Now(),
		"usedAt":    nil,
	})
	if err != nil {
		log.Printf("magic-link insert error: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	link := strings.TrimRight(s.cfg.AppBaseURL, "/") + "/auth/callback?token=" + token
	// In production, send this link via email. In dev we log it.
	log.Printf("MAGIC LINK for %s -> %s", email, link)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func (s *Service) handleVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body verifyBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	body.Token = strings.TrimSpace(body.Token)
	if body.Token == "" {
		http.Error(w, "missing token", http.StatusBadRequest)
		return
	}

	coll := s.cfg.DB.Collection(magicCollection)
	var doc struct {
		Email     string     `bson:"email"`
		Name      string     `bson:"name"`
		Token     string     `bson:"token"`
		ExpiresAt time.Time  `bson:"expiresAt"`
		UsedAt    *time.Time `bson:"usedAt"`
	}
	err := coll.FindOne(r.Context(), bson.M{"token": body.Token}).Decode(&doc)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	if doc.UsedAt != nil {
		http.Error(w, "token already used", http.StatusUnauthorized)
		return
	}
	if time.Now().After(doc.ExpiresAt) {
		http.Error(w, "token expired", http.StatusUnauthorized)
		return
	}

	now := time.Now()
	_, _ = coll.UpdateOne(r.Context(), bson.M{"token": body.Token}, bson.M{"$set": bson.M{"usedAt": now}})

	super := permissions.NewSuperActor()

	// Find or create the OwnerUser.
	existing, _, err := s.cfg.Client.OwnerUser().SelectByEmailUnique(
		r.Context(),
		super,
		owner_user.SelectByEmailUniqueQuery{Email: doc.Email},
		owneruserapi.NewProjection(true),
	)
	var ownerId string
	if err != nil {
		created, _, cErr := s.cfg.Client.OwnerUser().Create(
			r.Context(),
			super,
			owner_user.Model{
				Email: doc.Email,
				Name:  doc.Name,
			},
			owner_user.NewProjection(true),
		)
		if cErr != nil {
			log.Printf("owner create error: %v", cErr)
			http.Error(w, "could not create owner", http.StatusInternalServerError)
			return
		}
		ownerId = created.Id
	} else {
		ownerId = existing.Id
	}

	sessionToken, err := s.mintJWT(ownerId)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    sessionToken,
		Path:     "/",
		Expires:  time.Now().Add(sessionTTL),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "ownerId": ownerId})
}

func (s *Service) handleMe(w http.ResponseWriter, r *http.Request) {
	actor, err := s.ResolveOwnerFromRequest(r)
	if err != nil || actor == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	m, ok := actor.(*owner_user.Model)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(meResponse{
		Id:    m.Id,
		Email: m.Email,
		Name:  m.Name,
	})
}

func (s *Service) handleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	// Also clear any buyer verification cookies so "Sign out" fully revokes
	// registry-access sessions in this browser.
	for _, c := range r.Cookies() {
		if strings.HasPrefix(c.Name, "br_buyer_") {
			http.SetCookie(w, &http.Cookie{
				Name:     c.Name,
				Value:    "",
				Path:     "/",
				Expires:  time.Unix(0, 0),
				MaxAge:   -1,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			})
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) ResolveOwnerFromRequest(r *http.Request) (permissions.Actor, error) {
	c, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(c.Value, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.cfg.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	sub, _ := claims["sub"].(string)
	if sub == "" {
		return nil, errors.New("invalid subject")
	}

	super := permissions.NewSuperActor()
	user, _, err := s.cfg.Client.OwnerUser().SelectById(
		context.Background(),
		super,
		owner_user.SelectByIdQuery{Id: sub},
		owneruserapi.NewProjection(true),
	)
	if err != nil {
		return nil, err
	}
	user.Model.ActorRoles = []actor_role.Model{{
		Role:    enum_role.Owner,
		OwnerId: user.Model.Id,
	}}
	return &user.Model, nil
}

func (s *Service) mintJWT(ownerId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": ownerId,
		"exp": time.Now().Add(sessionTTL).Unix(),
		"iat": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(s.cfg.JWTSecret)
}

func randomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
