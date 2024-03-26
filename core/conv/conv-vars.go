/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package conv

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceConv github.com/andypangaribuan/gmod.iceConv
var iceConv ice.Conv

//go:linkname iceConvTime github.com/andypangaribuan/gmod.iceConvTime
var iceConvTime ice.ConvTime
