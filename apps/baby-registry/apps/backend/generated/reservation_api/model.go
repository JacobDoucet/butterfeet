package reservation_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query reservation.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query reservation.SelectByIdQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj reservation.Model, projection reservation.Projection) (reservation.Model, reservation.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj reservation.Model, projection reservation.Projection) (reservation.Model, reservation.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query reservation.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
	Aggregate(ctx context.Context, actor permissions.Actor, query reservation.WhereClause, options AggregateOptions) (AggregateResult, error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj reservation.Model, projection reservation.Projection) (reservation.Model, error)
	Update(ctx context.Context, obj reservation.Model, where reservation.WhereClause, projection reservation.Projection) (reservation.Model, error)
	Delete(ctx context.Context, id string) error
	Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error)
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	reservation.Model
	Item     *registry_item.Model
	Registry *registry.Model
}

type WhereClause struct {
	Reservation reservation.WhereClause
	Item        registry_item.WhereClause
	Registry    registry.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       reservation.SortParams
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
	Sort       reservation.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	reservation.Projection `json:",inline"`
	Item                   *registry_item.Projection `json:"Item,omitempty"`
	Registry               *registry.Projection      `json:"Registry,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	itemProjection := registry_item.NewProjection(defaultVal)
	registryProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection: reservation.NewProjection(defaultVal),
		Item:       &itemProjection,
		Registry:   &registryProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = reservation.ProjectReadPermissions(projection.Projection, actor)
	if projection.Item != nil {
		itemProjection := registry_item.ProjectReadPermissions(*projection.Item, actor)
		projection.Item = &itemProjection
	}
	if projection.Registry != nil {
		registryProjection := registry.ProjectReadPermissions(*projection.Registry, actor)
		projection.Registry = &registryProjection
	}

	return projection
}

func (m *Model) GetItem() registry_item.Model {
	if m.Item == nil {
		return registry_item.Model{}
	}
	return *m.Item
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

// Valid aggregatable fields for Reservation
const (
	AggregateFieldQuantity AggregateField = "quantity"
)

// ValidAggregateFields returns all valid aggregatable fields
func ValidAggregateFields() []AggregateField {
	return []AggregateField{
		AggregateFieldQuantity,
	}
}

// GroupByField represents a field that can be used for grouping
type GroupByField string

// Valid group-by fields for Reservation
const (
	GroupByFieldContactEmail GroupByField = "contactEmail"
	GroupByFieldIsAnonymous  GroupByField = "isAnonymous"
	GroupByFieldItemId       GroupByField = "itemId"
	GroupByFieldMessage      GroupByField = "message"
	GroupByFieldQuantity     GroupByField = "quantity"
	GroupByFieldRegistryId   GroupByField = "registryId"
	GroupByFieldReserverName GroupByField = "reserverName"
)

// ValidGroupByFields returns all valid group-by fields
func ValidGroupByFields() []GroupByField {
	return []GroupByField{
		GroupByFieldContactEmail,
		GroupByFieldIsAnonymous,
		GroupByFieldItemId,
		GroupByFieldMessage,
		GroupByFieldQuantity,
		GroupByFieldRegistryId,
		GroupByFieldReserverName,
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
	// Projection for Item ref field
	ItemProjection *registry_item.Projection `json:"itemProjection,omitempty"`
	// Projection for Registry ref field
	RegistryProjection *registry.Projection `json:"registryProjection,omitempty"`
}

// AggregateResultRow holds a single aggregation result row with a partial model structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	ContactEmail *string `json:"contactEmail,omitempty"`
	IsAnonymous  *bool   `json:"isAnonymous,omitempty"`
	ItemId       *string `json:"itemId,omitempty"`
	Message      *string `json:"message,omitempty"`
	Quantity     *int    `json:"quantity,omitempty"`
	RegistryId   *string `json:"registryId,omitempty"`
	ReserverName *string `json:"reserverName,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field Item
	Item *registry_item.Model `json:"item,omitempty"`
	// Ref field Registry
	Registry *registry.Model `json:"registry,omitempty"`
	// Metadata fields indicating which fields are populated
	GroupKeys     []string `json:"__groupKeys"`
	AggregateKeys []string `json:"__aggregateKeys"`
}

// AggregateResult holds the full aggregation query result
type AggregateResult struct {
	Data  []AggregateResultRow `json:"data"`
	Total int                  `json:"total"`
}
