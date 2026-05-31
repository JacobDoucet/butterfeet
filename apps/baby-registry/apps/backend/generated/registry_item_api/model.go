package registry_item_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query registry_item.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query registry_item.SelectByIdQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj registry_item.Model, projection registry_item.Projection) (registry_item.Model, registry_item.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj registry_item.Model, projection registry_item.Projection) (registry_item.Model, registry_item.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query registry_item.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
	Aggregate(ctx context.Context, actor permissions.Actor, query registry_item.WhereClause, options AggregateOptions) (AggregateResult, error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj registry_item.Model, projection registry_item.Projection) (registry_item.Model, error)
	Update(ctx context.Context, obj registry_item.Model, where registry_item.WhereClause, projection registry_item.Projection) (registry_item.Model, error)
	Delete(ctx context.Context, id string) error
	Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error)
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	registry_item.Model
	Reservations            *[]reservation.Model
	ShippingAddressRequests *[]shipping_address_request.Model
	Registry                *registry.Model
}

type WhereClause struct {
	RegistryItem            registry_item.WhereClause
	Reservations            reservation.WhereClause
	ShippingAddressRequests shipping_address_request.WhereClause
	Registry                registry.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       registry_item.SortParams
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
	Sort       registry_item.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	registry_item.Projection `json:",inline"`
	Reservations             *reservation.Projection              `json:"Reservations,omitempty"`
	ShippingAddressRequests  *shipping_address_request.Projection `json:"ShippingAddressRequests,omitempty"`
	Registry                 *registry.Projection                 `json:"Registry,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	reservationsProjection := reservation.NewProjection(defaultVal)
	shippingAddressRequestsProjection := shipping_address_request.NewProjection(defaultVal)
	registryProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection:              registry_item.NewProjection(defaultVal),
		Reservations:            &reservationsProjection,
		ShippingAddressRequests: &shippingAddressRequestsProjection,
		Registry:                &registryProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = registry_item.ProjectReadPermissions(projection.Projection, actor)
	if projection.Reservations != nil {
		reservationsProjection := reservation.ProjectReadPermissions(*projection.Reservations, actor)
		projection.Reservations = &reservationsProjection
	}
	if projection.ShippingAddressRequests != nil {
		shippingAddressRequestsProjection := shipping_address_request.ProjectReadPermissions(*projection.ShippingAddressRequests, actor)
		projection.ShippingAddressRequests = &shippingAddressRequestsProjection
	}
	if projection.Registry != nil {
		registryProjection := registry.ProjectReadPermissions(*projection.Registry, actor)
		projection.Registry = &registryProjection
	}

	return projection
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
func (m *Model) GetRegistry() registry.Model {
	if m.Registry == nil {
		return registry.Model{}
	}
	return *m.Registry
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

// Valid aggregatable fields for RegistryItem
const (
	AggregateFieldPosition   AggregateField = "position"
	AggregateFieldPriceCents AggregateField = "priceCents"
	AggregateFieldQuantity   AggregateField = "quantity"
)

// ValidAggregateFields returns all valid aggregatable fields
func ValidAggregateFields() []AggregateField {
	return []AggregateField{
		AggregateFieldPosition,
		AggregateFieldPriceCents,
		AggregateFieldQuantity,
	}
}

// GroupByField represents a field that can be used for grouping
type GroupByField string

// Valid group-by fields for RegistryItem
const (
	GroupByFieldAffiliateUrl      GroupByField = "affiliateUrl"
	GroupByFieldCanonicalUrl      GroupByField = "canonicalUrl"
	GroupByFieldCategory          GroupByField = "category"
	GroupByFieldCurrency          GroupByField = "currency"
	GroupByFieldDescription       GroupByField = "description"
	GroupByFieldImageBgColor      GroupByField = "imageBgColor"
	GroupByFieldImageUrl          GroupByField = "imageUrl"
	GroupByFieldNoSubstitutes     GroupByField = "noSubstitutes"
	GroupByFieldNotes             GroupByField = "notes"
	GroupByFieldOriginalUrl       GroupByField = "originalUrl"
	GroupByFieldOwnerPurchased    GroupByField = "ownerPurchased"
	GroupByFieldParentItemId      GroupByField = "parentItemId"
	GroupByFieldPosition          GroupByField = "position"
	GroupByFieldPriceCents        GroupByField = "priceCents"
	GroupByFieldProductUrl        GroupByField = "productUrl"
	GroupByFieldQuantity          GroupByField = "quantity"
	GroupByFieldQuantityUnlimited GroupByField = "quantityUnlimited"
	GroupByFieldRegistryId        GroupByField = "registryId"
	GroupByFieldRetailer          GroupByField = "retailer"
	GroupByFieldTitle             GroupByField = "title"
)

// ValidGroupByFields returns all valid group-by fields
func ValidGroupByFields() []GroupByField {
	return []GroupByField{
		GroupByFieldAffiliateUrl,
		GroupByFieldCanonicalUrl,
		GroupByFieldCategory,
		GroupByFieldCurrency,
		GroupByFieldDescription,
		GroupByFieldImageBgColor,
		GroupByFieldImageUrl,
		GroupByFieldNoSubstitutes,
		GroupByFieldNotes,
		GroupByFieldOriginalUrl,
		GroupByFieldOwnerPurchased,
		GroupByFieldParentItemId,
		GroupByFieldPosition,
		GroupByFieldPriceCents,
		GroupByFieldProductUrl,
		GroupByFieldQuantity,
		GroupByFieldQuantityUnlimited,
		GroupByFieldRegistryId,
		GroupByFieldRetailer,
		GroupByFieldTitle,
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
	// Projection for Reservations ref field
	ReservationsProjection *reservation.Projection `json:"reservationsProjection,omitempty"`
	// Projection for ShippingAddressRequests ref field
	ShippingAddressRequestsProjection *shipping_address_request.Projection `json:"shippingAddressRequestsProjection,omitempty"`
	// Projection for Registry ref field
	RegistryProjection *registry.Projection `json:"registryProjection,omitempty"`
}

// AggregateResultRow holds a single aggregation result row with a partial model structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	AffiliateUrl      *string `json:"affiliateUrl,omitempty"`
	CanonicalUrl      *string `json:"canonicalUrl,omitempty"`
	Category          *string `json:"category,omitempty"`
	Currency          *string `json:"currency,omitempty"`
	Description       *string `json:"description,omitempty"`
	ImageBgColor      *string `json:"imageBgColor,omitempty"`
	ImageUrl          *string `json:"imageUrl,omitempty"`
	NoSubstitutes     *bool   `json:"noSubstitutes,omitempty"`
	Notes             *string `json:"notes,omitempty"`
	OriginalUrl       *string `json:"originalUrl,omitempty"`
	OwnerPurchased    *bool   `json:"ownerPurchased,omitempty"`
	ParentItemId      *string `json:"parentItemId,omitempty"`
	Position          *int    `json:"position,omitempty"`
	PriceCents        *int    `json:"priceCents,omitempty"`
	ProductUrl        *string `json:"productUrl,omitempty"`
	Quantity          *int    `json:"quantity,omitempty"`
	QuantityUnlimited *bool   `json:"quantityUnlimited,omitempty"`
	RegistryId        *string `json:"registryId,omitempty"`
	Retailer          *string `json:"retailer,omitempty"`
	Title             *string `json:"title,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field Registry
	Registry *registry.Model `json:"registry,omitempty"`
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
