package registry_item_mongo

import (
	"context"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(ctx context.Context, db *mongo.Database, id primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if err = reservation_mongo.DeleteByRegistryItem(ctx, db, id); err != nil {
		return err
	}
	return nil
}

func DeleteBulk(ctx context.Context, db *mongo.Database, ids []primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	_, err := coll.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return err
	}

	if err = reservation_mongo.DeleteByManyRegistryItems(ctx, db, ids); err != nil {
		return err
	}
	return nil
}

func DeleteByRegistry(ctx context.Context, db *mongo.Database, registryId primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	_, err := coll.DeleteMany(ctx, bson.M{"registryId": registryId})
	if err != nil {
		return err
	}
	return nil
}

func DeleteByManyRegistrys(ctx context.Context, db *mongo.Database, registryId []primitive.ObjectID) error {
	coll := db.Collection(CollectionName)
	_, err := coll.DeleteMany(ctx, bson.M{"registryId": bson.M{"$in": registryId}})
	if err != nil {
		return err
	}
	return nil
}
