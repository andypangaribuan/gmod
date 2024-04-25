/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import (
	_ "unsafe"

	"github.com/andypangaribuan/gmod/ice"
)

var orderedInitLs [][]any

//go:linkname mainFmIceNet github.com/andypangaribuan/gmod.mainFmIceNet
var mainFmIceNet func() ice.Net