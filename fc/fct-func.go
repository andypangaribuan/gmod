/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fc

import (
	"github.com/shopspring/decimal"
)

func safeNew(value any) (*decimal.Decimal, error) {
	switch val := value.(type) {
	case string:
		return safeToDecimal(&val)

	case *string:
		return safeToDecimal(val)

	case int:
		return safeToDecimal(&val)

	case *int:
		return safeToDecimal(val)

	case int32:
		return safeToDecimal(&val)

	case *int32:
		return safeToDecimal(val)

	case int64:
		return safeToDecimal(&val)

	case *int64:
		return safeToDecimal(val)

	case float32:
		return safeToDecimal(&val)

	case *float32:
		return safeToDecimal(val)

	case float64:
		return safeToDecimal(&val)

	case *float64:
		return safeToDecimal(val)

	case decimal.Decimal:
		return &val, nil

	case *decimal.Decimal:
		if val == nil {
			return nil, errCastNil
		}

		return val, nil

	case FCT:
		return &val.vd, nil

	case *FCT:
		if val == nil {
			return nil, errCastNil
		}

		return &val.vd, nil
	}

	return nil, errCastUncovered
}

func safeToDecimal[T castType](value *T) (*decimal.Decimal, error) {
	switch val := any(value).(type) {
	case *string:
		if val == nil {
			return nil, errCastNil
		}

		v, err := decimal.NewFromString(*val)
		if err != nil {
			return nil, err
		}

		return &v, nil

	case *int:
		if val == nil {
			return nil, errCastNil
		}

		v64 := int64(*val)
		return safeToDecimal(&v64)

	case *int32:
		if val == nil {
			return nil, errCastNil
		}

		v := decimal.NewFromInt32(*val)
		return &v, nil

	case *int64:
		if val == nil {
			return nil, errCastNil
		}

		v := decimal.NewFromInt(*val)
		return &v, nil

	case *float32:
		if val == nil {
			return nil, errCastNil
		}

		v := decimal.NewFromFloat32(*val)
		return &v, nil

	case *float64:
		if val == nil {
			return nil, errCastNil
		}

		v := decimal.NewFromFloat(*val)
		return &v, nil
	}

	return nil, errCastUncovered
}
