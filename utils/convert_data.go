package utils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

func ConvertToInt(value interface{}) (uint64, bool) {
	switch v := value.(type) {
	case int:
		return uint64(v), true
	case int8:
		return uint64(v), true
	case int16:
		return uint64(v), true
	case int32:
		return uint64(v), true
	case int64:
		return uint64(v), true
	case uint:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case uint64:
		return v, true
	case float32:
		return uint64(v), true
	case float64:
		return uint64(v), true
	case string:
		i, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			return i, true
		}
		return 0, false
	default:
		return 0, false
	}
}

func ConvertToDecimal(value interface{}) (decimal.Decimal, bool) {
	switch v := value.(type) {
	case int:
		return decimal.NewFromInt(int64(v)), true
	case int8:
		return decimal.NewFromInt(int64(v)), true
	case int16:
		return decimal.NewFromInt(int64(v)), true
	case int32:
		return decimal.NewFromInt(int64(v)), true
	case int64:
		return decimal.NewFromInt(v), true
	case uint:
		return decimal.NewFromInt(int64(v)), true
	case uint8:
		return decimal.NewFromInt(int64(v)), true
	case uint16:
		return decimal.NewFromInt(int64(v)), true
	case uint32:
		return decimal.NewFromInt(int64(v)), true
	case uint64:
		return decimal.NewFromInt(int64(v)), true
	case float32:
		return decimal.NewFromFloat32(v), true
	case float64:
		return decimal.NewFromFloat(v), true
	case string:
		d, err := decimal.NewFromString(v)
		if err == nil {
			return d, true
		}
		return decimal.Zero, false
	default:
		return decimal.Zero, false
	}
}
