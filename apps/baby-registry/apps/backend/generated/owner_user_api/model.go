package owner_user_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
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
	AddressAccessSessions   *[]address_access_session.Model
	RegistryApprovedGuests  *[]registry_approved_guest.Model
	Registrys               *[]registry.Model
	ShippingAddressRequests *[]shipping_address_request.Model
}

type WhereClause struct {
	OwnerUser               owner_user.WhereClause
	AddressAccessSessions   address_access_session.WhereClause
	RegistryApprovedGuests  registry_approved_guest.WhereClause
	Registrys               registry.WhereClause
	ShippingAddressRequests shipping_address_request.WhereClause
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
	owner_user.Projection   `json:",inline"`
	AddressAccessSessions   *address_access_session.Projection   `json:"AddressAccessSessions,omitempty"`
	RegistryApprovedGuests  *registry_approved_guest.Projection  `json:"RegistryApprovedGuests,omitempty"`
	Registrys               *registry.Projection                 `json:"Registrys,omitempty"`
	ShippingAddressRequests *shipping_address_request.Projection `json:"ShippingAddressRequests,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	addressAccessSessionsProjection := address_access_session.NewProjection(defaultVal)
	registryApprovedGuestsProjection := registry_approved_guest.NewProjection(defaultVal)
	registrysProjection := registry.NewProjection(defaultVal)
	shippingAddressRequestsProjection := shipping_address_request.NewProjection(defaultVal)
	return Projection{
		Projection:              owner_user.NewProjection(defaultVal),
		AddressAccessSessions:   &addressAccessSessionsProjection,
		RegistryApprovedGuests:  &registryApprovedGuestsProjection,
		Registrys:               &registrysProjection,
		ShippingAddressRequests: &shippingAddressRequestsProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = owner_user.ProjectReadPermissions(projection.Projection, actor)
	if projection.AddressAccessSessions != nil {
		addressAccessSessionsProjection := address_access_session.ProjectReadPermissions(*projection.AddressAccessSessions, actor)
		projection.AddressAccessSessions = &addressAccessSessionsProjection
	}
	if projection.RegistryApprovedGuests != nil {
		registryApprovedGuestsProjection := registry_approved_guest.ProjectReadPermissions(*projection.RegistryApprovedGuests, actor)
		projection.RegistryApprovedGuests = &registryApprovedGuestsProjection
	}
	if projection.Registrys != nil {
		registrysProjection := registry.ProjectReadPermissions(*projection.Registrys, actor)
		projection.Registrys = &registrysProjection
	}
	if projection.ShippingAddressRequests != nil {
		shippingAddressRequestsProjection := shipping_address_request.ProjectReadPermissions(*projection.ShippingAddressRequests, actor)
		projection.ShippingAddressRequests = &shippingAddressRequestsProjection
	}

	return projection
}

func (m *Model) GetAddressAccessSessions() []address_access_session.Model {
	if m.AddressAccessSessions == nil {
		return []address_access_session.Model{}
	}
	return *m.AddressAccessSessions
}
func (m *Model) GetRegistryApprovedGuests() []registry_approved_guest.Model {
	if m.RegistryApprovedGuests == nil {
		return []registry_approved_guest.Model{}
	}
	return *m.RegistryApprovedGuests
}
func (m *Model) GetRegistrys() []registry.Model {
	if m.Registrys == nil {
		return []registry.Model{}
	}
	return *m.Registrys
}
func (m *Model) GetShippingAddressRequests() []shipping_address_request.Model {
	if m.ShippingAddressRequests == nil {
		return []shipping_address_request.Model{}
	}
	return *m.ShippingAddressRequests
}
