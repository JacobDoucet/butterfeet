package cart_http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type GetMetadataOnSearchHook func(ctx context.Context, actor permissions.Actor, queryResult cart_api.HTTPQueryResult) (map[string]any, error)
type GetMetadataOnCreateHook func(ctx context.Context, actor permissions.Actor, mutationResult cart_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnUpdateHook func(ctx context.Context, actor permissions.Actor, mutationResult cart_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnDeleteHook func(ctx context.Context, actor permissions.Actor, deleteResult cart_api.HTTPDeleteResult) (map[string]any, error)

type MetadataHooks struct {
	OnSearch GetMetadataOnSearchHook
	OnCreate GetMetadataOnCreateHook
	OnUpdate GetMetadataOnUpdateHook
	OnDelete GetMetadataOnDeleteHook
}

type HandlerProps struct {
	Api           cart_api.Client
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
		p.OnError(endpoint+"<Cart>", e)
	}
}

type SearchRequest struct {
	Query      cart.HTTPWhereClause `json:"query"`
	Sort       cart.HTTPSortParams  `json:"sort"`
	Projection *cart_api.Projection `json:"projection,omitempty"`
	Limit      int                  `json:"limit,omitempty"`
	Skip       int                  `json:"skip,omitempty"`
}

func (sr *SearchRequest) ResolveProjection() *cart_api.Projection {
	if sr.Projection != nil {
		return sr.Projection
	}
	projection := cart_api.NewProjection(true)
	projection.Reservations = nil

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

		queryResult, projectionResult, err := props.Api.Search(ctx, actor, searchQuery, cart_api.QueryOptions{
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

		response, err := cart_api.ToHTTPQueryResult(queryResult, projectionResult)
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
	Projection *cart_api.Projection `json:"projection,omitempty"`
}

func (sr *SelectRequest) ResolveProjection() cart_api.Projection {
	if sr.Projection != nil {
		return *sr.Projection
	}
	return cart_api.NewProjection(true)
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
	Query cart.HTTPSelectByIdQuery `json:"query"`
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
			Query: cart.HTTPSelectByIdQuery{
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

		response, err := cart_api.ToHTTPModel(getByIdResult, projectionResult)
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

type SelectByReferenceUniqueRequest struct {
	Query cart.HTTPSelectByReferenceUniqueQuery `json:"query"`
}

func GetSelectByReferenceUniqueHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		selectRequest, err := resolveSelectRequest(r)
		if err != nil {
			props.onError("SelectByReferenceUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select request")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("SelectByReferenceUnique", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		referenceCode := r.PathValue("referenceCode")
		if referenceCode == "" {
			props.onError("SelectByReferenceCode", errors.New("missing path parameter: referenceCode"))
			log.Debug().Err(err).Msg("missing path parameter")
			http.Error(w, "missing query parameter: referenceCode", http.StatusBadRequest)
			return
		}

		selectByReferenceUniqueRequest := SelectByReferenceUniqueRequest{
			Query: cart.HTTPSelectByReferenceUniqueQuery{
				ReferenceCode: referenceCode,
			},
		}

		selectByReferenceUniqueQuery, err := selectByReferenceUniqueRequest.Query.ToSelectByReferenceUniqueQuery()
		if err != nil {
			props.onError("SelectByReferenceUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		projection := selectRequest.ResolveProjection()

		getByReferenceUniqueResult, projectionResult, err := props.Api.SelectByReferenceUnique(ctx, actor, selectByReferenceUniqueQuery, projection)
		if err != nil {
			props.onError("SelectByReferenceUnique", err)
			log.Debug().Err(err).Msg("failed to execute select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := cart_api.ToHTTPModel(getByReferenceUniqueResult, projectionResult)
		if err != nil {
			props.onError("SelectByReferenceUnique", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("SelectByReferenceUnique", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

type SaveRequest struct {
	Data cart.HTTPRecord `json:"data"`
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

		projection := cart.NewProjection(true)
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

		response, err := cart_api.ToHTTPMutationResult(createResult, projectionResult)
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

		response, err := cart_api.ToHTTPMutationResult(updateResult, projectionResult)
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

		response := cart_api.ToHTTPDeleteResult(id)

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
			Query: cart.HTTPWhereClause{},
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
			case "amountCents":
				if len(values) == 1 {
					searchRequest.Query.AmountCentsEq = utils.StringSliceToIntPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.AmountCentsIn = utils.StringSliceToIntSlicePtr(values)
					continue
				}
			case "claimedAt":
				if len(values) == 1 {
					searchRequest.Query.ClaimedAtEq = utils.StringSliceToTimestampPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ClaimedAtIn = utils.StringSliceToTimestampSlicePtr(values)
					continue
				}
			case "contributorEmail":
				if len(values) == 1 {
					searchRequest.Query.ContributorEmailEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ContributorEmailIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "contributorName":
				if len(values) == 1 {
					searchRequest.Query.ContributorNameEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ContributorNameIn = utils.StringSliceToStringSlicePtr(values)
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
			case "decidedAt":
				if len(values) == 1 {
					searchRequest.Query.DecidedAtEq = utils.StringSliceToTimestampPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.DecidedAtIn = utils.StringSliceToTimestampSlicePtr(values)
					continue
				}
			case "decisionReason":
				if len(values) == 1 {
					searchRequest.Query.DecisionReasonEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.DecisionReasonIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "message":
				if len(values) == 1 {
					searchRequest.Query.MessageEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.MessageIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "methodDisplayName":
				if len(values) == 1 {
					searchRequest.Query.MethodDisplayNameEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.MethodDisplayNameIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "referenceCode":
				if len(values) == 1 {
					searchRequest.Query.ReferenceCodeEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ReferenceCodeIn = utils.StringSliceToStringSlicePtr(values)
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
	Query                   cart.HTTPWhereClause                `json:"query"`
	Fields                  []HTTPAggregateFieldSpec            `json:"fields"`
	GroupBy                 []string                            `json:"groupBy"`
	ReservationsProjection  *reservation.Projection             `json:"reservationsProjection,omitempty"`
	OwnerProjection         *owner_user.Projection              `json:"ownerProjection,omitempty"`
	PaymentMethodProjection *registry_payment_method.Projection `json:"paymentMethodProjection,omitempty"`
	RegistryProjection      *registry.Projection                `json:"registryProjection,omitempty"`
}

// AggregateResultRowHTTP is the HTTP response type for a single aggregate result row
type AggregateResultRowHTTP struct {
	AmountCents       any `json:"amountCents,omitempty"`
	ClaimedAt         any `json:"claimedAt,omitempty"`
	ContributorEmail  any `json:"contributorEmail,omitempty"`
	ContributorName   any `json:"contributorName,omitempty"`
	Currency          any `json:"currency,omitempty"`
	DecidedAt         any `json:"decidedAt,omitempty"`
	DecisionReason    any `json:"decisionReason,omitempty"`
	Message           any `json:"message,omitempty"`
	MethodDisplayName any `json:"methodDisplayName,omitempty"`
	OwnerId           any `json:"ownerId,omitempty"`
	PaymentMethodId   any `json:"paymentMethodId,omitempty"`
	ReferenceCode     any `json:"referenceCode,omitempty"`
	RegistryId        any `json:"registryId,omitempty"`
	// Ref field Owner
	Owner any `json:"owner,omitempty"`
	// Ref field PaymentMethod
	PaymentMethod any `json:"paymentMethod,omitempty"`
	// Ref field Registry
	Registry any `json:"registry,omitempty"`
	// Ref field Reservations
	Reservations any `json:"reservations,omitempty"`
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
		apiFields := make([]cart_api.AggregateFieldSpec, len(aggregateRequest.Fields))
		for i, f := range aggregateRequest.Fields {
			apiFields[i] = cart_api.AggregateFieldSpec{
				Field:  cart_api.AggregateField(f.Field),
				Method: cart_api.AggregateMethod(f.Method),
				Alias:  f.Alias,
			}
		}

		// Convert HTTP request group-by fields to typed API fields
		apiGroupBy := make([]cart_api.GroupByField, len(aggregateRequest.GroupBy))
		for i, g := range aggregateRequest.GroupBy {
			apiGroupBy[i] = cart_api.GroupByField(g)
		}

		aggregateResult, err := props.Api.Aggregate(ctx, actor, searchQuery, cart_api.AggregateOptions{
			Fields:                  apiFields,
			GroupBy:                 apiGroupBy,
			ReservationsProjection:  aggregateRequest.ReservationsProjection,
			OwnerProjection:         aggregateRequest.OwnerProjection,
			PaymentMethodProjection: aggregateRequest.PaymentMethodProjection,
			RegistryProjection:      aggregateRequest.RegistryProjection,
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
			httpRow.AmountCents = row.AmountCents
			httpRow.ClaimedAt = row.ClaimedAt
			httpRow.ContributorEmail = row.ContributorEmail
			httpRow.ContributorName = row.ContributorName
			httpRow.Currency = row.Currency
			httpRow.DecidedAt = row.DecidedAt
			httpRow.DecisionReason = row.DecisionReason
			httpRow.Message = row.Message
			httpRow.MethodDisplayName = row.MethodDisplayName
			httpRow.OwnerId = row.OwnerId
			httpRow.PaymentMethodId = row.PaymentMethodId
			httpRow.ReferenceCode = row.ReferenceCode
			httpRow.RegistryId = row.RegistryId
			// Convert ref fields to HTTP records
			if row.Owner != nil {
				if aggregateRequest.OwnerProjection != nil {
					httpRec, _ := row.Owner.ToHTTPRecord(*aggregateRequest.OwnerProjection)
					httpRow.Owner = httpRec
				}
			}
			if row.PaymentMethod != nil {
				if aggregateRequest.PaymentMethodProjection != nil {
					httpRec, _ := row.PaymentMethod.ToHTTPRecord(*aggregateRequest.PaymentMethodProjection)
					httpRow.PaymentMethod = httpRec
				}
			}
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
