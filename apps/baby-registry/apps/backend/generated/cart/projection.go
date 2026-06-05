package cart

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id                       bool                   `json:"id"`
	AmountCents              bool                   `json:"amountCents"`
	ClaimedAt                bool                   `json:"claimedAt"`
	ContributorEmail         bool                   `json:"contributorEmail"`
	ContributorName          bool                   `json:"contributorName"`
	Created                  bool                   `json:"created"`
	CreatedFields            actor_trace.Projection `json:"createdFields,omitempty"`
	Currency                 bool                   `json:"currency"`
	DecidedAt                bool                   `json:"decidedAt"`
	DecisionReason           bool                   `json:"decisionReason"`
	Message                  bool                   `json:"message"`
	MethodDisplayName        bool                   `json:"methodDisplayName"`
	MethodType               bool                   `json:"methodType"`
	OwnerId                  bool
	PaymentMethodId          bool
	ReferenceCode            bool `json:"referenceCode"`
	RegistryId               bool
	Status                   bool                   `json:"status"`
	Updated                  bool                   `json:"updated"`
	UpdatedFields            actor_trace.Projection `json:"updatedFields,omitempty"`
	UpdatedByOwnerUser       bool                   `json:"updatedByOwnerUser"`
	UpdatedByOwnerUserFields actor_trace.Projection `json:"updatedByOwnerUserFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:                       defaultVal,
		AmountCents:              defaultVal,
		ClaimedAt:                defaultVal,
		ContributorEmail:         defaultVal,
		ContributorName:          defaultVal,
		Created:                  defaultVal,
		CreatedFields:            actor_trace.NewProjection(defaultVal),
		Currency:                 defaultVal,
		DecidedAt:                defaultVal,
		DecisionReason:           defaultVal,
		Message:                  defaultVal,
		MethodDisplayName:        defaultVal,
		MethodType:               defaultVal,
		OwnerId:                  defaultVal,
		PaymentMethodId:          defaultVal,
		ReferenceCode:            defaultVal,
		RegistryId:               defaultVal,
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
	if p.AmountCents {
		projection["amountCents"] = 1
	}
	if p.ClaimedAt {
		projection["claimedAt"] = 1
	}
	if p.ContributorEmail {
		projection["contributorEmail"] = 1
	}
	if p.ContributorName {
		projection["contributorName"] = 1
	}
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
	if p.Currency {
		projection["currency"] = 1
	}
	if p.DecidedAt {
		projection["decidedAt"] = 1
	}
	if p.DecisionReason {
		projection["decisionReason"] = 1
	}
	if p.Message {
		projection["message"] = 1
	}
	if p.MethodDisplayName {
		projection["methodDisplayName"] = 1
	}
	if p.MethodType {
		projection["methodType"] = 1
	}
	if p.OwnerId {
		projection["ownerId"] = 1
	}
	if p.PaymentMethodId {
		projection["paymentMethodId"] = 1
	}
	if p.ReferenceCode {
		projection["referenceCode"] = 1
	}
	if p.RegistryId {
		projection["registryId"] = 1
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
