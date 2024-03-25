/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"sync"

	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceDb github.com/andypangaribuan/gmod.iceDb
var iceDb ice.Db

var (
	connWriteLocking sync.Mutex
	connReadLocking  sync.Mutex
)
