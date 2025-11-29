/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import (
	"time"

	"github.com/andypangaribuan/gmod/clog"
)

type Lock interface {
	NewOpt() LockOpt
	Tx(logc clog.Instance, key string, opt ...LockOpt) (LockInstance, error)
}

type LockOpt interface {
	SetTimeout(duration time.Duration) LockOpt
	TryFor(duration time.Duration) LockOpt
	SetPrefix(prefix string) LockOpt
}

type LockInstance interface {
	Release()
	IsLocked() (bool, error)
	Extend(duration time.Duration) error
}
