package registry_approved_guest_mongo

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LookupOptions struct {
	Projection                      registry_approved_guest.Projection
	Sort                            registry_approved_guest.MongoSortParams
	AddressAccessSessionsProjection *address_access_session.Projection
	OwnerProjection                 *owner_user.Projection
	RegistryProjection              *registry.Projection
	Limit                           int
	Skip                            int
}

func (lo *LookupOptions) limit() int {
	if lo.Limit > 0 {
		return lo.Limit
	}
	return 100
}

type WhereClause struct {
	RegistryApprovedGuest registry_approved_guest.MongoWhereClause
	AddressAccessSessions address_access_session.MongoWhereClause
	Owner                 owner_user.MongoWhereClause
	Registry              registry.MongoWhereClause
}

func aggregateWithRefs(ctx context.Context, where WhereClause, collection *mongo.Collection, lookupOptions LookupOptions) (QueryResult, error) {
	whereRegistryApprovedGuest, err := where.RegistryApprovedGuest.GetLookupQuery()
	if err != nil {
		return QueryResult{}, err
	}

	sortStage := bson.D{}

	if lookupOptions.Sort.CreatedAt > 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: 1})
	} else if lookupOptions.Sort.CreatedAt < 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: -1})
	}
	if lookupOptions.Sort.EmailHash > 0 {
		sortStage = append(sortStage, bson.E{Key: "emailHash", Value: 1})
	} else if lookupOptions.Sort.EmailHash < 0 {
		sortStage = append(sortStage, bson.E{Key: "emailHash", Value: -1})
	}
	if lookupOptions.Sort.OwnerId > 0 {
		sortStage = append(sortStage, bson.E{Key: "ownerId", Value: 1})
	} else if lookupOptions.Sort.OwnerId < 0 {
		sortStage = append(sortStage, bson.E{Key: "ownerId", Value: -1})
	}
	if lookupOptions.Sort.RegistryId > 0 {
		sortStage = append(sortStage, bson.E{Key: "registryId", Value: 1})
	} else if lookupOptions.Sort.RegistryId < 0 {
		sortStage = append(sortStage, bson.E{Key: "registryId", Value: -1})
	}
	if lookupOptions.Sort.UpdatedAt > 0 {
		sortStage = append(sortStage, bson.E{Key: "updated.at", Value: 1})
	} else if lookupOptions.Sort.UpdatedAt < 0 {
		sortStage = append(sortStage, bson.E{Key: "updated.at", Value: -1})
	}

	if lookupOptions.OwnerProjection != nil {
		lookupOptions.Projection.OwnerId = true
	}
	if lookupOptions.RegistryProjection != nil {
		lookupOptions.Projection.RegistryId = true
	}

	if lookupOptions.AddressAccessSessionsProjection != nil {
		lookupOptions.Projection.Id = true
	}

	dataPipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: whereRegistryApprovedGuest},
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

	// Add $lookup stage for Owner
	if lookupOptions.OwnerProjection != nil {
		// whereOwnerId, err := where.Owner.GetLookupQuery()
		// if err != nil {
		//     return QueryResult{}, err
		// }
		objectProject := bson.E{Key: "$project", Value: lookupOptions.OwnerProjection.ToBson()}
		objectPipeline := bson.D{objectProject}
		// if len(whereOwnerId) > 0 {
		//     objectPipeline = bson.D{
		//         {Key: "$match", Value: whereOwnerId},
		//         objectProject,
		//     }
		// }
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				// TODO get actual collection name
				{Key: "from", Value: "owner_users"},
				{Key: "localField", Value: "ownerUserId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "Owner"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})

		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$Owner"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		})
	}
	// Add $lookup stage for Registry
	if lookupOptions.RegistryProjection != nil {
		// whereRegistryId, err := where.Registry.GetLookupQuery()
		// if err != nil {
		//     return QueryResult{}, err
		// }
		objectProject := bson.E{Key: "$project", Value: lookupOptions.RegistryProjection.ToBson()}
		objectPipeline := bson.D{objectProject}
		// if len(whereRegistryId) > 0 {
		//     objectPipeline = bson.D{
		//         {Key: "$match", Value: whereRegistryId},
		//         objectProject,
		//     }
		// }
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				// TODO get actual collection name
				{Key: "from", Value: "registries"},
				{Key: "localField", Value: "registryId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "Registry"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})

		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$Registry"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		})
	}
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
				{Key: "foreignField", Value: "registryApprovedGuestId"},
				{Key: "as", Value: "AddressAccessSessions"},
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
							Value: whereRegistryApprovedGuest,
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
