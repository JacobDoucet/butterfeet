package owner_user_mongo

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LookupOptions struct {
	Projection                        owner_user.Projection
	Sort                              owner_user.MongoSortParams
	AddressAccessSessionsProjection   *address_access_session.Projection
	RegistryApprovedGuestsProjection  *registry_approved_guest.Projection
	RegistrysProjection               *registry.Projection
	ShippingAddressRequestsProjection *shipping_address_request.Projection
	Limit                             int
	Skip                              int
}

func (lo *LookupOptions) limit() int {
	if lo.Limit > 0 {
		return lo.Limit
	}
	return 100
}

type WhereClause struct {
	OwnerUser               owner_user.MongoWhereClause
	AddressAccessSessions   address_access_session.MongoWhereClause
	RegistryApprovedGuests  registry_approved_guest.MongoWhereClause
	Registrys               registry.MongoWhereClause
	ShippingAddressRequests shipping_address_request.MongoWhereClause
}

func aggregateWithRefs(ctx context.Context, where WhereClause, collection *mongo.Collection, lookupOptions LookupOptions) (QueryResult, error) {
	whereOwnerUser, err := where.OwnerUser.GetLookupQuery()
	if err != nil {
		return QueryResult{}, err
	}

	sortStage := bson.D{}

	if lookupOptions.Sort.CreatedAt > 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: 1})
	} else if lookupOptions.Sort.CreatedAt < 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: -1})
	}
	if lookupOptions.Sort.Email > 0 {
		sortStage = append(sortStage, bson.E{Key: "email", Value: 1})
	} else if lookupOptions.Sort.Email < 0 {
		sortStage = append(sortStage, bson.E{Key: "email", Value: -1})
	}
	if lookupOptions.Sort.UpdatedAt > 0 {
		sortStage = append(sortStage, bson.E{Key: "updated.at", Value: 1})
	} else if lookupOptions.Sort.UpdatedAt < 0 {
		sortStage = append(sortStage, bson.E{Key: "updated.at", Value: -1})
	}

	if lookupOptions.AddressAccessSessionsProjection != nil {
		lookupOptions.Projection.Id = true
	}
	if lookupOptions.RegistryApprovedGuestsProjection != nil {
		lookupOptions.Projection.Id = true
	}
	if lookupOptions.RegistrysProjection != nil {
		lookupOptions.Projection.Id = true
	}
	if lookupOptions.ShippingAddressRequestsProjection != nil {
		lookupOptions.Projection.Id = true
	}

	dataPipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: whereOwnerUser},
		},
		{
			{Key: "$project", Value: lookupOptions.Projection.ToBson()},
		},
	}

	if len(sortStage) > 0 {
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$sort", Value: sortStage},
		})
	}

	if lookupOptions.Skip > 0 {
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$skip", Value: lookupOptions.Skip},
		})
	}

	dataPipeline = append(dataPipeline, bson.D{
		{Key: "$limit", Value: lookupOptions.limit()},
	})

	// Add $lookup stage for AddressAccessSessions
	if lookupOptions.AddressAccessSessionsProjection != nil {
		// whereAddressAccessSessions, err := where.AddressAccessSessions.GetLookupQuery()
		// if err != nil {
		//     return QueryResult{}, err
		// }
		objectProject := bson.E{Key: "$project", Value: lookupOptions.AddressAccessSessionsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}
		// if len(whereAddressAccessSessions) > 0 {
		//     objectPipeline = bson.D{
		//         {Key: "$match", Value: whereAddressAccessSessions},
		//         objectProject,
		//     }
		// }
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "address_access_sessions"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "ownerUserId"},
				{Key: "as", Value: "AddressAccessSessions"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}
	// Add $lookup stage for RegistryApprovedGuests
	if lookupOptions.RegistryApprovedGuestsProjection != nil {
		// whereRegistryApprovedGuests, err := where.RegistryApprovedGuests.GetLookupQuery()
		// if err != nil {
		//     return QueryResult{}, err
		// }
		objectProject := bson.E{Key: "$project", Value: lookupOptions.RegistryApprovedGuestsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}
		// if len(whereRegistryApprovedGuests) > 0 {
		//     objectPipeline = bson.D{
		//         {Key: "$match", Value: whereRegistryApprovedGuests},
		//         objectProject,
		//     }
		// }
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "registry_approved_guests"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "ownerUserId"},
				{Key: "as", Value: "RegistryApprovedGuests"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}
	// Add $lookup stage for Registrys
	if lookupOptions.RegistrysProjection != nil {
		// whereRegistrys, err := where.Registrys.GetLookupQuery()
		// if err != nil {
		//     return QueryResult{}, err
		// }
		objectProject := bson.E{Key: "$project", Value: lookupOptions.RegistrysProjection.ToBson()}
		objectPipeline := bson.D{objectProject}
		// if len(whereRegistrys) > 0 {
		//     objectPipeline = bson.D{
		//         {Key: "$match", Value: whereRegistrys},
		//         objectProject,
		//     }
		// }
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "registries"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "ownerUserId"},
				{Key: "as", Value: "Registrys"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}
	// Add $lookup stage for ShippingAddressRequests
	if lookupOptions.ShippingAddressRequestsProjection != nil {
		// whereShippingAddressRequests, err := where.ShippingAddressRequests.GetLookupQuery()
		// if err != nil {
		//     return QueryResult{}, err
		// }
		objectProject := bson.E{Key: "$project", Value: lookupOptions.ShippingAddressRequestsProjection.ToBson()}
		objectPipeline := bson.D{objectProject}
		// if len(whereShippingAddressRequests) > 0 {
		//     objectPipeline = bson.D{
		//         {Key: "$match", Value: whereShippingAddressRequests},
		//         objectProject,
		//     }
		// }
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "shipping_address_requests"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "ownerUserId"},
				{Key: "as", Value: "ShippingAddressRequests"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})
	}

	// Add the $facet stage
	pipeline := mongo.Pipeline{
		{{
			Key: "$facet",
			Value: bson.D{
				{Key: "data", Value: dataPipeline},
				{Key: "count", Value: bson.A{
					bson.D{
						{
							Key:   "$match",
							Value: whereOwnerUser,
						},
					},
					bson.D{
						{
							Key:   "$count",
							Value: "total",
						},
					},
				}},
			},
		}},
	}

	// Execute the aggregation pipeline
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return QueryResult{}, err
	}

	// Decode the result
	var aggregatedResult []struct {
		Data  []Model `bson:"data"`
		Count []struct {
			Total int `bson:"total"`
		} `bson:"count"`
	}
	if err := cur.All(ctx, &aggregatedResult); err != nil {
		return QueryResult{}, errors.Join(cur.Close(ctx), err)
	}

	var totalCount int
	if len(aggregatedResult) > 0 && len(aggregatedResult[0].Count) > 0 {
		totalCount = aggregatedResult[0].Count[0].Total
	}

	return QueryResult{Data: aggregatedResult[0].Data, Count: totalCount}, cur.Close(ctx)
}
