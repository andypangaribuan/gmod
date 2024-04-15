/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package json

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceJson github.com/andypangaribuan/gmod.iceJson
var iceJson ice.Json

//go:linkname mainJsonCallback github.com/andypangaribuan/gmod.mainJsonCallback
var mainJsonCallback func()
