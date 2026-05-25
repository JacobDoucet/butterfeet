package address_access_session_http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/utils"
	"github.com/rs/zerolog/log"
	"net/http"
)

type GetMetadataOnSearchHook func(ctx context.Context, actor permissions.Actor, queryResult address_access_session_api.HTTPQueryResult) (map[string]any, error)
type GetMetadataOnCreateHook func(ctx context.Context, actor permissions.Actor, mutationResult address_access_session_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnUpdateHook func(ctx context.Context, actor permissions.Actor, mutationResult address_access_session_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnDeleteHook func(ctx context.Context, actor permissions.Actor, deleteResult address_access_session_api.HTTPDeleteResult) (map[string]any, error)

type MetadataHooks struct {
	OnSearch GetMetadataOnSearchHook
	OnCreate GetMetadataOnCreateHook
	OnUpdate GetMetadataOnUpdateHook
	OnDelete GetMetadataOnDeleteHook
}

type HandlerProps struct {
	Api           address_access_session_api.Client
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
		p.OnError(endpoint+"<AddressAccessSession>", e)
	}
}

type SearchRequest struct {
	Query      address_access_session.HTTPWhereClause `json:"query"`
	Sort       address_access_session.HTTPSortParams  `json:"sort"`
	Projection *address_access_session_api.Projection `json:"projection,omitempty"`
	Limit      int                                    `json:"limit,omitempty"`
	Skip       int                                    `json:"skip,omitempty"`
}

func (sr *SearchRequest) ResolveProjection() *address_access_session_api.Projection {
	if sr.Projection != nil {
		return sr.Projection
	}
	projection := address_access_session_api.NewProjection(true)

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

		queryResult, projectionResult, err := props.Api.Search(ctx, actor, searchQuery, address_access_session_api.QueryOptions{
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

		response, err := address_access_session_api.ToHTTPQueryResult(queryResult, projectionResult)
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
	Projection *address_access_session_api.Projection `json:"projection,omitempty"`
}

func (sr *SelectRequest) ResolveProjection() address_access_session_api.Projection {
	if sr.Projection != nil {
		return *sr.Projection
	}
	return address_access_session_api.NewProjection(true)
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
	Query address_access_session.HTTPSelectByIdQuery `json:"query"`
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
			Query: address_access_session.HTTPSelectByIdQuery{
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

		response, err := address_access_session_api.ToHTTPModel(getByIdResult, projectionResult)
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

type SelectByTokenUniqueRequest struct {
	Query address_access_session.HTTPSelectByTokenUniqueQuery `json:"query"`
}

func GetSelectByTokenUniqueHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		selectRequest, err := resolveSelectRequest(r)
		if err != nil {
			props.onError("SelectByTokenUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select request")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("SelectByTokenUnique", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		tokenHash := r.PathValue("tokenHash")
		if tokenHash == "" {
			props.onError("SelectByTokenHash", errors.New("missing path parameter: tokenHash"))
			log.Debug().Err(err).Msg("missing path parameter")
			http.Error(w, "missing query parameter: tokenHash", http.StatusBadRequest)
			return
		}

		selectByTokenUniqueRequest := SelectByTokenUniqueRequest{
			Query: address_access_session.HTTPSelectByTokenUniqueQuery{
				TokenHash: tokenHash,
			},
		}

		selectByTokenUniqueQuery, err := selectByTokenUniqueRequest.Query.ToSelectByTokenUniqueQuery()
		if err != nil {
			props.onError("SelectByTokenUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		projection := selectRequest.ResolveProjection()

		getByTokenUniqueResult, projectionResult, err := props.Api.SelectByTokenUnique(ctx, actor, selectByTokenUniqueQuery, projection)
		if err != nil {
			props.onError("SelectByTokenUnique", err)
			log.Debug().Err(err).Msg("failed to execute select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := address_access_session_api.ToHTTPModel(getByTokenUniqueResult, projectionResult)
		if err != nil {
			props.onError("SelectByTokenUnique", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("SelectByTokenUnique", err)
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
			Query: address_access_session.HTTPWhereClause{},
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
			case "emailHash":
				if len(values) == 1 {
					searchRequest.Query.EmailHashEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.EmailHashIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "expiresAt":
				if len(values) == 1 {
					searchRequest.Query.ExpiresAtEq = utils.StringSliceToTimestampPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.ExpiresAtIn = utils.StringSliceToTimestampSlicePtr(values)
					continue
				}
			case "policyVersionAtIssue":
				if len(values) == 1 {
					searchRequest.Query.PolicyVersionAtIssueEq = utils.StringSliceToIntPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.PolicyVersionAtIssueIn = utils.StringSliceToIntSlicePtr(values)
					continue
				}
			case "tokenHash":
				if len(values) == 1 {
					searchRequest.Query.TokenHashEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.TokenHashIn = utils.StringSliceToStringSlicePtr(values)
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
	Query                   address_access_session.HTTPWhereClause `json:"query"`
	Fields                  []HTTPAggregateFieldSpec               `json:"fields"`
	GroupBy                 []string                               `json:"groupBy"`
	ApprovedGuestProjection *registry_approved_guest.Projection    `json:"approvedGuestProjection,omitempty"`
	OwnerProjection         *owner_user.Projection                 `json:"ownerProjection,omitempty"`
	RegistryProjection      *registry.Projection                   `json:"registryProjection,omitempty"`
}

// AggregateResultRowHTTP is the HTTP response type for a single aggregate result row
type AggregateResultRowHTTP struct {
	ApprovedGuestId      any `json:"approvedGuestId,omitempty"`
	EmailHash            any `json:"emailHash,omitempty"`
	ExpiresAt            any `json:"expiresAt,omitempty"`
	OwnerId              any `json:"ownerId,omitempty"`
	PolicyVersionAtIssue any `json:"policyVersionAtIssue,omitempty"`
	RegistryId           any `json:"registryId,omitempty"`
	TokenHash            any `json:"tokenHash,omitempty"`
	// Ref field ApprovedGuest
	ApprovedGuest any `json:"approvedGuest,omitempty"`
	// Ref field Owner
	Owner any `json:"owner,omitempty"`
	// Ref field Registry
	Registry any `json:"registry,omitempty"`
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
		apiFields := make([]address_access_session_api.AggregateFieldSpec, len(aggregateRequest.Fields))
		for i, f := range aggregateRequest.Fields {
			apiFields[i] = address_access_session_api.AggregateFieldSpec{
				Field:  address_access_session_api.AggregateField(f.Field),
				Method: address_access_session_api.AggregateMethod(f.Method),
				Alias:  f.Alias,
			}
		}

		// Convert HTTP request group-by fields to typed API fields
		apiGroupBy := make([]address_access_session_api.GroupByField, len(aggregateRequest.GroupBy))
		for i, g := range aggregateRequest.GroupBy {
			apiGroupBy[i] = address_access_session_api.GroupByField(g)
		}

		aggregateResult, err := props.Api.Aggregate(ctx, actor, searchQuery, address_access_session_api.AggregateOptions{
			Fields:                  apiFields,
			GroupBy:                 apiGroupBy,
			ApprovedGuestProjection: aggregateRequest.ApprovedGuestProjection,
			OwnerProjection:         aggregateRequest.OwnerProjection,
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
			httpRow.ApprovedGuestId = row.ApprovedGuestId
			httpRow.EmailHash = row.EmailHash
			httpRow.ExpiresAt = row.ExpiresAt
			httpRow.OwnerId = row.OwnerId
			httpRow.PolicyVersionAtIssue = row.PolicyVersionAtIssue
			httpRow.RegistryId = row.RegistryId
			httpRow.TokenHash = row.TokenHash
			// Convert ref fields to HTTP records
			if row.ApprovedGuest != nil {
				if aggregateRequest.ApprovedGuestProjection != nil {
					httpRec, _ := row.ApprovedGuest.ToHTTPRecord(*aggregateRequest.ApprovedGuestProjection)
					httpRow.ApprovedGuest = httpRec
				}
			}
			if row.Owner != nil {
				if aggregateRequest.OwnerProjection != nil {
					httpRec, _ := row.Owner.ToHTTPRecord(*aggregateRequest.OwnerProjection)
					httpRow.Owner = httpRec
				}
			}
			if row.Registry != nil {
				if aggregateRequest.RegistryProjection != nil {
					httpRec, _ := row.Registry.ToHTTPRecord(*aggregateRequest.RegistryProjection)
					httpRow.Registry = httpRec
				}
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
