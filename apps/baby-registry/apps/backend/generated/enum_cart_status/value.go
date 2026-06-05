package enum_cart_status

import (
	"fmt"
)

type Value string

const (
	Pending              Value = "Pending"
	AwaitingConfirmation Value = "AwaitingConfirmation"
	Completed            Value = "Completed"
	Rejected             Value = "Rejected"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Pending:
		return string(v), nil
	case AwaitingConfirmation:
		return string(v), nil
	case Completed:
		return string(v), nil
	case Rejected:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_cart_status.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Pending:
		return nil
	case AwaitingConfirmation:
		return nil
	case Completed:
		return nil
	case Rejected:
		return nil
	default:
		return fmt.Errorf("invalid enum_cart_status.Value: %s", v)
	}
}
