package enum_payment_method_type

import (
	"fmt"
)

type Value string

const (
	PaymentLink      Value = "PaymentLink"
	InteracETransfer Value = "InteracETransfer"
	BankTransfer     Value = "BankTransfer"
)

func (v Value) ToString() (string, error) {
	switch v {
	case PaymentLink:
		return string(v), nil
	case InteracETransfer:
		return string(v), nil
	case BankTransfer:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_payment_method_type.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case PaymentLink:
		return nil
	case InteracETransfer:
		return nil
	case BankTransfer:
		return nil
	default:
		return fmt.Errorf("invalid enum_payment_method_type.Value: %s", v)
	}
}
