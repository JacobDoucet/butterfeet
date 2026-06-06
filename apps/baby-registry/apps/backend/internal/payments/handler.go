// Package payments implements the owner-facing endpoints for the manual
// (parent-to-parent) payment flow: configuring enabled payment methods on a
// registry and reviewing/approving the contribution payments guests claim to
// have sent. Stork Nest never processes or holds money — parent approval is
// the source of truth.
package payments

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	cartapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_cart_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_payment_method_type"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	registryitemapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
	registrypaymentmethodapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	reservationapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_api"
	"github.com/rs/zerolog/log"
)

// ActorResolver mirrors the auth package's resolver to avoid an import cycle.
type ActorResolver func(*http.Request) (permissions.Actor, error)

type Handler struct {
	mux      *http.ServeMux
	client   api.Client
	resolver ActorResolver
}

func NewHandler(client api.Client, resolver ActorResolver) *Handler {
	h := &Handler{client: client, resolver: resolver}
	mux := http.NewServeMux()
	mux.HandleFunc("/registries/", h.handleRegistryScoped)
	mux.HandleFunc("/payment-methods/", h.handleMethodById)
	mux.HandleFunc("/carts/", h.handleCartById)
	h.mux = mux
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// handleRegistryScoped covers:
//
//	GET  /registries/{registryId}/payment-methods
//	POST /registries/{registryId}/payment-methods
//	GET  /registries/{registryId}/carts
func (h *Handler) handleRegistryScoped(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/registries/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 {
		http.NotFound(w, r)
		return
	}
	registryId := parts[0]

	actor := h.requireOwner(w, r)
	if actor == nil {
		return
	}
	reg, err := h.requireRegistry(r.Context(), actor, registryId)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "registry not found")
		return
	}

	switch parts[1] {
	case "payment-methods":
		switch r.Method {
		case http.MethodGet:
			h.listMethods(w, r, actor, reg.Model.Id)
		case http.MethodPost:
			h.createMethod(w, r, actor, reg.Model.OwnerId, reg.Model.Id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	case "carts":
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.listCarts(w, r, actor, reg.Model.Id)
	default:
		http.NotFound(w, r)
	}
}

// handleMethodById covers:
//
//	PATCH  /payment-methods/{id}
//	DELETE /payment-methods/{id}
func (h *Handler) handleMethodById(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/payment-methods/")
	id := strings.Trim(rest, "/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	actor := h.requireOwner(w, r)
	if actor == nil {
		return
	}

	switch r.Method {
	case http.MethodPatch:
		h.updateMethod(w, r, actor, id)
	case http.MethodDelete:
		if err := h.client.RegistryPaymentMethod().Delete(r.Context(), actor, id); err != nil {
			writeJSONError(w, http.StatusForbidden, "could not delete")
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleCartById covers:
//
//	POST /carts/{id}/approve
//	POST /carts/{id}/reject
func (h *Handler) handleCartById(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/carts/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 || r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	cartId, action := parts[0], parts[1]

	actor := h.requireOwner(w, r)
	if actor == nil {
		return
	}

	var body struct {
		Reason string `json:"reason"`
	}
	_ = json.NewDecoder(r.Body).Decode(&body)

	current, _, err := h.client.Cart().SelectById(
		r.Context(), actor,
		cart.SelectByIdQuery{Id: cartId},
		cartapi.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "cart not found")
		return
	}

	switch action {
	case "approve":
		current.Model.Status = enum_cart_status.Completed
		current.Model.DecidedAt = time.Now().UTC()
		current.Model.DecisionReason = strings.TrimSpace(body.Reason)
		// Approval is the source of truth: mark the locked gifts purchased.
		h.setCartReservations(r.Context(), actor, current.Model.Id, enum_reservation_status.Purchased)
	case "reject":
		current.Model.Status = enum_cart_status.Rejected
		current.Model.DecidedAt = time.Now().UTC()
		current.Model.DecisionReason = strings.TrimSpace(body.Reason)
		// Rejection frees the gifts so other guests can claim them.
		h.setCartReservations(r.Context(), actor, current.Model.Id, enum_reservation_status.Cancelled)
	default:
		http.NotFound(w, r)
		return
	}

	updated, _, err := h.client.Cart().Update(
		r.Context(), actor, current.Model, cart.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusForbidden, "update failed")
		return
	}
	writeJSON(w, http.StatusOK, cartResponse(updated, h.cartItemsFor(r.Context(), actor, updated)))
}

// setCartReservations flips every reservation locked into a cart to a new
// status. Best-effort: a failure on one reservation should not block the
// owner's decision.
func (h *Handler) setCartReservations(ctx context.Context, actor permissions.Actor, cartId string, newStatus enum_reservation_status.Value) {
	res, _, err := h.client.Reservation().Search(
		ctx, actor, reservation.WhereClause{CartIdEq: &cartId}, reservationapi.QueryOptions{Limit: 1000},
	)
	if err != nil {
		log.Warn().Str("cartId", cartId).Err(err).Msg("decision: reservation lookup failed")
		return
	}
	for i := range res.Data {
		m := res.Data[i].Model
		if m.Status != enum_reservation_status.AwaitingConfirmation {
			continue
		}
		m.Status = newStatus
		if _, _, err := h.client.Reservation().Update(ctx, actor, m, reservation.NewProjection(true)); err != nil {
			log.Warn().Str("reservationId", m.Id).Err(err).Msg("decision: reservation update failed")
		}
	}
}

type methodBody struct {
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
	Enabled           *bool  `json:"enabled"`
	Position          *int   `json:"position"`
}

func (h *Handler) listMethods(w http.ResponseWriter, r *http.Request, actor permissions.Actor, registryId string) {
	res, _, err := h.client.RegistryPaymentMethod().Search(
		r.Context(), actor,
		registry_payment_method.WhereClause{RegistryIdEq: &registryId},
		registrypaymentmethodapi.QueryOptions{Limit: 100},
	)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	out := make([]map[string]any, 0, len(res.Data))
	for i := range res.Data {
		out = append(out, methodResponse(&res.Data[i]))
	}
	writeJSON(w, http.StatusOK, map[string]any{"data": out})
}

func (h *Handler) createMethod(w http.ResponseWriter, r *http.Request, actor permissions.Actor, ownerId, registryId string) {
	var body methodBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid body")
		return
	}
	t, ok := parseMethodType(body.Type)
	if !ok {
		writeJSONError(w, http.StatusBadRequest, "invalid payment method type")
		return
	}
	enabled := true
	if body.Enabled != nil {
		enabled = *body.Enabled
	}
	pos := 0
	if body.Position != nil {
		pos = *body.Position
	}
	created, _, err := h.client.RegistryPaymentMethod().Create(r.Context(), actor, registry_payment_method.Model{
		OwnerId:           ownerId,
		RegistryId:        registryId,
		Type:              t,
		DisplayName:       strings.TrimSpace(body.DisplayName),
		Instructions:      strings.TrimSpace(body.Instructions),
		PaymentUrl:        strings.TrimSpace(body.PaymentUrl),
		RecipientEmail:    strings.TrimSpace(body.RecipientEmail),
		RecipientPhone:    strings.TrimSpace(body.RecipientPhone),
		BankName:          strings.TrimSpace(body.BankName),
		BankAccountName:   strings.TrimSpace(body.BankAccountName),
		BankAccountNumber: strings.TrimSpace(body.BankAccountNumber),
		BankRoutingNumber: strings.TrimSpace(body.BankRoutingNumber),
		BankIban:          strings.TrimSpace(body.BankIban),
		BankSwift:         strings.TrimSpace(body.BankSwift),
		Enabled:           enabled,
		Position:          pos,
	}, registry_payment_method.NewProjection(true))
	if err != nil {
		writeJSONError(w, http.StatusForbidden, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, methodResponse(&registrypaymentmethodapi.Model{Model: created}))
}

func (h *Handler) updateMethod(w http.ResponseWriter, r *http.Request, actor permissions.Actor, id string) {
	var body methodBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid body")
		return
	}
	current, _, err := h.client.RegistryPaymentMethod().SelectById(
		r.Context(), actor,
		registry_payment_method.SelectByIdQuery{Id: id},
		registrypaymentmethodapi.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "method not found")
		return
	}
	if t, ok := parseMethodType(body.Type); ok {
		current.Model.Type = t
	}
	current.Model.DisplayName = strings.TrimSpace(body.DisplayName)
	current.Model.Instructions = strings.TrimSpace(body.Instructions)
	current.Model.PaymentUrl = strings.TrimSpace(body.PaymentUrl)
	current.Model.RecipientEmail = strings.TrimSpace(body.RecipientEmail)
	current.Model.RecipientPhone = strings.TrimSpace(body.RecipientPhone)
	current.Model.BankName = strings.TrimSpace(body.BankName)
	current.Model.BankAccountName = strings.TrimSpace(body.BankAccountName)
	current.Model.BankAccountNumber = strings.TrimSpace(body.BankAccountNumber)
	current.Model.BankRoutingNumber = strings.TrimSpace(body.BankRoutingNumber)
	current.Model.BankIban = strings.TrimSpace(body.BankIban)
	current.Model.BankSwift = strings.TrimSpace(body.BankSwift)
	if body.Enabled != nil {
		current.Model.Enabled = *body.Enabled
	}
	if body.Position != nil {
		current.Model.Position = *body.Position
	}
	updated, _, err := h.client.RegistryPaymentMethod().Update(
		r.Context(), actor, current.Model, registry_payment_method.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusForbidden, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, methodResponse(&registrypaymentmethodapi.Model{Model: updated}))
}

func (h *Handler) listCarts(w http.ResponseWriter, r *http.Request, actor permissions.Actor, registryId string) {
	where := cart.WhereClause{RegistryIdEq: &registryId}
	switch statuses := parseCartStatuses(r.URL.Query().Get("status")); len(statuses) {
	case 1:
		where.StatusEq = &statuses[0]
	default:
		if len(statuses) > 1 {
			where.StatusIn = &statuses
		}
	}
	res, _, err := h.client.Cart().Search(
		r.Context(), actor, where, cartapi.QueryOptions{Limit: 500},
	)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Load all reservations + items for this registry once, then group line
	// items by cart to avoid a query per cart.
	itemMap := map[string]registry_item.Model{}
	if ir, _, ierr := h.client.RegistryItem().Search(
		r.Context(), actor, registry_item.WhereClause{RegistryIdEq: &registryId}, registryitemapi.QueryOptions{Limit: 500},
	); ierr == nil {
		for i := range ir.Data {
			itemMap[ir.Data[i].Model.Id] = ir.Data[i].Model
		}
	}
	cartItems := map[string][]map[string]any{}
	if rr, _, rerr := h.client.Reservation().Search(
		r.Context(), actor, reservation.WhereClause{RegistryIdEq: &registryId}, reservationapi.QueryOptions{Limit: 1000},
	); rerr == nil {
		for i := range rr.Data {
			m := rr.Data[i].Model
			cid := strings.TrimSpace(m.CartId)
			if cid == "" {
				continue
			}
			cartItems[cid] = append(cartItems[cid], lineItem(m, itemMap))
		}
	}

	out := make([]map[string]any, 0, len(res.Data))
	for i := range res.Data {
		c := res.Data[i].Model
		out = append(out, cartResponse(c, cartItems[c.Id]))
	}
	writeJSON(w, http.StatusOK, map[string]any{"data": out})
}

// cartItemsFor builds a cart's line items by joining its reservations with the
// registry's items.
func (h *Handler) cartItemsFor(ctx context.Context, actor permissions.Actor, c cart.Model) []map[string]any {
	out := []map[string]any{}
	res, _, err := h.client.Reservation().Search(
		ctx, actor, reservation.WhereClause{CartIdEq: &c.Id}, reservationapi.QueryOptions{Limit: 1000},
	)
	if err != nil {
		return out
	}
	itemMap := map[string]registry_item.Model{}
	if ir, _, ierr := h.client.RegistryItem().Search(
		ctx, actor, registry_item.WhereClause{RegistryIdEq: &c.RegistryId}, registryitemapi.QueryOptions{Limit: 500},
	); ierr == nil {
		for i := range ir.Data {
			itemMap[ir.Data[i].Model.Id] = ir.Data[i].Model
		}
	}
	for i := range res.Data {
		out = append(out, lineItem(res.Data[i].Model, itemMap))
	}
	return out
}

func lineItem(m reservation.Model, itemMap map[string]registry_item.Model) map[string]any {
	qty := m.Quantity
	if qty <= 0 {
		qty = 1
	}
	it := itemMap[m.ItemId]
	return map[string]any{
		"reservationId": m.Id,
		"itemId":        m.ItemId,
		"title":         it.Title,
		"quantity":      qty,
		"priceCents":    it.PriceCents,
		"currency":      it.Currency,
	}
}

func (h *Handler) requireOwner(w http.ResponseWriter, r *http.Request) permissions.Actor {
	actor, err := h.resolver(r)
	if err != nil || actor == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return nil
	}
	if _, ok := actor.(*owner_user.Model); !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return nil
	}
	return actor
}

