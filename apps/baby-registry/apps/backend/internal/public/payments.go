package public

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	cartapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_cart_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	registryitemapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
	registrypaymentmethodapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	reservationapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_api"
)

// paymentMethodPublic is the approved-contributor-facing view of a configured
// payment method. It intentionally includes the recipient/banking details
// because the caller has already passed the approved-contributor gate.
type paymentMethodPublic struct {
	Id                string `json:"id"`
	Type              string `json:"type"`
	DisplayName       string `json:"displayName"`
	Instructions      string `json:"instructions"`
	PaymentUrl        string `json:"paymentUrl"`
	RecipientEmail    string `json:"recipientEmail"`
	RecipientPhone    string `json:"recipientPhone"`
	BankName          string `json:"bankName"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountNumber string `json:"bankAccountNumber"`
	BankRoutingNumber string `json:"bankRoutingNumber"`
	BankIban          string `json:"bankIban"`
	BankSwift         string `json:"bankSwift"`
}

// cartItemSnapshot is a line item in a cart, joined live from the linked
// reservation and its registry item.
type cartItemSnapshot struct {
	ReservationId string `json:"reservationId"`
	ItemId        string `json:"itemId"`
	Title         string `json:"title"`
	Quantity      int    `json:"quantity"`
	PriceCents    int    `json:"priceCents"`
	Currency      string `json:"currency"`
}

// resolveContributor verifies the buyer is email-verified for the registry and,
// when the registry requires approval, that they are an Active approved guest.
// Returns the registry and the lowercased buyer email.
func (h *Handler) resolveContributor(r *http.Request, slug string) (*registry.Model, string, bool, int) {
	if h.resolveBuyer == nil {
		return nil, "", false, http.StatusUnauthorized
	}
	email, err := h.resolveBuyer(r, slug)
	if err != nil || strings.TrimSpace(email) == "" {
		return nil, "", false, http.StatusUnauthorized
	}

	super := permissions.NewSuperActor()
	regResult, _, err := h.client.Registry().Search(r.Context(), super, registry.WhereClause{
		SlugEq: &slug,
	}, registryapi.QueryOptions{Limit: 1})
	if err != nil || len(regResult.Data) == 0 {
		return nil, "", false, http.StatusNotFound
	}
	reg := regResult.Data[0].Model
	if !reg.IsPublic {
		return nil, "", false, http.StatusNotFound
	}

	if !reg.AllowOpenAccess {
		guest, gErr := h.fetchApprovedGuest(r.Context(), super, reg.Id, email)
		if gErr != nil {
			return nil, "", false, http.StatusInternalServerError
		}
		if guest == nil || guest.Status != enum_guest_status.Active {
			return nil, "", false, http.StatusForbidden
		}
	}
	return &reg, strings.ToLower(strings.TrimSpace(email)), true, http.StatusOK
}

// handlePaymentMethods returns the registry's enabled payment methods to an
// approved contributor. GET /payments/methods?slug=<slug>
func (h *Handler) handlePaymentMethods(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	slug := strings.TrimSpace(r.URL.Query().Get("slug"))
	if slug == "" {
		writeResolveError(w, http.StatusBadRequest, "missing slug")
		return
	}
	reg, _, ok, code := h.resolveContributor(r, slug)
	if !ok {
		writeResolveError(w, code, "not allowed")
		return
	}

	super := permissions.NewSuperActor()
	res, _, err := h.client.RegistryPaymentMethod().Search(r.Context(), super,
		registry_payment_method.WhereClause{RegistryIdEq: &reg.Id},
		registrypaymentmethodapi.QueryOptions{Limit: 100},
	)
	if err != nil {
		writeResolveError(w, http.StatusInternalServerError, "lookup error")
		return
	}
	out := make([]paymentMethodPublic, 0, len(res.Data))
	for i := range res.Data {
		m := res.Data[i].Model
		if !m.Enabled {
			continue
		}
		out = append(out, paymentMethodPublic{
			Id:                m.Id,
			Type:              string(m.Type),
			DisplayName:       m.DisplayName,
			Instructions:      m.Instructions,
			PaymentUrl:        m.PaymentUrl,
			RecipientEmail:    m.RecipientEmail,
			RecipientPhone:    m.RecipientPhone,
			BankName:          m.BankName,
			BankAccountName:   m.BankAccountName,
			BankAccountNumber: m.BankAccountNumber,
			BankRoutingNumber: m.BankRoutingNumber,
			BankIban:          m.BankIban,
			BankSwift:         m.BankSwift,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"data": out})
}

type createPaymentIntentBody struct {
	Slug            string `json:"slug"`
	PaymentMethodId string `json:"paymentMethodId"`
}

// handleCreatePaymentIntent locks the buyer's active reservations into a new
// Cart (status Pending) for the chosen payment method. This happens *before*
// we show the buyer where to send the money, so concurrent checkouts cannot
// double-claim the same reservations. Any half-finished Pending cart the buyer
// already had is released first.
//
// POST /payments/intent
func (h *Handler) handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body createPaymentIntentBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeResolveError(w, http.StatusBadRequest, "invalid body")
		return
	}
	slug := strings.TrimSpace(body.Slug)
	reg, email, ok, code := h.resolveContributor(r, slug)
	if !ok {
		writeResolveError(w, code, "not allowed")
		return
	}

	super := permissions.NewSuperActor()

	// Validate the chosen method belongs to this registry and is enabled.
	method, _, err := h.client.RegistryPaymentMethod().SelectById(r.Context(), super,
		registry_payment_method.SelectByIdQuery{Id: strings.TrimSpace(body.PaymentMethodId)},
		registrypaymentmethodapi.NewProjection(true),
	)
	if err != nil || method.Model.RegistryId != reg.Id || !method.Model.Enabled {
		writeResolveError(w, http.StatusBadRequest, "invalid payment method")
		return
	}

	// Release any abandoned Pending cart this buyer still holds so its
	// reservations return to the open cart before we re-lock them.
	h.releasePendingCarts(r.Context(), super, reg.Id, email)

	// Gather the buyer's open reservations (Reserved, not yet in any cart).
	items := h.loadItemsByRegistry(r, super, reg.Id)
	resvResult, _, err := h.client.Reservation().Search(r.Context(), super,
		reservation.WhereClause{RegistryIdEq: &reg.Id},
		reservationapi.QueryOptions{Limit: 1000},
	)
	if err != nil {
		writeResolveError(w, http.StatusInternalServerError, "lookup error")
		return
	}

	now := time.Now().UTC()
	eligible := []reservation.Model{}
	contributorName := ""
	amountCents := 0
	currency := ""
	for _, rsv := range resvResult.Data {
		m := rsv.Model
		// Reserved is the source of truth for "open / available to pay". A
		// leftover CartId on a Reserved reservation is stale (the generated
		// Update omits empty Refs, so a released cart cannot blank it) and must
		// be ignored here, otherwise the buyer could never re-checkout it.
		if m.Status != enum_reservation_status.Reserved {
			continue
		}
		if strings.ToLower(strings.TrimSpace(m.ContactEmail)) != email {
			continue
		}
		if !m.ExpiresAt.IsZero() && m.ExpiresAt.Before(now) {
			continue
		}
		qty := m.Quantity
		if qty <= 0 {
			qty = 1
		}
		if it, found := items[m.ItemId]; found && it.PriceCents > 0 {
			amountCents += it.PriceCents * qty
			if currency == "" {
				currency = it.Currency
			}
		}
		if contributorName == "" {
			contributorName = m.ReserverName
		}
		eligible = append(eligible, m)
	}

	if len(eligible) == 0 {
		writeResolveError(w, http.StatusBadRequest, "your cart is empty")
		return
	}
	if currency == "" {
		currency = "USD"
	}

	// Create the cart first so reservations point at a real id.
	created, _, err := h.client.Cart().Create(r.Context(), super, cart.Model{
		OwnerId:           reg.OwnerId,
		RegistryId:        reg.Id,
		PaymentMethodId:   method.Model.Id,
		MethodType:        method.Model.Type,
		MethodDisplayName: firstNonEmpty(method.Model.DisplayName, string(method.Model.Type)),
		ReferenceCode:     generateReferenceCode(),
		AmountCents:       amountCents,
		Currency:          currency,
		ContributorEmail:  email,
		ContributorName:   contributorName,
		Status:            enum_cart_status.Pending,
	}, cart.NewProjection(true))
	if err != nil {
		writeResolveError(w, http.StatusInternalServerError, "could not start payment: "+err.Error())
		return
	}

	// Lock each reservation into the cart. Re-read to guard against a race
	// where another request claimed it between our search and now.
	lockedAny := false
	for _, m := range eligible {
		fresh, _, sErr := h.client.Reservation().SelectById(r.Context(), super,
			reservation.SelectByIdQuery{Id: m.Id}, reservationapi.NewProjection(true))
		if sErr != nil {
			continue
		}
		// Only the Reserved status gates locking; a stale CartId is ignored and
		// will be overwritten with the new cart id below.
		if fresh.Model.Status != enum_reservation_status.Reserved {
			continue
		}
		fresh.Model.CartId = created.Id
		fresh.Model.Status = enum_reservation_status.AwaitingConfirmation
		if _, _, uErr := h.client.Reservation().Update(r.Context(), super, fresh.Model, reservation.NewProjection(true)); uErr == nil {
			lockedAny = true
		}
	}

	if !lockedAny {
		// Everything got grabbed by another request; roll back the empty cart.
		_ = h.client.Cart().Delete(r.Context(), super, created.Id)
		writeResolveError(w, http.StatusConflict, "those gifts were just claimed by someone else")
		return
	}

	snapshots := h.cartLineItems(r.Context(), super, created.Id, items)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":            true,
		"id":            created.Id,
		"referenceCode": created.ReferenceCode,
		"amountCents":   created.AmountCents,
		"currency":      created.Currency,
		"status":        string(created.Status),
		"items":         snapshots,
	})
}

type claimPaymentBody struct {
	Slug    string `json:"slug"`
	Message string `json:"message"`
}

// handlePaymentRoute handles cart sub-actions:
//
//	POST /payments/{id}/claim  — buyer asserts they have sent the money.
//	POST /payments/{id}/cancel — buyer backs out; release the locked gifts.
func (h *Handler) handlePaymentRoute(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/payments/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 || r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	cartId, action := parts[0], parts[1]
	if action != "claim" && action != "cancel" {
		http.NotFound(w, r)
		return
	}

	var body claimPaymentBody
	_ = json.NewDecoder(r.Body).Decode(&body)
	slug := strings.TrimSpace(body.Slug)
	_, email, ok, code := h.resolveContributor(r, slug)
	if !ok {
		writeResolveError(w, code, "not allowed")
		return
	}

	super := permissions.NewSuperActor()
	current, _, err := h.client.Cart().SelectById(r.Context(), super,
		cart.SelectByIdQuery{Id: cartId}, cartapi.NewProjection(true))
	if err != nil {
		writeResolveError(w, http.StatusNotFound, "cart not found")
		return
	}
	// Only the contributor who created it may act on it.
	if strings.ToLower(strings.TrimSpace(current.Model.ContributorEmail)) != email {
		writeResolveError(w, http.StatusForbidden, "forbidden")
		return
	}

	switch action {
	case "claim":
		if current.Model.Status == enum_cart_status.Pending {
			current.Model.Status = enum_cart_status.AwaitingConfirmation
			current.Model.ClaimedAt = time.Now().UTC()
			if msg := strings.TrimSpace(body.Message); msg != "" {
				current.Model.Message = msg
			}
			if updated, _, uErr := h.client.Cart().Update(r.Context(), super, current.Model, cart.NewProjection(true)); uErr == nil {
				current.Model = updated
			} else {
				writeResolveError(w, http.StatusInternalServerError, "update failed")
				return
			}
		}
	case "cancel":
		if current.Model.Status == enum_cart_status.Pending {
			h.releaseCart(r.Context(), super, current.Model.Id)
			_ = h.client.Cart().Delete(r.Context(), super, current.Model.Id)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": "Cancelled"})
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "status": string(current.Model.Status)})
}

// releasePendingCarts releases every Pending cart held by the buyer for a
// registry, returning their reservations to the open cart.
func (h *Handler) releasePendingCarts(ctx context.Context, super permissions.Actor, registryId, email string) {
	pending := enum_cart_status.Pending
	res, _, err := h.client.Cart().Search(ctx, super, cart.WhereClause{
		RegistryIdEq:       &registryId,
		ContributorEmailEq: &email,
		StatusEq:           &pending,
	}, cartapi.QueryOptions{Limit: 50})
	if err != nil {
		return
	}
	for i := range res.Data {
		c := res.Data[i].Model
		h.releaseCart(ctx, super, c.Id)
		_ = h.client.Cart().Delete(ctx, super, c.Id)
	}
}

// releaseCart returns a cart's reservations to the open cart (Reserved, no
// cartId). Best effort.
func (h *Handler) releaseCart(ctx context.Context, super permissions.Actor, cartId string) {
	res, _, err := h.client.Reservation().Search(ctx, super,
		reservation.WhereClause{CartIdEq: &cartId}, reservationapi.QueryOptions{Limit: 1000})
	if err != nil {
		return
	}
	for i := range res.Data {
		m := res.Data[i].Model
		m.CartId = ""
		if m.Status == enum_reservation_status.AwaitingConfirmation {
			m.Status = enum_reservation_status.Reserved
		}
		_, _, _ = h.client.Reservation().Update(ctx, super, m, reservation.NewProjection(true))
	}
}

// cartLineItems joins a cart's reservations with their items into snapshots.
func (h *Handler) cartLineItems(ctx context.Context, super permissions.Actor, cartId string, items map[string]registry_item.Model) []cartItemSnapshot {
	out := []cartItemSnapshot{}
	res, _, err := h.client.Reservation().Search(ctx, super,
		reservation.WhereClause{CartIdEq: &cartId}, reservationapi.QueryOptions{Limit: 1000})
	if err != nil {
		return out
	}
	for i := range res.Data {
		m := res.Data[i].Model
		qty := m.Quantity
		if qty <= 0 {
			qty = 1
		}
		it := items[m.ItemId]
		out = append(out, cartItemSnapshot{
			ReservationId: m.Id,
			ItemId:        m.ItemId,
			Title:         it.Title,
			Quantity:      qty,
			PriceCents:    it.PriceCents,
			Currency:      it.Currency,
		})
	}
	return out
}

// loadItemsByRegistry returns a map of itemId -> item model for a registry.
func (h *Handler) loadItemsByRegistry(r *http.Request, super permissions.Actor, registryId string) map[string]registry_item.Model {
	out := map[string]registry_item.Model{}
	res, _, err := h.client.RegistryItem().Search(r.Context(), super,
		registry_item.WhereClause{RegistryIdEq: &registryId},
		registryitemapi.QueryOptions{Limit: 500},
	)
	if err != nil {
		return out
	}
	for i := range res.Data {
		out[res.Data[i].Model.Id] = res.Data[i].Model
	}
	return out
}

// generateReferenceCode returns a code like "STORK-8F4K2M" using an
// unambiguous alphabet (no 0/O/1/I).
func generateReferenceCode() string {
	const alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	const n = 6
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		// Fall back to a time-derived suffix; collisions are guarded by the
		// unique index on referenceCode anyway.
		return "STORK-" + strings.ToUpper(time.Now().UTC().Format("150405"))
	}
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = alphabet[int(buf[i])%len(alphabet)]
	}
	return "STORK-" + string(out)
}
