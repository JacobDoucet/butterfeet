package address_access_session_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
)

type Model struct {
	address_access_session.MongoRecord `bson:",inline"`
	ApprovedGuest                      *registry_approved_guest.MongoRecord `bson:"ApprovedGuest,omitempty"`
	Owner                              *owner_user.MongoRecord              `bson:"Owner,omitempty"`
	Registry                           *registry.MongoRecord                `bson:"Registry,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
