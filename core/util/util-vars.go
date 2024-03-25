/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceUtil github.com/andypangaribuan/gmod.iceUtil
var iceUtil ice.Util

//go:linkname iceUtilEnv github.com/andypangaribuan/gmod.iceUtilEnv
var iceUtilEnv ice.UtilEnv
