package event_mongo

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/event"
)

type Model struct {
	event.MongoRecord `bson:",inline"`
}

type QueryResult struct {
	Data  []Model `bson:"data"`
	Count int     `bson:"count"`
	Skip  int     `bson:"skip"`
}
