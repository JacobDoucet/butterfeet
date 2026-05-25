package enum_address_access_mode

import (
	"fmt"
)

type Value string

const (
	RequestApproval    Value = "RequestApproval"
	ApprovedGuestsOnly Value = "ApprovedGuestsOnly"
	Disabled           Value = "Disabled"
)

func (v Value) ToString() (string, error) {
	switch v {
	case RequestApproval:
		return string(v), nil
	case ApprovedGuestsOnly:
		return string(v), nil
	case Disabled:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_address_access_mode.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case RequestApproval:
		return nil
	case ApprovedGuestsOnly:
		return nil
	case Disabled:
		return nil
	default:
		return fmt.Errorf("invalid enum_address_access_mode.Value: %s", v)
	}
}
