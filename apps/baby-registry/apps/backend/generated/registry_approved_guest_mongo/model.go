package registry_approved_guest_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
)

type Model struct {
	registry_approved_guest.MongoRecord `bson:",inline"`
	AddressAccessSessions               *[]address_access_session.MongoRecord `bson:"AddressAccessSessions,omitempty"`
	Owner                               *owner_user.MongoRecord               `bson:"Owner,omitempty"`
	Registry                            *registry.MongoRecord                 `bson:"Registry,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
