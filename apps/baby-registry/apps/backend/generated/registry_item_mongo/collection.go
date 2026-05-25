package registry_item_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName = "registry_items"

func CreateIndexes(ctx context.Context, db *mongo.Database) error {
	if err := createCreatedAtIndex(ctx, db); err != nil {
		return err
	}
	if err := createRegistryIdxIndex(ctx, db); err != nil {
		return err
	}
	if err := createRegistryPositionIndex(ctx, db); err != nil {
		return err
	}
	if err := createUpdatedAtIndex(ctx, db); err != nil {
		return err
	}
	return nil
}

func createCreatedAtIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "created.at", Value: -1},
		},
		Options: options.Index().SetName("createdAt"),
	})
	return err
}

func createRegistryIdxIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "registryId", Value: 1},
		},
		Options: options.Index().SetName("registry_idx"),
	})
	return err
}

func createRegistryPositionIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "registryId", Value: 1},
			{Key: "position", Value: 1},
		},
		Options: options.Index().SetName("registry_position"),
	})
	return err
}

func createUpdatedAtIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "updated.at", Value: -1},
		},
		Options: options.Index().SetName("updatedAt"),
	})
	return err
}
