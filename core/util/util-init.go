/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package util

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func xinit() {
	timezoneLocking = &sync.Mutex{}
	timezones = make(map[string]*time.Location, 0)

	mainUtilCallback = func() {
		iceUtil.Timenow()
	}

	xRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	l3uid, l3uidN, l3uidK = l3uidGenerate()
	l3uidCheckUnique(l3uid)
	l3uidLength = len(fmt.Sprint(len(l3uid)))
}
