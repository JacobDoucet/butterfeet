package shipping

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_request_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_access_level"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	owneruserapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	registryapprovedguestapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
	shippingaddressrequestapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/mailer"
	"github.com/rs/zerolog/log"
)

// ActorResolver mirrors the auth package's resolver to avoid an import cycle.
type ActorResolver func(*http.Request) (permissions.Actor, error)

type Handler struct {
	mux        *http.ServeMux
	client     api.Client
	resolver   ActorResolver
	mailer     mailer.Mailer
	appBaseURL string
}

func NewHandler(client api.Client, resolver ActorResolver, notificationMailer mailer.Mailer, appBaseURL string) *Handler {
	h := &Handler{
		client:     client,
		resolver:   resolver,
		mailer:     notificationMailer,
		appBaseURL: strings.TrimRight(appBaseURL, "/"),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/registries/", h.handleRegistryScoped)
	mux.HandleFunc("/approved-guests/", h.handleGuestById)
	mux.HandleFunc("/requests/", h.handleRequestById)
	h.mux = mux
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// handleRegistryScoped covers:
//
//	GET  /registries/{registryId}/approved-guests
//	POST /registries/{registryId}/approved-guests
func (h *Handler) handleRegistryScoped(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/registries/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 {
		http.NotFound(w, r)
		return
	}
	registryId := parts[0]

	actor, _ := h.requireOwner(w, r)
	if actor == nil {
		return
	}
	reg, err := h.requireRegistry(r.Context(), actor, registryId)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "registry not found")
		return
	}

	switch parts[1] {
	case "approved-guests":
		switch r.Method {
		case http.MethodGet:
			h.listGuests(w, r, actor, reg)
		case http.MethodPost:
			h.addGuest(w, r, actor, reg)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	case "requests":
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.listRequests(w, r, actor, reg)
	default:
		http.NotFound(w, r)
	}
}

// handleGuestById covers:
//
//	POST /approved-guests/{id}/revoke
//	POST /approved-guests/{id}/block
//	POST /approved-guests/{id}/reactivate
//	DELETE /approved-guests/{id}
func (h *Handler) handleGuestById(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/approved-guests/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) < 1 || parts[0] == "" {
		http.NotFound(w, r)
		return
	}
	guestId := parts[0]

	actor, _ := h.requireOwner(w, r)
	if actor == nil {
		return
	}

	if r.Method == http.MethodDelete && len(parts) == 1 {
		if err := h.client.RegistryApprovedGuest().Delete(r.Context(), actor, guestId); err != nil {
			writeJSONError(w, http.StatusForbidden, "could not delete")
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
		return
	}

	if r.Method != http.MethodPost || len(parts) != 2 {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if parts[1] == "issue-link" {
		h.issueLinkForGuest(w, r, actor, guestId)
		return
	}

	var newStatus enum_guest_status.Value
	switch parts[1] {
	case "revoke":
		newStatus = enum_guest_status.Revoked
	case "block":
		newStatus = enum_guest_status.Blocked
	case "reactivate":
		newStatus = enum_guest_status.Active
	default:
		http.NotFound(w, r)
		return
	}

	current, _, err := h.client.RegistryApprovedGuest().SelectById(
		r.Context(), actor,
		registry_approved_guest.SelectByIdQuery{Id: guestId},
		registryapprovedguestapi.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "guest not found")
		return
	}
	current.Model.Status = newStatus
	updated, _, err := h.client.RegistryApprovedGuest().Update(
		r.Context(), actor, current.Model, registry_approved_guest.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusForbidden, "update failed")
		return
	}
	writeJSON(w, http.StatusOK, guestResponse(&registryapprovedguestapi.Model{Model: updated}))
}

func (h *Handler) listGuests(w http.ResponseWriter, r *http.Request, actor permissions.Actor, reg *registryapi.Model) {
	regId := reg.Model.Id
	res, _, err := h.client.RegistryApprovedGuest().Search(
		r.Context(), actor,
		registry_approved_guest.WhereClause{RegistryIdEq: &regId},
		registryapprovedguestapi.QueryOptions{Limit: 500},
	)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	out := make([]map[string]any, 0, len(res.Data))
	for i := range res.Data {
		out = append(out, guestResponse(&res.Data[i]))
	}
	writeJSON(w, http.StatusOK, map[string]any{"data": out})
}

