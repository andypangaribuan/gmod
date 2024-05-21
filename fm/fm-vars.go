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

	"github.com/andypangaribuan/gmod/clog"
)

var orderedInitLs [][]any

//go:linkname mainReflection github.com/andypangaribuan/gmod.mainReflection
var mainReflection func(key string, arg ...any) []any

//go:linkname clogGetId github.com/andypangaribuan/gmod/clog.clogGetId
var clogGetId func(clog clog.Instance) (string, *string, *string)
