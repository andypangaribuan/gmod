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
	ToStrDateTime(val time.Time) string
	ToStrDate(val time.Time) string
	// layout using "yyyy-MM-dd HH:mm:ss.SSSSSS TZ"
	ToStr(val time.Time, layout string) string

	ToTimeFull(val string) (*time.Time, error)
	ToTimeDateTime(val string) (*time.Time, error)
	ToTimeDate(val string) (*time.Time, error)
	// layout using "yyyy-MM-dd HH:mm:ss.SSSSSS TZ"
	ToTime(val string, layout string) (*time.Time, error)
}
