package http_server

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session_http"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/event_http"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_http"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest_http"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_http"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_http"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_http"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request_http"
	"net/http"
)

type ServeMuxProps struct {
	ResolveActor                        func(r *http.Request) (permissions.Actor, error)
	AddressAccessSessionMetadataHooks   []address_access_session_http.MetadataHooks
	EventMetadataHooks                  []event_http.MetadataHooks
	OwnerUserMetadataHooks              []owner_user_http.MetadataHooks
	RegistryMetadataHooks               []registry_http.MetadataHooks
	RegistryApprovedGuestMetadataHooks  []registry_approved_guest_http.MetadataHooks
	RegistryItemMetadataHooks           []registry_item_http.MetadataHooks
	ReservationMetadataHooks            []reservation_http.MetadataHooks
	ShippingAddressRequestMetadataHooks []shipping_address_request_http.MetadataHooks
	OnError                             func(handler string, e error)
}

func ServeMux(client api.Client, props ServeMuxProps) (*http.ServeMux, error) {
	serveMux := http.NewServeMux()

	addressAccessSessionApi := client.AddressAccessSession()
	addressAccessSessionServeMux, err := address_access_session_http.RegisterRoutes(address_access_session_http.HandlerProps{
		Api:           addressAccessSessionApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.AddressAccessSessionMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/address-access-sessions/", http.StripPrefix("/address-access-sessions", addressAccessSessionServeMux))

	eventApi := client.Event()
	eventServeMux, err := event_http.RegisterRoutes(event_http.HandlerProps{
		Api:           eventApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.EventMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/events/", http.StripPrefix("/events", eventServeMux))

	ownerUserApi := client.OwnerUser()
	ownerUserServeMux, err := owner_user_http.RegisterRoutes(owner_user_http.HandlerProps{
		Api:           ownerUserApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.OwnerUserMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/owner-users/", http.StripPrefix("/owner-users", ownerUserServeMux))

	registryApi := client.Registry()
	registryServeMux, err := registry_http.RegisterRoutes(registry_http.HandlerProps{
		Api:           registryApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.RegistryMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/registries/", http.StripPrefix("/registries", registryServeMux))

	registryApprovedGuestApi := client.RegistryApprovedGuest()
	registryApprovedGuestServeMux, err := registry_approved_guest_http.RegisterRoutes(registry_approved_guest_http.HandlerProps{
		Api:           registryApprovedGuestApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.RegistryApprovedGuestMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/registry-approved-guests/", http.StripPrefix("/registry-approved-guests", registryApprovedGuestServeMux))

	registryItemApi := client.RegistryItem()
	registryItemServeMux, err := registry_item_http.RegisterRoutes(registry_item_http.HandlerProps{
		Api:           registryItemApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.RegistryItemMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/registry-items/", http.StripPrefix("/registry-items", registryItemServeMux))

	reservationApi := client.Reservation()
	reservationServeMux, err := reservation_http.RegisterRoutes(reservation_http.HandlerProps{
		Api:           reservationApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.ReservationMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/reservations/", http.StripPrefix("/reservations", reservationServeMux))

	shippingAddressRequestApi := client.ShippingAddressRequest()
	shippingAddressRequestServeMux, err := shipping_address_request_http.RegisterRoutes(shipping_address_request_http.HandlerProps{
		Api:           shippingAddressRequestApi,
		ResolveActor:  props.ResolveActor,
		MetadataHooks: props.ShippingAddressRequestMetadataHooks,
		OnError:       props.OnError,
	})
	if err != nil {
		return nil, err
	}
	serveMux.Handle("/shipping-address-requests/", http.StripPrefix("/shipping-address-requests", shippingAddressRequestServeMux))

	return serveMux, nil
}
