/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/andypangaribuan/gmod/gm"
)

func xinit() {
	timezoneLocking = &sync.Mutex{}
	timezones = make(map[string]*time.Location, 0)

	val, err := iceUtil.ReflectionGet(gm.Conf, "timeZone")
	if err == nil {
		val, _ := val.(string)
		dvalTimezone = val
	}

	xRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	l3uid, l3uidN, l3uidK = l3uidGenerate()
	l3uidCheckUnique(l3uid)
	l3uidLength = len(fmt.Sprint(len(l3uid)))
}
