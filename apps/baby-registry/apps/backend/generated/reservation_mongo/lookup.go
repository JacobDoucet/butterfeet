package reservation_mongo

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LookupOptions struct {
	Projection         reservation.Projection
	Sort               reservation.MongoSortParams
	ItemProjection     *registry_item.Projection
	RegistryProjection *registry.Projection
	Limit              int
	Skip               int
}

func (lo *LookupOptions) limit() int {
	if lo.Limit > 0 {
		return lo.Limit
	}
	return 100
}

type WhereClause struct {
	Reservation reservation.MongoWhereClause
	Item        registry_item.MongoWhereClause
	Registry    registry.MongoWhereClause
}

func aggregateWithRefs(ctx context.Context, where WhereClause, collection *mongo.Collection, lookupOptions LookupOptions) (QueryResult, error) {
	whereReservation, err := where.Reservation.GetLookupQuery()
	if err != nil {
		return QueryResult{}, err
	}

	sortStage := bson.D{}

	if lookupOptions.Sort.CreatedAt > 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: 1})
	} else if lookupOptions.Sort.CreatedAt < 0 {
		sortStage = append(sortStage, bson.E{Key: "created.at", Value: -1})
	}
	if lookupOptions.Sort.ItemId > 0 {
		sortStage = append(sortStage, bson.E{Key: "itemId", Value: 1})
	} else if lookupOptions.Sort.ItemId < 0 {
		sortStage = append(sortStage, bson.E{Key: "itemId", Value: -1})
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

	if lookupOptions.ItemProjection != nil {
		lookupOptions.Projection.ItemId = true
	}
	if lookupOptions.RegistryProjection != nil {
		lookupOptions.Projection.RegistryId = true
	}

	dataPipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: whereReservation},
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

	// Add $lookup stage for Item
	if lookupOptions.ItemProjection != nil {
		// whereItemId, err := where.Item.GetLookupQuery()
		// if err != nil {
		//     return QueryResult{}, err
		// }
		objectProject := bson.E{Key: "$project", Value: lookupOptions.ItemProjection.ToBson()}
		objectPipeline := bson.D{objectProject}
		// if len(whereItemId) > 0 {
		//     objectPipeline = bson.D{
		//         {Key: "$match", Value: whereItemId},
		//         objectProject,
		//     }
		// }
		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$lookup", Value: bson.D{
				// TODO get actual collection name
				{Key: "from", Value: "registry_items"},
				{Key: "localField", Value: "registryItemId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "Item"},
				{Key: "pipeline", Value: bson.A{objectPipeline}},
			}},
		})

		dataPipeline = append(dataPipeline, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$Item"},
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
							Value: whereReservation,
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
