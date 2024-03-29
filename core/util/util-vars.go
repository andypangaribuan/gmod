/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"math/rand"
	"sync"
	"time"

	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceUtil github.com/andypangaribuan/gmod.iceUtil
var iceUtil ice.Util

//go:linkname iceUtilEnv github.com/andypangaribuan/gmod.iceUtilEnv
var iceUtilEnv ice.UtilEnv

const (
	alphabetLower = "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric       = "0123456789"
)

var (
	dvalTimezone      string
	isGetDvalTimezone bool
	timezoneLocking   *sync.Mutex
	timezones         map[string]*time.Location
	xRand             *rand.Rand
	l3uidLength       int
	l3uid             []string
	l3uidN            map[string]string
	l3uidK            map[string]string
)
