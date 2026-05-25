package registry_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
	"time"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query registry.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query registry.SelectByIdQuery, projection Projection) (Model, Projection, error)
	SelectBySlugUnique(ctx context.Context, actor permissions.Actor, query registry.SelectBySlugUniqueQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj registry.Model, projection registry.Projection) (registry.Model, registry.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj registry.Model, projection registry.Projection) (registry.Model, registry.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query registry.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
	Aggregate(ctx context.Context, actor permissions.Actor, query registry.WhereClause, options AggregateOptions) (AggregateResult, error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj registry.Model, projection registry.Projection) (registry.Model, error)
	Update(ctx context.Context, obj registry.Model, where registry.WhereClause, projection registry.Projection) (registry.Model, error)
	Delete(ctx context.Context, id string) error
	Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error)
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	registry.Model
	AddressAccessSessions   *[]address_access_session.Model
	RegistryApprovedGuests  *[]registry_approved_guest.Model
	RegistryItems           *[]registry_item.Model
	Reservations            *[]reservation.Model
	ShippingAddressRequests *[]shipping_address_request.Model
	Owner                   *owner_user.Model
}

type WhereClause struct {
	Registry                registry.WhereClause
	AddressAccessSessions   address_access_session.WhereClause
	RegistryApprovedGuests  registry_approved_guest.WhereClause
	RegistryItems           registry_item.WhereClause
	Reservations            reservation.WhereClause
	ShippingAddressRequests shipping_address_request.WhereClause
	Owner                   owner_user.WhereClause
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
	registry.Projection     `json:",inline"`
	AddressAccessSessions   *address_access_session.Projection   `json:"AddressAccessSessions,omitempty"`
	RegistryApprovedGuests  *registry_approved_guest.Projection  `json:"RegistryApprovedGuests,omitempty"`
	RegistryItems           *registry_item.Projection            `json:"RegistryItems,omitempty"`
	Reservations            *reservation.Projection              `json:"Reservations,omitempty"`
	ShippingAddressRequests *shipping_address_request.Projection `json:"ShippingAddressRequests,omitempty"`
	Owner                   *owner_user.Projection               `json:"Owner,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	addressAccessSessionsProjection := address_access_session.NewProjection(defaultVal)
	registryApprovedGuestsProjection := registry_approved_guest.NewProjection(defaultVal)
	registryItemsProjection := registry_item.NewProjection(defaultVal)
	reservationsProjection := reservation.NewProjection(defaultVal)
	shippingAddressRequestsProjection := shipping_address_request.NewProjection(defaultVal)
	ownerProjection := owner_user.NewProjection(defaultVal)
	return Projection{
		Projection:              registry.NewProjection(defaultVal),
		AddressAccessSessions:   &addressAccessSessionsProjection,
		RegistryApprovedGuests:  &registryApprovedGuestsProjection,
		RegistryItems:           &registryItemsProjection,
		Reservations:            &reservationsProjection,
		ShippingAddressRequests: &shippingAddressRequestsProjection,
		Owner:                   &ownerProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = registry.ProjectReadPermissions(projection.Projection, actor)
	if projection.AddressAccessSessions != nil {
		addressAccessSessionsProjection := address_access_session.ProjectReadPermissions(*projection.AddressAccessSessions, actor)
		projection.AddressAccessSessions = &addressAccessSessionsProjection
	}
	if projection.RegistryApprovedGuests != nil {
		registryApprovedGuestsProjection := registry_approved_guest.ProjectReadPermissions(*projection.RegistryApprovedGuests, actor)
		projection.RegistryApprovedGuests = &registryApprovedGuestsProjection
	}
	if projection.RegistryItems != nil {
		registryItemsProjection := registry_item.ProjectReadPermissions(*projection.RegistryItems, actor)
		projection.RegistryItems = &registryItemsProjection
	}
	if projection.Reservations != nil {
		reservationsProjection := reservation.ProjectReadPermissions(*projection.Reservations, actor)
		projection.Reservations = &reservationsProjection
	}
	if projection.ShippingAddressRequests != nil {
		shippingAddressRequestsProjection := shipping_address_request.ProjectReadPermissions(*projection.ShippingAddressRequests, actor)
		projection.ShippingAddressRequests = &shippingAddressRequestsProjection
	}
	if projection.Owner != nil {
		ownerProjection := owner_user.ProjectReadPermissions(*projection.Owner, actor)
		projection.Owner = &ownerProjection
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
func (m *Model) GetShippingAddressRequests() []shipping_address_request.Model {
	if m.ShippingAddressRequests == nil {
		return []shipping_address_request.Model{}
	}
	return *m.ShippingAddressRequests
}
func (m *Model) GetOwner() owner_user.Model {
	if m.Owner == nil {
		return owner_user.Model{}
	}
	return *m.Owner
}

// AggregateMethod represents the type of aggregation operation
type AggregateMethod string

const (
	AggregateSum   AggregateMethod = "sum"
	AggregateAvg   AggregateMethod = "avg"
	AggregateMin   AggregateMethod = "min"
	AggregateMax   AggregateMethod = "max"
	AggregateCount AggregateMethod = "count"
	AggregateFirst AggregateMethod = "first"
	AggregateLast  AggregateMethod = "last"
)

// AggregateField represents a field that can be aggregated
type AggregateField string

// Valid aggregatable fields for Registry
const (
	AggregateFieldShippingPolicyVersion AggregateField = "shippingPolicyVersion"
)

// ValidAggregateFields returns all valid aggregatable fields
func ValidAggregateFields() []AggregateField {
	return []AggregateField{
		AggregateFieldShippingPolicyVersion,
	}
}

// GroupByField represents a field that can be used for grouping
type GroupByField string

// Valid group-by fields for Registry
const (
	GroupByFieldCoverImageUrl         GroupByField = "coverImageUrl"
	GroupByFieldDueDate               GroupByField = "dueDate"
	GroupByFieldIsPublic              GroupByField = "isPublic"
	GroupByFieldOwnerId               GroupByField = "ownerId"
	GroupByFieldParentNames           GroupByField = "parentNames"
	GroupByFieldShippingCity          GroupByField = "shippingCity"
	GroupByFieldShippingCountry       GroupByField = "shippingCountry"
	GroupByFieldShippingDeliveryNotes GroupByField = "shippingDeliveryNotes"
	GroupByFieldShippingLine1         GroupByField = "shippingLine1"
	GroupByFieldShippingLine2         GroupByField = "shippingLine2"
	GroupByFieldShippingPolicyVersion GroupByField = "shippingPolicyVersion"
	GroupByFieldShippingPostalCode    GroupByField = "shippingPostalCode"
	GroupByFieldShippingRecipientName GroupByField = "shippingRecipientName"
	GroupByFieldShippingRegion        GroupByField = "shippingRegion"
	GroupByFieldSlug                  GroupByField = "slug"
	GroupByFieldThemeColor            GroupByField = "themeColor"
	GroupByFieldTitle                 GroupByField = "title"
	GroupByFieldWelcomeMessage        GroupByField = "welcomeMessage"
)

// ValidGroupByFields returns all valid group-by fields
func ValidGroupByFields() []GroupByField {
	return []GroupByField{
		GroupByFieldCoverImageUrl,
		GroupByFieldDueDate,
		GroupByFieldIsPublic,
		GroupByFieldOwnerId,
		GroupByFieldParentNames,
		GroupByFieldShippingCity,
		GroupByFieldShippingCountry,
		GroupByFieldShippingDeliveryNotes,
		GroupByFieldShippingLine1,
		GroupByFieldShippingLine2,
		GroupByFieldShippingPolicyVersion,
		GroupByFieldShippingPostalCode,
		GroupByFieldShippingRecipientName,
		GroupByFieldShippingRegion,
		GroupByFieldSlug,
		GroupByFieldThemeColor,
		GroupByFieldTitle,
		GroupByFieldWelcomeMessage,
	}
}

// AggregateFieldSpec specifies which field to aggregate and how
type AggregateFieldSpec struct {
	Field  AggregateField  `json:"field"`
	Method AggregateMethod `json:"method"`
	Alias  string          `json:"alias,omitempty"`
}

// Sum creates an aggregation spec for summing this field
func (f AggregateField) Sum() AggregateFieldSpec {
	return AggregateFieldSpec{Field: f, Method: AggregateSum}
}

// Avg creates an aggregation spec for averaging this field
func (f AggregateField) Avg() AggregateFieldSpec {
	return AggregateFieldSpec{Field: f, Method: AggregateAvg}
}

// Min creates an aggregation spec for finding the minimum of this field
func (f AggregateField) Min() AggregateFieldSpec {
	return AggregateFieldSpec{Field: f, Method: AggregateMin}
}

// Max creates an aggregation spec for finding the maximum of this field
func (f AggregateField) Max() AggregateFieldSpec {
	return AggregateFieldSpec{Field: f, Method: AggregateMax}
}

// Count creates an aggregation spec for counting records
func (f AggregateField) Count() AggregateFieldSpec {
	return AggregateFieldSpec{Field: f, Method: AggregateCount}
}

// First creates an aggregation spec for getting the first value
func (f AggregateField) First() AggregateFieldSpec {
	return AggregateFieldSpec{Field: f, Method: AggregateFirst}
}

// Last creates an aggregation spec for getting the last value
func (f AggregateField) Last() AggregateFieldSpec {
	return AggregateFieldSpec{Field: f, Method: AggregateLast}
}

// WithAlias sets a custom alias for the aggregation result
func (a AggregateFieldSpec) WithAlias(alias string) AggregateFieldSpec {
	a.Alias = alias
	return a
}

// AggregateOptions defines the aggregation query options
type AggregateOptions struct {
	// Fields to aggregate with their methods
	Fields []AggregateFieldSpec `json:"fields"`
	// Fields to group by
	GroupBy []GroupByField `json:"groupBy"`
	// Projection for AddressAccessSessions ref field
	AddressAccessSessionsProjection *address_access_session.Projection `json:"addressAccessSessionsProjection,omitempty"`
	// Projection for RegistryApprovedGuests ref field
	RegistryApprovedGuestsProjection *registry_approved_guest.Projection `json:"registryApprovedGuestsProjection,omitempty"`
	// Projection for RegistryItems ref field
	RegistryItemsProjection *registry_item.Projection `json:"registryItemsProjection,omitempty"`
	// Projection for Reservations ref field
	ReservationsProjection *reservation.Projection `json:"reservationsProjection,omitempty"`
	// Projection for ShippingAddressRequests ref field
	ShippingAddressRequestsProjection *shipping_address_request.Projection `json:"shippingAddressRequestsProjection,omitempty"`
	// Projection for Owner ref field
	OwnerProjection *owner_user.Projection `json:"ownerProjection,omitempty"`
}

// AggregateResultRow holds a single aggregation result row with a partial model structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	CoverImageUrl         *string    `json:"coverImageUrl,omitempty"`
	DueDate               *time.Time `json:"dueDate,omitempty"`
	IsPublic              *bool      `json:"isPublic,omitempty"`
	OwnerId               *string    `json:"ownerId,omitempty"`
	ParentNames           *string    `json:"parentNames,omitempty"`
	ShippingCity          *string    `json:"shippingCity,omitempty"`
	ShippingCountry       *string    `json:"shippingCountry,omitempty"`
	ShippingDeliveryNotes *string    `json:"shippingDeliveryNotes,omitempty"`
	ShippingLine1         *string    `json:"shippingLine1,omitempty"`
	ShippingLine2         *string    `json:"shippingLine2,omitempty"`
	ShippingPolicyVersion *int       `json:"shippingPolicyVersion,omitempty"`
	ShippingPostalCode    *string    `json:"shippingPostalCode,omitempty"`
	ShippingRecipientName *string    `json:"shippingRecipientName,omitempty"`
	ShippingRegion        *string    `json:"shippingRegion,omitempty"`
	Slug                  *string    `json:"slug,omitempty"`
	ThemeColor            *string    `json:"themeColor,omitempty"`
	Title                 *string    `json:"title,omitempty"`
	WelcomeMessage        *string    `json:"welcomeMessage,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field Owner
	Owner *owner_user.Model `json:"owner,omitempty"`
	// Ref field AddressAccessSessions
	AddressAccessSessions []address_access_session.Model `json:"addressAccessSessions,omitempty"`
	// Ref field RegistryApprovedGuests
	RegistryApprovedGuests []registry_approved_guest.Model `json:"registryApprovedGuests,omitempty"`
	// Ref field RegistryItems
	RegistryItems []registry_item.Model `json:"registryItems,omitempty"`
	// Ref field Reservations
	Reservations []reservation.Model `json:"reservations,omitempty"`
	// Ref field ShippingAddressRequests
	ShippingAddressRequests []shipping_address_request.Model `json:"shippingAddressRequests,omitempty"`
	// Metadata fields indicating which fields are populated
	GroupKeys     []string `json:"__groupKeys"`
	AggregateKeys []string `json:"__aggregateKeys"`
}

// AggregateResult holds the full aggregation query result
type AggregateResult struct {
	Data  []AggregateResultRow `json:"data"`
	Total int                  `json:"total"`
}
