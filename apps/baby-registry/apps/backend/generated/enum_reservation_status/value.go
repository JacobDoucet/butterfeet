package enum_reservation_status

import (
	"fmt"
)

type Value string

const (
	Reserved  Value = "Reserved"
	Purchased Value = "Purchased"
	Received  Value = "Received"
	Cancelled Value = "Cancelled"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Reserved:
		return string(v), nil
	case Purchased:
		return string(v), nil
	case Received:
		return string(v), nil
	case Cancelled:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_reservation_status.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Reserved:
		return nil
	case Purchased:
		return nil
	case Received:
		return nil
	case Cancelled:
		return nil
	default:
		return fmt.Errorf("invalid enum_reservation_status.Value: %s", v)
	}
}
