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
	"context"
	"time"

	"github.com/bsm/redislock"
)

type stuLock struct{}

type stuLockOpt struct {
	timeout *time.Duration
	tryFor  *time.Duration
	prefix  *string
}

type stuLockInstance struct {
	ctx  context.Context
	lock *redislock.Lock
}
