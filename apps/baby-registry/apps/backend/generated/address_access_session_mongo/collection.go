package address_access_session_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName = "address_access_sessions"

func CreateIndexes(ctx context.Context, db *mongo.Database) error {
	if err := createCreatedAtIndex(ctx, db); err != nil {
		return err
	}
	if err := createOwnerIdxIndex(ctx, db); err != nil {
		return err
	}
	if err := createRegistryEmailIdxIndex(ctx, db); err != nil {
		return err
	}
	if err := createTokenUniqueIndex(ctx, db); err != nil {
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

func createRegistryEmailIdxIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "registryId", Value: 1},
			{Key: "emailHash", Value: 1},
		},
		Options: options.Index().SetName("registry_email_idx"),
	})
	return err
}

func createTokenUniqueIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(CollectionName)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "tokenHash", Value: 1},
		},
		Options: options.Index().SetName("token_unique").SetUnique(true).
			SetPartialFilterExpression(bson.D{{Key: "tokenHash", Value: bson.D{{Key: "$exists", Value: true}}}}),
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
