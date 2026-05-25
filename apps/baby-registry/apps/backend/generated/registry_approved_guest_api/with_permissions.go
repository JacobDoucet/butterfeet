package registry_approved_guest_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_access_level"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_status"
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

type OnBeforeSearchHook func(ctx context.Context, actor permissions.Actor, query registry_approved_guest.WhereClause, options QueryOptions) (registry_approved_guest.WhereClause, QueryOptions, error)
type OnSearchHook func(ctx context.Context, actor permissions.Actor, r QueryResult, p Projection, err error) error
type OnBeforeSelectByIdHook func(ctx context.Context, actor permissions.Actor, query registry_approved_guest.SelectByIdQuery, projection Projection) (registry_approved_guest.SelectByIdQuery, Projection, error)
type OnSelectByIdHook func(ctx context.Context, actor permissions.Actor, m Model, p Projection, err error) error
type OnBeforeSelectByRegistryEmailUniqueHook func(ctx context.Context, actor permissions.Actor, query registry_approved_guest.SelectByRegistryEmailUniqueQuery, projection Projection) (registry_approved_guest.SelectByRegistryEmailUniqueQuery, Projection, error)
type OnSelectByRegistryEmailUniqueHook func(ctx context.Context, actor permissions.Actor, m Model, p Projection, err error) error
type OnBeforeCreateHook func(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, registry_approved_guest.Projection, error)
type OnCreateHook func(ctx context.Context, actor permissions.Actor, m registry_approved_guest.Model, p registry_approved_guest.Projection, err error) error
type OnBeforeUpdateHook func(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, registry_approved_guest.Projection, error)
type OnUpdateHook func(ctx context.Context, actor permissions.Actor, m registry_approved_guest.Model, p registry_approved_guest.Projection, err error) error
type OnBeforeDeleteHook func(ctx context.Context, actor permissions.Actor, id string) (string, error)
type OnDeleteHook func(ctx context.Context, actor permissions.Actor, id string, err error) error

type Hooks struct {
	OnBeforeSearch                      OnBeforeSearchHook
	OnSearch                            OnSearchHook
	OnBeforeSelectById                  OnBeforeSelectByIdHook
	OnSelectById                        OnSelectByIdHook
	OnBeforeSelectByRegistryEmailUnique OnBeforeSelectByRegistryEmailUniqueHook
	OnSelectByRegistryEmailUnique       OnSelectByRegistryEmailUniqueHook
	OnBeforeCreate                      OnBeforeCreateHook
	OnCreate                            OnCreateHook
	OnBeforeUpdate                      OnBeforeUpdateHook
	OnUpdate                            OnUpdateHook
	OnBeforeDelete                      OnBeforeDeleteHook
	OnDelete                            OnDeleteHook
}

