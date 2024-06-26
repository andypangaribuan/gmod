/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import "time"

type ConvTime interface {
	ToStrFull(val time.Time) string
	ToStrDT(val time.Time) string

	ToTimeFull(val string) (*time.Time, error)
	ToTimeDT(val string) (*time.Time, error)
	ToTimeD(val string) (*time.Time, error)
	ToTime(val string, layout string) (*time.Time, error)
}
