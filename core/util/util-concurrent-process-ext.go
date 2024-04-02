/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"time"

	"github.com/andypangaribuan/gmod/fm"
)

func (*srUtil) concurrentProcess(total, max int, fn func(index int)) {
	c := &srConcurrency{
		active:        0,
		total:         total,
		max:           max,
		fn:            fn,
		sleepDuration: time.Millisecond * 10,
	}

	c.start()
}

func (slf *srConcurrency) start() {
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

func (slf *srConcurrency) execute(index int) {
	slf.fn(index)
	slf.addActive(-1)
}

func (slf *srConcurrency) addActive(add int) {
	slf.mx.Lock()
	defer slf.mx.Unlock()
	slf.active += add
}

func (slf *srConcurrency) sleep() {
	time.Sleep(slf.sleepDuration)
}
