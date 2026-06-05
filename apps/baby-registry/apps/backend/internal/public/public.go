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
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	cartapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_access_mode"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_request_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_cart_status"
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
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
	shippingaddressrequestapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/exchangerates"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/mailer"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/shipping"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	mux          *http.ServeMux
	client       api.Client
	db           *mongo.Database
	mailer       mailer.Mailer
	appBaseURL   string
	resolveBuyer func(r *http.Request, slug string) (string, error)
	rates        *exchangerates.Store
}

func NewHandler(client api.Client, db *mongo.Database, notificationMailer mailer.Mailer, appBaseURL string) *Handler {
	h := &Handler{client: client, db: db, mailer: notificationMailer, appBaseURL: strings.TrimRight(appBaseURL, "/"), rates: exchangerates.NewStore(db)}
	mux := http.NewServeMux()
	mux.HandleFunc("/r/", h.handleRegistryBySlug)
	mux.HandleFunc("/items/", h.handleItemRoute)               // /items/:id/reserve and /items/:id/click
	mux.HandleFunc("/reservations/", h.handleReservationRoute) // /reservations/:id/confirm and /reservations/:id/cancel
	mux.HandleFunc("/shipping/resolve", h.handleShippingResolve)
	mux.HandleFunc("/address-requests", h.handleAddressRequestCreate)
	mux.HandleFunc("/registry-access/request", h.handleRegistryAccessRequest)
	mux.HandleFunc("/exchange-rates", h.handleExchangeRates)        // GET cached FX rates for currency conversion
	mux.HandleFunc("/payments/methods", h.handlePaymentMethods)     // GET ?slug= -> enabled methods
	mux.HandleFunc("/payments/intent", h.handleCreatePaymentIntent) // POST create contribution payment
	mux.HandleFunc("/payments/", h.handlePaymentRoute)              // /payments/:id/claim
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

// handleExchangeRates returns the cached foreign-exchange rates so the public
// page can convert item prices into a viewer-chosen currency. The response is
// served entirely from the cached Mongo document; the upstream provider is only
// contacted by the background poller. GET /exchange-rates
func (h *Handler) handleExchangeRates(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	snap, ok := h.rates.Get(r.Context())
	if !ok {
		http.Error(w, "rates unavailable", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"base":      snap.Base,
		"rates":     snap.Rates,
		"date":      snap.Date,
		"fetchedAt": snap.FetchedAt,
	})
}

func (h *Handler) brand() mailer.Brand {
	return mailer.Brand{AppName: "Stork Nest", AppBaseURL: h.appBaseURL}
}

func reservationDetailsHTML(quantity int, buyerEmail, message string) string {
	var b strings.Builder
	b.WriteString(`<div style="margin:18px 0;padding:16px 18px;background:#fbf7f2;border:1px solid #ecdfce;border-radius:14px;">`)
	if quantity > 1 {
		fmt.Fprintf(&b, `<div style="font-size:13px;color:#6b665f;margin-bottom:4px;">Quantity</div><div style="font-size:15px;font-weight:600;margin-bottom:10px;">%d</div>`, quantity)
	}
	if buyerEmail != "" {
		fmt.Fprintf(&b, `<div style="font-size:13px;color:#6b665f;margin-bottom:4px;">Buyer email</div><div style="font-size:15px;margin-bottom:10px;">%s</div>`, mailer.Esc(buyerEmail))
	}
	if message != "" {
		fmt.Fprintf(&b, `<div style="font-size:13px;color:#6b665f;margin-bottom:4px;">Message</div><div style="font-size:15px;white-space:pre-wrap;">%s</div>`, mailer.Esc(message))
	}
	b.WriteString(`</div>`)
	return b.String()
}

func requestNoteHTML(note string) string {
	note = strings.TrimSpace(note)
	if note == "" {
		return ""
	}
	return `<div style="margin:18px 0;padding:16px 18px;background:#fbf7f2;border:1px solid #ecdfce;border-radius:14px;"><div style="font-size:13px;color:#6b665f;margin-bottom:4px;">Their message</div><div style="font-size:15px;white-space:pre-wrap;">` +
		mailer.Esc(note) + `</div></div>`
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

type publicItem struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	ImageUrl          string `json:"imageUrl"`
	ImageBgColor      string `json:"imageBgColor"`
	ProductUrl        string `json:"productUrl"`
	AffiliateUrl      string `json:"affiliateUrl"`
	Retailer          string `json:"retailer"`
	Source            string `json:"source"`
	PriceCents        int    `json:"priceCents"`
	Currency          string `json:"currency"`
	Quantity          int    `json:"quantity"`
	QuantityUnlimited bool   `json:"quantityUnlimited"`
	Category          string `json:"category"`
	NoSubstitutes     bool   `json:"noSubstitutes"`
	ParentItemId      string `json:"parentItemId,omitempty"`
	Notes             string `json:"notes"`
	Position          int    `json:"position"`
	Reserved          int    `json:"reserved"`
}

type publicRegistry struct {
	Id                    string          `json:"id"`
	Slug                  string          `json:"slug"`
	Title                 string          `json:"title"`
	ParentNames           string          `json:"parentNames"`
	WelcomeMessage        string          `json:"welcomeMessage"`
	ThemeColor            string          `json:"themeColor"`
	CoverImageUrl         string          `json:"coverImageUrl"`
	ShippingRecipientName string          `json:"shippingRecipientName,omitempty"`
	ShippingLine1         string          `json:"shippingLine1,omitempty"`
	ShippingLine2         string          `json:"shippingLine2,omitempty"`
	ShippingCity          string          `json:"shippingCity,omitempty"`
	ShippingRegion        string          `json:"shippingRegion,omitempty"`
	ShippingPostalCode    string          `json:"shippingPostalCode,omitempty"`
	ShippingCountry       string          `json:"shippingCountry,omitempty"`
	ShippingDeliveryNotes string          `json:"shippingDeliveryNotes,omitempty"`
	Items                 []publicItem    `json:"items"`
	MyReservations        []myReservation `json:"myReservations,omitempty"`
	MyCarts               []myCart        `json:"myCarts,omitempty"`
}

// myReservation surfaces the buyer's own active reservations so the UI can
// prompt them to confirm or cancel after returning from the retailer.
type myReservation struct {
	Id        string `json:"id"`
	ItemId    string `json:"itemId"`
	ItemTitle string `json:"itemTitle"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"createdAt"`
	ExpiresAt string `json:"expiresAt"`
}

// myCart surfaces a cart the buyer has started (and possibly claimed as sent)
// so the UI can show its status while parents confirm receipt.
type myCart struct {
	Id                string             `json:"id"`
	ReferenceCode     string             `json:"referenceCode"`
	Status            string             `json:"status"`
	AmountCents       int                `json:"amountCents"`
	Currency          string             `json:"currency"`
	MethodDisplayName string             `json:"methodDisplayName"`
	CreatedAt         string             `json:"createdAt"`
	Items             []cartItemSnapshot `json:"items"`
}

// gatedRegistry is the minimal payload returned when the viewer has not been
// granted access to a registry that requires owner approval. It deliberately
// excludes items, shipping details, and any other sensitive content.
type gatedRegistry struct {
	AccessGated         bool   `json:"accessGated"`
	Slug                string `json:"slug"`
	Title               string `json:"title"`
	ParentNames         string `json:"parentNames,omitempty"`
	ThemeColor          string `json:"themeColor,omitempty"`
	CoverImageUrl       string `json:"coverImageUrl,omitempty"`
	WelcomeMessage      string `json:"welcomeMessage,omitempty"`
	OwnerDisplayName    string `json:"ownerDisplayName,omitempty"`
	AccessRequestStatus string `json:"accessRequestStatus"` // none | pending | rejected
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

	// Gate 1: buyer must have verified their email for this registry.
	canViewShippingAddress := false
	var buyerEmail string
	if h.resolveBuyer != nil {
		email, err := h.resolveBuyer(r, slug)
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

		buyerEmail = email
		guest, gErr := h.fetchApprovedGuest(r.Context(), super, reg.Id, buyerEmail)
		if gErr != nil {
			http.Error(w, "lookup error", http.StatusInternalServerError)
			return
		}

		// Gate 2: if the registry requires owner approval, the buyer must be
		// an Active approved guest to see any contents. Otherwise return a
		// minimal payload that drives the request-access UI.
		if !reg.AllowOpenAccess && (guest == nil || guest.Status != enum_guest_status.Active) {
			status := "none"
			if guest != nil {
				switch guest.Status {
				case enum_guest_status.Pending:
					status = "pending"
				case enum_guest_status.Blocked, enum_guest_status.Revoked:
					status = "rejected"
				}
			}
			ownerName := ""
			if owner, _, oErr := h.client.OwnerUser().SelectById(r.Context(), super,
				owner_user.SelectByIdQuery{Id: reg.OwnerId}, owneruserapi.NewProjection(true)); oErr == nil {
				ownerName = strings.TrimSpace(owner.Name)
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(gatedRegistry{
				AccessGated:         true,
				Slug:                reg.Slug,
				Title:               reg.Title,
				ParentNames:         reg.ParentNames,
				ThemeColor:          reg.ThemeColor,
				CoverImageUrl:       reg.CoverImageUrl,
				WelcomeMessage:      reg.WelcomeMessage,
				OwnerDisplayName:    ownerName,
				AccessRequestStatus: status,
			})
			return
		}

		if guest != nil && guest.Status == enum_guest_status.Active {
			canViewShippingAddress = true
		}
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
	itemByID := map[string]registryitemapi.Model{}
	for _, it := range itemsResult.Data {
		itemByID[it.Id] = it
	}

	reservedByGroup := map[string]int{}
	now := time.Now().UTC()
	for _, rsv := range resvResult.Data {
		if rsv.Status == enum_reservation_status.Cancelled {
			continue
		}
		if !rsv.ExpiresAt.IsZero() && rsv.ExpiresAt.Before(now) && rsv.Status == enum_reservation_status.Reserved {
			continue
		}
		it, ok := itemByID[rsv.ItemId]
		if !ok {
			continue
		}
		q := rsv.Quantity
		if q <= 0 {
			q = 1
		}
		groupID := groupRootID(it, itemByID)
		reservedByGroup[groupID] += q
	}

	publicItems := make([]publicItem, 0, len(itemsResult.Data))
	for _, it := range itemsResult.Data {
		groupID := groupRootID(it, itemByID)
		publicItems = append(publicItems, publicItem{
			Id:                it.Id,
			Title:             it.Title,
			Description:       it.Description,
			ImageUrl:          it.ImageUrl,
			ImageBgColor:      it.ImageBgColor,
			ProductUrl:        it.ProductUrl,
			AffiliateUrl:      firstNonEmpty(it.AffiliateUrl, it.ProductUrl),
			Retailer:          it.Retailer,
			Source:            string(it.Source),
			PriceCents:        it.PriceCents,
			Currency:          it.Currency,
			Quantity:          it.Quantity,
			QuantityUnlimited: it.QuantityUnlimited,
			Category:          it.Category,
			NoSubstitutes:     it.NoSubstitutes,
			ParentItemId:      it.ParentItemId,
			Notes:             it.Notes,
			Position:          it.Position,
			Reserved:          reservedByGroup[groupID],
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

	if buyerEmail != "" {
		lowerEmail := strings.ToLower(buyerEmail)

		// Group the buyer's reservations by the cart they belong to so each
		// cart can show its line items without an extra query per cart.
		cartItems := map[string][]cartItemSnapshot{}
		for _, rsv := range resvResult.Data {
			cid := strings.TrimSpace(rsv.CartId)
			// Reserved/Cancelled reservations can carry a stale CartId (empty
			// Refs cannot be unset on update), so only reservations actually
			// committed to a cart should surface as that cart's line items.
			if cid == "" ||
				rsv.Status == enum_reservation_status.Reserved ||
				rsv.Status == enum_reservation_status.Cancelled {
				continue
			}
			qty := rsv.Quantity
			if qty <= 0 {
				qty = 1
			}
			it := itemByID[rsv.ItemId]
			cartItems[cid] = append(cartItems[cid], cartItemSnapshot{
				ReservationId: rsv.Id,
				ItemId:        rsv.ItemId,
				Title:         it.Title,
				Quantity:      qty,
				PriceCents:    it.PriceCents,
				Currency:      it.Currency,
			})
		}

		// Surface the buyer's active carts (anything not rejected) so the UI
		// can show their status while parents confirm receipt. Reservations in
		// these carts are already AwaitingConfirmation, so they are naturally
		// excluded from the open-cart list below.
		cartResult, _, cartErr := h.client.Cart().Search(r.Context(), super, cart.WhereClause{
			RegistryIdEq:       &reg.Id,
			ContributorEmailEq: &lowerEmail,
		}, cartapi.QueryOptions{Limit: 200})
		if cartErr == nil {
			for _, c := range cartResult.Data {
				if c.Status == enum_cart_status.Rejected {
					continue
				}
				resp.MyCarts = append(resp.MyCarts, myCart{
					Id:                c.Id,
					ReferenceCode:     c.ReferenceCode,
					Status:            string(c.Status),
					AmountCents:       c.AmountCents,
					Currency:          c.Currency,
					MethodDisplayName: c.MethodDisplayName,
					CreatedAt:         c.Created.At.UTC().Format(time.RFC3339),
					Items:             cartItems[c.Id],
				})
			}
		}

		for _, rsv := range resvResult.Data {
			if rsv.Status != enum_reservation_status.Reserved {
				continue
			}
			if strings.ToLower(strings.TrimSpace(rsv.ContactEmail)) != lowerEmail {
				continue
			}
			if !rsv.ExpiresAt.IsZero() && rsv.ExpiresAt.Before(now) {
				continue
			}
			// A Reserved reservation is by definition open; any CartId it still
			// carries is stale and must not hide it from the buyer's cart.
			title := ""
			if it, ok := itemByID[rsv.ItemId]; ok {
				title = it.Title
			}
			resp.MyReservations = append(resp.MyReservations, myReservation{
				Id:        rsv.Id,
				ItemId:    rsv.ItemId,
				ItemTitle: title,
				Quantity:  rsv.Quantity,
				CreatedAt: rsv.Created.At.UTC().Format(time.RFC3339),
				ExpiresAt: rsv.ExpiresAt.UTC().Format(time.RFC3339),
			})
		}
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
	if len(parts) != 2 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	switch parts[1] {
	case "reserve":
		h.handleItemReserve(w, r, parts[0])
	case "click":
		h.handleItemClick(w, r, parts[0])
	default:
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func (h *Handler) handleItemClick(w http.ResponseWriter, r *http.Request, itemId string) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	super := permissions.NewSuperActor()
	item, _, err := h.client.RegistryItem().SelectById(r.Context(), super, registry_item.SelectByIdQuery{Id: itemId}, registryitemapi.NewProjection(true))
	if err != nil {
		// Don't leak whether the item exists; the click is fire-and-forget.
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Gate access for registries that require owner approval.
	if h.resolveBuyer != nil {
		reg, _, regErr := h.client.Registry().SelectById(r.Context(), super,
			registry.SelectByIdQuery{Id: item.RegistryId}, registryapi.NewProjection(true))
		if regErr == nil && !reg.AllowOpenAccess {
			email, bErr := h.resolveBuyer(r, reg.Slug)
			if bErr != nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			guest, gErr := h.fetchApprovedGuest(r.Context(), super, reg.Id, email)
			if gErr != nil || guest == nil || guest.Status != enum_guest_status.Active {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
	}
	event := bson.M{
		"event":      "registry_item_purchase_click",
		"registryId": item.RegistryId,
		"itemId":     item.Id,
		"retailer":   item.Retailer,
		"clickedAt":  time.Now().UTC(),
	}
	if h.db != nil {
		if _, err := h.db.Collection("purchase_clicks").InsertOne(r.Context(), event); err != nil {
			log.Warn().Err(err).Msg("purchase click insert failed")
		}
	}
	log.Info().
		Str("event", "registry_item_purchase_click").
		Str("registryId", item.RegistryId).
		Str("itemId", item.Id).
		Str("retailer", item.Retailer).
		Msg("purchase click")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) handleItemReserve(w http.ResponseWriter, r *http.Request, itemId string) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
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

	itemsResult, _, err := h.client.RegistryItem().Search(r.Context(), super, registry_item.WhereClause{
		RegistryIdEq: &item.RegistryId,
	}, registryitemapi.QueryOptions{Limit: 500})
	if err != nil {
		http.Error(w, "lookup error", http.StatusInternalServerError)
		return
	}

	itemByID := map[string]registryitemapi.Model{}
	for _, it := range itemsResult.Data {
		itemByID[it.Id] = it
	}
	groupID := groupRootID(item, itemByID)
	rootItem, ok := itemByID[groupID]
	if !ok {
		rootItem = item
	}

	// Gate: reserver must be email-verified for this registry. When the
	// registry requires owner approval, the reserver must also be an Active
	// approved guest.
	var buyerEmail string
	if h.resolveBuyer != nil {
		email, err := h.resolveBuyer(r, reg.Slug)
		if err != nil {
			http.Error(w, "verification required", http.StatusUnauthorized)
			return
		}
		buyerEmail = email
		if !reg.AllowOpenAccess {
			guest, gErr := h.fetchApprovedGuest(r.Context(), super, reg.Id, email)
			if gErr != nil {
				http.Error(w, "lookup error", http.StatusInternalServerError)
				return
			}
			if guest == nil || guest.Status != enum_guest_status.Active {
				http.Error(w, "access not approved", http.StatusForbidden)
				return
			}
		}
	}
	if buyerEmail == "" {
		buyerEmail = strings.TrimSpace(body.ContactEmail)
	}

	if !rootItem.QuantityUnlimited {
		resvResult, _, err := h.client.Reservation().Search(r.Context(), super, reservation.WhereClause{
			RegistryIdEq: &item.RegistryId,
		}, reservationapi.QueryOptions{Limit: 1000})
		if err != nil {
			http.Error(w, "lookup error", http.StatusInternalServerError)
			return
		}

		reservedByGroup := map[string]int{}
		now := time.Now().UTC()
		for _, rsv := range resvResult.Data {
			if rsv.Status == enum_reservation_status.Cancelled {
				continue
			}
			if !rsv.ExpiresAt.IsZero() && rsv.ExpiresAt.Before(now) && rsv.Status == enum_reservation_status.Reserved {
				continue
			}
			it, ok := itemByID[rsv.ItemId]
			if !ok {
				continue
			}
			q := rsv.Quantity
			if q <= 0 {
				q = 1
			}
			reservedByGroup[groupRootID(it, itemByID)] += q
		}

		remaining := rootItem.Quantity - reservedByGroup[groupID]
		if remaining <= 0 {
			http.Error(w, "item is fully reserved", http.StatusConflict)
			return
		}
		if body.Quantity > remaining {
			http.Error(w, fmt.Sprintf("only %d remaining", remaining), http.StatusConflict)
			return
		}
	}

	created, _, err := h.client.Reservation().Create(r.Context(), super, reservation.Model{
		ItemId:       item.Id,
		RegistryId:   item.RegistryId,
		ReserverName: name,
		IsAnonymous:  body.IsAnonymous,
		Message:      strings.TrimSpace(body.Message),
		ContactEmail: buyerEmail,
		Quantity:     body.Quantity,
		Status:       enum_reservation_status.Reserved,
		ExpiresAt:    time.Now().UTC().Add(24 * time.Hour),
	}, reservation.NewProjection(true))
	if err != nil {
		http.Error(w, "could not reserve: "+err.Error(), http.StatusInternalServerError)
		return
	}

	h.sendOwnerReservationNotification(reg.OwnerId, reg.Slug, reg.Title, item.Title, body.Quantity, name, body.IsAnonymous, buyerEmail, strings.TrimSpace(body.Message))

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "id": created.Id})
}

// handleReservationRoute handles POST /reservations/:id/confirm and
// /reservations/:id/cancel — used by Phase 3 return-visit reminders so a
// buyer can resolve their own pending reservation.
func (h *Handler) handleReservationRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	rest := strings.TrimPrefix(r.URL.Path, "/reservations/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	id := parts[0]
	var next enum_reservation_status.Value
	switch parts[1] {
	case "confirm":
		next = enum_reservation_status.Purchased
	case "cancel":
		next = enum_reservation_status.Cancelled
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	super := permissions.NewSuperActor()
	rsv, _, err := h.client.Reservation().SelectById(r.Context(), super,
		reservation.SelectByIdQuery{Id: id}, reservationapi.NewProjection(true))
	if err != nil {
		http.Error(w, "reservation not found", http.StatusNotFound)
		return
	}

	if h.resolveBuyer == nil {
		http.Error(w, "verification required", http.StatusUnauthorized)
		return
	}
	reg, _, regErr := h.client.Registry().SelectById(r.Context(), super,
		registry.SelectByIdQuery{Id: rsv.RegistryId}, registryapi.NewProjection(true))
	if regErr != nil {
		http.Error(w, "registry not found", http.StatusNotFound)
		return
	}
	buyerEmail, bErr := h.resolveBuyer(r, reg.Slug)
	if bErr != nil {
		http.Error(w, "verification required", http.StatusUnauthorized)
		return
	}
	if strings.ToLower(strings.TrimSpace(rsv.ContactEmail)) != strings.ToLower(strings.TrimSpace(buyerEmail)) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	if rsv.Status != enum_reservation_status.Reserved {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": string(rsv.Status)})
		return
	}

	inner := rsv.Model
	inner.Status = next
	if _, _, err := h.client.Reservation().Update(r.Context(), super, inner, reservation.NewProjection(true)); err != nil {
		http.Error(w, "update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": next})
}

func (h *Handler) sendOwnerReservationNotification(ownerID, registrySlug, registryTitle, itemTitle string, quantity int, reserverName string, isAnonymous bool, buyerEmail, message string) {
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
			HTML: h.brand().Render(mailer.Email{
				Preheader: buyerName + " claimed \"" + itemTitle + "\".",
				Heading:   "Someone claimed a gift",
				Intro:     "Hi " + ownerName + ",\n\n" + buyerName + " marked \"" + itemTitle + "\" as claimed on your \"" + registryTitle + "\" registry.",
				BodyHTML:  reservationDetailsHTML(quantity, fallbackString(strings.TrimSpace(buyerEmail), "not available"), message),
				CTAText:   "Open your registry",
				CTAHref:   h.appBaseURL + "/owner/r/" + registrySlug,
			}),
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

func groupRootID(item registryitemapi.Model, byID map[string]registryitemapi.Model) string {
	if item.ParentItemId == "" {
		return item.Id
	}
	parent, ok := byID[item.ParentItemId]
	if !ok {
		return item.Id
	}
	if parent.ParentItemId != "" {
		return groupRootID(parent, byID)
	}
	return parent.Id
}

func (h *Handler) resolveApprovedGuest(ctx context.Context, super permissions.Actor, registryId, email string) (*registryapprovedguestapi.Model, error) {
	guest, err := h.fetchApprovedGuest(ctx, super, registryId, email)
	if err != nil || guest == nil {
		return nil, err
	}
	if guest.Status != enum_guest_status.Active {
		return nil, nil
	}
	return guest, nil
}

// fetchApprovedGuest returns the approved-guest row for (registry, email) regardless
// of status, so callers can distinguish Pending / Active / Blocked.
func (h *Handler) fetchApprovedGuest(ctx context.Context, super permissions.Actor, registryId, email string) (*registryapprovedguestapi.Model, error) {
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
	return &res.Data[0], nil
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

type addressRequestBody struct {
	Slug   string `json:"slug"`
	ItemId string `json:"itemId"`
	Name   string `json:"name"`
	Note   string `json:"note"`
}

// handleAddressRequestCreate lets a verified buyer ask the registry owner
// for the shipping address. Creates a Pending ShippingAddressRequest and
// emails the owner. Idempotent: if a Pending request already exists for the
// same buyer + registry, it is returned without creating a duplicate.
func (h *Handler) handleAddressRequestCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body addressRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	slug := strings.TrimSpace(body.Slug)
	if slug == "" {
		http.Error(w, "slug required", http.StatusBadRequest)
		return
	}

	if h.resolveBuyer == nil {
		http.Error(w, "verification required", http.StatusUnauthorized)
		return
	}
	buyerEmail, err := h.resolveBuyer(r, slug)
	if err != nil || strings.TrimSpace(buyerEmail) == "" {
		http.Error(w, "verification required", http.StatusUnauthorized)
		return
	}

	super := permissions.NewSuperActor()

	regResult, _, err := h.client.Registry().Search(r.Context(), super,
		registry.WhereClause{SlugEq: &slug},
		registryapi.QueryOptions{Limit: 1},
	)
	if err != nil || len(regResult.Data) == 0 {
		http.Error(w, "registry not found", http.StatusNotFound)
		return
	}
	reg := regResult.Data[0]

	if reg.AddressAccessMode == enum_address_access_mode.Disabled {
		http.Error(w, "the owner has disabled address sharing", http.StatusForbidden)
		return
	}

	// When the registry requires owner approval, the buyer must be an Active
	// approved guest to even ask for the address.
	if !reg.AllowOpenAccess {
		guest, gErr := h.fetchApprovedGuest(r.Context(), super, reg.Id, buyerEmail)
		if gErr != nil {
			http.Error(w, "lookup error", http.StatusInternalServerError)
			return
		}
		if guest == nil || guest.Status != enum_guest_status.Active {
			http.Error(w, "access not approved", http.StatusForbidden)
			return
		}
	}

	emailHash := shipping.HashEmail(buyerEmail)
	pending := enum_address_request_status.Pending

	// Dedupe: return the existing Pending request if one already exists.
	existing, _, err := h.client.ShippingAddressRequest().Search(r.Context(), super,
		shipping_address_request.WhereClause{
			RegistryIdEq: &reg.Id,
			EmailHashEq:  &emailHash,
			StatusEq:     &pending,
		},
		shippingaddressrequestapi.QueryOptions{Limit: 1},
	)
	if err == nil && len(existing.Data) > 0 {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"ok":     true,
			"status": "pending",
			"id":     existing.Data[0].Id,
		})
		return
	}

	created, _, err := h.client.ShippingAddressRequest().Create(r.Context(), super,
		shipping_address_request.Model{
			OwnerId:        reg.OwnerId,
			RegistryId:     reg.Id,
			RegistryItemId: strings.TrimSpace(body.ItemId),
			EmailHash:      emailHash,
			EmailEnc:       shipping.EncryptEmail(buyerEmail),
			Name:           strings.TrimSpace(body.Name),
			Note:           strings.TrimSpace(body.Note),
			Status:         pending,
			PolicyVersion:  reg.ShippingPolicyVersion,
		},
		shipping_address_request.NewProjection(true),
	)
	if err != nil {
		log.Error().Err(err).Str("slug", slug).Msg("create address request failed")
		http.Error(w, "could not create request", http.StatusInternalServerError)
		return
	}

	h.sendOwnerAddressRequestNotification(reg.OwnerId, reg.Title, reg.Slug, strings.TrimSpace(body.Name), buyerEmail, strings.TrimSpace(body.Note))

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":     true,
		"status": "pending",
		"id":     created.Id,
	})
}

func (h *Handler) sendOwnerAddressRequestNotification(ownerID, registryTitle, registrySlug, buyerName, buyerEmail, note string) {
	if h.mailer == nil || ownerID == "" {
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		super := permissions.NewSuperActor()
		owner, _, err := h.client.OwnerUser().SelectById(ctx, super,
			owner_user.SelectByIdQuery{Id: ownerID},
			owneruserapi.NewProjection(true),
		)
		if err != nil {
			log.Error().Err(err).Str("ownerId", ownerID).Msg("owner address-request notification lookup failed")
			return
		}
		if strings.TrimSpace(owner.Email) == "" {
			return
		}
		ownerName := fallbackString(strings.TrimSpace(owner.Name), "there")

		displayName := strings.TrimSpace(buyerName)
		if displayName == "" {
			displayName = "A guest"
		}
		noteLine := ""
		if note != "" {
			noteLine = "Their message: \"" + note + "\"\n\n"
		}

		dashboardLink := h.appBaseURL + "/owner/r/" + registrySlug + "?tab=access"

		err = h.mailer.Send(ctx, mailer.Message{
			To:      owner.Email,
			Subject: "Someone is asking for your shipping address on Stork Nest",
			Text: fmt.Sprintf(
				"Hi %s,\n\n%s (%s) requested your shipping address for your \"%s\" registry.\n\n%sReview the request and approve or decline here:\n%s\n\nIf you approve, we'll generate a private link for them to view the address. They won't see it until you do.\n",
				ownerName,
				displayName,
				fallbackString(strings.TrimSpace(buyerEmail), "email not available"),
				registryTitle,
				noteLine,
				dashboardLink,
			),
			HTML: h.brand().Render(mailer.Email{
				Preheader: displayName + " requested your shipping address.",
				Heading:   "Address request",
				Intro:     "Hi " + ownerName + ",\n\n" + displayName + " (" + fallbackString(strings.TrimSpace(buyerEmail), "email not available") + ") requested your shipping address for your \"" + registryTitle + "\" registry.",
				BodyHTML:  requestNoteHTML(note),
				CTAText:   "Review the request",
				CTAHref:   dashboardLink,
				Footnote:  "If you approve, we'll generate a private link for them to view the address. They won't see it until you do.",
			}),
		})
		if err != nil {
			log.Error().Err(err).Str("ownerId", ownerID).Str("email", owner.Email).Msg("owner address-request notification send failed")
		}
	}()
}

type registryAccessRequestBody struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Note string `json:"note"`
}

// handleRegistryAccessRequest lets a verified buyer ask the registry owner
// for access to view the registry. Creates (or reuses) a RegistryApprovedGuest
// row with status=Pending and emails the owner. Idempotent.
func (h *Handler) handleRegistryAccessRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body registryAccessRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	slug := strings.TrimSpace(body.Slug)
	if slug == "" {
		http.Error(w, "slug required", http.StatusBadRequest)
		return
	}

	if h.resolveBuyer == nil {
		http.Error(w, "verification required", http.StatusUnauthorized)
		return
	}
	buyerEmail, err := h.resolveBuyer(r, slug)
	if err != nil || strings.TrimSpace(buyerEmail) == "" {
		http.Error(w, "verification required", http.StatusUnauthorized)
		return
	}

	super := permissions.NewSuperActor()

	regResult, _, err := h.client.Registry().Search(r.Context(), super,
		registry.WhereClause{SlugEq: &slug},
		registryapi.QueryOptions{Limit: 1},
	)
	if err != nil || len(regResult.Data) == 0 {
		http.Error(w, "registry not found", http.StatusNotFound)
		return
	}
	reg := regResult.Data[0]

	hash := shipping.HashEmail(buyerEmail)
	name := strings.TrimSpace(body.Name)
	note := strings.TrimSpace(body.Note)

	existing, _, err := h.client.RegistryApprovedGuest().Search(r.Context(), super,
		registry_approved_guest.WhereClause{RegistryIdEq: &reg.Id, EmailHashEq: &hash},
		registryapprovedguestapi.QueryOptions{Limit: 1},
	)
	if err != nil {
		http.Error(w, "lookup error", http.StatusInternalServerError)
		return
	}

	if len(existing.Data) > 0 {
		guest := existing.Data[0]
		switch guest.Status {
		case enum_guest_status.Active:
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": "approved"})
			return
		case enum_guest_status.Blocked, enum_guest_status.Revoked:
			// Don't allow re-requesting if the owner already declined.
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": "rejected"})
			return
		case enum_guest_status.Pending:
			// Refresh name if provided, but keep status as Pending.
			if name != "" && guest.Name == "" {
				guest.Model.Name = name
				if _, _, uErr := h.client.RegistryApprovedGuest().Update(r.Context(), super,
					guest.Model, registry_approved_guest.NewProjection(true)); uErr != nil {
					log.Warn().Err(uErr).Msg("update pending guest name failed")
				}
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": "pending"})
			return
		}
	}

	_, _, err = h.client.RegistryApprovedGuest().Create(r.Context(), super, registry_approved_guest.Model{
		OwnerId:     reg.OwnerId,
		RegistryId:  reg.Id,
		EmailHash:   hash,
		EmailEnc:    shipping.EncryptEmail(buyerEmail),
		Name:        name,
		AccessLevel: enum_guest_access_level.ViewShippingAddress,
		Status:      enum_guest_status.Pending,
	}, registry_approved_guest.NewProjection(true))
	if err != nil {
		log.Error().Err(err).Str("slug", slug).Msg("create registry access request failed")
		http.Error(w, "could not create request", http.StatusInternalServerError)
		return
	}

	h.sendOwnerRegistryAccessNotification(reg.OwnerId, reg.Title, reg.Slug, name, buyerEmail, note)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": "pending"})
}

func (h *Handler) sendOwnerRegistryAccessNotification(ownerID, registryTitle, registrySlug, buyerName, buyerEmail, note string) {
	if h.mailer == nil || ownerID == "" {
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		super := permissions.NewSuperActor()
		owner, _, err := h.client.OwnerUser().SelectById(ctx, super,
			owner_user.SelectByIdQuery{Id: ownerID},
			owneruserapi.NewProjection(true),
		)
		if err != nil {
			log.Error().Err(err).Str("ownerId", ownerID).Msg("owner registry-access notification lookup failed")
			return
		}
		if strings.TrimSpace(owner.Email) == "" {
			return
		}
		ownerName := fallbackString(strings.TrimSpace(owner.Name), "there")

		displayName := strings.TrimSpace(buyerName)
		if displayName == "" {
			displayName = "A guest"
		}
		noteLine := ""
		if note != "" {
			noteLine = "Their message: \"" + note + "\"\n\n"
		}

		dashboardLink := h.appBaseURL + "/owner/r/" + registrySlug + "?tab=access"

		err = h.mailer.Send(ctx, mailer.Message{
			To:      owner.Email,
			Subject: "Someone is asking to view your Stork Nest registry",
			Text: fmt.Sprintf(
				"Hi %s,\n\n%s (%s) is asking for access to view your \"%s\" registry.\n\n%sReview the request and approve or decline here:\n%s\n\nThey will not see any of your registry contents until you approve them.\n",
				ownerName,
				displayName,
				fallbackString(strings.TrimSpace(buyerEmail), "email not available"),
				registryTitle,
				noteLine,
				dashboardLink,
			),
			HTML: h.brand().Render(mailer.Email{
				Preheader: displayName + " wants to view your registry.",
				Heading:   "Access request",
				Intro:     "Hi " + ownerName + ",\n\n" + displayName + " (" + fallbackString(strings.TrimSpace(buyerEmail), "email not available") + ") is asking for access to view your \"" + registryTitle + "\" registry.",
				BodyHTML:  requestNoteHTML(note),
				CTAText:   "Review the request",
				CTAHref:   dashboardLink,
				Footnote:  "They will not see any of your registry contents until you approve them.",
			}),
		})
		if err != nil {
			log.Error().Err(err).Str("ownerId", ownerID).Str("email", owner.Email).Msg("owner registry-access notification send failed")
		}
	}()
}