type addGuestBody struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessLevel string `json:"accessLevel"`
}

func (h *Handler) addGuest(w http.ResponseWriter, r *http.Request, actor permissions.Actor, reg *registryapi.Model) {
	var body addGuestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid body")
		return
	}
	email := strings.TrimSpace(body.Email)
	if email == "" || !strings.ContainsRune(email, '@') {
		writeJSONError(w, http.StatusBadRequest, "valid email required")
		return
	}
	accessLevel := enum_guest_access_level.ViewShippingAddress
	if body.AccessLevel == string(enum_guest_access_level.ReserveOnly) {
		accessLevel = enum_guest_access_level.ReserveOnly
	}

	hash := HashEmail(email)
	regId := reg.Model.Id

	// Upsert: if a row already exists with the same (registry, emailHash),
	// reactivate / update access level rather than failing on the unique index.
	existing, _, err := h.client.RegistryApprovedGuest().Search(
		r.Context(), actor,
		registry_approved_guest.WhereClause{RegistryIdEq: &regId, EmailHashEq: &hash},
		registryapprovedguestapi.QueryOptions{Limit: 1},
	)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(existing.Data) > 0 {
		existing.Data[0].Model.Status = enum_guest_status.Active
		existing.Data[0].Model.AccessLevel = accessLevel
		if strings.TrimSpace(body.Name) != "" {
			existing.Data[0].Model.Name = strings.TrimSpace(body.Name)
		}
		updated, _, err := h.client.RegistryApprovedGuest().Update(
			r.Context(), actor, existing.Data[0].Model, registry_approved_guest.NewProjection(true),
		)
		if err != nil {
			writeJSONError(w, http.StatusForbidden, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, guestResponse(&registryapprovedguestapi.Model{Model: updated}))
		return
	}

	created, _, err := h.client.RegistryApprovedGuest().Create(r.Context(), actor, registry_approved_guest.Model{
		OwnerId:     reg.Model.OwnerId,
		RegistryId:  regId,
		EmailHash:   hash,
		EmailEnc:    EncryptEmail(email),
		Name:        strings.TrimSpace(body.Name),
		AccessLevel: accessLevel,
		Status:      enum_guest_status.Active,
	}, registry_approved_guest.NewProjection(true))
	if err != nil {
		writeJSONError(w, http.StatusForbidden, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, guestResponse(&registryapprovedguestapi.Model{Model: created}))
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

func (h *Handler) requireRegistry(ctx context.Context, actor permissions.Actor, registryId string) (*registryapi.Model, error) {
	reg, _, err := h.client.Registry().SelectById(
		ctx, actor, registry.SelectByIdQuery{Id: registryId}, registryapi.NewProjection(true),
	)
	if err != nil {
		return nil, err
	}
	return &reg, nil
}

func guestResponse(g *registryapprovedguestapi.Model) map[string]any {
	return map[string]any{
		"id":          g.Model.Id,
		"registryId":  g.Model.RegistryId,
		"email":       DecryptEmail(g.Model.EmailEnc),
		"name":        g.Model.Name,
		"accessLevel": string(g.Model.AccessLevel),
		"status":      string(g.Model.Status),
	}
}

func requestResponse(req *shippingaddressrequestapi.Model) map[string]any {
	return map[string]any{
		"id":             req.Model.Id,
		"registryId":     req.Model.RegistryId,
		"registryItemId": req.Model.RegistryItemId,
		"email":          DecryptEmail(req.Model.EmailEnc),
		"name":           req.Model.Name,
		"note":           req.Model.Note,
		"status":         string(req.Model.Status),
		"decisionReason": req.Model.DecisionReason,
		"policyVersion":  req.Model.PolicyVersion,
		"createdAt":      req.Model.Created.At,
	}
}

func (h *Handler) listRequests(w http.ResponseWriter, r *http.Request, actor permissions.Actor, reg *registryapi.Model) {
	regId := reg.Model.Id
	res, _, err := h.client.ShippingAddressRequest().Search(
		r.Context(), actor,
		shipping_address_request.WhereClause{RegistryIdEq: &regId},
		shippingaddressrequestapi.QueryOptions{Limit: 500},
	)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	out := make([]map[string]any, 0, len(res.Data))
	for i := range res.Data {
		out = append(out, requestResponse(&res.Data[i]))
	}
	writeJSON(w, http.StatusOK, map[string]any{"data": out})
}

type requestDecisionBody struct {
	Permanent bool   `json:"permanent"`
	Reason    string `json:"reason"`
}

func (h *Handler) handleRequestById(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/requests/")
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	if len(parts) != 2 || r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	reqId, action := parts[0], parts[1]

	actor, _ := h.requireOwner(w, r)
	if actor == nil {
		return
	}

	var body requestDecisionBody
	_ = json.NewDecoder(r.Body).Decode(&body)

	current, _, err := h.client.ShippingAddressRequest().SelectById(
		r.Context(), actor,
		shipping_address_request.SelectByIdQuery{Id: reqId},
		shippingaddressrequestapi.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "request not found")
		return
	}

	switch action {
	case "approve":
		current.Model.Status = enum_address_request_status.Approved
		current.Model.DecisionReason = strings.TrimSpace(body.Reason)
		if body.Permanent {
			if err := h.upsertApprovedGuest(r.Context(), actor, current.Model); err != nil {
				writeJSONError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	case "reject":
		current.Model.Status = enum_address_request_status.Rejected
		current.Model.DecisionReason = strings.TrimSpace(body.Reason)
	case "block":
		current.Model.Status = enum_address_request_status.Blocked
		current.Model.DecisionReason = strings.TrimSpace(body.Reason)
		if err := h.blockGuest(r.Context(), actor, current.Model); err != nil {
			writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		http.NotFound(w, r)
		return
	}

	updated, _, err := h.client.ShippingAddressRequest().Update(
		r.Context(), actor, current.Model, shipping_address_request.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusForbidden, err.Error())
		return
	}

	resp := requestResponse(&shippingaddressrequestapi.Model{Model: updated})
	if action == "approve" {
		reg, regErr := h.requireRegistry(r.Context(), actor, updated.RegistryId)
		if regErr == nil {
			raw, exp, err := h.createAccessSession(
				r.Context(), actor,
				updated.OwnerId, updated.RegistryId, updated.EmailHash, "",
				reg.Model.ShippingPolicyVersion,
			)
			if err == nil {
				resp["token"] = raw
				resp["tokenExpiresAt"] = exp
				h.sendBuyerApprovedNotification(updated.OwnerId, reg.Model.Title, DecryptEmail(updated.EmailEnc), updated.Name, raw, exp)
			}
		}
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) upsertApprovedGuest(ctx context.Context, actor permissions.Actor, req shipping_address_request.Model) error {
	existing, _, err := h.client.RegistryApprovedGuest().Search(ctx, actor,
		registry_approved_guest.WhereClause{RegistryIdEq: &req.RegistryId, EmailHashEq: &req.EmailHash},
		registryapprovedguestapi.QueryOptions{Limit: 1})
	if err != nil {
		return err
	}
	if len(existing.Data) > 0 {
		existing.Data[0].Model.Status = enum_guest_status.Active
		if strings.TrimSpace(existing.Data[0].Model.Name) == "" {
			existing.Data[0].Model.Name = req.Name
		}
		_, _, err := h.client.RegistryApprovedGuest().Update(ctx, actor, existing.Data[0].Model, registry_approved_guest.NewProjection(true))
		return err
	}
	_, _, err = h.client.RegistryApprovedGuest().Create(ctx, actor, registry_approved_guest.Model{
		OwnerId:     req.OwnerId,
		RegistryId:  req.RegistryId,
		EmailHash:   req.EmailHash,
		EmailEnc:    req.EmailEnc,
		Name:        req.Name,
		AccessLevel: enum_guest_access_level.ViewShippingAddress,
		Status:      enum_guest_status.Active,
	}, registry_approved_guest.NewProjection(true))
	return err
}

func (h *Handler) blockGuest(ctx context.Context, actor permissions.Actor, req shipping_address_request.Model) error {
	existing, _, err := h.client.RegistryApprovedGuest().Search(ctx, actor,
		registry_approved_guest.WhereClause{RegistryIdEq: &req.RegistryId, EmailHashEq: &req.EmailHash},
		registryapprovedguestapi.QueryOptions{Limit: 1})
	if err != nil {
		return err
	}
	if len(existing.Data) > 0 {
		existing.Data[0].Model.Status = enum_guest_status.Blocked
		_, _, err := h.client.RegistryApprovedGuest().Update(ctx, actor, existing.Data[0].Model, registry_approved_guest.NewProjection(true))
		return err
	}
	_, _, err = h.client.RegistryApprovedGuest().Create(ctx, actor, registry_approved_guest.Model{
		OwnerId:     req.OwnerId,
		RegistryId:  req.RegistryId,
		EmailHash:   req.EmailHash,
		EmailEnc:    req.EmailEnc,
		Name:        req.Name,
		AccessLevel: enum_guest_access_level.ViewShippingAddress,
		Status:      enum_guest_status.Blocked,
	}, registry_approved_guest.NewProjection(true))
	return err
}

const accessTokenTTL = 7 * 24 * time.Hour

// createAccessSession persists a new AddressAccessSession for the given guest
// identity and returns the raw token (only ever returned to the owner; never stored).
func (h *Handler) createAccessSession(
	ctx context.Context, actor permissions.Actor,
	ownerId, registryId, emailHash, approvedGuestId string, policyVersion int,
) (rawToken string, expiresAt time.Time, err error) {
	raw, hash, err := NewToken()
	if err != nil {
		return "", time.Time{}, err
	}
	exp := time.Now().Add(accessTokenTTL).UTC()
	_, _, err = h.client.AddressAccessSession().Create(ctx, actor, address_access_session.Model{
		OwnerId:              ownerId,
		RegistryId:           registryId,
		ApprovedGuestId:      approvedGuestId,
		EmailHash:            emailHash,
		TokenHash:            hash,
		ExpiresAt:            exp,
		PolicyVersionAtIssue: policyVersion,
	}, address_access_session.NewProjection(true))
	if err != nil {
		return "", time.Time{}, err
	}
	return raw, exp, nil
}

func (h *Handler) issueLinkForGuest(w http.ResponseWriter, r *http.Request, actor permissions.Actor, guestId string) {
	guest, _, err := h.client.RegistryApprovedGuest().SelectById(
		r.Context(), actor,
		registry_approved_guest.SelectByIdQuery{Id: guestId},
		registryapprovedguestapi.NewProjection(true),
	)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "guest not found")
		return
	}
	if guest.Model.Status != enum_guest_status.Active {
		writeJSONError(w, http.StatusForbidden, "guest is not active")
		return
	}
	reg, err := h.requireRegistry(r.Context(), actor, guest.Model.RegistryId)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "registry not found")
		return
	}
	raw, exp, err := h.createAccessSession(
		r.Context(), actor,
		guest.Model.OwnerId, guest.Model.RegistryId, guest.Model.EmailHash,
		guest.Model.Id, reg.Model.ShippingPolicyVersion,
	)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"token":     raw,
		"expiresAt": exp,
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]any{"error": msg})
}

