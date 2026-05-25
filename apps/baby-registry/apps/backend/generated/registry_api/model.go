package registry_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query registry.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query registry.SelectByIdQuery, projection Projection) (Model, Projection, error)
	SelectBySlugUnique(ctx context.Context, actor permissions.Actor, query registry.SelectBySlugUniqueQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj registry.Model, projection registry.Projection) (registry.Model, registry.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj registry.Model, projection registry.Projection) (registry.Model, registry.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query registry.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj registry.Model, projection registry.Projection) (registry.Model, error)
	Update(ctx context.Context, obj registry.Model, where registry.WhereClause, projection registry.Projection) (registry.Model, error)
	Delete(ctx context.Context, id string) error
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	registry.Model
	RegistryItems *[]registry_item.Model
	Reservations  *[]reservation.Model
	Owner         *owner_user.Model
}

type WhereClause struct {
	Registry      registry.WhereClause
	RegistryItems registry_item.WhereClause
	Reservations  reservation.WhereClause
	Owner         owner_user.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       registry.SortParams
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
	Sort       registry.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	registry.Projection `json:",inline"`
	RegistryItems       *registry_item.Projection `json:"RegistryItems,omitempty"`
	Reservations        *reservation.Projection   `json:"Reservations,omitempty"`
	Owner               *owner_user.Projection    `json:"Owner,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	registryItemsProjection := registry_item.NewProjection(defaultVal)
	reservationsProjection := reservation.NewProjection(defaultVal)
	ownerProjection := owner_user.NewProjection(defaultVal)
	return Projection{
		Projection:    registry.NewProjection(defaultVal),
		RegistryItems: &registryItemsProjection,
		Reservations:  &reservationsProjection,
		Owner:         &ownerProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = registry.ProjectReadPermissions(projection.Projection, actor)
	if projection.RegistryItems != nil {
		registryItemsProjection := registry_item.ProjectReadPermissions(*projection.RegistryItems, actor)
		projection.RegistryItems = &registryItemsProjection
	}
	if projection.Reservations != nil {
		reservationsProjection := reservation.ProjectReadPermissions(*projection.Reservations, actor)
		projection.Reservations = &reservationsProjection
	}
	if projection.Owner != nil {
		ownerProjection := owner_user.ProjectReadPermissions(*projection.Owner, actor)
		projection.Owner = &ownerProjection
	}

	return projection
}

func (m *Model) GetRegistryItems() []registry_item.Model {
	if m.RegistryItems == nil {
		return []registry_item.Model{}
	}
	return *m.RegistryItems
}
func (m *Model) GetReservations() []reservation.Model {
	if m.Reservations == nil {
		return []reservation.Model{}
	}
	return *m.Reservations
}
func (m *Model) GetOwner() owner_user.Model {
	if m.Owner == nil {
		return owner_user.Model{}
	}
	return *m.Owner
}
