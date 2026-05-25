package registry_item_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
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
	Reservations *[]reservation.Model
	Registry     *registry.Model
}

type WhereClause struct {
	RegistryItem registry_item.WhereClause
	Reservations reservation.WhereClause
	Registry     registry.WhereClause
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
	Reservations             *reservation.Projection `json:"Reservations,omitempty"`
	Registry                 *registry.Projection    `json:"Registry,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	reservationsProjection := reservation.NewProjection(defaultVal)
	registryProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection:   registry_item.NewProjection(defaultVal),
		Reservations: &reservationsProjection,
		Registry:     &registryProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = registry_item.ProjectReadPermissions(projection.Projection, actor)
	if projection.Reservations != nil {
		reservationsProjection := reservation.ProjectReadPermissions(*projection.Reservations, actor)
		projection.Reservations = &reservationsProjection
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
	GroupByFieldCurrency    GroupByField = "currency"
	GroupByFieldDescription GroupByField = "description"
	GroupByFieldImageUrl    GroupByField = "imageUrl"
	GroupByFieldNotes       GroupByField = "notes"
	GroupByFieldPosition    GroupByField = "position"
	GroupByFieldPriceCents  GroupByField = "priceCents"
	GroupByFieldProductUrl  GroupByField = "productUrl"
	GroupByFieldQuantity    GroupByField = "quantity"
	GroupByFieldRegistryId  GroupByField = "registryId"
	GroupByFieldTitle       GroupByField = "title"
)

// ValidGroupByFields returns all valid group-by fields
func ValidGroupByFields() []GroupByField {
	return []GroupByField{
		GroupByFieldCurrency,
		GroupByFieldDescription,
		GroupByFieldImageUrl,
		GroupByFieldNotes,
		GroupByFieldPosition,
		GroupByFieldPriceCents,
		GroupByFieldProductUrl,
		GroupByFieldQuantity,
		GroupByFieldRegistryId,
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
	// Projection for Registry ref field
	RegistryProjection *registry.Projection `json:"registryProjection,omitempty"`
}

// AggregateResultRow holds a single aggregation result row with a partial model structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	Currency    *string `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	ImageUrl    *string `json:"imageUrl,omitempty"`
	Notes       *string `json:"notes,omitempty"`
	Position    *int    `json:"position,omitempty"`
	PriceCents  *int    `json:"priceCents,omitempty"`
	ProductUrl  *string `json:"productUrl,omitempty"`
	Quantity    *int    `json:"quantity,omitempty"`
	RegistryId  *string `json:"registryId,omitempty"`
	Title       *string `json:"title,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field Registry
	Registry *registry.Model `json:"registry,omitempty"`
	// Ref field Reservations
	Reservations []reservation.Model `json:"reservations,omitempty"`
	// Metadata fields indicating which fields are populated
	GroupKeys     []string `json:"__groupKeys"`
	AggregateKeys []string `json:"__aggregateKeys"`
}

// AggregateResult holds the full aggregation query result
type AggregateResult struct {
	Data  []AggregateResultRow `json:"data"`
	Total int                  `json:"total"`
}
