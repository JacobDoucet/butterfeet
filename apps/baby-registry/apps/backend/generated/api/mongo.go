package api

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/event_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request_api"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoBackedClient(db *mongo.Database) Client {
	return &mongoClient{
		addressAccessSession:   address_access_session_api.NewMongoBackedClient(db),
		event:                  event_api.NewMongoBackedClient(db),
		ownerUser:              owner_user_api.NewMongoBackedClient(db),
		registry:               registry_api.NewMongoBackedClient(db),
		registryApprovedGuest:  registry_approved_guest_api.NewMongoBackedClient(db),
		registryItem:           registry_item_api.NewMongoBackedClient(db),
		reservation:            reservation_api.NewMongoBackedClient(db),
		shippingAddressRequest: shipping_address_request_api.NewMongoBackedClient(db),
	}
}

type mongoClient struct {
	addressAccessSession   address_access_session_api.Client
	event                  event_api.Client
	ownerUser              owner_user_api.Client
	registry               registry_api.Client
	registryApprovedGuest  registry_approved_guest_api.Client
	registryItem           registry_item_api.Client
	reservation            reservation_api.Client
	shippingAddressRequest shipping_address_request_api.Client
}

func (m *mongoClient) ValidateClients() error {
	return nil
}
func (c *mongoClient) AddressAccessSession() address_access_session_api.Client {
	return c.addressAccessSession
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
func (c *mongoClient) RegistryApprovedGuest() registry_approved_guest_api.Client {
	return c.registryApprovedGuest
}
func (c *mongoClient) RegistryItem() registry_item_api.Client {
	return c.registryItem
}
func (c *mongoClient) Reservation() reservation_api.Client {
	return c.reservation
}
func (c *mongoClient) ShippingAddressRequest() shipping_address_request_api.Client {
	return c.shippingAddressRequest
}
