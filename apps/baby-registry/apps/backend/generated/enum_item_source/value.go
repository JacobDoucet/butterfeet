package enum_item_source

import (
	"fmt"
)

type Value string

const (
	Amazon        Value = "Amazon"
	MamasAndPapas Value = "MamasAndPapas"
	Etsy          Value = "Etsy"
	JohnLewis     Value = "JohnLewis"
	IKEA          Value = "IKEA"
	Other         Value = "Other"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Amazon:
		return string(v), nil
	case MamasAndPapas:
		return string(v), nil
	case Etsy:
		return string(v), nil
	case JohnLewis:
		return string(v), nil
	case IKEA:
		return string(v), nil
	case Other:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_item_source.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Amazon:
		return nil
	case MamasAndPapas:
		return nil
	case Etsy:
		return nil
	case JohnLewis:
		return nil
	case IKEA:
		return nil
	case Other:
		return nil
	default:
		return fmt.Errorf("invalid enum_item_source.Value: %s", v)
	}
}
