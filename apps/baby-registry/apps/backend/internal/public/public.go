package public

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	addressaccesssessionapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_access_mode"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_access_level"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_item_source"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	owneruserapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	registryapprovedguestapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	registryitemapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	reservationapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/mailer"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/shipping"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	mux          *http.ServeMux
	client       api.Client
	mailer       mailer.Mailer
	resolveBuyer func(r *http.Request, slug string) (string, error)
}

func NewHandler(client api.Client, notificationMailer mailer.Mailer) *Handler {
	h := &Handler{client: client, mailer: notificationMailer}
	mux := http.NewServeMux()
	mux.HandleFunc("/r/", h.handleRegistryBySlug)
	mux.HandleFunc("/items/", h.handleItemRoute) // /items/:id/reserve
	mux.HandleFunc("/shipping/resolve", h.handleShippingResolve)
	h.mux = mux
	return h
}

// Mux exposes the internal mux so other modules (e.g. buyer auth) can register
// additional public routes under the same /api/public/ prefix.
func (h *Handler) Mux() *http.ServeMux { return h.mux }

// SetBuyerResolver wires the buyer-auth resolver so the public handler can
// gate access by verified-email cookie.
func (h *Handler) SetBuyerResolver(f func(r *http.Request, slug string) (string, error)) {
	h.resolveBuyer = f
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

type publicItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	ProductUrl  string `json:"productUrl"`
	Source      string `json:"source"`
	PriceCents  int    `json:"priceCents"`
	Currency    string `json:"currency"`
	Quantity    int    `json:"quantity"`
	Notes       string `json:"notes"`
	Position    int    `json:"position"`
	Reserved    int    `json:"reserved"`
}

