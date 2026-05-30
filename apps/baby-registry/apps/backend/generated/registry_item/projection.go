package registry_item

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id                       bool                   `json:"id"`
	Category                 bool                   `json:"category"`
	Created                  bool                   `json:"created"`
	CreatedFields            actor_trace.Projection `json:"createdFields,omitempty"`
	Currency                 bool                   `json:"currency"`
	Description              bool                   `json:"description"`
	ImageBgColor             bool                   `json:"imageBgColor"`
	ImageUrl                 bool                   `json:"imageUrl"`
	NoSubstitutes            bool                   `json:"noSubstitutes"`
	Notes                    bool                   `json:"notes"`
	OwnerPurchased           bool                   `json:"ownerPurchased"`
	ParentItemId             bool
	Position                 bool `json:"position"`
	PriceCents               bool `json:"priceCents"`
	ProductUrl               bool `json:"productUrl"`
	Quantity                 bool `json:"quantity"`
	QuantityUnlimited        bool `json:"quantityUnlimited"`
	RegistryId               bool
	Source                   bool                   `json:"source"`
	Title                    bool                   `json:"title"`
	Updated                  bool                   `json:"updated"`
	UpdatedFields            actor_trace.Projection `json:"updatedFields,omitempty"`
	UpdatedByOwnerUser       bool                   `json:"updatedByOwnerUser"`
	UpdatedByOwnerUserFields actor_trace.Projection `json:"updatedByOwnerUserFields,omitempty"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:                       defaultVal,
		Category:                 defaultVal,
		Created:                  defaultVal,
		CreatedFields:            actor_trace.NewProjection(defaultVal),
		Currency:                 defaultVal,
		Description:              defaultVal,
		ImageBgColor:             defaultVal,
		ImageUrl:                 defaultVal,
		NoSubstitutes:            defaultVal,
		Notes:                    defaultVal,
		OwnerPurchased:           defaultVal,
		ParentItemId:             defaultVal,
		Position:                 defaultVal,
		PriceCents:               defaultVal,
		ProductUrl:               defaultVal,
		Quantity:                 defaultVal,
		QuantityUnlimited:        defaultVal,
		RegistryId:               defaultVal,
		Source:                   defaultVal,
		Title:                    defaultVal,
		Updated:                  defaultVal,
		UpdatedFields:            actor_trace.NewProjection(defaultVal),
		UpdatedByOwnerUser:       defaultVal,
		UpdatedByOwnerUserFields: actor_trace.NewProjection(defaultVal),
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	projection["_id"] = 1
	if p.Category {
		projection["category"] = 1
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
	if p.Description {
		projection["description"] = 1
	}
	if p.ImageBgColor {
		projection["imageBgColor"] = 1
	}
	if p.ImageUrl {
		projection["imageUrl"] = 1
	}
	if p.NoSubstitutes {
		projection["noSubstitutes"] = 1
	}
	if p.Notes {
		projection["notes"] = 1
	}
	if p.OwnerPurchased {
		projection["ownerPurchased"] = 1
	}
	if p.ParentItemId {
		projection["parentItemId"] = 1
	}
	if p.Position {
		projection["position"] = 1
	}
	if p.PriceCents {
		projection["priceCents"] = 1
	}
	if p.ProductUrl {
		projection["productUrl"] = 1
	}
	if p.Quantity {
		projection["quantity"] = 1
	}
	if p.QuantityUnlimited {
		projection["quantityUnlimited"] = 1
	}
	if p.RegistryId {
		projection["registryId"] = 1
	}
	if p.Source {
		projection["source"] = 1
	}
	if p.Title {
		projection["title"] = 1
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
