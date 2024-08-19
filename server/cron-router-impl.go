/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"time"

	"github.com/andypangaribuan/gmod/gm"
)

func (slf *stuCronRouter) Every(duration string, fn func(), startUpDelayed *time.Duration, allowParallel ...bool) {
	withParallel := false
	if len(allowParallel) > 0 {
		withParallel = allowParallel[0]
	}

	slf.everyItems = append(slf.everyItems, &stuCronEveryItem{
		uid:            gm.Util.UID(50),
		duration:       duration,
		fn:             fn,
		startUpDelayed: startUpDelayed,
		allowParallel:  withParallel,
	})
}

func (slf *stuCronRouter) NEvery(duration string, fns []func(), startUpDelayed *time.Duration, allowParallel ...bool) {
	withParallel := false
	if len(allowParallel) > 0 {
		withParallel = allowParallel[0]
	}

	slf.everyNItems = append(slf.everyNItems, &stuCronNEveryItem{
		uid:            gm.Util.UID(50),
		duration:       duration,
		fns:            fns,
		startUpDelayed: startUpDelayed,
		allowParallel:  withParallel,
	})
}

func (slf *stuCronRouter) EveryDay(at string, fn func(), allowParallel ...bool) {
	withParallel := false
	if len(allowParallel) > 0 {
		withParallel = allowParallel[0]
	}

	slf.everyDayItems = append(slf.everyDayItems, &stuCronEveryDayItem{
		uid:           gm.Util.UID(50),
		at:            at,
		fn:            fn,
		allowParallel: withParallel,
	})
}

func (slf *stuCronRouter) NEveryDay(at string, fns []func(), allowParallel ...bool) {
	withParallel := false
	if len(allowParallel) > 0 {
		withParallel = allowParallel[0]
	}

	slf.everyDayNItems = append(slf.everyDayNItems, &stuCronNEveryDayItem{
		uid:           gm.Util.UID(50),
		at:            at,
		fns:           fns,
		allowParallel: withParallel,
	})
}
