package enum_guest_status

import (
	"fmt"
)

type Value string

const (
	Pending Value = "Pending"
	Active  Value = "Active"
	Revoked Value = "Revoked"
	Blocked Value = "Blocked"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Pending:
		return string(v), nil
	case Active:
		return string(v), nil
	case Revoked:
		return string(v), nil
	case Blocked:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_guest_status.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Pending:
		return nil
	case Active:
		return nil
	case Revoked:
		return nil
	case Blocked:
		return nil
	default:
		return fmt.Errorf("invalid enum_guest_status.Value: %s", v)
	}
}
