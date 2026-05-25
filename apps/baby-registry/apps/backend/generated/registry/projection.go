package registry

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
)

type Projection struct {
	Id                       bool                   `json:"id"`
	AddressAccessMode        bool                   `json:"addressAccessMode"`
	CoverImageUrl            bool                   `json:"coverImageUrl"`
	Created                  bool                   `json:"created"`
	CreatedFields            actor_trace.Projection `json:"createdFields,omitempty"`
	DueDate                  bool                   `json:"dueDate"`
	IsPublic                 bool                   `json:"isPublic"`
	OwnerId                  bool
	ParentNames              bool                   `json:"parentNames"`
	ShippingCity             bool                   `json:"shippingCity"`
	ShippingCountry          bool                   `json:"shippingCountry"`
	ShippingDeliveryNotes    bool                   `json:"shippingDeliveryNotes"`
	ShippingLine1            bool                   `json:"shippingLine1"`
	ShippingLine2            bool                   `json:"shippingLine2"`
	ShippingPolicyVersion    bool                   `json:"shippingPolicyVersion"`
	ShippingPostalCode       bool                   `json:"shippingPostalCode"`
	ShippingRecipientName    bool                   `json:"shippingRecipientName"`
	ShippingRegion           bool                   `json:"shippingRegion"`
	Slug                     bool                   `json:"slug"`
	ThemeColor               bool                   `json:"themeColor"`
	Title                    bool                   `json:"title"`
	Updated                  bool                   `json:"updated"`
	UpdatedFields            actor_trace.Projection `json:"updatedFields,omitempty"`
	UpdatedByOwnerUser       bool                   `json:"updatedByOwnerUser"`
	UpdatedByOwnerUserFields actor_trace.Projection `json:"updatedByOwnerUserFields,omitempty"`
	WelcomeMessage           bool                   `json:"welcomeMessage"`
}

func NewProjection(defaultVal bool) Projection {
	return Projection{
		Id:                       defaultVal,
		AddressAccessMode:        defaultVal,
		CoverImageUrl:            defaultVal,
		Created:                  defaultVal,
		CreatedFields:            actor_trace.NewProjection(defaultVal),
		DueDate:                  defaultVal,
		IsPublic:                 defaultVal,
		OwnerId:                  defaultVal,
		ParentNames:              defaultVal,
		ShippingCity:             defaultVal,
		ShippingCountry:          defaultVal,
		ShippingDeliveryNotes:    defaultVal,
		ShippingLine1:            defaultVal,
		ShippingLine2:            defaultVal,
		ShippingPolicyVersion:    defaultVal,
		ShippingPostalCode:       defaultVal,
		ShippingRecipientName:    defaultVal,
		ShippingRegion:           defaultVal,
		Slug:                     defaultVal,
		ThemeColor:               defaultVal,
		Title:                    defaultVal,
		Updated:                  defaultVal,
		UpdatedFields:            actor_trace.NewProjection(defaultVal),
		UpdatedByOwnerUser:       defaultVal,
		UpdatedByOwnerUserFields: actor_trace.NewProjection(defaultVal),
		WelcomeMessage:           defaultVal,
	}
}

func (p Projection) ToBson() bson.M {
	projection := bson.M{}
	projection["_id"] = 1
	if p.AddressAccessMode {
		projection["addressAccessMode"] = 1
	}
	if p.CoverImageUrl {
		projection["coverImageUrl"] = 1
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
	if p.DueDate {
		projection["dueDate"] = 1
	}
	if p.IsPublic {
		projection["isPublic"] = 1
	}
	if p.OwnerId {
		projection["ownerId"] = 1
	}
	if p.ParentNames {
		projection["parentNames"] = 1
	}
	if p.ShippingCity {
		projection["shippingCity"] = 1
	}
	if p.ShippingCountry {
		projection["shippingCountry"] = 1
	}
	if p.ShippingDeliveryNotes {
		projection["shippingDeliveryNotes"] = 1
	}
	if p.ShippingLine1 {
		projection["shippingLine1"] = 1
	}
	if p.ShippingLine2 {
		projection["shippingLine2"] = 1
	}
	if p.ShippingPolicyVersion {
		projection["shippingPolicyVersion"] = 1
	}
	if p.ShippingPostalCode {
		projection["shippingPostalCode"] = 1
	}
	if p.ShippingRecipientName {
		projection["shippingRecipientName"] = 1
	}
	if p.ShippingRegion {
		projection["shippingRegion"] = 1
	}
	if p.Slug {
		projection["slug"] = 1
	}
	if p.ThemeColor {
		projection["themeColor"] = 1
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
	if p.WelcomeMessage {
		projection["welcomeMessage"] = 1
	}
	return projection
}
