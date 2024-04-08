/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package conv

import "time"

const (
	layoutFull = "yyyy-MM-dd HH:mm:ss.SSSSSS TZ"
	layoutDT   = "yyyy-MM-dd HH:mm:ss TZ"
)

func (slf *stuConvTime) ToStrFull(val time.Time) string {
	return slf.toStr(val, layoutFull)
}

func (slf *stuConvTime) ToStrDT(val time.Time) string {
	return slf.toStr(val, layoutDT)
}
