package public

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_item_source"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	registryitemapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	reservationapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_api"
)

type Handler struct {
	mux    *http.ServeMux
	client api.Client
}

func NewHandler(client api.Client) *Handler {
	h := &Handler{client: client}
	mux := http.NewServeMux()
	mux.HandleFunc("/r/", h.handleRegistryBySlug)
	mux.HandleFunc("/items/", h.handleItemRoute) // /items/:id/reserve
	h.mux = mux
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

type publicItem struct {
	Id         string  `json:"id"`
	Title      string  `json:"title"`
	Description string `json:"description"`
	ImageUrl   string  `json:"imageUrl"`
	ProductUrl string  `json:"productUrl"`
	Source     string  `json:"source"`
	PriceCents int     `json:"priceCents"`
	Currency   string  `json:"currency"`
	Quantity   int     `json:"quantity"`
	Notes      string  `json:"notes"`
	Position   int     `json:"position"`
	Reserved   int     `json:"reserved"`
}

type publicRegistry struct {
	Id             string       `json:"id"`
	Slug           string       `json:"slug"`
	Title          string       `json:"title"`
	ParentNames    string       `json:"parentNames"`
	WelcomeMessage string       `json:"welcomeMessage"`
	ThemeColor     string       `json:"themeColor"`
	CoverImageUrl  string       `json:"coverImageUrl"`
	Items          []publicItem `json:"items"`
}

func (h *Handler) handleRegistryBySlug(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	slug := strings.TrimPrefix(r.URL.Path, "/r/")
	slug = strings.Trim(slug, "/")
	if slug == "" {
		http.Error(w, "missing slug", http.StatusBadRequest)
		return
	}

	super := permissions.NewSuperActor()

	regResult, _, err := h.client.Registry().Search(r.Context(), super, registry.WhereClause{
		SlugEq: &slug,
	}, registryapi.QueryOptions{Limit: 1})
	if err != nil || len(regResult.Data) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	reg := regResult.Data[0]
	if !reg.IsPublic {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	itemsResult, _, err := h.client.RegistryItem().Search(r.Context(), super, registry_item.WhereClause{
		RegistryIdEq: &reg.Id,
	}, registryitemapi.QueryOptions{Limit: 500})
	if err != nil {
		http.Error(w, "lookup error", http.StatusInternalServerError)
		return
	}

	// Count reservations per item.
	resvResult, _, err := h.client.Reservation().Search(r.Context(), super, reservation.WhereClause{
		RegistryIdEq: &reg.Id,
	}, reservationapi.QueryOptions{Limit: 1000})
	if err != nil {
		http.Error(w, "lookup error", http.StatusInternalServerError)
		return
	}
	reservedByItem := map[string]int{}
	for _, rsv := range resvResult.Data {
		if rsv.Status == enum_reservation_status.Cancelled {
			continue
		}
		q := rsv.Quantity
		if q <= 0 {
			q = 1
		}
		reservedByItem[rsv.ItemId] += q
	}

	publicItems := make([]publicItem, 0, len(itemsResult.Data))
	for _, it := range itemsResult.Data {
		publicItems = append(publicItems, publicItem{
			Id:          it.Id,
			Title:       it.Title,
			Description: it.Description,
			ImageUrl:    it.ImageUrl,
			ProductUrl:  it.ProductUrl,
			Source:      string(it.Source),
			PriceCents:  it.PriceCents,
			Currency:    it.Currency,
			Quantity:    it.Quantity,
			Notes:       it.Notes,
			Position:    it.Position,
			Reserved:    reservedByItem[it.Id],
		})
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(publicRegistry{
		Id:             reg.Id,
		Slug:           reg.Slug,
		Title:          reg.Title,
		ParentNames:    reg.ParentNames,
		WelcomeMessage: reg.WelcomeMessage,
		ThemeColor:     reg.ThemeColor,
		CoverImageUrl:  reg.CoverImageUrl,
		Items:          publicItems,
	})
}

type reserveBody struct {
	ReserverName string `json:"reserverName"`
	IsAnonymous  bool   `json:"isAnonymous"`
	Message      string `json:"message"`
	ContactEmail string `json:"contactEmail"`
	Quantity     int    `json:"quantity"`
}

// handleItemRoute handles /items/:id/reserve
func (h *Handler) handleItemRoute(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/items/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 || parts[1] != "reserve" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	itemId := parts[0]

	var body reserveBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	if body.Quantity <= 0 {
		body.Quantity = 1
	}
	name := strings.TrimSpace(body.ReserverName)
	if !body.IsAnonymous && name == "" {
		http.Error(w, "name required", http.StatusBadRequest)
		return
	}

	super := permissions.NewSuperActor()

	item, _, err := h.client.RegistryItem().SelectById(r.Context(), super, registry_item.SelectByIdQuery{Id: itemId}, registryitemapi.NewProjection(true))
	if err != nil {
		http.Error(w, "item not found", http.StatusNotFound)
		return
	}

	_, _, err = h.client.Reservation().Create(r.Context(), super, reservation.Model{
		ItemId:       item.Id,
		RegistryId:   item.RegistryId,
		ReserverName: name,
		IsAnonymous:  body.IsAnonymous,
		Message:      strings.TrimSpace(body.Message),
		ContactEmail: strings.TrimSpace(body.ContactEmail),
		Quantity:     body.Quantity,
		Status:       enum_reservation_status.Reserved,
	}, reservation.NewProjection(true))
	if err != nil {
		http.Error(w, "could not reserve: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

// Unused import suppression for items that may not be referenced in some builds.
var _ = enum_item_source.Other
