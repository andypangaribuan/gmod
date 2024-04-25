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
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func (slf *FCT) set(vd decimal.Decimal) {
	exp := int(vd.Exponent())
	if exp < 0 {
		exp *= -1
	}

	if exp < 1 {
		exp = 1
	}

	format := "%." + strconv.Itoa(exp) + "f"

	slf.vd = vd
	slf.V1 = fmt.Sprintf(format, slf.vd.InexactFloat64())
	slf.V2 = printer.Sprintf(format, slf.vd.InexactFloat64())
}
