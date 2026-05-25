package enum_model

import (
	"fmt"
)

type Value string

const (
	OwnerUser    Value = "OwnerUser"
	Registry     Value = "Registry"
	RegistryItem Value = "RegistryItem"
	Reservation  Value = "Reservation"
)

func (v Value) ToString() (string, error) {
	switch v {
	case OwnerUser:
		return string(v), nil
	case Registry:
		return string(v), nil
	case RegistryItem:
		return string(v), nil
	case Reservation:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_model.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case OwnerUser:
		return nil
	case Registry:
		return nil
	case RegistryItem:
		return nil
	case Reservation:
		return nil
	default:
		return fmt.Errorf("invalid enum_model.Value: %s", v)
	}
}
