package owner_user_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query owner_user.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query owner_user.SelectByIdQuery, projection Projection) (Model, Projection, error)
	SelectByEmailUnique(ctx context.Context, actor permissions.Actor, query owner_user.SelectByEmailUniqueQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj owner_user.Model, projection owner_user.Projection) (owner_user.Model, owner_user.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj owner_user.Model, projection owner_user.Projection) (owner_user.Model, owner_user.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query owner_user.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj owner_user.Model, projection owner_user.Projection) (owner_user.Model, error)
	Update(ctx context.Context, obj owner_user.Model, where owner_user.WhereClause, projection owner_user.Projection) (owner_user.Model, error)
	Delete(ctx context.Context, id string) error
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	owner_user.Model
	Registrys *[]registry.Model
}

type WhereClause struct {
	OwnerUser owner_user.WhereClause
	Registrys registry.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       owner_user.SortParams
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
	Sort       owner_user.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	owner_user.Projection `json:",inline"`
	Registrys             *registry.Projection `json:"Registrys,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	registrysProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection: owner_user.NewProjection(defaultVal),
		Registrys:  &registrysProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = owner_user.ProjectReadPermissions(projection.Projection, actor)
	if projection.Registrys != nil {
		registrysProjection := registry.ProjectReadPermissions(*projection.Registrys, actor)
		projection.Registrys = &registrysProjection
	}

	return projection
}

func (m *Model) GetRegistrys() []registry.Model {
	if m.Registrys == nil {
		return []registry.Model{}
	}
	return *m.Registrys
}
