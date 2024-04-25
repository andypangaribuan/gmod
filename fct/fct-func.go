/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fct

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func create(deci decimal.Decimal) *FCT {
	return new(FCT).set(deci)
}

func dvalFCT(dval ...FCT) *FCT {
	if len(dval) > 0 {
		return &dval[0]
	}

	return nil
}

func getFCT(val *FCT, dval ...FCT) (*FCT, error) {
	if val == nil {
		if len(dval) > 0 {
			return &dval[0], nil
		}

		return nil, errors.New("fct is nil")
	}

	return val, nil
}

func convert(value any) (*decimal.Decimal, error) {
	switch val := value.(type) {
	case string:
		v, err := decimal.NewFromString(val)
		if err != nil {
			return nil, errors.Wrap(err, "fct converter: invalid value")
		}

		return &v, nil

	case *string:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		v, err := decimal.NewFromString(*val)
		if err != nil {
			return nil, errors.Wrap(err, "fct converter: invalid value")
		}

		return &v, nil

	case int:
		v := decimal.NewFromInt(int64(val))
		return &v, nil

	case *int:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		v := decimal.NewFromInt(int64(*val))
		return &v, nil

	case int32:
		v := decimal.NewFromInt32(val)
		return &v, nil

	case *int32:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		v := decimal.NewFromInt32(*val)
		return &v, nil

	case int64:
		v := decimal.NewFromInt(val)
		return &v, nil

	case *int64:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		v := decimal.NewFromInt(*val)
		return &v, nil

	case float32:
		v := decimal.NewFromFloat32(val)
		return &v, nil

	case *float32:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		v := decimal.NewFromFloat32(*val)
		return &v, nil

	case float64:
		v := decimal.NewFromFloat(val)
		return &v, nil

	case *float64:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		v := decimal.NewFromFloat(*val)
		return &v, nil

	case decimal.Decimal:
		return &val, nil

	case *decimal.Decimal:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		return val, nil

	case FCT:
		return &val.deci, nil

	case *FCT:
		if val == nil {
			return nil, errors.New("fct converter: value cannot nil")
		}

		return &val.deci, nil
	}

	return nil, errors.New("fct converter: unhandled value conversion")
}
