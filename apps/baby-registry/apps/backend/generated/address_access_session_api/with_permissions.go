package address_access_session_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"strings"
)

const (
	ErrorCodeNotFound = "NOT_FOUND"
)

type clientWithPermissions struct {
	client clientImpl
	hooks  []Hooks
}

func New(client clientImpl, hooks ...Hooks) Client {
	return &clientWithPermissions{
		client: client,
		hooks:  hooks,
	}
}

type OnBeforeSearchHook func(ctx context.Context, actor permissions.Actor, query address_access_session.WhereClause, options QueryOptions) (address_access_session.WhereClause, QueryOptions, error)
type OnSearchHook func(ctx context.Context, actor permissions.Actor, r QueryResult, p Projection, err error) error
type OnBeforeSelectByIdHook func(ctx context.Context, actor permissions.Actor, query address_access_session.SelectByIdQuery, projection Projection) (address_access_session.SelectByIdQuery, Projection, error)
type OnSelectByIdHook func(ctx context.Context, actor permissions.Actor, m Model, p Projection, err error) error
type OnBeforeSelectByTokenUniqueHook func(ctx context.Context, actor permissions.Actor, query address_access_session.SelectByTokenUniqueQuery, projection Projection) (address_access_session.SelectByTokenUniqueQuery, Projection, error)
type OnSelectByTokenUniqueHook func(ctx context.Context, actor permissions.Actor, m Model, p Projection, err error) error
type OnBeforeCreateHook func(ctx context.Context, actor permissions.Actor, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, address_access_session.Projection, error)
type OnCreateHook func(ctx context.Context, actor permissions.Actor, m address_access_session.Model, p address_access_session.Projection, err error) error
type OnBeforeUpdateHook func(ctx context.Context, actor permissions.Actor, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, address_access_session.Projection, error)
type OnUpdateHook func(ctx context.Context, actor permissions.Actor, m address_access_session.Model, p address_access_session.Projection, err error) error
type OnBeforeDeleteHook func(ctx context.Context, actor permissions.Actor, id string) (string, error)
type OnDeleteHook func(ctx context.Context, actor permissions.Actor, id string, err error) error

type Hooks struct {
	OnBeforeSearch              OnBeforeSearchHook
	OnSearch                    OnSearchHook
	OnBeforeSelectById          OnBeforeSelectByIdHook
	OnSelectById                OnSelectByIdHook
	OnBeforeSelectByTokenUnique OnBeforeSelectByTokenUniqueHook
	OnSelectByTokenUnique       OnSelectByTokenUniqueHook
	OnBeforeCreate              OnBeforeCreateHook
	OnCreate                    OnCreateHook
	OnBeforeUpdate              OnBeforeUpdateHook
	OnUpdate                    OnUpdateHook
	OnBeforeDelete              OnBeforeDeleteHook
	OnDelete                    OnDeleteHook
}

func (c *clientWithPermissions) Search(ctx context.Context, actor permissions.Actor, query address_access_session.WhereClause, options QueryOptions) (QueryResult, Projection, error) {
	for _, hook := range c.hooks {
		if hook.OnBeforeSearch != nil {
			var err error
			query, options, err = hook.OnBeforeSearch(ctx, actor, query, options)
			if err != nil {
				return QueryResult{}, Projection{}, err
			}
		}
	}

	projection := projectReadPermissions(actor, options.GetProjection())
	where, err := address_access_session.ApplyActorReadPermissionsToWhereClause(actor, query)
	if err != nil {
		return QueryResult{}, Projection{}, err
	}
	whereRegistryApprovedGuest, err := registry_approved_guest.ApplyActorReadPermissionsToWhereClause(actor, registry_approved_guest.WhereClause{})
	if err != nil {
		projection.ApprovedGuest = nil
	}
	whereOwnerUser, err := owner_user.ApplyActorReadPermissionsToWhereClause(actor, owner_user.WhereClause{})
	if err != nil {
		projection.Owner = nil
	}
	whereRegistry, err := registry.ApplyActorReadPermissionsToWhereClause(actor, registry.WhereClause{})
	if err != nil {
		projection.Registry = nil
	}

	options.Projection = &projection
	result, err := c.client.Search(ctx, WhereClause{
		AddressAccessSession: where,
		ApprovedGuest:        whereRegistryApprovedGuest,
		Owner:                whereOwnerUser,
		Registry:             whereRegistry,
	}, options)

	for _, hook := range c.hooks {
		if hook.OnSearch != nil {
			err = hook.OnSearch(ctx, actor, result, projection, err)
			if err != nil {
				return QueryResult{}, Projection{}, err
			}
		}
	}

	return result, projection, err
}

