package shipping_address_request

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id                       bool                   `json:"id"`
	Created                  bool                   `json:"created"`
	CreatedFields            actor_trace.Projection `json:"createdFields,omitempty"`
	DecisionReason           bool                   `json:"decisionReason"`
	EmailEnc                 bool                   `json:"emailEnc"`
	EmailHash                bool                   `json:"emailHash"`
	Name                     bool                   `json:"name"`
	Note                     bool                   `json:"note"`
	OwnerId                  bool
	PolicyVersion            bool `json:"policyVersion"`
	RegistryId               bool
	RegistryItemId           bool
	Status                   bool                   `json:"status"`
	Updated                  bool                   `json:"updated"`
	UpdatedFields            actor_trace.Projection `json:"updatedFields,omitempty"`
	UpdatedByOwnerUser       bool                   `json:"updatedByOwnerUser"`
	UpdatedByOwnerUserFields actor_trace.Projection `json:"updatedByOwnerUserFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:                       defaultVal,
		Created:                  defaultVal,
		CreatedFields:            actor_trace.NewProjection(defaultVal),
		DecisionReason:           defaultVal,
		EmailEnc:                 defaultVal,
		EmailHash:                defaultVal,
		Name:                     defaultVal,
		Note:                     defaultVal,
		OwnerId:                  defaultVal,
		PolicyVersion:            defaultVal,
		RegistryId:               defaultVal,
		RegistryItemId:           defaultVal,
		Status:                   defaultVal,
		Updated:                  defaultVal,
		UpdatedFields:            actor_trace.NewProjection(defaultVal),
		UpdatedByOwnerUser:       defaultVal,
		UpdatedByOwnerUserFields: actor_trace.NewProjection(defaultVal),
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	projection["_id"] = 1
	if p.Created {
		if p.CreatedFields.ActorId {
			projection["created.actorId"] = 1
		}
		if p.CreatedFields.ActorName {
			projection["created.actorName"] = 1
		}
		if p.CreatedFields.ActorType {
			projection["created.actorType"] = 1
		}
		if p.CreatedFields.At {
			projection["created.at"] = 1
		}
	}
	if p.DecisionReason {
		projection["decisionReason"] = 1
	}
	if p.EmailEnc {
		projection["emailEnc"] = 1
	}
	if p.EmailHash {
		projection["emailHash"] = 1
	}
	if p.Name {
		projection["name"] = 1
	}
	if p.Note {
		projection["note"] = 1
	}
	if p.OwnerId {
		projection["ownerId"] = 1
	}
	if p.PolicyVersion {
		projection["policyVersion"] = 1
	}
	if p.RegistryId {
		projection["registryId"] = 1
	}
	if p.RegistryItemId {
		projection["registryItemId"] = 1
	}
	if p.Status {
		projection["status"] = 1
	}
	if p.Updated {
		if p.UpdatedFields.ActorId {
			projection["updated.actorId"] = 1
		}
		if p.UpdatedFields.ActorName {
			projection["updated.actorName"] = 1
		}
		if p.UpdatedFields.ActorType {
			projection["updated.actorType"] = 1
		}
		if p.UpdatedFields.At {
			projection["updated.at"] = 1
		}
	}
	if p.UpdatedByOwnerUser {
		if p.UpdatedByOwnerUserFields.ActorId {
			projection["updatedByOwnerUser.actorId"] = 1
		}
		if p.UpdatedByOwnerUserFields.ActorName {
			projection["updatedByOwnerUser.actorName"] = 1
		}
		if p.UpdatedByOwnerUserFields.ActorType {
			projection["updatedByOwnerUser.actorType"] = 1
		}
		if p.UpdatedByOwnerUserFields.At {
			projection["updatedByOwnerUser.at"] = 1
		}
	}
	return projection
}
