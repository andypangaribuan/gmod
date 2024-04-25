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

	"github.com/shopspring/decimal"
)

func (slf *FCT) panic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (slf *FCT) set(deci decimal.Decimal) *FCT {
	slf.deci = deci
	slf.v1, slf.v2 = getString(deci)

	return slf
}
