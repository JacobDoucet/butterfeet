package registry_payment_method

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id                       bool                   `json:"id"`
	BankAccountName          bool                   `json:"bankAccountName"`
	BankAccountNumber        bool                   `json:"bankAccountNumber"`
	BankIban                 bool                   `json:"bankIban"`
	BankName                 bool                   `json:"bankName"`
	BankRoutingNumber        bool                   `json:"bankRoutingNumber"`
	BankSwift                bool                   `json:"bankSwift"`
	Created                  bool                   `json:"created"`
	CreatedFields            actor_trace.Projection `json:"createdFields,omitempty"`
	DisplayName              bool                   `json:"displayName"`
	Enabled                  bool                   `json:"enabled"`
	Instructions             bool                   `json:"instructions"`
	OwnerId                  bool
	PaymentUrl               bool `json:"paymentUrl"`
	Position                 bool `json:"position"`
	RecipientEmail           bool `json:"recipientEmail"`
	RecipientPhone           bool `json:"recipientPhone"`
	RegistryId               bool
	Type                     bool                   `json:"type"`
	Updated                  bool                   `json:"updated"`
	UpdatedFields            actor_trace.Projection `json:"updatedFields,omitempty"`
	UpdatedByOwnerUser       bool                   `json:"updatedByOwnerUser"`
	UpdatedByOwnerUserFields actor_trace.Projection `json:"updatedByOwnerUserFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:                       defaultVal,
		BankAccountName:          defaultVal,
		BankAccountNumber:        defaultVal,
		BankIban:                 defaultVal,
		BankName:                 defaultVal,
		BankRoutingNumber:        defaultVal,
		BankSwift:                defaultVal,
		Created:                  defaultVal,
		CreatedFields:            actor_trace.NewProjection(defaultVal),
		DisplayName:              defaultVal,
		Enabled:                  defaultVal,
		Instructions:             defaultVal,
		OwnerId:                  defaultVal,
		PaymentUrl:               defaultVal,
		Position:                 defaultVal,
		RecipientEmail:           defaultVal,
		RecipientPhone:           defaultVal,
		RegistryId:               defaultVal,
		Type:                     defaultVal,
		Updated:                  defaultVal,
		UpdatedFields:            actor_trace.NewProjection(defaultVal),
		UpdatedByOwnerUser:       defaultVal,
		UpdatedByOwnerUserFields: actor_trace.NewProjection(defaultVal),
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	projection["_id"] = 1
	if p.BankAccountName {
		projection["bankAccountName"] = 1
	}
	if p.BankAccountNumber {
		projection["bankAccountNumber"] = 1
	}
	if p.BankIban {
		projection["bankIban"] = 1
	}
	if p.BankName {
		projection["bankName"] = 1
	}
	if p.BankRoutingNumber {
		projection["bankRoutingNumber"] = 1
	}
	if p.BankSwift {
		projection["bankSwift"] = 1
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
	if p.DisplayName {
		projection["displayName"] = 1
	}
	if p.Enabled {
		projection["enabled"] = 1
	}
	if p.Instructions {
		projection["instructions"] = 1
	}
	if p.OwnerId {
		projection["ownerId"] = 1
	}
	if p.PaymentUrl {
		projection["paymentUrl"] = 1
	}
	if p.Position {
		projection["position"] = 1
	}
	if p.RecipientEmail {
		projection["recipientEmail"] = 1
	}
	if p.RecipientPhone {
		projection["recipientPhone"] = 1
	}
	if p.RegistryId {
		projection["registryId"] = 1
	}
	if p.Type {
		projection["type"] = 1
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
