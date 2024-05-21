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
	"strconv"
	"strings"

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

func getString(deci decimal.Decimal) (string, string) {
	exp := int(deci.Exponent())
	if exp < 0 {
		exp *= -1
	}

	if exp == 0 {
		exp = 1
	}

	v1 := deci.StringFixedBank(int32(exp))
	v2 := printer.Sprintf("%."+strconv.Itoa(exp)+"f", deci.InexactFloat64())

	ls := strings.Split(v2, ".")
	if len(ls) > 1 {
		decimalValue := ls[1]
		for {
			if len(decimalValue) == 0 {
				break
			}

			if decimalValue[len(decimalValue)-1:] == "0" {
				decimalValue = decimalValue[:len(decimalValue)-1]
			} else {
				break
			}
		}

		if len(decimalValue) > 0 {
			v2 = ls[0] + "." + decimalValue
		} else {
			v2 = ls[0] + ".0"
		}
	}

	return v1, v2
}

func ifHaveIn[T comparable](val T, in ...T) bool {
	for _, v := range in {
		if val == v {
			return true
		}
	}

	return false
}

func isOperator(val any) (string, bool) {
	switch v := val.(type) {
	case string:
		if v == "+" || v == "-" || v == "*" || v == "/" || v == "%" {
			return v, true
		}
	}

	return "", false
}

func removeIndex[T any](ls []T, index int) []T {
	return append(ls[:index], ls[index+1:]...)
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