func (h *Handler) requireRegistry(ctx context.Context, actor permissions.Actor, registryId string) (*registryapi.Model, error) {
	reg, _, err := h.client.Registry().SelectById(
		ctx, actor, registry.SelectByIdQuery{Id: registryId}, registryapi.NewProjection(true),
	)
	if err != nil {
		return nil, err
	}
	return &reg, nil
}

func parseMethodType(raw string) (enum_payment_method_type.Value, bool) {
	switch enum_payment_method_type.Value(strings.TrimSpace(raw)) {
	case enum_payment_method_type.PaymentLink:
		return enum_payment_method_type.PaymentLink, true
	case enum_payment_method_type.InteracETransfer:
		return enum_payment_method_type.InteracETransfer, true
	case enum_payment_method_type.BankTransfer:
		return enum_payment_method_type.BankTransfer, true
	default:
		return "", false
	}
}

func parseCartStatus(raw string) (enum_cart_status.Value, bool) {
	switch enum_cart_status.Value(strings.TrimSpace(raw)) {
	case enum_cart_status.Pending:
		return enum_cart_status.Pending, true
	case enum_cart_status.AwaitingConfirmation:
		return enum_cart_status.AwaitingConfirmation, true
	case enum_cart_status.Completed:
		return enum_cart_status.Completed, true
	case enum_cart_status.Rejected:
		return enum_cart_status.Rejected, true
	default:
		return "", false
	}
}

