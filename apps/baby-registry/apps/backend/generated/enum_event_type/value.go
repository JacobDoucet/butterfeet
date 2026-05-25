package enum_event_type

import (
	"fmt"
)

type Value string

const (
	ItemAdded            Value = "ItemAdded"
	ItemDeleted          Value = "ItemDeleted"
	ItemReserved         Value = "ItemReserved"
	ItemUpdated          Value = "ItemUpdated"
	OwnerLoggedIn        Value = "OwnerLoggedIn"
	RegistryCreated      Value = "RegistryCreated"
	RegistryDeleted      Value = "RegistryDeleted"
	RegistryUpdated      Value = "RegistryUpdated"
	ReservationCancelled Value = "ReservationCancelled"
)

func (v Value) ToString() (string, error) {
	switch v {
	case ItemAdded:
		return string(v), nil
	case ItemDeleted:
		return string(v), nil
	case ItemReserved:
		return string(v), nil
	case ItemUpdated:
		return string(v), nil
	case OwnerLoggedIn:
		return string(v), nil
	case RegistryCreated:
		return string(v), nil
	case RegistryDeleted:
		return string(v), nil
	case RegistryUpdated:
		return string(v), nil
	case ReservationCancelled:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_event_type.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case ItemAdded:
		return nil
	case ItemDeleted:
		return nil
	case ItemReserved:
		return nil
	case ItemUpdated:
		return nil
	case OwnerLoggedIn:
		return nil
	case RegistryCreated:
		return nil
	case RegistryDeleted:
		return nil
	case RegistryUpdated:
		return nil
	case ReservationCancelled:
		return nil
	default:
		return fmt.Errorf("invalid enum_event_type.Value: %s", v)
	}
}
