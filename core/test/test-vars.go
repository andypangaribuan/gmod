/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceTest github.com/andypangaribuan/gmod.iceTest
var iceTest ice.Test
