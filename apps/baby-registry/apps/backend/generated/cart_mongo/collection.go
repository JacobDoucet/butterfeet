package cart_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName = "carts"

func CreateIndexes(ctx context.Context, db *mongo.Database) error {
	if err := createContributorIdxIndex(ctx, db); err != nil {
		return err
	}
	if err := createCreatedAtIndex(ctx, db); err != nil {
		return err
	}
	if err := createOwnerIdxIndex(ctx, db); err != nil {
		return err
	}
	if err := createReferenceUniqueIndex(ctx, db); err != nil {
		return err
	}
	if err := createRegistryIdxIndex(ctx, db); err != nil {
		return err
	}
	if err := createUpdatedAtIndex(ctx, db); err != nil {
		return err
	}
	return nil
}

func createContributorIdxIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "registryId", Value: 1},
			{Key: "contributorEmail", Value: 1},
		},
		Options: options.Index().SetName("contributor_idx"),
	})
	return err
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

func createOwnerIdxIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "ownerId", Value: 1},
		},
		Options: options.Index().SetName("owner_idx"),
	})
	return err
}

func createReferenceUniqueIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "referenceCode", Value: 1},
		},
		Options: options.Index().SetName("reference_unique").SetUnique(true).
			SetPartialFilterExpression(bson.D{{Key: "referenceCode", Value: bson.D{{Key: "$exists", Value: true}}}}),
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