type publicRegistry struct {
	Id                    string       `json:"id"`
	Slug                  string       `json:"slug"`
	Title                 string       `json:"title"`
	ParentNames           string       `json:"parentNames"`
	WelcomeMessage        string       `json:"welcomeMessage"`
	ThemeColor            string       `json:"themeColor"`
	CoverImageUrl         string       `json:"coverImageUrl"`
	ShippingRecipientName string       `json:"shippingRecipientName,omitempty"`
	ShippingLine1         string       `json:"shippingLine1,omitempty"`
	ShippingLine2         string       `json:"shippingLine2,omitempty"`
	ShippingCity          string       `json:"shippingCity,omitempty"`
	ShippingRegion        string       `json:"shippingRegion,omitempty"`
	ShippingPostalCode    string       `json:"shippingPostalCode,omitempty"`
	ShippingCountry       string       `json:"shippingCountry,omitempty"`
	ShippingDeliveryNotes string       `json:"shippingDeliveryNotes,omitempty"`
	Items                 []publicItem `json:"items"`
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

	// Gate: buyer must have verified their email for this registry.
	canViewShippingAddress := false
	if h.resolveBuyer != nil {
		buyerEmail, err := h.resolveBuyer(r, slug)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"error":                "verification required",
				"verificationRequired": true,
				"title":                reg.Title,
				"parentNames":          reg.ParentNames,
				"themeColor":           reg.ThemeColor,
			})
			return
		}

		guest, gErr := h.resolveApprovedGuest(r.Context(), super, reg.Id, buyerEmail)
		if gErr != nil {
			http.Error(w, "lookup error", http.StatusInternalServerError)
			return
		}
		if guest == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"error":                 "approved guest access required",
				"approvedGuestRequired": true,
				"title":                 reg.Title,
				"parentNames":           reg.ParentNames,
				"themeColor":            reg.ThemeColor,
			})
			return
		}
		canViewShippingAddress = guest.AccessLevel == enum_guest_access_level.ViewShippingAddress
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

	resp := publicRegistry{
		Id:             reg.Id,
		Slug:           reg.Slug,
		Title:          reg.Title,
		ParentNames:    reg.ParentNames,
		WelcomeMessage: reg.WelcomeMessage,
		ThemeColor:     reg.ThemeColor,
		CoverImageUrl:  reg.CoverImageUrl,
		Items:          publicItems,
	}
	if canViewShippingAddress {
		resp.ShippingRecipientName = reg.ShippingRecipientName
		resp.ShippingLine1 = reg.ShippingLine1
		resp.ShippingLine2 = reg.ShippingLine2
		resp.ShippingCity = reg.ShippingCity
		resp.ShippingRegion = reg.ShippingRegion
		resp.ShippingPostalCode = reg.ShippingPostalCode
		resp.ShippingCountry = reg.ShippingCountry
		resp.ShippingDeliveryNotes = reg.ShippingDeliveryNotes
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
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

	reg, _, regErr := h.client.Registry().SelectById(r.Context(), super,
		registry.SelectByIdQuery{Id: item.RegistryId}, registryapi.NewProjection(true))
	if regErr != nil {
		http.Error(w, "registry not found", http.StatusNotFound)
		return
	}

	// Gate: reserver must be email-verified for this registry.
	var buyerEmail string
	if h.resolveBuyer != nil {
		email, err := h.resolveBuyer(r, reg.Slug)
		if err != nil {
			http.Error(w, "verification required", http.StatusUnauthorized)
			return
		}
		buyerEmail = email

		guest, gErr := h.resolveApprovedGuest(r.Context(), super, item.RegistryId, buyerEmail)
		if gErr != nil {
			http.Error(w, "lookup error", http.StatusInternalServerError)
			return
		}
		if guest == nil {
			http.Error(w, "approved guest access required", http.StatusForbidden)
			return
		}
	}
	if buyerEmail == "" {
		buyerEmail = strings.TrimSpace(body.ContactEmail)
	}

	_, _, err = h.client.Reservation().Create(r.Context(), super, reservation.Model{
		ItemId:       item.Id,
		RegistryId:   item.RegistryId,
		ReserverName: name,
		IsAnonymous:  body.IsAnonymous,
		Message:      strings.TrimSpace(body.Message),
		ContactEmail: buyerEmail,
		Quantity:     body.Quantity,
		Status:       enum_reservation_status.Reserved,
	}, reservation.NewProjection(true))
	if err != nil {
		http.Error(w, "could not reserve: "+err.Error(), http.StatusInternalServerError)
		return
	}

	h.sendOwnerReservationNotification(reg.OwnerId, reg.Title, item.Title, body.Quantity, name, body.IsAnonymous, buyerEmail, strings.TrimSpace(body.Message))

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func (h *Handler) sendOwnerReservationNotification(ownerID, registryTitle, itemTitle string, quantity int, reserverName string, isAnonymous bool, buyerEmail, message string) {
	if h.mailer == nil || ownerID == "" {
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		super := permissions.NewSuperActor()
		owner, _, err := h.client.OwnerUser().SelectById(ctx, super, owner_user.SelectByIdQuery{Id: ownerID}, owneruserapi.NewProjection(true))
		if err != nil {
			log.Error().Err(err).Str("ownerId", ownerID).Msg("owner reservation notification lookup failed")
			return
		}
		if strings.TrimSpace(owner.Email) == "" {
			return
		}
		ownerName := fallbackString(strings.TrimSpace(owner.Name), "there")

		buyerName := reserverName
		if isAnonymous {
			buyerName = "Anonymous"
		} else if buyerName == "" {
			buyerName = "Unknown buyer"
		}
		qtyLine := ""
		if quantity > 1 {
			qtyLine = fmt.Sprintf("Quantity: %d\n", quantity)
		}
		messageLine := ""
		if message != "" {
			messageLine = "Message: " + message + "\n"
		}

		err = h.mailer.Send(ctx, mailer.Message{
			To:      owner.Email,
			Subject: "Someone claimed an item on your Stork Nest registry",
			Text: fmt.Sprintf(
				"Hi %s,\n\n%s marked \"%s\" as claimed on your \"%s\" registry.\n%sBuyer email: %s\n%s\nYou can review your registry in Stork Nest.\n",
				ownerName,
				buyerName,
				itemTitle,
				registryTitle,
				qtyLine,
				fallbackString(strings.TrimSpace(buyerEmail), "not available"),
				messageLine,
			),
		})
		if err != nil {
			log.Error().Err(err).Str("ownerId", ownerID).Str("email", owner.Email).Msg("owner reservation notification send failed")
		}
	}()
}

func fallbackString(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

func (h *Handler) resolveApprovedGuest(ctx context.Context, super permissions.Actor, registryId, email string) (*registryapprovedguestapi.Model, error) {
	hash := shipping.HashEmail(email)
	res, _, err := h.client.RegistryApprovedGuest().Search(ctx, super,
		registry_approved_guest.WhereClause{RegistryIdEq: &registryId, EmailHashEq: &hash},
		registryapprovedguestapi.QueryOptions{Limit: 1},
	)
	if err != nil {
		return nil, err
	}
	if len(res.Data) == 0 {
		return nil, nil
	}
	guest := res.Data[0]
	if guest.Status != enum_guest_status.Active {
		return nil, nil
	}
	return &guest, nil
}

// Unused import suppression for items that may not be referenced in some builds.
var _ = enum_item_source.Other

type resolveBody struct {
	Token string `json:"token"`
}

func (h *Handler) handleShippingResolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body resolveBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || strings.TrimSpace(body.Token) == "" {
		writeResolveError(w, http.StatusBadRequest, "invalid token")
		return
	}

	super := permissions.NewSuperActor()
	hash := shipping.HashToken(strings.TrimSpace(body.Token))

	res, _, err := h.client.AddressAccessSession().Search(r.Context(), super,
		address_access_session.WhereClause{TokenHashEq: &hash},
		addressaccesssessionapi.QueryOptions{Limit: 1},
	)
	if err != nil || len(res.Data) == 0 {
		writeResolveError(w, http.StatusNotFound, "link is invalid or expired")
		return
	}
	sess := res.Data[0]
	if sess.ExpiresAt.Before(time.Now()) {
		writeResolveError(w, http.StatusGone, "link has expired")
		return
	}

	reg, _, err := h.client.Registry().SelectById(r.Context(), super,
		registry.SelectByIdQuery{Id: sess.RegistryId}, registryapi.NewProjection(true))
	if err != nil {
		writeResolveError(w, http.StatusNotFound, "registry not found")
		return
	}
	if reg.AddressAccessMode == enum_address_access_mode.Disabled {
		writeResolveError(w, http.StatusForbidden, "the owner has disabled address sharing")
		return
	}
	if sess.PolicyVersionAtIssue != reg.ShippingPolicyVersion {
		writeResolveError(w, http.StatusForbidden, "the owner's privacy settings have changed; ask for a new link")
		return
	}

	// If the session traces back to an approved guest row, that row must be Active.
	if sess.ApprovedGuestId != "" {
		guest, _, err := h.client.RegistryApprovedGuest().SelectById(r.Context(), super,
			registry_approved_guest.SelectByIdQuery{Id: sess.ApprovedGuestId},
			registryapprovedguestapi.NewProjection(true),
		)
		if err != nil || guest.Status != enum_guest_status.Active {
			writeResolveError(w, http.StatusForbidden, "access has been revoked")
			return
		}
	} else {
		// Per-request session: still honor block list against the email hash.
		guests, _, err := h.client.RegistryApprovedGuest().Search(r.Context(), super,
			registry_approved_guest.WhereClause{RegistryIdEq: &sess.RegistryId, EmailHashEq: &sess.EmailHash},
			registryapprovedguestapi.QueryOptions{Limit: 1},
		)
		if err == nil && len(guests.Data) > 0 && guests.Data[0].Status != enum_guest_status.Active {
			writeResolveError(w, http.StatusForbidden, "access has been revoked")
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"registryTitle": reg.Title,
		"recipientName": reg.ShippingRecipientName,
		"line1":         reg.ShippingLine1,
		"line2":         reg.ShippingLine2,
		"city":          reg.ShippingCity,
		"region":        reg.ShippingRegion,
		"postalCode":    reg.ShippingPostalCode,
		"country":       reg.ShippingCountry,
		"deliveryNotes": reg.ShippingDeliveryNotes,
		"expiresAt":     sess.ExpiresAt,
	})
}

func writeResolveError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"error": msg})
}
