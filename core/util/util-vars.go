/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"sync"
	"time"

	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceUtil github.com/andypangaribuan/gmod.iceUtil
var iceUtil ice.Util

//go:linkname iceUtilEnv github.com/andypangaribuan/gmod.iceUtilEnv
var iceUtilEnv ice.UtilEnv

var (
	dvalTimezone      string
	isGetDvalTimezone bool
	timezoneLocking   *sync.Mutex
	timezones         map[string]*time.Location
)
