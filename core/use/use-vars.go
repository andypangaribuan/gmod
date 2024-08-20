/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package use

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceUse github.com/andypangaribuan/gmod.iceUse
var iceUse ice.Use

//go:linkname iceUseGcs github.com/andypangaribuan/gmod.iceUseGcs
var iceUseGcs ice.UseGcs