func (c *clientWithPermissions) Search(ctx context.Context, actor permissions.Actor, query registry_approved_guest.WhereClause, options QueryOptions) (QueryResult, Projection, error) {
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
	where, err := registry_approved_guest.ApplyActorReadPermissionsToWhereClause(actor, query)
	if err != nil {
		return QueryResult{}, Projection{}, err
	}
	whereAddressAccessSession, err := address_access_session.ApplyActorReadPermissionsToWhereClause(actor, address_access_session.WhereClause{})
	if err != nil {
		projection.AddressAccessSessions = nil
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
		RegistryApprovedGuest: where,
		AddressAccessSessions: whereAddressAccessSession,
		Owner:                 whereOwnerUser,
		Registry:              whereRegistry,
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

func (c *clientWithPermissions) SelectById(ctx context.Context, actor permissions.Actor, query registry_approved_guest.SelectByIdQuery, projection Projection) (Model, Projection, error) {
	for _, hook := range c.hooks {
		if hook.OnBeforeSelectById != nil {
			var err error
			query, projection, err = hook.OnBeforeSelectById(ctx, actor, query, projection)
			if err != nil {
				return Model{}, Projection{}, err
			}
		}
	}

	result, resultProjection, err := c.Search(ctx, actor, registry_approved_guest.WhereClause{
		IdEq: &query.Id,
	}, QueryOptions{
		Projection: &projection,
		Limit:      1,
	})
	if err != nil {
		return Model{}, Projection{}, err
	}
	if len(result.Data) == 0 {
		return Model{}, Projection{}, coded_error.NewNotFoundError("registryApprovedGuest")
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

func (c *clientWithPermissions) SelectByRegistryEmailUnique(ctx context.Context, actor permissions.Actor, query registry_approved_guest.SelectByRegistryEmailUniqueQuery, projection Projection) (Model, Projection, error) {
	for _, hook := range c.hooks {
		if hook.OnBeforeSelectByRegistryEmailUnique != nil {
			var err error
			query, projection, err = hook.OnBeforeSelectByRegistryEmailUnique(ctx, actor, query, projection)
			if err != nil {
				return Model{}, Projection{}, err
			}
		}
	}

	result, resultProjection, err := c.Search(ctx, actor, registry_approved_guest.WhereClause{
		RegistryIdEq: &query.RegistryId,
		EmailHashEq:  &query.EmailHash,
	}, QueryOptions{
		Projection: &projection,
		Limit:      1,
	})
	if err != nil {
		return Model{}, Projection{}, err
	}
	if len(result.Data) == 0 {
		return Model{}, Projection{}, coded_error.NewNotFoundError("registryApprovedGuest")
	}

	for _, hook := range c.hooks {
		if hook.OnSelectByRegistryEmailUnique != nil {
			err = hook.OnSelectByRegistryEmailUnique(ctx, actor, result.Data[0], resultProjection, err)
			if err != nil {
				return Model{}, Projection{}, err
			}
		}
	}

	return result.Data[0], resultProjection, nil
}

func (c *clientWithPermissions) CanWrite(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model) error {
	ok := registry_approved_guest.HasWritePermissions(obj, actor)
	if !ok {
		return coded_error.NewUnauthorizedError("no write permissions")
	}
	if obj.Id == "" {
		return nil
	}
	abacProjection := Projection{Projection: registry_approved_guest.GetAbacProjection(actor)}
	dbRecord, _, err := c.SelectById(ctx, actor, registry_approved_guest.SelectByIdQuery{Id: obj.Id}, abacProjection)
	if err != nil {
		return coded_error.NewNotFoundError("registryApprovedGuest")
	}
	ok = registry_approved_guest.HasWritePermissions(dbRecord.Model, actor)
	if !ok {
		return coded_error.NewUnauthorizedError("no write permissions")
	}

	return nil
}

func ValidateBeforeCreate(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, registry_approved_guest.Projection, error) {
	var errs []string
	if !projection.AccessLevel {
		errs = append(errs, "missing required field accessLevel")
	} else if valErr := enum_guest_access_level.Validate(obj.AccessLevel); valErr != nil {
		errs = append(errs, "AccessLevel must be a valid GuestAccessLevel")
	}
	if !projection.EmailEnc {
		errs = append(errs, "missing required field emailEnc")
	} else if obj.EmailEnc == "" {
		errs = append(errs, "EmailEnc cannot be empty")
	}
	if !projection.EmailHash {
		errs = append(errs, "missing required field emailHash")
	} else if obj.EmailHash == "" {
		errs = append(errs, "EmailHash cannot be empty")
	}
	if !projection.OwnerId {
		errs = append(errs, "missing required field ownerId")
	}
	if !projection.RegistryId {
		errs = append(errs, "missing required field registryId")
	}
	if !projection.Status {
		errs = append(errs, "missing required field status")
	} else if valErr := enum_guest_status.Validate(obj.Status); valErr != nil {
		errs = append(errs, "Status must be a valid GuestStatus")
	}
	if len(errs) > 0 {
		return registry_approved_guest.Model{},
			registry_approved_guest.Projection{},
			coded_error.NewInvalidRequestError("failed creation validation: " + strings.Join(errs, ", "))
	}
	return obj, projection, nil
}

func (c *clientWithPermissions) Create(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, registry_approved_guest.Projection, error) {
	var err error
	obj, projection, err = ValidateBeforeCreate(ctx, actor, obj, projection)
	if err != nil {
		return registry_approved_guest.Model{}, registry_approved_guest.Projection{}, err
	}

	for _, hook := range c.hooks {
		if hook.OnBeforeCreate != nil {
			var err error
			obj, projection, err = hook.OnBeforeCreate(ctx, actor, obj, projection)
			if err != nil {
				return registry_approved_guest.Model{}, registry_approved_guest.Projection{}, err
			}
		}
	}

	err = c.CanWrite(ctx, actor, obj)
	if err != nil {
		return registry_approved_guest.Model{}, registry_approved_guest.Projection{}, err
	}

	projection = registry_approved_guest.ProjectWritePermissions(projection, actor)

	obj.Created = permissions.Trace(actor)
	projection.Updated = false
	projection.Created = true
	projection.CreatedFields = actor_trace.NewProjection(true)

	result, err := c.client.Create(ctx, obj, projection)

	for _, hook := range c.hooks {
		if hook.OnCreate != nil {
			err = hook.OnCreate(ctx, actor, result, projection, err)
			if err != nil {
				return registry_approved_guest.Model{}, registry_approved_guest.Projection{}, err
			}
		}
	}

	return result, projection, err
}

func (c *clientWithPermissions) Update(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, registry_approved_guest.Projection, error) {
	for _, hook := range c.hooks {
		if hook.OnBeforeUpdate != nil {
			var err error
			obj, projection, err = hook.OnBeforeUpdate(ctx, actor, obj, projection)
			if err != nil {
				return registry_approved_guest.Model{}, registry_approved_guest.Projection{}, err
			}
		}
	}

	where := registry_approved_guest.WhereClause{}
	if obj.Id != "" {
		where.IdEq = &obj.Id
	}

	var err error
	where, err = registry_approved_guest.ApplyActorWritePermissionsToWhereClause(actor, where)
	if err != nil {
		return registry_approved_guest.Model{}, registry_approved_guest.Projection{}, err
	}

	projection = registry_approved_guest.ProjectWritePermissions(projection, actor)

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
				return registry_approved_guest.Model{}, registry_approved_guest.Projection{}, err
			}
		}
	}

	return result, projection, err
}

func (c *clientWithPermissions) CanDelete(ctx context.Context, actor permissions.Actor, id string) error {
	if id == "" {
		return coded_error.NewInvalidRequestError("no id")
	}
	abacProjection := Projection{Projection: registry_approved_guest.GetAbacProjection(actor)}
	// TODO this should be c.client.Search to avoid permission read checks
	dbRecord, _, err := c.SelectById(ctx, actor, registry_approved_guest.SelectByIdQuery{Id: id}, abacProjection)
	if err != nil {
		return coded_error.NewNotFoundError("registryApprovedGuest")
	}
	ok := registry_approved_guest.HasWritePermissions(dbRecord.Model, actor)
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

func (c *clientWithPermissions) PaginateAll(ctx context.Context, actor permissions.Actor, query registry_approved_guest.WhereClause, options PaginationOptions) (<-chan Model, <-chan error) {
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
