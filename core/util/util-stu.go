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
	"sync"
	// "time"
)

type stuUtil struct {
	initFuncMx         sync.Mutex
	initFuncExecuteMap map[string]any
}

type stuUtilEnv struct{}

type stuUtilEnvAppEnv struct {
	val string
}

// type stuConcurrency struct {
// 	mx            sync.Mutex
// 	max           int
// 	total         int
// 	active        int
// 	fn            func(index int)
// 	sleepDuration time.Duration
// }

type stuXConcurrency struct {
	hasInit       bool
	maxConcurrent int
	job           chan int
	waiter        chan int
	callback      func(index int)
}

type stuEnvBase64 struct {
	key  string
	data []byte
}
