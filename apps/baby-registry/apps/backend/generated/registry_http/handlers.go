package registry_http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type GetMetadataOnSearchHook func(ctx context.Context, actor permissions.Actor, queryResult registry_api.HTTPQueryResult) (map[string]any, error)
type GetMetadataOnCreateHook func(ctx context.Context, actor permissions.Actor, mutationResult registry_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnUpdateHook func(ctx context.Context, actor permissions.Actor, mutationResult registry_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnDeleteHook func(ctx context.Context, actor permissions.Actor, deleteResult registry_api.HTTPDeleteResult) (map[string]any, error)

type MetadataHooks struct {
	OnSearch GetMetadataOnSearchHook
	OnCreate GetMetadataOnCreateHook
	OnUpdate GetMetadataOnUpdateHook
	OnDelete GetMetadataOnDeleteHook
}

type HandlerProps struct {
	Api           registry_api.Client
	ResolveActor  func(r *http.Request) (permissions.Actor, error)
	MetadataHooks []MetadataHooks
	OnError       func(handler string, e error)
}

func (p *HandlerProps) Validate() error {
	var err error
	if p.Api == nil {
		err = errors.Join(err, errors.New("api is required"))
	}
	if p.ResolveActor == nil {
		err = errors.Join(err, errors.New("resolveActor is required"))
	}
	return err
}

func (p *HandlerProps) onError(endpoint string, e error) {
	if p.OnError != nil {
		p.OnError(endpoint+"<Registry>", e)
	}
}

type SearchRequest struct {
	Query      registry.HTTPWhereClause `json:"query"`
	Sort       registry.HTTPSortParams  `json:"sort"`
	Projection *registry_api.Projection `json:"projection,omitempty"`
	Limit      int                      `json:"limit,omitempty"`
	Skip       int                      `json:"skip,omitempty"`
}

func (sr *SearchRequest) ResolveProjection() *registry_api.Projection {
	if sr.Projection != nil {
		return sr.Projection
	}
	projection := registry_api.NewProjection(true)
	projection.AddressAccessSessions = nil
	projection.RegistryApprovedGuests = nil
	projection.RegistryItems = nil
	projection.Reservations = nil
	projection.ShippingAddressRequests = nil

	return &projection
}

func GetSearchHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		searchRequest, err := resolveSearchRequest(r)
		if err != nil {
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		if searchRequest.Limit > 1000 {
			props.onError("Search", errors.New("search limit exceeded"))
			log.Debug().Msg("limit exceeded")
			http.Error(w, "MAX_LIMIT_1000", http.StatusBadRequest)
			return
		}

		searchQuery, err := searchRequest.Query.ToWhereClause()
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to resolve search query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		queryResult, projectionResult, err := props.Api.Search(ctx, actor, searchQuery, registry_api.QueryOptions{
			Projection: searchRequest.ResolveProjection(),
			Sort:       searchRequest.Sort.ToSortParams(),
			Limit:      searchRequest.Limit,
			Skip:       searchRequest.Skip,
		})
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to execute search query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := registry_api.ToHTTPQueryResult(queryResult, projectionResult)
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		for i, metadataHook := range props.MetadataHooks {
			if metadataHook.OnSearch != nil {
				metadata, err := metadataHook.OnSearch(ctx, actor, response)
				if err != nil {
					log.Debug().Err(err).Int("idx", i).Msg("failed to execute metadata hook")
					http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
					return
				}
				response.Metadata = metadata
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusInternalServerError)
			return
		}
	}, nil
}

type SelectRequest struct {
	Projection *registry_api.Projection `json:"projection,omitempty"`
}

func (sr *SelectRequest) ResolveProjection() registry_api.Projection {
	if sr.Projection != nil {
		return *sr.Projection
	}
	return registry_api.NewProjection(true)
}

func resolveSelectRequest(r *http.Request) (SelectRequest, error) {
	switch r.Method {
	case http.MethodPost:
		selectRequest := SelectRequest{}
		err := json.NewDecoder(r.Body).Decode(&selectRequest)
		if err != nil {
			return SelectRequest{}, coded_error.NewInvalidRequestError()
		}
		return selectRequest, nil
	case http.MethodGet:
		return SelectRequest{}, nil
	default:
		return SelectRequest{}, coded_error.NewMethodNotAllowedError(r.Method)
	}
}

type SelectByIdRequest struct {
	Query registry.HTTPSelectByIdQuery `json:"query"`
}

func GetSelectByIdHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		selectRequest, err := resolveSelectRequest(r)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to resolve select request")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		id := r.PathValue("id")
		if id == "" {
			props.onError("SelectById", errors.New("missing path parameter: id"))
			log.Debug().Err(err).Msg("missing path parameter")
			http.Error(w, "missing query parameter: id", http.StatusBadRequest)
			return
		}

		selectByIdRequest := SelectByIdRequest{
			Query: registry.HTTPSelectByIdQuery{
				Id: id,
			},
		}

		selectByIdQuery, err := selectByIdRequest.Query.ToSelectByIdQuery()
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to resolve select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		projection := selectRequest.ResolveProjection()

		getByIdResult, projectionResult, err := props.Api.SelectById(ctx, actor, selectByIdQuery, projection)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to execute select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := registry_api.ToHTTPModel(getByIdResult, projectionResult)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

type SelectBySlugUniqueRequest struct {
	Query registry.HTTPSelectBySlugUniqueQuery `json:"query"`
}

func GetSelectBySlugUniqueHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		selectRequest, err := resolveSelectRequest(r)
		if err != nil {
			props.onError("SelectBySlugUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select request")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("SelectBySlugUnique", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		slug := r.PathValue("slug")
		if slug == "" {
			props.onError("SelectBySlug", errors.New("missing path parameter: slug"))
			log.Debug().Err(err).Msg("missing path parameter")
			http.Error(w, "missing query parameter: slug", http.StatusBadRequest)
			return
		}

		selectBySlugUniqueRequest := SelectBySlugUniqueRequest{
			Query: registry.HTTPSelectBySlugUniqueQuery{
				Slug: slug,
			},
		}

		selectBySlugUniqueQuery, err := selectBySlugUniqueRequest.Query.ToSelectBySlugUniqueQuery()
		if err != nil {
			props.onError("SelectBySlugUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		projection := selectRequest.ResolveProjection()

		getBySlugUniqueResult, projectionResult, err := props.Api.SelectBySlugUnique(ctx, actor, selectBySlugUniqueQuery, projection)
		if err != nil {
			props.onError("SelectBySlugUnique", err)
			log.Debug().Err(err).Msg("failed to execute select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := registry_api.ToHTTPModel(getBySlugUniqueResult, projectionResult)
		if err != nil {
			props.onError("SelectBySlugUnique", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("SelectBySlugUnique", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

type SaveRequest struct {
	Data registry.HTTPRecord `json:"data"`
}

func GetCreateHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			props.onError("Create", errors.New("method not allowed"))
			log.Debug().Msg("method not allowed")
			http.Error(w, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to read request body")
			http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
			return
		}

		projection := registry.NewProjection(true)
		projection.Id = false

		var createRequest SaveRequest
		err = json.Unmarshal(body, &createRequest)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to unmarshal request body")
			http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
			return
		}

		create, err := createRequest.Data.ToModel()
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to convert request body to model")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusBadRequest)
			return
		}

		createResult, projectionResult, err := props.Api.Create(ctx, actor, create, projection)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to execute create query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := registry_api.ToHTTPMutationResult(createResult, projectionResult)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to convert create result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		for i, metadataHook := range props.MetadataHooks {
			if metadataHook.OnCreate != nil {
				metadata, err := metadataHook.OnCreate(ctx, actor, response)
				if err != nil {
					props.onError("Create", err)
					log.Debug().Err(err).Int("idx", i).Msg("failed to execute metadata hook")
					http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
					return
				}
				response.Metadata = metadata
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("Create", err)
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

func GetUpdateHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			props.onError("Update", errors.New("method not allowed"))
			log.Debug().Msg("method not allowed")
			http.Error(w, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Update", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			props.onError("Update", err)
			log.Debug().Err(err).Msg("failed to read request body")
			http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
			return
		}

		var updateRequest SaveRequest
		err = json.Unmarshal(body, &updateRequest)
		if err != nil {
			props.onError("Update", err)
			log.Debug().Err(err).Msg("failed to unmarshal request body")
			http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
			return
		}

		update, err := updateRequest.Data.ToModel()
		if err != nil {
			props.onError("Update", err)
			log.Debug().Err(err).Msg("failed to convert request body to model")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusBadRequest)
			return
		}

		projection, err := updateRequest.Data.ToProjection()
		if err != nil {
			props.onError("Update", err)
			log.Debug().Err(err).Msg("failed to convert request body to projection")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusBadRequest)
			return
		}

		updateResult, projectionResult, err := props.Api.Update(ctx, actor, update, projection)
		if err != nil {
			props.onError("Update", err)
			log.Debug().Err(err).Msg("failed to execute update query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := registry_api.ToHTTPMutationResult(updateResult, projectionResult)
		if err != nil {
			props.onError("Update", err)
			log.Debug().Err(err).Msg("failed to convert update result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		for i, metadataHook := range props.MetadataHooks {
			if metadataHook.OnUpdate != nil {
				metadata, err := metadataHook.OnUpdate(ctx, actor, response)
				if err != nil {
					props.onError("Update", err)
					log.Debug().Err(err).Int("idx", i).Msg("failed to execute metadata hook")
					http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
					return
				}
				response.Metadata = metadata
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("Update", err)
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

func GetDeleteHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			props.onError("Delete", errors.New("method not allowed"))
			log.Debug().Msg("method not allowed")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Delete", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		id := r.PathValue("id")
		if id == "" {
			props.onError("Delete", errors.New("missing path parameter: id"))
			log.Debug().Msg("missing path parameter")
			http.Error(w, "missing path parameter: id", http.StatusBadRequest)
			return
		}

		err = props.Api.Delete(ctx, actor, id)
		if err != nil {
			props.onError("Delete", err)
			log.Debug().Err(err).Msg("failed to execute delete query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response := registry_api.ToHTTPDeleteResult(id)

		for i, metadataHook := range props.MetadataHooks {
			if metadataHook.OnDelete != nil {
				metadata, err := metadataHook.OnDelete(ctx, actor, response)
				if err != nil {
					props.onError("Delete", err)
					log.Debug().Err(err).Int("idx", i).Msg("failed to execute metadata hook")
					http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
					return
				}
				response.Metadata = metadata
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("Delete", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

func resolveSearchRequest(r *http.Request) (SearchRequest, error) {
	switch r.Method {
	case http.MethodPost:
		searchRequest := SearchRequest{
			Limit: 100,
		}
		err := json.NewDecoder(r.Body).Decode(&searchRequest)
		if err != nil {
			return SearchRequest{}, err
		}
		return searchRequest, nil
	case http.MethodGet:
		searchRequest := SearchRequest{
			Query: registry.HTTPWhereClause{},
		}
		for key, values := range r.URL.Query() {
			switch key {
			case "Limit":
				limit := utils.StringSliceToIntPtr(values)
				if limit == nil {
					return SearchRequest{}, coded_error.NewInvalidRequestError("invalid limit")
				}
				searchRequest.Limit = *limit
				continue
			case "Skip":
				skip := utils.StringSliceToIntPtr(values)
				if skip == nil {
					return SearchRequest{}, coded_error.NewInvalidRequestError("invalid skip")
				}
				searchRequest.Skip = *skip
				continue
			case "id":
				if len(values) == 1 {
					searchRequest.Query.IdEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.IdIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "coverImageUrl":
				if len(values) == 1 {
					searchRequest.Query.CoverImageUrlEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.CoverImageUrlIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "dueDate":
				if len(values) == 1 {
					searchRequest.Query.DueDateEq = utils.StringSliceToTimestampPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.DueDateIn = utils.StringSliceToTimestampSlicePtr(values)
					continue
				}
			case "isPublic":
				if len(values) == 1 {
					searchRequest.Query.IsPublicEq = utils.StringSliceToBoolPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.IsPublicIn = utils.StringSliceToBoolSlicePtr(values)
					continue
				}
			case "parentNames":
				if len(values) == 1 {
					searchRequest.Query.ParentNamesEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ParentNamesIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingCity":
				if len(values) == 1 {
					searchRequest.Query.ShippingCityEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingCityIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingCountry":
				if len(values) == 1 {
					searchRequest.Query.ShippingCountryEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingCountryIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingDeliveryNotes":
				if len(values) == 1 {
					searchRequest.Query.ShippingDeliveryNotesEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingDeliveryNotesIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingLine1":
				if len(values) == 1 {
					searchRequest.Query.ShippingLine1Eq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingLine1In = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingLine2":
				if len(values) == 1 {
					searchRequest.Query.ShippingLine2Eq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingLine2In = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingPolicyVersion":
				if len(values) == 1 {
					searchRequest.Query.ShippingPolicyVersionEq = utils.StringSliceToIntPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingPolicyVersionIn = utils.StringSliceToIntSlicePtr(values)
					continue
				}
			case "shippingPostalCode":
				if len(values) == 1 {
					searchRequest.Query.ShippingPostalCodeEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingPostalCodeIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingRecipientName":
				if len(values) == 1 {
					searchRequest.Query.ShippingRecipientNameEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingRecipientNameIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "shippingRegion":
				if len(values) == 1 {
					searchRequest.Query.ShippingRegionEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ShippingRegionIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "slug":
				if len(values) == 1 {
					searchRequest.Query.SlugEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.SlugIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "themeColor":
				if len(values) == 1 {
					searchRequest.Query.ThemeColorEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ThemeColorIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "title":
				if len(values) == 1 {
					searchRequest.Query.TitleEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.TitleIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "welcomeMessage":
				if len(values) == 1 {
					searchRequest.Query.WelcomeMessageEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.WelcomeMessageIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			}
		}
		return searchRequest, nil
	default:
		return SearchRequest{}, coded_error.NewMethodNotAllowedError(r.Method)
	}
}

// HTTPAggregateFieldSpec is the JSON representation of an aggregate field spec
type HTTPAggregateFieldSpec struct {
	Field  string `json:"field"`
	Method string `json:"method"`
	Alias  string `json:"alias,omitempty"`
}

type AggregateRequest struct {
	Query                             registry.HTTPWhereClause             `json:"query"`
	Fields                            []HTTPAggregateFieldSpec             `json:"fields"`
	GroupBy                           []string                             `json:"groupBy"`
	AddressAccessSessionsProjection   *address_access_session.Projection   `json:"addressAccessSessionsProjection,omitempty"`
	RegistryApprovedGuestsProjection  *registry_approved_guest.Projection  `json:"registryApprovedGuestsProjection,omitempty"`
	RegistryItemsProjection           *registry_item.Projection            `json:"registryItemsProjection,omitempty"`
	ReservationsProjection            *reservation.Projection              `json:"reservationsProjection,omitempty"`
	ShippingAddressRequestsProjection *shipping_address_request.Projection `json:"shippingAddressRequestsProjection,omitempty"`
	OwnerProjection                   *owner_user.Projection               `json:"ownerProjection,omitempty"`
}

// AggregateResultRowHTTP is the HTTP response type for a single aggregate result row
type AggregateResultRowHTTP struct {
	CoverImageUrl         any `json:"coverImageUrl,omitempty"`
	DueDate               any `json:"dueDate,omitempty"`
	IsPublic              any `json:"isPublic,omitempty"`
	OwnerId               any `json:"ownerId,omitempty"`
	ParentNames           any `json:"parentNames,omitempty"`
	ShippingCity          any `json:"shippingCity,omitempty"`
	ShippingCountry       any `json:"shippingCountry,omitempty"`
	ShippingDeliveryNotes any `json:"shippingDeliveryNotes,omitempty"`
	ShippingLine1         any `json:"shippingLine1,omitempty"`
	ShippingLine2         any `json:"shippingLine2,omitempty"`
	ShippingPolicyVersion any `json:"shippingPolicyVersion,omitempty"`
	ShippingPostalCode    any `json:"shippingPostalCode,omitempty"`
	ShippingRecipientName any `json:"shippingRecipientName,omitempty"`
	ShippingRegion        any `json:"shippingRegion,omitempty"`
	Slug                  any `json:"slug,omitempty"`
	ThemeColor            any `json:"themeColor,omitempty"`
	Title                 any `json:"title,omitempty"`
	WelcomeMessage        any `json:"welcomeMessage,omitempty"`
	// Ref field Owner
	Owner any `json:"owner,omitempty"`
	// Ref field AddressAccessSessions
	AddressAccessSessions any `json:"addressAccessSessions,omitempty"`
	// Ref field RegistryApprovedGuests
	RegistryApprovedGuests any `json:"registryApprovedGuests,omitempty"`
	// Ref field RegistryItems
	RegistryItems any `json:"registryItems,omitempty"`
	// Ref field Reservations
	Reservations any `json:"reservations,omitempty"`
	// Ref field ShippingAddressRequests
	ShippingAddressRequests any `json:"shippingAddressRequests,omitempty"`
	// Metadata
	GroupKeys     []string `json:"__groupKeys"`
	AggregateKeys []string `json:"__aggregateKeys"`
}

// AggregateResponseHTTP is the HTTP response type for aggregate results
type AggregateResponseHTTP struct {
	Data  []AggregateResultRowHTTP `json:"data"`
	Total int                      `json:"total"`
}

func GetAggregateHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			props.onError("Aggregate", errors.New("method not allowed"))
			log.Debug().Msg("method not allowed")
			http.Error(w, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Aggregate", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		var aggregateRequest AggregateRequest
		err = json.NewDecoder(r.Body).Decode(&aggregateRequest)
		if err != nil {
			props.onError("Aggregate", err)
			log.Debug().Err(err).Msg("failed to decode request body")
			http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
			return
		}

		searchQuery, err := aggregateRequest.Query.ToWhereClause()
		if err != nil {
			props.onError("Aggregate", err)
			log.Debug().Err(err).Msg("failed to resolve aggregate query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		// Convert HTTP request fields to typed API fields
		apiFields := make([]registry_api.AggregateFieldSpec, len(aggregateRequest.Fields))
		for i, f := range aggregateRequest.Fields {
			apiFields[i] = registry_api.AggregateFieldSpec{
				Field:  registry_api.AggregateField(f.Field),
				Method: registry_api.AggregateMethod(f.Method),
				Alias:  f.Alias,
			}
		}

		// Convert HTTP request group-by fields to typed API fields
		apiGroupBy := make([]registry_api.GroupByField, len(aggregateRequest.GroupBy))
		for i, g := range aggregateRequest.GroupBy {
			apiGroupBy[i] = registry_api.GroupByField(g)
		}

		aggregateResult, err := props.Api.Aggregate(ctx, actor, searchQuery, registry_api.AggregateOptions{
			Fields:                            apiFields,
			GroupBy:                           apiGroupBy,
			AddressAccessSessionsProjection:   aggregateRequest.AddressAccessSessionsProjection,
			RegistryApprovedGuestsProjection:  aggregateRequest.RegistryApprovedGuestsProjection,
			RegistryItemsProjection:           aggregateRequest.RegistryItemsProjection,
			ReservationsProjection:            aggregateRequest.ReservationsProjection,
			ShippingAddressRequestsProjection: aggregateRequest.ShippingAddressRequestsProjection,
			OwnerProjection:                   aggregateRequest.OwnerProjection,
		})
		if err != nil {
			props.onError("Aggregate", err)
			log.Debug().Err(err).Msg("failed to execute aggregate query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		// Convert API result to HTTP response
		httpRows := make([]AggregateResultRowHTTP, len(aggregateResult.Data))
		for i, row := range aggregateResult.Data {
			httpRow := AggregateResultRowHTTP{
				GroupKeys:     row.GroupKeys,
				AggregateKeys: row.AggregateKeys,
			}
			// Copy group-by fields
			httpRow.CoverImageUrl = row.CoverImageUrl
			httpRow.DueDate = row.DueDate
			httpRow.IsPublic = row.IsPublic
			httpRow.OwnerId = row.OwnerId
			httpRow.ParentNames = row.ParentNames
			httpRow.ShippingCity = row.ShippingCity
			httpRow.ShippingCountry = row.ShippingCountry
			httpRow.ShippingDeliveryNotes = row.ShippingDeliveryNotes
			httpRow.ShippingLine1 = row.ShippingLine1
			httpRow.ShippingLine2 = row.ShippingLine2
			httpRow.ShippingPolicyVersion = row.ShippingPolicyVersion
			httpRow.ShippingPostalCode = row.ShippingPostalCode
			httpRow.ShippingRecipientName = row.ShippingRecipientName
			httpRow.ShippingRegion = row.ShippingRegion
			httpRow.Slug = row.Slug
			httpRow.ThemeColor = row.ThemeColor
			httpRow.Title = row.Title
			httpRow.WelcomeMessage = row.WelcomeMessage
			// Convert ref fields to HTTP records
			if row.Owner != nil {
				if aggregateRequest.OwnerProjection != nil {
					httpRec, _ := row.Owner.ToHTTPRecord(*aggregateRequest.OwnerProjection)
					httpRow.Owner = httpRec
				}
			}
			if row.AddressAccessSessions != nil && aggregateRequest.AddressAccessSessionsProjection != nil {
				httpRecs := make([]any, len(row.AddressAccessSessions))
				for j, rec := range row.AddressAccessSessions {
					httpRec, _ := rec.ToHTTPRecord(*aggregateRequest.AddressAccessSessionsProjection)
					httpRecs[j] = httpRec
				}
				httpRow.AddressAccessSessions = httpRecs
			}
			if row.RegistryApprovedGuests != nil && aggregateRequest.RegistryApprovedGuestsProjection != nil {
				httpRecs := make([]any, len(row.RegistryApprovedGuests))
				for j, rec := range row.RegistryApprovedGuests {
					httpRec, _ := rec.ToHTTPRecord(*aggregateRequest.RegistryApprovedGuestsProjection)
					httpRecs[j] = httpRec
				}
				httpRow.RegistryApprovedGuests = httpRecs
			}
			if row.RegistryItems != nil && aggregateRequest.RegistryItemsProjection != nil {
				httpRecs := make([]any, len(row.RegistryItems))
				for j, rec := range row.RegistryItems {
					httpRec, _ := rec.ToHTTPRecord(*aggregateRequest.RegistryItemsProjection)
					httpRecs[j] = httpRec
				}
				httpRow.RegistryItems = httpRecs
			}
			if row.Reservations != nil && aggregateRequest.ReservationsProjection != nil {
				httpRecs := make([]any, len(row.Reservations))
				for j, rec := range row.Reservations {
					httpRec, _ := rec.ToHTTPRecord(*aggregateRequest.ReservationsProjection)
					httpRecs[j] = httpRec
				}
				httpRow.Reservations = httpRecs
			}
			if row.ShippingAddressRequests != nil && aggregateRequest.ShippingAddressRequestsProjection != nil {
				httpRecs := make([]any, len(row.ShippingAddressRequests))
				for j, rec := range row.ShippingAddressRequests {
					httpRec, _ := rec.ToHTTPRecord(*aggregateRequest.ShippingAddressRequestsProjection)
					httpRecs[j] = httpRec
				}
				httpRow.ShippingAddressRequests = httpRecs
			}
			httpRows[i] = httpRow
		}

		httpResponse := AggregateResponseHTTP{
			Data:  httpRows,
			Total: aggregateResult.Total,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(httpResponse)
		if err != nil {
			props.onError("Aggregate", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusInternalServerError)
			return
		}
	}, nil
}
