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
	"fmt"
	"log"
	"math/big"
	"runtime/debug"
	"strings"

	"github.com/shopspring/decimal"
)

func (slf *FCT) get(dval FCT) FCT {
	if slf != nil {
		return *slf
	}
	return dval
}

func (slf *FCT) float64(dval ...FCT) (float64, error) {
	fv, err := getFCT(slf, dval...)
	if err != nil {
		return 0, err
	}

	return fv.deci.InexactFloat64(), nil
}

func (slf *FCT) ptrRound(places int, dval ...FCT) (*FCT, error) {
	fv, err := getFCT(slf, dval...)
	if err != nil {
		return nil, err
	}

	return create(fv.deci.Round(int32(places))), nil
}

func (slf *FCT) ptrFloor(places int, dval ...FCT) (*FCT, error) {
	fv, err := getFCT(slf, dval...)
	if err != nil {
		return nil, err
	}

	if places < 1 {
		return create(fv.deci.Floor()), nil
	}

	exp := fv.deci.Exponent()
	if exp < 0 {
		exp *= -1
		if exp > int32(places) {
			var (
				sub                = int(exp) - places
				div                = "1"
				thousandDivDecimal = big.NewInt(1)
			)

			for i := 0; i < sub; i++ {
				div = fmt.Sprintf("%v0", div)
				v, ok := new(big.Int).SetString(div, 10)
				if !ok {
					debug.PrintStack()
					log.Panicf("error when converting to big.int, value: %v\n", div)
				}

				thousandDivDecimal = v
			}

			currentValue := fv.deci.Coefficient()
			newValue := new(big.Int).Div(currentValue, thousandDivDecimal)
			deci := decimal.NewFromBigInt(newValue, int32(places*-1))
			return create(deci), nil
		}
	}

	return slf, nil
}

func (slf *FCT) ptrCeil(places int, dval ...FCT) (*FCT, error) {
	fv, err := getFCT(slf, dval...)
	if err != nil {
		return nil, err
	}

	if places < 1 {
		return create(fv.deci.Ceil()), nil
	}

	exp := fv.deci.Exponent()
	if exp < 0 {
		exp *= -1
		if exp > int32(places) {
			var (
				sub                = int(exp) - places
				div                = "1"
				thousandDivDecimal = big.NewInt(1)
			)

			for i := 0; i < sub; i++ {
				div = fmt.Sprintf("%v0", div)
				v, ok := new(big.Int).SetString(div, 10)
				if !ok {
					debug.PrintStack()
					log.Panicf("error when converting to big.int, value: %v\n", div)
				}

				thousandDivDecimal = v
			}

			currentValue := fv.deci.Coefficient()
			newValue := new(big.Int).Div(currentValue, thousandDivDecimal)
			newValue = new(big.Int).Add(newValue, big.NewInt(1))
			deci := decimal.NewFromBigInt(newValue, int32(places*-1))
			return create(deci), nil
		}
	}

	return slf, nil
}

func (slf *FCT) ptrTruncate(places int, dval ...FCT) (*FCT, error) {
	fv, err := getFCT(slf, dval...)
	if err != nil {
		return nil, err
	}

	deci := fv.deci.Truncate(int32(places))
	return create(deci), nil
}

func (slf *FCT) ptrPow(val any, dval ...FCT) (*FCT, error) {
	v, err := New(val)
	if err != nil {
		return nil, err
	}

	fv, err := getFCT(slf, dval...)
	if err != nil {
		return nil, err
	}

	deci := fv.deci.Pow(v.deci)
	return create(deci), nil
}

func (slf *FCT) toString() (string, error) {
	fv, err := getFCT(slf)
	if err != nil {
		return "", err
	}

	v1, _ := getString(fv.deci)
	return v1, nil
}

func (slf *FCT) toStringTrim() (string, error) {
	fv, err := getFCT(slf)
	if err != nil {
		return "", err
	}

	v1, _ := getString(fv.deci)
	ls := strings.Split(v1, ".")
	if len(ls) > 1 {
		var (
			intValue     = ls[0]
			decimalValue = ls[1]
		)

		for {
			lastDigit := decimalValue[len(decimalValue)-1:]
			if len(decimalValue) == 1 || lastDigit != "0" {
				break
			}

			decimalValue = decimalValue[:len(decimalValue)-1]
		}

		v1 = fmt.Sprintf("%v.%v", intValue, decimalValue)
	}

	return v1, nil
}
