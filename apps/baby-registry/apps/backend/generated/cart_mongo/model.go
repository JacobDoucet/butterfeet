package cart_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
)

type Model struct {
	cart.MongoRecord `bson:",inline"`
	Reservations     *[]reservation.MongoRecord           `bson:"Reservations,omitempty"`
	Owner            *owner_user.MongoRecord              `bson:"Owner,omitempty"`
	PaymentMethod    *registry_payment_method.MongoRecord `bson:"PaymentMethod,omitempty"`
	Registry         *registry.MongoRecord                `bson:"Registry,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
