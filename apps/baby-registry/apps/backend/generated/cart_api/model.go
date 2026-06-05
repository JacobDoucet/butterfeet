package cart_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"time"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query cart.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query cart.SelectByIdQuery, projection Projection) (Model, Projection, error)
	SelectByReferenceUnique(ctx context.Context, actor permissions.Actor, query cart.SelectByReferenceUniqueQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj cart.Model, projection cart.Projection) (cart.Model, cart.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj cart.Model, projection cart.Projection) (cart.Model, cart.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query cart.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
	Aggregate(ctx context.Context, actor permissions.Actor, query cart.WhereClause, options AggregateOptions) (AggregateResult, error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj cart.Model, projection cart.Projection) (cart.Model, error)
	Update(ctx context.Context, obj cart.Model, where cart.WhereClause, projection cart.Projection) (cart.Model, error)
	Delete(ctx context.Context, id string) error
	Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error)
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	cart.Model
	Reservations  *[]reservation.Model
	Owner         *owner_user.Model
	PaymentMethod *registry_payment_method.Model
	Registry      *registry.Model
}

type WhereClause struct {
	Cart          cart.WhereClause
	Reservations  reservation.WhereClause
	Owner         owner_user.WhereClause
	PaymentMethod registry_payment_method.WhereClause
	Registry      registry.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       cart.SortParams
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
	Sort       cart.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	cart.Projection `json:",inline"`
	Reservations    *reservation.Projection             `json:"Reservations,omitempty"`
	Owner           *owner_user.Projection              `json:"Owner,omitempty"`
	PaymentMethod   *registry_payment_method.Projection `json:"PaymentMethod,omitempty"`
	Registry        *registry.Projection                `json:"Registry,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	reservationsProjection := reservation.NewProjection(defaultVal)
	ownerProjection := owner_user.NewProjection(defaultVal)
	paymentMethodProjection := registry_payment_method.NewProjection(defaultVal)
	registryProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection:    cart.NewProjection(defaultVal),
		Reservations:  &reservationsProjection,
		Owner:         &ownerProjection,
		PaymentMethod: &paymentMethodProjection,
		Registry:      &registryProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = cart.ProjectReadPermissions(projection.Projection, actor)
	if projection.Reservations != nil {
		reservationsProjection := reservation.ProjectReadPermissions(*projection.Reservations, actor)
		projection.Reservations = &reservationsProjection
	}
	if projection.Owner != nil {
		ownerProjection := owner_user.ProjectReadPermissions(*projection.Owner, actor)
		projection.Owner = &ownerProjection
	}
	if projection.PaymentMethod != nil {
		paymentMethodProjection := registry_payment_method.ProjectReadPermissions(*projection.PaymentMethod, actor)
		projection.PaymentMethod = &paymentMethodProjection
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
func (m *Model) GetOwner() owner_user.Model {
	if m.Owner == nil {
		return owner_user.Model{}
	}
	return *m.Owner
}
func (m *Model) GetPaymentMethod() registry_payment_method.Model {
	if m.PaymentMethod == nil {
		return registry_payment_method.Model{}
	}
	return *m.PaymentMethod
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

// Valid aggregatable fields for Cart
const (
	AggregateFieldAmountCents AggregateField = "amountCents"
)

// ValidAggregateFields returns all valid aggregatable fields
func ValidAggregateFields() []AggregateField {
	return []AggregateField{
		AggregateFieldAmountCents,
	}
}

// GroupByField represents a field that can be used for grouping
type GroupByField string

// Valid group-by fields for Cart
const (
	GroupByFieldAmountCents       GroupByField = "amountCents"
	GroupByFieldClaimedAt         GroupByField = "claimedAt"
	GroupByFieldContributorEmail  GroupByField = "contributorEmail"
	GroupByFieldContributorName   GroupByField = "contributorName"
	GroupByFieldCurrency          GroupByField = "currency"
	GroupByFieldDecidedAt         GroupByField = "decidedAt"
	GroupByFieldDecisionReason    GroupByField = "decisionReason"
	GroupByFieldMessage           GroupByField = "message"
	GroupByFieldMethodDisplayName GroupByField = "methodDisplayName"
	GroupByFieldOwnerId           GroupByField = "ownerId"
	GroupByFieldPaymentMethodId   GroupByField = "paymentMethodId"
	GroupByFieldReferenceCode     GroupByField = "referenceCode"
	GroupByFieldRegistryId        GroupByField = "registryId"
)

// ValidGroupByFields returns all valid group-by fields
func ValidGroupByFields() []GroupByField {
	return []GroupByField{
		GroupByFieldAmountCents,
		GroupByFieldClaimedAt,
		GroupByFieldContributorEmail,
		GroupByFieldContributorName,
		GroupByFieldCurrency,
		GroupByFieldDecidedAt,
		GroupByFieldDecisionReason,
		GroupByFieldMessage,
		GroupByFieldMethodDisplayName,
		GroupByFieldOwnerId,
		GroupByFieldPaymentMethodId,
		GroupByFieldReferenceCode,
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
	// Projection for Reservations ref field
	ReservationsProjection *reservation.Projection `json:"reservationsProjection,omitempty"`
	// Projection for Owner ref field
	OwnerProjection *owner_user.Projection `json:"ownerProjection,omitempty"`
	// Projection for PaymentMethod ref field
	PaymentMethodProjection *registry_payment_method.Projection `json:"paymentMethodProjection,omitempty"`
	// Projection for Registry ref field
	RegistryProjection *registry.Projection `json:"registryProjection,omitempty"`
}

// AggregateResultRow holds a single aggregation result row with a partial model structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	AmountCents       *int       `json:"amountCents,omitempty"`
	ClaimedAt         *time.Time `json:"claimedAt,omitempty"`
	ContributorEmail  *string    `json:"contributorEmail,omitempty"`
	ContributorName   *string    `json:"contributorName,omitempty"`
	Currency          *string    `json:"currency,omitempty"`
	DecidedAt         *time.Time `json:"decidedAt,omitempty"`
	DecisionReason    *string    `json:"decisionReason,omitempty"`
	Message           *string    `json:"message,omitempty"`
	MethodDisplayName *string    `json:"methodDisplayName,omitempty"`
	OwnerId           *string    `json:"ownerId,omitempty"`
	PaymentMethodId   *string    `json:"paymentMethodId,omitempty"`
	ReferenceCode     *string    `json:"referenceCode,omitempty"`
	RegistryId        *string    `json:"registryId,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field Owner
	Owner *owner_user.Model `json:"owner,omitempty"`
	// Ref field PaymentMethod
	PaymentMethod *registry_payment_method.Model `json:"paymentMethod,omitempty"`
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
