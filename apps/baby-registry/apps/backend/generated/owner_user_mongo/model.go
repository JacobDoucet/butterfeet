package owner_user_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
)

type Model struct {
	owner_user.MongoRecord `bson:",inline"`
	Registrys              *[]registry.MongoRecord `bson:"Registrys,omitempty"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
