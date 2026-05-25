package registry_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
)

type Model struct {
	registry.MongoRecord `bson:",inline"`
	RegistryItems        *[]registry_item.MongoRecord `bson:"RegistryItems,omitempty"`
	Reservations         *[]reservation.MongoRecord   `bson:"Reservations,omitempty"`
	Owner                *owner_user.MongoRecord      `bson:"Owner,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
