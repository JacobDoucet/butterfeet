package address_access_session_mongo

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

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

func (m AggregateMethod) ToMongoOperator() string {
	switch m {
	case AggregateSum:
		return "$sum"
	case AggregateAvg:
		return "$avg"
	case AggregateMin:
		return "$min"
	case AggregateMax:
		return "$max"
	case AggregateCount:
		return "$sum" // count uses $sum: 1
	case AggregateFirst:
		return "$first"
	case AggregateLast:
		return "$last"
	default:
		return ""
	}
}

// AggregateFieldSpec specifies which field to aggregate and how
type AggregateFieldSpec struct {
	Field  string          `json:"field"`
	Method AggregateMethod `json:"method"`
	Alias  string          `json:"alias,omitempty"`
}

func (a AggregateFieldSpec) GetAlias() string {
	if a.Alias != "" {
		return a.Alias
	}
	return a.Field
}

// AggregateOptions defines the aggregation query options
type AggregateOptions struct {
	// Fields to aggregate with their methods
	Fields []AggregateFieldSpec
	// Fields to group by
	GroupBy []string
	// Projection for ApprovedGuest ref field
	ApprovedGuestProjection *registry_approved_guest.Projection
	// Projection for Owner ref field
	OwnerProjection *owner_user.Projection
	// Projection for Registry ref field
	RegistryProjection *registry.Projection
}

// AggregateResultRow holds a single aggregation result row with flat structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	ApprovedGuestId      *primitive.ObjectID `bson:"approvedGuestId" json:"approvedGuestId,omitempty"`
	EmailHash            *string             `bson:"emailHash" json:"emailHash,omitempty"`
	ExpiresAt            *time.Time          `bson:"expiresAt" json:"expiresAt,omitempty"`
	OwnerId              *primitive.ObjectID `bson:"ownerId" json:"ownerId,omitempty"`
	PolicyVersionAtIssue *int                `bson:"policyVersionAtIssue" json:"policyVersionAtIssue,omitempty"`
	RegistryId           *primitive.ObjectID `bson:"registryId" json:"registryId,omitempty"`
	TokenHash            *string             `bson:"tokenHash" json:"tokenHash,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field ApprovedGuest
	ApprovedGuest *registry_approved_guest.MongoRecord `bson:"approvedGuest,omitempty" json:"approvedGuest,omitempty"`
	// Ref field Owner
	Owner *owner_user.MongoRecord `bson:"owner,omitempty" json:"owner,omitempty"`
	// Ref field Registry
	Registry *registry.MongoRecord `bson:"registry,omitempty" json:"registry,omitempty"`
	// Metadata
	GroupKeys     []string `bson:"-" json:"__groupKeys"`
	AggregateKeys []string `bson:"-" json:"__aggregateKeys"`
}

// AggregateQueryResult holds the full aggregation query result
type AggregateQueryResult struct {
	Data  []AggregateResultRow
	Total int
}

func Aggregate(ctx context.Context, db *mongo.Database, where WhereClause, options AggregateOptions) (AggregateQueryResult, error) {
	collection := db.Collection(CollectionName)
	return executeAggregation(ctx, where, collection, options)
}

func executeAggregation(ctx context.Context, where WhereClause, collection *mongo.Collection, options AggregateOptions) (AggregateQueryResult, error) {
	whereAddressAccessSession, err := where.AddressAccessSession.GetLookupQuery()
	if err != nil {
		return AggregateQueryResult{}, err
	}

	// Build the $group stage
	groupStage := bson.D{}

	// Build _id for grouping
	if len(options.GroupBy) == 0 {
		// No grouping - aggregate all documents
		groupStage = append(groupStage, bson.E{Key: "_id", Value: nil})
	} else if len(options.GroupBy) == 1 {
		// Single field grouping
		groupStage = append(groupStage, bson.E{Key: "_id", Value: "$" + options.GroupBy[0]})
	} else {
		// Multiple field grouping
		groupId := bson.D{}
		for _, field := range options.GroupBy {
			groupId = append(groupId, bson.E{Key: field, Value: "$" + field})
		}
		groupStage = append(groupStage, bson.E{Key: "_id", Value: groupId})
	}

	// Add aggregation fields
	for _, aggField := range options.Fields {
		alias := aggField.GetAlias()
		if aggField.Method == AggregateCount {
			countVal := bson.D{}
			countVal = append(countVal, bson.E{Key: "$sum", Value: 1})
			groupStage = append(groupStage, bson.E{Key: alias, Value: countVal})
		} else {
			aggVal := bson.D{}
			aggVal = append(aggVal, bson.E{Key: aggField.Method.ToMongoOperator(), Value: "$" + aggField.Field})
			groupStage = append(groupStage, bson.E{Key: alias, Value: aggVal})
		}
	}

	// Build $project stage to flatten group-by fields
	projectStage := bson.D{}
	projectStage = append(projectStage, bson.E{Key: "_id", Value: 0})

	// Project group-by fields from _id
	for _, field := range options.GroupBy {
		if len(options.GroupBy) == 1 {
			projectStage = append(projectStage, bson.E{Key: field, Value: "$_id"})
		} else {
			projectStage = append(projectStage, bson.E{Key: field, Value: "$_id." + field})
		}
	}

	// Project aggregate fields
	for _, aggField := range options.Fields {
		alias := aggField.GetAlias()
		projectStage = append(projectStage, bson.E{Key: alias, Value: 1})
	}

	// Build match stage
	matchStage := bson.D{}
	matchStage = append(matchStage, bson.E{Key: "$match", Value: whereAddressAccessSession})

	// Build group stage wrapper
	groupStageWrapper := bson.D{}
	groupStageWrapper = append(groupStageWrapper, bson.E{Key: "$group", Value: groupStage})

	// Build project stage wrapper
	projectStageWrapper := bson.D{}
	projectStageWrapper = append(projectStageWrapper, bson.E{Key: "$project", Value: projectStage})

	pipeline := mongo.Pipeline{matchStage, groupStageWrapper, projectStageWrapper}

	// Add $lookup stage for ApprovedGuest if projection is specified
	if options.ApprovedGuestProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.ApprovedGuestProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "registry_approved_guests"},
				{Key: "localField", Value: "approvedGuestId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "approvedGuest"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})

		pipeline = append(pipeline, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$approvedGuest"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		})
	}
	// Add $lookup stage for Owner if projection is specified
	if options.OwnerProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.OwnerProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "owner_users"},
				{Key: "localField", Value: "ownerId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "owner"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})

		pipeline = append(pipeline, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$owner"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		})
	}
	// Add $lookup stage for Registry if projection is specified
	if options.RegistryProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.RegistryProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "registries"},
				{Key: "localField", Value: "registryId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "registry"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})

		pipeline = append(pipeline, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$registry"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		})
	}

	// Execute the aggregation
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return AggregateQueryResult{}, err
	}
	defer cur.Close(ctx)

	// Parse results
	var results []AggregateResultRow
	if err := cur.All(ctx, &results); err != nil {
		return AggregateQueryResult{}, err
	}

	// Build metadata keys
	groupKeys := make([]string, len(options.GroupBy))
	copy(groupKeys, options.GroupBy)

	aggregateKeys := make([]string, len(options.Fields))
	for i, f := range options.Fields {
		aggregateKeys[i] = f.GetAlias()
	}

	// Add metadata to each result
	for i := range results {
		results[i].GroupKeys = groupKeys
		results[i].AggregateKeys = aggregateKeys
	}

	return AggregateQueryResult{
		Data:  results,
		Total: len(results),
	}, nil
}
