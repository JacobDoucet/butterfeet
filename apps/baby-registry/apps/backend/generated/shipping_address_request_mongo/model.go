package shipping_address_request_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
)

type Model struct {
	shipping_address_request.MongoRecord `bson:",inline"`
	Owner                                *owner_user.MongoRecord    `bson:"Owner,omitempty"`
	Registry                             *registry.MongoRecord      `bson:"Registry,omitempty"`
	RegistryItem                         *registry_item.MongoRecord `bson:"RegistryItem,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
