package utils

import (
	"fmt"
	"math/big"
	"strconv"
)

func String(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func Int(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return i, nil
	default:
		return 0, fmt.Errorf("Unsupported data type: %T", value)
	}
}

func Int64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		parsedValue, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return parsedValue, nil
	default:
		return 0, fmt.Errorf("Unsupported type: %T", value)
	}
}

func Float64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	default:
		return 0, fmt.Errorf("Unsupported type: %T", value)
	}
}

func Boolean(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		return v == "true"
	case int:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	default:
		// 모든 다른 값은 false로 처리
		return false
	}
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func NewBigInt(amount string, decimal string) *big.Int {

	amountBigInt := new(big.Int)
	amountBigInt.SetString(amount, 10)

	decimalBigInt := new(big.Int)
	decimalBigInt.SetString(decimal, 10)

	// amountBigInt를 10^decimalBigInt 만큼 곱하여 큰 정수를 생성합니다.
	result := new(big.Int).Mul(amountBigInt, new(big.Int).Exp(big.NewInt(10), decimalBigInt, nil))
	return result
}
