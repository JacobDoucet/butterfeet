package enum_guest_access_level

import (
	"fmt"
)

type Value string

const (
	ViewShippingAddress Value = "ViewShippingAddress"
	ReserveOnly         Value = "ReserveOnly"
)

func (v Value) ToString() (string, error) {
	switch v {
	case ViewShippingAddress:
		return string(v), nil
	case ReserveOnly:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_guest_access_level.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case ViewShippingAddress:
		return nil
	case ReserveOnly:
		return nil
	default:
		return fmt.Errorf("invalid enum_guest_access_level.Value: %s", v)
	}
}
