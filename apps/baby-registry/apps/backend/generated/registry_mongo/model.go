package registry_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
)

type Model struct {
	registry.MongoRecord    `bson:",inline"`
	AddressAccessSessions   *[]address_access_session.MongoRecord   `bson:"AddressAccessSessions,omitempty"`
	Carts                   *[]cart.MongoRecord                     `bson:"Carts,omitempty"`
	RegistryApprovedGuests  *[]registry_approved_guest.MongoRecord  `bson:"RegistryApprovedGuests,omitempty"`
	RegistryItems           *[]registry_item.MongoRecord            `bson:"RegistryItems,omitempty"`
	RegistryPaymentMethods  *[]registry_payment_method.MongoRecord  `bson:"RegistryPaymentMethods,omitempty"`
	Reservations            *[]reservation.MongoRecord              `bson:"Reservations,omitempty"`
	ShippingAddressRequests *[]shipping_address_request.MongoRecord `bson:"ShippingAddressRequests,omitempty"`
	Owner                   *owner_user.MongoRecord                 `bson:"Owner,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
