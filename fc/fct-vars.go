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
	"errors"

	"github.com/shopspring/decimal"
)

var (
	Zero FCT
)

var (
	deciZero         = decimal.NewFromInt(0)
	errCastNil       = errors.New("value cannot nil")
	errCastUncovered = errors.New("uncovered the value casting")
)