// parseCartStatuses parses a comma-separated list of cart statuses, ignoring
// any unrecognised values. Used to let one tab show multiple statuses (e.g.
// "To review" = Pending + AwaitingConfirmation).
func parseCartStatuses(raw string) []enum_cart_status.Value {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	var out []enum_cart_status.Value
	for _, part := range strings.Split(raw, ",") {
		if s, ok := parseCartStatus(part); ok {
			out = append(out, s)
		}
	}
	return out
}

func methodResponse(m *registrypaymentmethodapi.Model) map[string]any {
	return map[string]any{
		"id":                m.Model.Id,
		"registryId":        m.Model.RegistryId,
		"type":              string(m.Model.Type),
		"displayName":       m.Model.DisplayName,
		"instructions":      m.Model.Instructions,
		"paymentUrl":        m.Model.PaymentUrl,
		"recipientEmail":    m.Model.RecipientEmail,
		"recipientPhone":    m.Model.RecipientPhone,
		"bankName":          m.Model.BankName,
		"bankAccountName":   m.Model.BankAccountName,
		"bankAccountNumber": m.Model.BankAccountNumber,
		"bankRoutingNumber": m.Model.BankRoutingNumber,
		"bankIban":          m.Model.BankIban,
		"bankSwift":         m.Model.BankSwift,
		"enabled":           m.Model.Enabled,
		"position":          m.Model.Position,
	}
}

func cartResponse(c cart.Model, items []map[string]any) map[string]any {
	resp := map[string]any{
		"id":                c.Id,
		"registryId":        c.RegistryId,
		"paymentMethodId":   c.PaymentMethodId,
		"methodType":        string(c.MethodType),
		"methodDisplayName": c.MethodDisplayName,
		"referenceCode":     c.ReferenceCode,
		"amountCents":       c.AmountCents,
		"currency":          c.Currency,
		"contributorEmail":  c.ContributorEmail,
		"contributorName":   c.ContributorName,
		"message":           c.Message,
		"status":            string(c.Status),
		"decisionReason":    c.DecisionReason,
		"createdAt":         c.Created.At,
	}
	if !c.ClaimedAt.IsZero() {
		resp["claimedAt"] = c.ClaimedAt
	}
	if !c.DecidedAt.IsZero() {
		resp["decidedAt"] = c.DecidedAt
	}
	if items == nil {
		items = []map[string]any{}
	}
	resp["items"] = items
	return resp
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]any{"error": msg})
}

var _ = errors.New
