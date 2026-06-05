package enum_payment_method_type

import (
	"fmt"
)

type Value string

const (
	PayPal           Value = "PayPal"
	Revolut          Value = "Revolut"
	Wise             Value = "Wise"
	InteracETransfer Value = "InteracETransfer"
	BankTransfer     Value = "BankTransfer"
	Other            Value = "Other"
)

func (v Value) ToString() (string, error) {
	switch v {
	case PayPal:
		return string(v), nil
	case Revolut:
		return string(v), nil
	case Wise:
		return string(v), nil
	case InteracETransfer:
		return string(v), nil
	case BankTransfer:
		return string(v), nil
	case Other:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_payment_method_type.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case PayPal:
		return nil
	case Revolut:
		return nil
	case Wise:
		return nil
	case InteracETransfer:
		return nil
	case BankTransfer:
		return nil
	case Other:
		return nil
	default:
		return fmt.Errorf("invalid enum_payment_method_type.Value: %s", v)
	}
}
