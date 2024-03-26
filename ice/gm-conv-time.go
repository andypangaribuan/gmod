/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package ice

import "time"

type ConvTime interface {
	ToStrFull(val time.Time) string
	ToStrDT(val time.Time) string
}
