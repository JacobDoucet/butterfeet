package registry_item_http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type GetMetadataOnSearchHook func(ctx context.Context, actor permissions.Actor, queryResult registry_item_api.HTTPQueryResult) (map[string]any, error)
type GetMetadataOnCreateHook func(ctx context.Context, actor permissions.Actor, mutationResult registry_item_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnUpdateHook func(ctx context.Context, actor permissions.Actor, mutationResult registry_item_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnDeleteHook func(ctx context.Context, actor permissions.Actor, deleteResult registry_item_api.HTTPDeleteResult) (map[string]any, error)

type MetadataHooks struct {
	OnSearch GetMetadataOnSearchHook
	OnCreate GetMetadataOnCreateHook
	OnUpdate GetMetadataOnUpdateHook
	OnDelete GetMetadataOnDeleteHook
}

type HandlerProps struct {
	Api           registry_item_api.Client
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
		p.OnError(endpoint+"<RegistryItem>", e)
	}
}

type SearchRequest struct {
	Query      registry_item.HTTPWhereClause `json:"query"`
	Sort       registry_item.HTTPSortParams  `json:"sort"`
	Projection *registry_item_api.Projection `json:"projection,omitempty"`
	Limit      int                           `json:"limit,omitempty"`
	Skip       int                           `json:"skip,omitempty"`
}

func (sr *SearchRequest) ResolveProjection() *registry_item_api.Projection {
	if sr.Projection != nil {
		return sr.Projection
	}
	projection := registry_item_api.NewProjection(true)
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

		queryResult, projectionResult, err := props.Api.Search(ctx, actor, searchQuery, registry_item_api.QueryOptions{
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

		response, err := registry_item_api.ToHTTPQueryResult(queryResult, projectionResult)
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
	Projection *registry_item_api.Projection `json:"projection,omitempty"`
}

func (sr *SelectRequest) ResolveProjection() registry_item_api.Projection {
	if sr.Projection != nil {
		return *sr.Projection
	}
	return registry_item_api.NewProjection(true)
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
	Query registry_item.HTTPSelectByIdQuery `json:"query"`
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
			Query: registry_item.HTTPSelectByIdQuery{
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

		response, err := registry_item_api.ToHTTPModel(getByIdResult, projectionResult)
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

type SaveRequest struct {
	Data registry_item.HTTPRecord `json:"data"`
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

		projection := registry_item.NewProjection(true)
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

		response, err := registry_item_api.ToHTTPMutationResult(createResult, projectionResult)
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

		response, err := registry_item_api.ToHTTPMutationResult(updateResult, projectionResult)
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

		response := registry_item_api.ToHTTPDeleteResult(id)

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
			Query: registry_item.HTTPWhereClause{},
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
			case "category":
				if len(values) == 1 {
					searchRequest.Query.CategoryEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.CategoryIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "currency":
				if len(values) == 1 {
					searchRequest.Query.CurrencyEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.CurrencyIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "description":
				if len(values) == 1 {
					searchRequest.Query.DescriptionEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.DescriptionIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "imageBgColor":
				if len(values) == 1 {
					searchRequest.Query.ImageBgColorEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ImageBgColorIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "imageUrl":
				if len(values) == 1 {
					searchRequest.Query.ImageUrlEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ImageUrlIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "noSubstitutes":
				if len(values) == 1 {
					searchRequest.Query.NoSubstitutesEq = utils.StringSliceToBoolPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.NoSubstitutesIn = utils.StringSliceToBoolSlicePtr(values)
					continue
				}
			case "notes":
				if len(values) == 1 {
					searchRequest.Query.NotesEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.NotesIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "ownerPurchased":
				if len(values) == 1 {
					searchRequest.Query.OwnerPurchasedEq = utils.StringSliceToBoolPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.OwnerPurchasedIn = utils.StringSliceToBoolSlicePtr(values)
					continue
				}
			case "position":
				if len(values) == 1 {
					searchRequest.Query.PositionEq = utils.StringSliceToIntPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.PositionIn = utils.StringSliceToIntSlicePtr(values)
					continue
				}
			case "priceCents":
				if len(values) == 1 {
					searchRequest.Query.PriceCentsEq = utils.StringSliceToIntPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.PriceCentsIn = utils.StringSliceToIntSlicePtr(values)
					continue
				}
			case "productUrl":
				if len(values) == 1 {
					searchRequest.Query.ProductUrlEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ProductUrlIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "quantity":
				if len(values) == 1 {
					searchRequest.Query.QuantityEq = utils.StringSliceToIntPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.QuantityIn = utils.StringSliceToIntSlicePtr(values)
					continue
				}
			case "quantityUnlimited":
				if len(values) == 1 {
					searchRequest.Query.QuantityUnlimitedEq = utils.StringSliceToBoolPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.QuantityUnlimitedIn = utils.StringSliceToBoolSlicePtr(values)
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
	Query                             registry_item.HTTPWhereClause        `json:"query"`
	Fields                            []HTTPAggregateFieldSpec             `json:"fields"`
	GroupBy                           []string                             `json:"groupBy"`
	ReservationsProjection            *reservation.Projection              `json:"reservationsProjection,omitempty"`
	ShippingAddressRequestsProjection *shipping_address_request.Projection `json:"shippingAddressRequestsProjection,omitempty"`
	RegistryProjection                *registry.Projection                 `json:"registryProjection,omitempty"`
}

// AggregateResultRowHTTP is the HTTP response type for a single aggregate result row
type AggregateResultRowHTTP struct {
	Category          any `json:"category,omitempty"`
	Currency          any `json:"currency,omitempty"`
	Description       any `json:"description,omitempty"`
	ImageBgColor      any `json:"imageBgColor,omitempty"`
	ImageUrl          any `json:"imageUrl,omitempty"`
	NoSubstitutes     any `json:"noSubstitutes,omitempty"`
	Notes             any `json:"notes,omitempty"`
	OwnerPurchased    any `json:"ownerPurchased,omitempty"`
	ParentItemId      any `json:"parentItemId,omitempty"`
	Position          any `json:"position,omitempty"`
	PriceCents        any `json:"priceCents,omitempty"`
	ProductUrl        any `json:"productUrl,omitempty"`
	Quantity          any `json:"quantity,omitempty"`
	QuantityUnlimited any `json:"quantityUnlimited,omitempty"`
	RegistryId        any `json:"registryId,omitempty"`
	Title             any `json:"title,omitempty"`
	// Ref field Registry
	Registry any `json:"registry,omitempty"`
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
		apiFields := make([]registry_item_api.AggregateFieldSpec, len(aggregateRequest.Fields))
		for i, f := range aggregateRequest.Fields {
			apiFields[i] = registry_item_api.AggregateFieldSpec{
				Field:  registry_item_api.AggregateField(f.Field),
				Method: registry_item_api.AggregateMethod(f.Method),
				Alias:  f.Alias,
			}
		}

		// Convert HTTP request group-by fields to typed API fields
		apiGroupBy := make([]registry_item_api.GroupByField, len(aggregateRequest.GroupBy))
		for i, g := range aggregateRequest.GroupBy {
			apiGroupBy[i] = registry_item_api.GroupByField(g)
		}

		aggregateResult, err := props.Api.Aggregate(ctx, actor, searchQuery, registry_item_api.AggregateOptions{
			Fields:                            apiFields,
			GroupBy:                           apiGroupBy,
			ReservationsProjection:            aggregateRequest.ReservationsProjection,
			ShippingAddressRequestsProjection: aggregateRequest.ShippingAddressRequestsProjection,
			RegistryProjection:                aggregateRequest.RegistryProjection,
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
			httpRow.Category = row.Category
			httpRow.Currency = row.Currency
			httpRow.Description = row.Description
			httpRow.ImageBgColor = row.ImageBgColor
			httpRow.ImageUrl = row.ImageUrl
			httpRow.NoSubstitutes = row.NoSubstitutes
			httpRow.Notes = row.Notes
			httpRow.OwnerPurchased = row.OwnerPurchased
			httpRow.ParentItemId = row.ParentItemId
			httpRow.Position = row.Position
			httpRow.PriceCents = row.PriceCents
			httpRow.ProductUrl = row.ProductUrl
			httpRow.Quantity = row.Quantity
			httpRow.QuantityUnlimited = row.QuantityUnlimited
			httpRow.RegistryId = row.RegistryId
			httpRow.Title = row.Title
			// Convert ref fields to HTTP records
			if row.Registry != nil {
				if aggregateRequest.RegistryProjection != nil {
					httpRec, _ := row.Registry.ToHTTPRecord(*aggregateRequest.RegistryProjection)
					httpRow.Registry = httpRec
				}
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
