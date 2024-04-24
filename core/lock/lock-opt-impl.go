/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import (
	"time"

	"github.com/andypangaribuan/gmod/ice"
)

func (slf *stuLockOpt) SetTimeout(duration time.Duration) ice.LockOpt {
	slf.timeout = &duration
	return slf
}

func (slf *stuLockOpt) TryFor(duration time.Duration) ice.LockOpt {
	slf.tryFor = &duration
	return slf
}

func (slf *stuLockOpt) SetPrefix(prefix string) ice.LockOpt {
	slf.prefix = &prefix
	return slf
}