func (c *clientWithPermissions) SelectById(ctx context.Context, actor permissions.Actor, query address_access_session.SelectByIdQuery, projection Projection) (Model, Projection, error) {
	for _, hook := range c.hooks {
		if hook.OnBeforeSelectById != nil {
			var err error
			query, projection, err = hook.OnBeforeSelectById(ctx, actor, query, projection)
			if err != nil {
				return Model{}, Projection{}, err
			}
		}
	}

	result, resultProjection, err := c.Search(ctx, actor, address_access_session.WhereClause{
		IdEq: &query.Id,
	}, QueryOptions{
		Projection: &projection,
		Limit:      1,
	})
	if err != nil {
		return Model{}, Projection{}, err
	}
	if len(result.Data) == 0 {
		return Model{}, Projection{}, coded_error.NewNotFoundError("addressAccessSession")
	}

	for _, hook := range c.hooks {
		if hook.OnSelectById != nil {
			err = hook.OnSelectById(ctx, actor, result.Data[0], resultProjection, err)
			if err != nil {
				return Model{}, Projection{}, err
			}
		}
	}

	return result.Data[0], resultProjection, nil
}

func (c *clientWithPermissions) SelectByTokenUnique(ctx context.Context, actor permissions.Actor, query address_access_session.SelectByTokenUniqueQuery, projection Projection) (Model, Projection, error) {
	for _, hook := range c.hooks {
		if hook.OnBeforeSelectByTokenUnique != nil {
			var err error
			query, projection, err = hook.OnBeforeSelectByTokenUnique(ctx, actor, query, projection)
			if err != nil {
				return Model{}, Projection{}, err
			}
		}
	}

	result, resultProjection, err := c.Search(ctx, actor, address_access_session.WhereClause{
		TokenHashEq: &query.TokenHash,
	}, QueryOptions{
		Projection: &projection,
		Limit:      1,
	})
	if err != nil {
		return Model{}, Projection{}, err
	}
	if len(result.Data) == 0 {
		return Model{}, Projection{}, coded_error.NewNotFoundError("addressAccessSession")
	}

	for _, hook := range c.hooks {
		if hook.OnSelectByTokenUnique != nil {
			err = hook.OnSelectByTokenUnique(ctx, actor, result.Data[0], resultProjection, err)
			if err != nil {
				return Model{}, Projection{}, err
			}
		}
	}

	return result.Data[0], resultProjection, nil
}

func (c *clientWithPermissions) CanWrite(ctx context.Context, actor permissions.Actor, obj address_access_session.Model) error {
	ok := address_access_session.HasWritePermissions(obj, actor)
	if !ok {
		return coded_error.NewUnauthorizedError("no write permissions")
	}
	if obj.Id == "" {
		return nil
	}
	abacProjection := Projection{Projection: address_access_session.GetAbacProjection(actor)}
	dbRecord, _, err := c.SelectById(ctx, actor, address_access_session.SelectByIdQuery{Id: obj.Id}, abacProjection)
	if err != nil {
		return coded_error.NewNotFoundError("addressAccessSession")
	}
	ok = address_access_session.HasWritePermissions(dbRecord.Model, actor)
	if !ok {
		return coded_error.NewUnauthorizedError("no write permissions")
	}

	return nil
}

func ValidateBeforeCreate(ctx context.Context, actor permissions.Actor, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, address_access_session.Projection, error) {
	var errs []string
	if !projection.EmailHash {
		errs = append(errs, "missing required field emailHash")
	} else if obj.EmailHash == "" {
		errs = append(errs, "EmailHash cannot be empty")
	}
	if !projection.ExpiresAt {
		errs = append(errs, "missing required field expiresAt")
	}
	if !projection.OwnerId {
		errs = append(errs, "missing required field ownerId")
	}
	if !projection.RegistryId {
		errs = append(errs, "missing required field registryId")
	}
	if !projection.TokenHash {
		errs = append(errs, "missing required field tokenHash")
	} else if obj.TokenHash == "" {
		errs = append(errs, "TokenHash cannot be empty")
	}
	if len(errs) > 0 {
		return address_access_session.Model{},
			address_access_session.Projection{},
			coded_error.NewInvalidRequestError("failed creation validation: " + strings.Join(errs, ", "))
	}
	return obj, projection, nil
}

func (c *clientWithPermissions) Create(ctx context.Context, actor permissions.Actor, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, address_access_session.Projection, error) {
	var err error
	obj, projection, err = ValidateBeforeCreate(ctx, actor, obj, projection)
	if err != nil {
		return address_access_session.Model{}, address_access_session.Projection{}, err
	}

	for _, hook := range c.hooks {
		if hook.OnBeforeCreate != nil {
			var err error
			obj, projection, err = hook.OnBeforeCreate(ctx, actor, obj, projection)
			if err != nil {
				return address_access_session.Model{}, address_access_session.Projection{}, err
			}
		}
	}

	err = c.CanWrite(ctx, actor, obj)
	if err != nil {
		return address_access_session.Model{}, address_access_session.Projection{}, err
	}

	projection = address_access_session.ProjectWritePermissions(projection, actor)

	obj.Created = permissions.Trace(actor)
	projection.Updated = false
	projection.Created = true
	projection.CreatedFields = actor_trace.NewProjection(true)

	result, err := c.client.Create(ctx, obj, projection)

	for _, hook := range c.hooks {
		if hook.OnCreate != nil {
			err = hook.OnCreate(ctx, actor, result, projection, err)
			if err != nil {
				return address_access_session.Model{}, address_access_session.Projection{}, err
			}
		}
	}

	return result, projection, err
}

