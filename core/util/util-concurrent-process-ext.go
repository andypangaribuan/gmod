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
	"time"

	"github.com/andypangaribuan/gmod/fm"
)

func (*stuUtil) concurrentProcess(total, max int, fn func(index int)) {
	if total < 1 || max < 1 {
		return
	}

	c := &stuConcurrency{
		active:        0,
		total:         total,
		max:           max,
		fn:            fn,
		sleepDuration: time.Millisecond * 5,
	}

	c.start()
}

func (slf *stuConcurrency) start() {
	n := 0
	for i := 0; i < slf.total; i++ {
		if slf.active >= slf.max {
			for {
				slf.sleep()
				if slf.active < slf.max {
					break
				}
			}
		}

		n++
		slf.addActive(1)
		idx := fm.Ptr(i)
		go slf.execute(*idx)
	}

	for {
		slf.sleep()
		if slf.active == 0 {
			break
		}
	}
}

func (slf *stuConcurrency) execute(index int) {
	slf.fn(index)
	slf.addActive(-1)
}

func (slf *stuConcurrency) addActive(add int) {
	slf.mx.Lock()
	defer slf.mx.Unlock()
	slf.active += add
}

func (slf *stuConcurrency) sleep() {
	time.Sleep(slf.sleepDuration)
}
