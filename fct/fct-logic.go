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
	"strconv"

	"github.com/shopspring/decimal"
)

func (slf *FCT) panic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (slf *FCT) set(deci decimal.Decimal) *FCT {
	exp := int(deci.Exponent())
	if exp < 0 {
		exp *= -1
	}

	if exp < 1 {
		exp = 1
	}

	format := "%." + strconv.Itoa(exp) + "f"

	slf.deci = deci
	slf.v1 = fmt.Sprintf(format, deci.InexactFloat64())
	slf.v2 = printer.Sprintf(format, deci.InexactFloat64())

	return slf
}
