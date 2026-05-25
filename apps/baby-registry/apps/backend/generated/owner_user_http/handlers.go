package owner_user_http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type GetMetadataOnSearchHook func(ctx context.Context, actor permissions.Actor, queryResult owner_user_api.HTTPQueryResult) (map[string]any, error)
type GetMetadataOnCreateHook func(ctx context.Context, actor permissions.Actor, mutationResult owner_user_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnUpdateHook func(ctx context.Context, actor permissions.Actor, mutationResult owner_user_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnDeleteHook func(ctx context.Context, actor permissions.Actor, deleteResult owner_user_api.HTTPDeleteResult) (map[string]any, error)

type MetadataHooks struct {
	OnSearch GetMetadataOnSearchHook
	OnCreate GetMetadataOnCreateHook
	OnUpdate GetMetadataOnUpdateHook
	OnDelete GetMetadataOnDeleteHook
}

type HandlerProps struct {
	Api           owner_user_api.Client
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
		p.OnError(endpoint+"<OwnerUser>", e)
	}
}

type SearchRequest struct {
	Query      owner_user.HTTPWhereClause `json:"query"`
	Sort       owner_user.HTTPSortParams  `json:"sort"`
	Projection *owner_user_api.Projection `json:"projection,omitempty"`
	Limit      int                        `json:"limit,omitempty"`
	Skip       int                        `json:"skip,omitempty"`
}

func (sr *SearchRequest) ResolveProjection() *owner_user_api.Projection {
	if sr.Projection != nil {
		return sr.Projection
	}
	projection := owner_user_api.NewProjection(true)
	projection.AddressAccessSessions = nil
	projection.RegistryApprovedGuests = nil
	projection.Registrys = nil
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

		queryResult, projectionResult, err := props.Api.Search(ctx, actor, searchQuery, owner_user_api.QueryOptions{
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

		response, err := owner_user_api.ToHTTPQueryResult(queryResult, projectionResult)
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
	Projection *owner_user_api.Projection `json:"projection,omitempty"`
}

func (sr *SelectRequest) ResolveProjection() owner_user_api.Projection {
	if sr.Projection != nil {
		return *sr.Projection
	}
	return owner_user_api.NewProjection(true)
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
	Query owner_user.HTTPSelectByIdQuery `json:"query"`
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
			Query: owner_user.HTTPSelectByIdQuery{
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

		response, err := owner_user_api.ToHTTPModel(getByIdResult, projectionResult)
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

type SelectByEmailUniqueRequest struct {
	Query owner_user.HTTPSelectByEmailUniqueQuery `json:"query"`
}

func GetSelectByEmailUniqueHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		selectRequest, err := resolveSelectRequest(r)
		if err != nil {
			props.onError("SelectByEmailUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select request")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("SelectByEmailUnique", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		email := r.PathValue("email")
		if email == "" {
			props.onError("SelectByEmail", errors.New("missing path parameter: email"))
			log.Debug().Err(err).Msg("missing path parameter")
			http.Error(w, "missing query parameter: email", http.StatusBadRequest)
			return
		}

		selectByEmailUniqueRequest := SelectByEmailUniqueRequest{
			Query: owner_user.HTTPSelectByEmailUniqueQuery{
				Email: email,
			},
		}

		selectByEmailUniqueQuery, err := selectByEmailUniqueRequest.Query.ToSelectByEmailUniqueQuery()
		if err != nil {
			props.onError("SelectByEmailUnique", err)
			log.Debug().Err(err).Msg("failed to resolve select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		projection := selectRequest.ResolveProjection()

		getByEmailUniqueResult, projectionResult, err := props.Api.SelectByEmailUnique(ctx, actor, selectByEmailUniqueQuery, projection)
		if err != nil {
			props.onError("SelectByEmailUnique", err)
			log.Debug().Err(err).Msg("failed to execute select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := owner_user_api.ToHTTPModel(getByEmailUniqueResult, projectionResult)
		if err != nil {
			props.onError("SelectByEmailUnique", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("SelectByEmailUnique", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

type SaveRequest struct {
	Data owner_user.HTTPRecord `json:"data"`
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

		response, err := owner_user_api.ToHTTPMutationResult(updateResult, projectionResult)
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
			Query: owner_user.HTTPWhereClause{},
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
			case "email":
				if len(values) == 1 {
					searchRequest.Query.EmailEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.EmailIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "name":
				if len(values) == 1 {
					searchRequest.Query.NameEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.NameIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			}
		}
		return searchRequest, nil
	default:
		return SearchRequest{}, coded_error.NewMethodNotAllowedError(r.Method)
	}
}
