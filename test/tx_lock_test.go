/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/andypangaribuan/gmod/gm"
)

func TestTxLock(t *testing.T) {
	gm.Test.Start(t, testTxLock)
}

func testTxLock(t *testing.T) {
	var (
		wg        = new(sync.WaitGroup)
		startedAt = time.Now().Add(time.Second * 3)
	)

	for i := range 100 {
		wg.Add(1)
		// key := fmt.Sprint(i) // all unique
		key := fmt.Sprint(i % 10) // only 10 unique key
		go txLockIns(t, wg, startedAt, key)
	}

	wg.Wait()
}

func txLockIns(t *testing.T, wg *sync.WaitGroup, startedAt time.Time, key string) {
	defer wg.Done()

	for {
		if time.Now().After(startedAt) {
			break
		} else {
			time.Sleep(time.Millisecond)
		}
	}

	lock, err := gm.Lock.Tx(nil, key)
	defer lock.Release()

	if err != nil {
		gm.Test.Printf(t, "%v: have an error when get the lock\n", key)
	} else {
		time.Sleep(time.Millisecond * 3100)
		gm.Test.Printf(t, "%v: done\n", key)
	}
}
