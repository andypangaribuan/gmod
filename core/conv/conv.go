/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package conv

func init() {
	var (
		conv     = new(srConv)
		convTime = new(srConvTime)
	)

	iceConv = conv
	iceConvTime = convTime
}
