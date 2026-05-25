package address_access_session_api

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"time"
)

type Client interface {
	Search(ctx context.Context, actor permissions.Actor, query address_access_session.WhereClause, options QueryOptions) (QueryResult, Projection, error)
	SelectById(ctx context.Context, actor permissions.Actor, query address_access_session.SelectByIdQuery, projection Projection) (Model, Projection, error)
	SelectByTokenUnique(ctx context.Context, actor permissions.Actor, query address_access_session.SelectByTokenUniqueQuery, projection Projection) (Model, Projection, error)
	Create(ctx context.Context, actor permissions.Actor, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, address_access_session.Projection, error)
	Update(ctx context.Context, actor permissions.Actor, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, address_access_session.Projection, error)
	Delete(ctx context.Context, actor permissions.Actor, id string) error
	PaginateAll(ctx context.Context, actor permissions.Actor, query address_access_session.WhereClause, options PaginationOptions) (<-chan Model, <-chan error)
	Aggregate(ctx context.Context, actor permissions.Actor, query address_access_session.WhereClause, options AggregateOptions) (AggregateResult, error)
}

type clientImpl interface {
	Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error)
	Create(ctx context.Context, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, error)
	Update(ctx context.Context, obj address_access_session.Model, where address_access_session.WhereClause, projection address_access_session.Projection) (address_access_session.Model, error)
	Delete(ctx context.Context, id string) error
	Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error)
}

type QueryResult struct {
	Data  []Model
	Total int
	Skip  int
}

type Model struct {
	address_access_session.Model
	ApprovedGuest *registry_approved_guest.Model
	Owner         *owner_user.Model
	Registry      *registry.Model
}

type WhereClause struct {
	AddressAccessSession address_access_session.WhereClause
	ApprovedGuest        registry_approved_guest.WhereClause
	Owner                owner_user.WhereClause
	Registry             registry.WhereClause
}

type QueryOptions struct {
	Projection *Projection
	Sort       address_access_session.SortParams
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
	Sort       address_access_session.SortParams
	BatchSize  int
}

func (qo *PaginationOptions) GetProjection() Projection {
	if qo.Projection == nil {
		return NewProjection(true)
	}
	return *qo.Projection
}

type Projection struct {
	address_access_session.Projection `json:",inline"`
	ApprovedGuest                     *registry_approved_guest.Projection `json:"ApprovedGuest,omitempty"`
	Owner                             *owner_user.Projection              `json:"Owner,omitempty"`
	Registry                          *registry.Projection                `json:"Registry,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	approvedGuestProjection := registry_approved_guest.NewProjection(defaultVal)
	ownerProjection := owner_user.NewProjection(defaultVal)
	registryProjection := registry.NewProjection(defaultVal)
	return Projection{
		Projection:    address_access_session.NewProjection(defaultVal),
		ApprovedGuest: &approvedGuestProjection,
		Owner:         &ownerProjection,
		Registry:      &registryProjection,
	}
}

func projectReadPermissions(actor permissions.Actor, projection Projection) Projection {
	projection.Projection = address_access_session.ProjectReadPermissions(projection.Projection, actor)
	if projection.ApprovedGuest != nil {
		approvedGuestProjection := registry_approved_guest.ProjectReadPermissions(*projection.ApprovedGuest, actor)
		projection.ApprovedGuest = &approvedGuestProjection
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

func (m *Model) GetApprovedGuest() registry_approved_guest.Model {
	if m.ApprovedGuest == nil {
		return registry_approved_guest.Model{}
	}
	return *m.ApprovedGuest
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

// Valid aggregatable fields for AddressAccessSession
const (
	AggregateFieldPolicyVersionAtIssue AggregateField = "policyVersionAtIssue"
)

// ValidAggregateFields returns all valid aggregatable fields
func ValidAggregateFields() []AggregateField {
	return []AggregateField{
		AggregateFieldPolicyVersionAtIssue,
	}
}

// GroupByField represents a field that can be used for grouping
type GroupByField string

// Valid group-by fields for AddressAccessSession
const (
	GroupByFieldApprovedGuestId      GroupByField = "approvedGuestId"
	GroupByFieldEmailHash            GroupByField = "emailHash"
	GroupByFieldExpiresAt            GroupByField = "expiresAt"
	GroupByFieldOwnerId              GroupByField = "ownerId"
	GroupByFieldPolicyVersionAtIssue GroupByField = "policyVersionAtIssue"
	GroupByFieldRegistryId           GroupByField = "registryId"
	GroupByFieldTokenHash            GroupByField = "tokenHash"
)

// ValidGroupByFields returns all valid group-by fields
func ValidGroupByFields() []GroupByField {
	return []GroupByField{
		GroupByFieldApprovedGuestId,
		GroupByFieldEmailHash,
		GroupByFieldExpiresAt,
		GroupByFieldOwnerId,
		GroupByFieldPolicyVersionAtIssue,
		GroupByFieldRegistryId,
		GroupByFieldTokenHash,
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
	// Projection for ApprovedGuest ref field
	ApprovedGuestProjection *registry_approved_guest.Projection `json:"approvedGuestProjection,omitempty"`
	// Projection for Owner ref field
	OwnerProjection *owner_user.Projection `json:"ownerProjection,omitempty"`
	// Projection for Registry ref field
	RegistryProjection *registry.Projection `json:"registryProjection,omitempty"`
}

// AggregateResultRow holds a single aggregation result row with a partial model structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	ApprovedGuestId      *string    `json:"approvedGuestId,omitempty"`
	EmailHash            *string    `json:"emailHash,omitempty"`
	ExpiresAt            *time.Time `json:"expiresAt,omitempty"`
	OwnerId              *string    `json:"ownerId,omitempty"`
	PolicyVersionAtIssue *int       `json:"policyVersionAtIssue,omitempty"`
	RegistryId           *string    `json:"registryId,omitempty"`
	TokenHash            *string    `json:"tokenHash,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field ApprovedGuest
	ApprovedGuest *registry_approved_guest.Model `json:"approvedGuest,omitempty"`
	// Ref field Owner
	Owner *owner_user.Model `json:"owner,omitempty"`
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
