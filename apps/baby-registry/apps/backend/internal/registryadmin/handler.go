package registryadmin

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	registryitemmongo "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActorResolver func(*http.Request) (permissions.Actor, error)

type Handler struct {
	mux      *http.ServeMux
	client   api.Client
	db       *mongo.Database
	resolver ActorResolver
}

func NewHandler(client api.Client, db *mongo.Database, resolver ActorResolver) *Handler {
	h := &Handler{client: client, db: db, resolver: resolver}
	mux := http.NewServeMux()
	mux.HandleFunc("/rename-category", h.handleRenameCategory)
	h.mux = mux
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

type renameCategoryBody struct {
	RegistryId  string `json:"registryId"`
	OldCategory string `json:"oldCategory"`
	NewCategory string `json:"newCategory"`
}

func (h *Handler) handleRenameCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	actor, err := h.requireOwner(w, r)
	if err != nil {
		return
	}
	var body renameCategoryBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid body")
		return
	}
	body.RegistryId = strings.TrimSpace(body.RegistryId)
	oldCat := strings.TrimSpace(body.OldCategory)
	newCat := strings.TrimSpace(body.NewCategory)
	if body.RegistryId == "" {
		writeJSONError(w, http.StatusBadRequest, "registryId required")
		return
	}
	if oldCat == newCat {
		writeJSON(w, http.StatusOK, map[string]any{"ok": true, "modifiedCount": 0})
		return
	}

	// Authorise: ensure the actor can read this registry (owner ACL via forge).
	if _, _, err := h.client.Registry().SelectById(
		r.Context(), actor,
		registry.SelectByIdQuery{Id: body.RegistryId},
		registryapi.NewProjection(true),
	); err != nil {
		writeJSONError(w, http.StatusNotFound, "registry not found")
		return
	}

	coll := h.db.Collection(registryitemmongo.CollectionName)

	regOID, err := primitive.ObjectIDFromHex(body.RegistryId)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid registryId")
		return
	}

	var filter bson.M
	if oldCat == "" {
		filter = bson.M{
			"registryId": regOID,
			"$or": []bson.M{
				{"category": ""},
				{"category": nil},
				{"category": bson.M{"$exists": false}},
			},
		}
	} else {
		filter = bson.M{"registryId": regOID, "category": oldCat}
	}

	var update bson.M
	if newCat == "" {
		update = bson.M{"$unset": bson.M{"category": ""}}
	} else {
		update = bson.M{"$set": bson.M{"category": newCat}}
	}

	res, err := coll.UpdateMany(r.Context(), filter, update)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "modifiedCount": res.ModifiedCount})
}

func (h *Handler) requireOwner(w http.ResponseWriter, r *http.Request) (permissions.Actor, error) {
	actor, err := h.resolver(r)
	if err != nil || actor == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}
	if _, ok := actor.(*owner_user.Model); !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}
	return actor, nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]any{"error": msg})
}
