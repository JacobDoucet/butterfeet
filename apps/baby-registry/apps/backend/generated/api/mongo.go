package api

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/event_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_api"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoBackedClient(db *mongo.Database) Client {
	return &mongoClient{
		event:        event_api.NewMongoBackedClient(db),
		ownerUser:    owner_user_api.NewMongoBackedClient(db),
		registry:     registry_api.NewMongoBackedClient(db),
		registryItem: registry_item_api.NewMongoBackedClient(db),
		reservation:  reservation_api.NewMongoBackedClient(db),
	}
}

type mongoClient struct {
	event        event_api.Client
	ownerUser    owner_user_api.Client
	registry     registry_api.Client
	registryItem registry_item_api.Client
	reservation  reservation_api.Client
}

func (m *mongoClient) ValidateClients() error {
	return nil
}
func (c *mongoClient) Event() event_api.Client {
	return c.event
}
func (c *mongoClient) OwnerUser() owner_user_api.Client {
	return c.ownerUser
}
func (c *mongoClient) Registry() registry_api.Client {
	return c.registry
}
func (c *mongoClient) RegistryItem() registry_item_api.Client {
	return c.registryItem
}
func (c *mongoClient) Reservation() reservation_api.Client {
	return c.reservation
}
