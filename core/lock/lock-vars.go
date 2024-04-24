/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import (
	"time"

	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceLock github.com/andypangaribuan/gmod.iceLock
var iceLock ice.Lock

var (
	dvalTxTimeout = time.Second * 3
)
