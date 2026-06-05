package registry_payment_method_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query registry_payment_method.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query registry_payment_method.SelectByIdQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj registry_payment_method.Model, projection registry_payment_method.Projection) (registry_payment_method.Model, registry_payment_method.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj registry_payment_method.Model, projection registry_payment_method.Projection) (registry_payment_method.Model, registry_payment_method.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query registry_payment_method.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
	Aggregate(ctx context.Context, actor permissions.Actor, query registry_payment_method.WhereClause, options AggregateOptions) (AggregateResult, error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj registry_payment_method.Model, projection registry_payment_method.Projection) (registry_payment_method.Model, error)
	Update(ctx context.Context, obj registry_payment_method.Model, where registry_payment_method.WhereClause, projection registry_payment_method.Projection) (registry_payment_method.Model, error)
	Delete(ctx context.Context, id string) error
	Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error)
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	registry_payment_method.Model
	Carts    *[]cart.Model
	Owner    *owner_user.Model
	Registry *registry.Model
}

type WhereClause struct {
	RegistryPaymentMethod registry_payment_method.WhereClause
	Carts                 cart.WhereClause
	Owner                 owner_user.WhereClause
	Registry              registry.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       registry_payment_method.SortParams
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
	Sort       registry_payment_method.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	registry_payment_method.Projection `json:",inline"`
	Carts                              *cart.Projection       `json:"Carts,omitempty"`
	Owner                              *owner_user.Projection `json:"Owner,omitempty"`
	Registry                           *registry.Projection   `json:"Registry,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	cartsProjection := cart.NewProjection(defaultVal)
	ownerProjection := owner_user.NewProjection(defaultVal)
	registryProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection: registry_payment_method.NewProjection(defaultVal),
		Carts:      &cartsProjection,
		Owner:      &ownerProjection,
		Registry:   &registryProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = registry_payment_method.ProjectReadPermissions(projection.Projection, actor)
	if projection.Carts != nil {
		cartsProjection := cart.ProjectReadPermissions(*projection.Carts, actor)
		projection.Carts = &cartsProjection
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

func (m *Model) GetCarts() []cart.Model {
	if m.Carts == nil {
		return []cart.Model{}
	}
	return *m.Carts
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

// Valid aggregatable fields for RegistryPaymentMethod
const (
	AggregateFieldPosition AggregateField = "position"
)

// ValidAggregateFields returns all valid aggregatable fields
func ValidAggregateFields() []AggregateField {
	return []AggregateField{
		AggregateFieldPosition,
	}
}

// GroupByField represents a field that can be used for grouping
type GroupByField string

// Valid group-by fields for RegistryPaymentMethod
const (
	GroupByFieldBankAccountName   GroupByField = "bankAccountName"
	GroupByFieldBankAccountNumber GroupByField = "bankAccountNumber"
	GroupByFieldBankIban          GroupByField = "bankIban"
	GroupByFieldBankName          GroupByField = "bankName"
	GroupByFieldBankRoutingNumber GroupByField = "bankRoutingNumber"
	GroupByFieldBankSwift         GroupByField = "bankSwift"
	GroupByFieldDisplayName       GroupByField = "displayName"
	GroupByFieldEnabled           GroupByField = "enabled"
	GroupByFieldInstructions      GroupByField = "instructions"
	GroupByFieldOwnerId           GroupByField = "ownerId"
	GroupByFieldPaymentUrl        GroupByField = "paymentUrl"
	GroupByFieldPosition          GroupByField = "position"
	GroupByFieldRecipientEmail    GroupByField = "recipientEmail"
	GroupByFieldRecipientPhone    GroupByField = "recipientPhone"
	GroupByFieldRegistryId        GroupByField = "registryId"
)

// ValidGroupByFields returns all valid group-by fields
func ValidGroupByFields() []GroupByField {
	return []GroupByField{
		GroupByFieldBankAccountName,
		GroupByFieldBankAccountNumber,
		GroupByFieldBankIban,
		GroupByFieldBankName,
		GroupByFieldBankRoutingNumber,
		GroupByFieldBankSwift,
		GroupByFieldDisplayName,
		GroupByFieldEnabled,
		GroupByFieldInstructions,
		GroupByFieldOwnerId,
		GroupByFieldPaymentUrl,
		GroupByFieldPosition,
		GroupByFieldRecipientEmail,
		GroupByFieldRecipientPhone,
		GroupByFieldRegistryId,
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
	// Projection for Carts ref field
	CartsProjection *cart.Projection `json:"cartsProjection,omitempty"`
	// Projection for Owner ref field
	OwnerProjection *owner_user.Projection `json:"ownerProjection,omitempty"`
	// Projection for Registry ref field
	RegistryProjection *registry.Projection `json:"registryProjection,omitempty"`
}

// AggregateResultRow holds a single aggregation result row with a partial model structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	BankAccountName   *string `json:"bankAccountName,omitempty"`
	BankAccountNumber *string `json:"bankAccountNumber,omitempty"`
	BankIban          *string `json:"bankIban,omitempty"`
	BankName          *string `json:"bankName,omitempty"`
	BankRoutingNumber *string `json:"bankRoutingNumber,omitempty"`
	BankSwift         *string `json:"bankSwift,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	Enabled           *bool   `json:"enabled,omitempty"`
	Instructions      *string `json:"instructions,omitempty"`
	OwnerId           *string `json:"ownerId,omitempty"`
	PaymentUrl        *string `json:"paymentUrl,omitempty"`
	Position          *int    `json:"position,omitempty"`
	RecipientEmail    *string `json:"recipientEmail,omitempty"`
	RecipientPhone    *string `json:"recipientPhone,omitempty"`
	RegistryId        *string `json:"registryId,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field Owner
	Owner *owner_user.Model `json:"owner,omitempty"`
	// Ref field Registry
	Registry *registry.Model `json:"registry,omitempty"`
	// Ref field Carts
	Carts []cart.Model `json:"carts,omitempty"`
	// Metadata fields indicating which fields are populated
	GroupKeys     []string `json:"__groupKeys"`
	AggregateKeys []string `json:"__aggregateKeys"`
}

// AggregateResult holds the full aggregation query result
type AggregateResult struct {
	Data  []AggregateResultRow `json:"data"`
	Total int                  `json:"total"`
}