func (c *clientWithPermissions) Update(ctx context.Context, actor permissions.Actor, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, address_access_session.Projection, error) {
	for _, hook := range c.hooks {
		if hook.OnBeforeUpdate != nil {
			var err error
			obj, projection, err = hook.OnBeforeUpdate(ctx, actor, obj, projection)
			if err != nil {
				return address_access_session.Model{}, address_access_session.Projection{}, err
			}
		}
	}

	where := address_access_session.WhereClause{}
	if obj.Id != "" {
		where.IdEq = &obj.Id
	}

	var err error
	where, err = address_access_session.ApplyActorWritePermissionsToWhereClause(actor, where)
	if err != nil {
		return address_access_session.Model{}, address_access_session.Projection{}, err
	}

	projection = address_access_session.ProjectWritePermissions(projection, actor)

	trace := permissions.Trace(actor)
	obj.Updated = trace
	projection.Updated = true
	projection.UpdatedFields = actor_trace.NewProjection(true)

	switch trace.ActorType {
	case string(permissions.ActorTypeOwnerUser):
		obj.UpdatedByOwnerUser = trace
		projection.UpdatedByOwnerUser = true
		projection.UpdatedByOwnerUserFields = actor_trace.NewProjection(true)
	}

	projection.Created = false

	result, err := c.client.Update(ctx, obj, where, projection)

	for _, hook := range c.hooks {
		if hook.OnUpdate != nil {
			err = hook.OnUpdate(ctx, actor, result, projection, err)
			if err != nil {
				return address_access_session.Model{}, address_access_session.Projection{}, err
			}
		}
	}

	return result, projection, err
}

func (c *clientWithPermissions) CanDelete(ctx context.Context, actor permissions.Actor, id string) error {
	if id == "" {
		return coded_error.NewInvalidRequestError("no id")
	}
	abacProjection := Projection{Projection: address_access_session.GetAbacProjection(actor)}
	// TODO this should be c.client.Search to avoid permission read checks
	dbRecord, _, err := c.SelectById(ctx, actor, address_access_session.SelectByIdQuery{Id: id}, abacProjection)
	if err != nil {
		return coded_error.NewNotFoundError("addressAccessSession")
	}
	ok := address_access_session.HasWritePermissions(dbRecord.Model, actor)
	if !ok {
		return coded_error.NewUnauthorizedError("no delete permissions")
	}
	return nil
}

func (c *clientWithPermissions) Delete(ctx context.Context, actor permissions.Actor, id string) error {
	for _, hook := range c.hooks {
		if hook.OnBeforeDelete != nil {
			var deleteHookErr error
			id, deleteHookErr = hook.OnBeforeDelete(ctx, actor, id)
			if deleteHookErr != nil {
				return deleteHookErr
			}
		}
	}

	err := c.CanDelete(ctx, actor, id)
	if err != nil {
		return err
	}

	for _, hook := range c.hooks {
		if hook.OnDelete != nil {
			deleteHookErr := hook.OnDelete(ctx, actor, id, err)
			if deleteHookErr != nil {
				return deleteHookErr
			}
		}
	}

	return c.client.Delete(ctx, id)
}

func (c *clientWithPermissions) PaginateAll(ctx context.Context, actor permissions.Actor, query address_access_session.WhereClause, options PaginationOptions) (<-chan Model, <-chan error) {
	modelCh := make(chan Model)
	errCh := make(chan error, 1)

	projection := options.GetProjection()

	go func() {
		defer close(modelCh)
		defer close(errCh)

		paginationOptions := QueryOptions{
			Projection: &projection,
			Sort:       options.Sort,
			Limit:      options.BatchSize,
			Skip:       0,
		}

		for {
			result, _, err := c.Search(ctx, actor, query, paginationOptions)
			if err != nil {
				errCh <- err
				return
			}

			for _, model := range result.Data {
				select {
				case modelCh <- model:
				case <-ctx.Done():
					return
				}
			}

			if len(result.Data) < paginationOptions.Limit {
				return
			}
			paginationOptions.Skip += len(result.Data)
		}
	}()

	return modelCh, errCh
}

func (c *clientWithPermissions) Aggregate(ctx context.Context, actor permissions.Actor, query address_access_session.WhereClause, options AggregateOptions) (AggregateResult, error) {
	// Apply read permissions to the where clause
	where, err := address_access_session.ApplyActorReadPermissionsToWhereClause(actor, query)
	if err != nil {
		return AggregateResult{}, err
	}

	result, err := c.client.Aggregate(ctx, WhereClause{
		AddressAccessSession: where,
	}, options)
	if err != nil {
		return AggregateResult{}, err
	}

	return result, nil
}
