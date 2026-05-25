package registry_item_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
)

type Model struct {
	registry_item.MongoRecord `bson:",inline"`
	Reservations              *[]reservation.MongoRecord              `bson:"Reservations,omitempty"`
	ShippingAddressRequests   *[]shipping_address_request.MongoRecord `bson:"ShippingAddressRequests,omitempty"`
	Registry                  *registry.MongoRecord                   `bson:"Registry,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
