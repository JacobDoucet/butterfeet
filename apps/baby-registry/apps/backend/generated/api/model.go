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
)

type Client interface {
	AddressAccessSession() address_access_session_api.Client
	Event() event_api.Client
	OwnerUser() owner_user_api.Client
	Registry() registry_api.Client
	RegistryApprovedGuest() registry_approved_guest_api.Client
	RegistryItem() registry_item_api.Client
	Reservation() reservation_api.Client
	ShippingAddressRequest() shipping_address_request_api.Client
	ValidateClients() error
}

type CustomClient struct {
	addressAccessSession   address_access_session_api.Client
	event                  event_api.Client
	ownerUser              owner_user_api.Client
	registry               registry_api.Client
	registryApprovedGuest  registry_approved_guest_api.Client
	registryItem           registry_item_api.Client
	reservation            reservation_api.Client
	shippingAddressRequest shipping_address_request_api.Client
}

func NewUnimplementedClient() CustomClient {
	c := &CustomClient{}
	c.ValidateClients()
	return *c
}

func (c *CustomClient) ValidateClients() error {
	if c.addressAccessSession == nil {
		c.addressAccessSession = address_access_session_api.New(&address_access_session_api.UnimplementedClient{})
	}
	if c.event == nil {
		c.event = event_api.New(&event_api.UnimplementedClient{})
	}
	if c.ownerUser == nil {
		c.ownerUser = owner_user_api.New(&owner_user_api.UnimplementedClient{})
	}
	if c.registry == nil {
		c.registry = registry_api.New(&registry_api.UnimplementedClient{})
	}
	if c.registryApprovedGuest == nil {
		c.registryApprovedGuest = registry_approved_guest_api.New(&registry_approved_guest_api.UnimplementedClient{})
	}
	if c.registryItem == nil {
		c.registryItem = registry_item_api.New(&registry_item_api.UnimplementedClient{})
	}
	if c.reservation == nil {
		c.reservation = reservation_api.New(&reservation_api.UnimplementedClient{})
	}
	if c.shippingAddressRequest == nil {
		c.shippingAddressRequest = shipping_address_request_api.New(&shipping_address_request_api.UnimplementedClient{})
	}
	return nil
}
func (c *CustomClient) UseAddressAccessSessionClient(client address_access_session_api.Client) *CustomClient {
	if client == nil {
		c.addressAccessSession = address_access_session_api.New(&address_access_session_api.UnimplementedClient{})
		return c
	}
	c.addressAccessSession = client
	return c
}

func (c *CustomClient) AddressAccessSession() address_access_session_api.Client {
	return c.addressAccessSession
}
func (c *CustomClient) UseEventClient(client event_api.Client) *CustomClient {
	if client == nil {
		c.event = event_api.New(&event_api.UnimplementedClient{})
		return c
	}
	c.event = client
	return c
}

func (c *CustomClient) Event() event_api.Client {
	return c.event
}
func (c *CustomClient) UseOwnerUserClient(client owner_user_api.Client) *CustomClient {
	if client == nil {
		c.ownerUser = owner_user_api.New(&owner_user_api.UnimplementedClient{})
		return c
	}
	c.ownerUser = client
	return c
}

func (c *CustomClient) OwnerUser() owner_user_api.Client {
	return c.ownerUser
}
func (c *CustomClient) UseRegistryClient(client registry_api.Client) *CustomClient {
	if client == nil {
		c.registry = registry_api.New(&registry_api.UnimplementedClient{})
		return c
	}
	c.registry = client
	return c
}

func (c *CustomClient) Registry() registry_api.Client {
	return c.registry
}
func (c *CustomClient) UseRegistryApprovedGuestClient(client registry_approved_guest_api.Client) *CustomClient {
	if client == nil {
		c.registryApprovedGuest = registry_approved_guest_api.New(&registry_approved_guest_api.UnimplementedClient{})
		return c
	}
	c.registryApprovedGuest = client
	return c
}

func (c *CustomClient) RegistryApprovedGuest() registry_approved_guest_api.Client {
	return c.registryApprovedGuest
}
func (c *CustomClient) UseRegistryItemClient(client registry_item_api.Client) *CustomClient {
	if client == nil {
		c.registryItem = registry_item_api.New(&registry_item_api.UnimplementedClient{})
		return c
	}
	c.registryItem = client
	return c
}

func (c *CustomClient) RegistryItem() registry_item_api.Client {
	return c.registryItem
}
func (c *CustomClient) UseReservationClient(client reservation_api.Client) *CustomClient {
	if client == nil {
		c.reservation = reservation_api.New(&reservation_api.UnimplementedClient{})
		return c
	}
	c.reservation = client
	return c
}

func (c *CustomClient) Reservation() reservation_api.Client {
	return c.reservation
}
func (c *CustomClient) UseShippingAddressRequestClient(client shipping_address_request_api.Client) *CustomClient {
	if client == nil {
		c.shippingAddressRequest = shipping_address_request_api.New(&shipping_address_request_api.UnimplementedClient{})
		return c
	}
	c.shippingAddressRequest = client
	return c
}

func (c *CustomClient) ShippingAddressRequest() shipping_address_request_api.Client {
	return c.shippingAddressRequest
}