func (h *Handler) sendBuyerApprovedNotification(ownerId, registryTitle, buyerEmail, buyerName, rawToken string, expiresAt time.Time) {
	if h.mailer == nil || strings.TrimSpace(buyerEmail) == "" || strings.TrimSpace(rawToken) == "" {
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var ownerName string
		if strings.TrimSpace(ownerId) != "" {
			super := permissions.NewSuperActor()
			owner, _, err := h.client.OwnerUser().SelectById(ctx, super,
				owner_user.SelectByIdQuery{Id: ownerId},
				owneruserapi.NewProjection(true),
			)
			if err == nil {
				ownerName = strings.TrimSpace(owner.Name)
			}
		}
		if ownerName == "" {
			ownerName = "the parents"
		}

		greeting := strings.TrimSpace(buyerName)
		if greeting == "" {
			greeting = "there"
		}

		link := h.appBaseURL + "/ship#tok=" + rawToken

		err := h.mailer.Send(ctx, mailer.Message{
			To:      buyerEmail,
			Subject: "Your shipping address request was approved",
			Text: "Hi " + greeting + ",\n\n" +
				ownerName + " approved your request to view the shipping address for the \"" + registryTitle + "\" registry.\n\n" +
				"View the address here (link expires " + expiresAt.UTC().Format("Jan 2, 2006 15:04 UTC") + "):\n" +
				link + "\n\n" +
				"Keep this link private — anyone with it can view the address until it expires.\n",
		})
		if err != nil {
			log.Error().Err(err).Str("ownerId", ownerId).Str("to", buyerEmail).Msg("buyer approved notification send failed")
		}
	}()
}
