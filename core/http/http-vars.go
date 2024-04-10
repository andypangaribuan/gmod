/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package http

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceHttp github.com/andypangaribuan/gmod.iceHttp
var iceHttp ice.Http