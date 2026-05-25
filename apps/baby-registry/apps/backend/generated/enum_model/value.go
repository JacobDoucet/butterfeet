package enum_model

import (
	"fmt"
)

type Value string

const (
	AddressAccessSession   Value = "AddressAccessSession"
	OwnerUser              Value = "OwnerUser"
	Registry               Value = "Registry"
	RegistryApprovedGuest  Value = "RegistryApprovedGuest"
	RegistryItem           Value = "RegistryItem"
	Reservation            Value = "Reservation"
	ShippingAddressRequest Value = "ShippingAddressRequest"
)

func (v Value) ToString() (string, error) {
	switch v {
	case AddressAccessSession:
		return string(v), nil
	case OwnerUser:
		return string(v), nil
	case Registry:
		return string(v), nil
	case RegistryApprovedGuest:
		return string(v), nil
	case RegistryItem:
		return string(v), nil
	case Reservation:
		return string(v), nil
	case ShippingAddressRequest:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_model.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case AddressAccessSession:
		return nil
	case OwnerUser:
		return nil
	case Registry:
		return nil
	case RegistryApprovedGuest:
		return nil
	case RegistryItem:
		return nil
	case Reservation:
		return nil
	case ShippingAddressRequest:
		return nil
	default:
		return fmt.Errorf("invalid enum_model.Value: %s", v)
	}
}
