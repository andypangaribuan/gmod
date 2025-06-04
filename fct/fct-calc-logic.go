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
	"log"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func unsafeCalc(val ...any) FCT {
	f, err := calc(val...)
	if err != nil {
		log.Panicf("fct calc: found some error\n%+v", err)
	}
	return f
}

// supported operator: +, -, *, /, %
func calc(val ...any) (FCT, error) {
	fv := FCT{
		v1: "0",
		v2: "0",
	}

	length := len(val)

	if length%2 == 0 || length < 2 {
		return fv, errors.New("fct calc: wrong implementation")
	}

	lsv := make([]any, 0)

	for i := range length {
		if i%2 == 0 {
			vd, err := convert(val[i])
			if err != nil {
				return fv, err
			}

			lsv = append(lsv, vd)
		} else {
			operator, ok := isOperator(val[i])
			if !ok {
				return fv, errors.New("fct calc: invalid operator")
			}

			lsv = append(lsv, operator)
		}
	}

	for i := 0; i < len(lsv); i++ {
		if i%2 == 0 && i < len(lsv)-1 {
			operator := lsv[i+1].(string)

			if operator == "*" || operator == "/" || operator == "%" {
				vd1 := lsv[i].(*decimal.Decimal)
				vd2 := lsv[i+2].(*decimal.Decimal)

				switch operator {
				case "*":
					val := vd1.Mul(*vd2)
					lsv[i] = &val

				case "/":
					val := vd1.Div(*vd2)
					lsv[i] = &val

				case "%":
					val := vd1.Mod(*vd2)
					lsv[i] = &val
				}

				lsv = removeIndex(lsv, i+2)
				lsv = removeIndex(lsv, i+1)
				i--
			}
		}
	}

	for i := 0; i < len(lsv); i++ {
		if i%2 == 0 && i < len(lsv)-1 {
			operator := lsv[i+1].(string)

			if operator == "+" || operator == "-" {
				vd1 := lsv[i].(*decimal.Decimal)
				vd2 := lsv[i+2].(*decimal.Decimal)

				switch operator {
				case "+":
					val := vd1.Add(*vd2)
					lsv[i] = &val

				case "-":
					val := vd1.Sub(*vd2)
					lsv[i] = &val
				}

				lsv = removeIndex(lsv, i+2)
				lsv = removeIndex(lsv, i+1)
				i--
			}
		}
	}

	if len(lsv) != 1 {
		return fv, errors.New("fct calc: something went wrong")
	}

	finalVal := lsv[0].(*decimal.Decimal)
	fv.set(*finalVal)
	return fv, nil
}
