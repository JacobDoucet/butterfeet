package registry_mongo

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
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
	// Projection for AddressAccessSessions ref field
	AddressAccessSessionsProjection *address_access_session.Projection
	// Projection for RegistryApprovedGuests ref field
	RegistryApprovedGuestsProjection *registry_approved_guest.Projection
	// Projection for RegistryItems ref field
	RegistryItemsProjection *registry_item.Projection
	// Projection for Reservations ref field
	ReservationsProjection *reservation.Projection
	// Projection for ShippingAddressRequests ref field
	ShippingAddressRequestsProjection *shipping_address_request.Projection
	// Projection for Owner ref field
	OwnerProjection *owner_user.Projection
}

// AggregateResultRow holds a single aggregation result row with flat structure
type AggregateResultRow struct {
	// Group-by fields (original types)
	CoverImageUrl         *string             `bson:"coverImageUrl" json:"coverImageUrl,omitempty"`
	DueDate               *time.Time          `bson:"dueDate" json:"dueDate,omitempty"`
	IsPublic              *bool               `bson:"isPublic" json:"isPublic,omitempty"`
	OwnerId               *primitive.ObjectID `bson:"ownerId" json:"ownerId,omitempty"`
	ParentNames           *string             `bson:"parentNames" json:"parentNames,omitempty"`
	ShippingCity          *string             `bson:"shippingCity" json:"shippingCity,omitempty"`
	ShippingCountry       *string             `bson:"shippingCountry" json:"shippingCountry,omitempty"`
	ShippingDeliveryNotes *string             `bson:"shippingDeliveryNotes" json:"shippingDeliveryNotes,omitempty"`
	ShippingLine1         *string             `bson:"shippingLine1" json:"shippingLine1,omitempty"`
	ShippingLine2         *string             `bson:"shippingLine2" json:"shippingLine2,omitempty"`
	ShippingPolicyVersion *int                `bson:"shippingPolicyVersion" json:"shippingPolicyVersion,omitempty"`
	ShippingPostalCode    *string             `bson:"shippingPostalCode" json:"shippingPostalCode,omitempty"`
	ShippingRecipientName *string             `bson:"shippingRecipientName" json:"shippingRecipientName,omitempty"`
	ShippingRegion        *string             `bson:"shippingRegion" json:"shippingRegion,omitempty"`
	Slug                  *string             `bson:"slug" json:"slug,omitempty"`
	ThemeColor            *string             `bson:"themeColor" json:"themeColor,omitempty"`
	Title                 *string             `bson:"title" json:"title,omitempty"`
	WelcomeMessage        *string             `bson:"welcomeMessage" json:"welcomeMessage,omitempty"`
	// Aggregate fields - always float64 since they're results of sum/avg/etc
	// Ref field Owner
	Owner *owner_user.MongoRecord `bson:"owner,omitempty" json:"owner,omitempty"`
	// Ref field AddressAccessSessions
	AddressAccessSessions []address_access_session.MongoRecord `bson:"addressAccessSessions,omitempty" json:"addressAccessSessions,omitempty"`
	// Ref field RegistryApprovedGuests
	RegistryApprovedGuests []registry_approved_guest.MongoRecord `bson:"registryApprovedGuests,omitempty" json:"registryApprovedGuests,omitempty"`
	// Ref field RegistryItems
	RegistryItems []registry_item.MongoRecord `bson:"registryItems,omitempty" json:"registryItems,omitempty"`
	// Ref field Reservations
	Reservations []reservation.MongoRecord `bson:"reservations,omitempty" json:"reservations,omitempty"`
	// Ref field ShippingAddressRequests
	ShippingAddressRequests []shipping_address_request.MongoRecord `bson:"shippingAddressRequests,omitempty" json:"shippingAddressRequests,omitempty"`
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
	whereRegistry, err := where.Registry.GetLookupQuery()
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
	matchStage = append(matchStage, bson.E{Key: "$match", Value: whereRegistry})

	// Build group stage wrapper
	groupStageWrapper := bson.D{}
	groupStageWrapper = append(groupStageWrapper, bson.E{Key: "$group", Value: groupStage})

	// Build project stage wrapper
	projectStageWrapper := bson.D{}
	projectStageWrapper = append(projectStageWrapper, bson.E{Key: "$project", Value: projectStage})

	pipeline := mongo.Pipeline{matchStage, groupStageWrapper, projectStageWrapper}

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
	// Add $lookup stage for AddressAccessSessions if projection is specified
	if options.AddressAccessSessionsProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.AddressAccessSessionsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "address_access_sessions"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "registryId"},
				{Key: "as", Value: "addressAccessSessions"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}
	// Add $lookup stage for RegistryApprovedGuests if projection is specified
	if options.RegistryApprovedGuestsProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.RegistryApprovedGuestsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "registry_approved_guests"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "registryId"},
				{Key: "as", Value: "registryApprovedGuests"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}
	// Add $lookup stage for RegistryItems if projection is specified
	if options.RegistryItemsProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.RegistryItemsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "registry_items"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "registryId"},
				{Key: "as", Value: "registryItems"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}
	// Add $lookup stage for Reservations if projection is specified
	if options.ReservationsProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.ReservationsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "reservations"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "registryId"},
				{Key: "as", Value: "reservations"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}
	// Add $lookup stage for ShippingAddressRequests if projection is specified
	if options.ShippingAddressRequestsProjection != nil {
		objectProject := bson.E{Key: "$project", Value: options.ShippingAddressRequestsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}

		pipeline = append(pipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "shipping_address_requests"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "registryId"},
				{Key: "as", Value: "shippingAddressRequests"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
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
