/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import "time"

type Lock interface {
	NewOpt() LockOpt

	// code
	// -1: not obtained
	//  0: have an error
	//  1: locked
	Tx(id string, opt ...LockOpt) (code int, err error)
}

type LockOpt interface {
	SetTimeout(duration time.Duration) LockOpt
	TryFor(duration time.Duration) LockOpt
	SetPrefix(prefix string) LockOpt
}

type LockInstance interface {
	Release()
}