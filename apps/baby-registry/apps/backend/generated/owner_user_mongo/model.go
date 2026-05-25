package owner_user_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
)

type Model struct {
	owner_user.MongoRecord  `bson:",inline"`
	AddressAccessSessions   *[]address_access_session.MongoRecord   `bson:"AddressAccessSessions,omitempty"`
	RegistryApprovedGuests  *[]registry_approved_guest.MongoRecord  `bson:"RegistryApprovedGuests,omitempty"`
	Registrys               *[]registry.MongoRecord                 `bson:"Registrys,omitempty"`
	ShippingAddressRequests *[]shipping_address_request.MongoRecord `bson:"ShippingAddressRequests,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
