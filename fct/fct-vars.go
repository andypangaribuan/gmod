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
	"github.com/shopspring/decimal"
	"golang.org/x/text/message"
)

var (
	Zero FCT
)

var (
	printer  *message.Printer
	emptyFCT FCT
	deciZero = decimal.NewFromInt(0)
)
