/*
* Copyright (c) 2024.
* Created by Andy Pangaribuan <https://github.com/apangaribuan>.
* All Rights Reserved.
 */

package util

import (
	"sync"
	"time"

	"github.com/andypangaribuan/gmod/gm"
)

func init() {
	util := new(srUtil)
	utilEnv := new(srUtilEnv)

	iceUtil = util
	iceUtilEnv = utilEnv

	adv()
}

func adv() {
	timezoneLocking = &sync.Mutex{}
	timezones = make(map[string]*time.Location, 0)

	val, err := iceUtil.ReflectionGet(gm.Conf, "timeZone")
	if err == nil {
		val, _ := val.(string)
		dvalTimezone = val
	}
}
