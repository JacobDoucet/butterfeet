package registry_approved_guest_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query registry_approved_guest.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query registry_approved_guest.SelectByIdQuery, projection Projection) (Model, Projection, error)
	SelectByRegistryEmailUnique(ctx context.Context, actor permissions.Actor, query registry_approved_guest.SelectByRegistryEmailUniqueQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, registry_approved_guest.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, registry_approved_guest.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query registry_approved_guest.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, error)
	Update(ctx context.Context, obj registry_approved_guest.Model, where registry_approved_guest.WhereClause, projection registry_approved_guest.Projection) (registry_approved_guest.Model, error)
	Delete(ctx context.Context, id string) error
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	registry_approved_guest.Model
	AddressAccessSessions *[]address_access_session.Model
	Owner                 *owner_user.Model
	Registry              *registry.Model
}

type WhereClause struct {
	RegistryApprovedGuest registry_approved_guest.WhereClause
	AddressAccessSessions address_access_session.WhereClause
	Owner                 owner_user.WhereClause
	Registry              registry.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       registry_approved_guest.SortParams
	Limit      int
	Skip       int
}

func (qo *QueryOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type PaginationOptions struct {
	Projection *Projection
	Sort       registry_approved_guest.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	registry_approved_guest.Projection `json:",inline"`
	AddressAccessSessions              *address_access_session.Projection `json:"AddressAccessSessions,omitempty"`
	Owner                              *owner_user.Projection             `json:"Owner,omitempty"`
	Registry                           *registry.Projection               `json:"Registry,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	addressAccessSessionsProjection := address_access_session.NewProjection(defaultVal)
	ownerProjection := owner_user.NewProjection(defaultVal)
	registryProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection:            registry_approved_guest.NewProjection(defaultVal),
		AddressAccessSessions: &addressAccessSessionsProjection,
		Owner:                 &ownerProjection,
		Registry:              &registryProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = registry_approved_guest.ProjectReadPermissions(projection.Projection, actor)
	if projection.AddressAccessSessions != nil {
		addressAccessSessionsProjection := address_access_session.ProjectReadPermissions(*projection.AddressAccessSessions, actor)
		projection.AddressAccessSessions = &addressAccessSessionsProjection
	}
	if projection.Owner != nil {
		ownerProjection := owner_user.ProjectReadPermissions(*projection.Owner, actor)
		projection.Owner = &ownerProjection
	}
	if projection.Registry != nil {
		registryProjection := registry.ProjectReadPermissions(*projection.Registry, actor)
		projection.Registry = &registryProjection
	}

	return projection
}

func (m *Model) GetAddressAccessSessions() []address_access_session.Model {
	if m.AddressAccessSessions == nil {
		return []address_access_session.Model{}
	}
	return *m.AddressAccessSessions
}
func (m *Model) GetOwner() owner_user.Model {
	if m.Owner == nil {
		return owner_user.Model{}
	}
	return *m.Owner
}
func (m *Model) GetRegistry() registry.Model {
	if m.Registry == nil {
		return registry.Model{}
	}
	return *m.Registry
}
