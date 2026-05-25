package enum_role

import (
	"fmt"
)

type Value string

const (
	Super  Value = "Super"
	Owner  Value = "Owner"
	Public Value = "Public"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Super:
		return string(v), nil
	case Owner:
		return string(v), nil
	case Public:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_role.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Super:
		return nil
	case Owner:
		return nil
	case Public:
		return nil
	default:
		return fmt.Errorf("invalid enum_role.Value: %s", v)
	}
}
