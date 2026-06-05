package registry_payment_method_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
)

type Model struct {
	registry_payment_method.MongoRecord `bson:",inline"`
	Carts                               *[]cart.MongoRecord     `bson:"Carts,omitempty"`
	Owner                               *owner_user.MongoRecord `bson:"Owner,omitempty"`
	Registry                            *registry.MongoRecord   `bson:"Registry,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
