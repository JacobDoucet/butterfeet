package actor_role

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	OwnerId bool `json:"ownerId"`
	Role    bool `json:"role"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		OwnerId: defaultVal,
		Role:    defaultVal,
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	if p.OwnerId {
		projection["ownerId"] = 1
	}
	if p.Role {
		projection["role"] = 1
	}
	return projection
}
