package enum_address_request_status

import (
	"fmt"
)

type Value string

const (
	Pending      Value = "Pending"
	Approved     Value = "Approved"
	AutoApproved Value = "AutoApproved"
	Rejected     Value = "Rejected"
	Blocked      Value = "Blocked"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Pending:
		return string(v), nil
	case Approved:
		return string(v), nil
	case AutoApproved:
		return string(v), nil
	case Rejected:
		return string(v), nil
	case Blocked:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_address_request_status.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Pending:
		return nil
	case Approved:
		return nil
	case AutoApproved:
		return nil
	case Rejected:
		return nil
	case Blocked:
		return nil
	default:
		return fmt.Errorf("invalid enum_address_request_status.Value: %s", v)
	}
}
