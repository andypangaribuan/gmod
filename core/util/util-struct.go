/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"sync"
	"time"
)

type srUtil struct{}

type srUtilEnv struct{}

type srUtilEnvAppEnv struct {
	val string
}

type srConcurrency struct {
	mx            sync.Mutex
	max           int
	total         int
	active        int
	fn            func(index int)
	sleepDuration time.Duration
}
