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

	"github.com/andypangaribuan/gmod/clog"
	"github.com/bsm/redislock"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type stuLock struct{}

type stuLockOpt struct {
	timeout *time.Duration
	tryFor  *time.Duration
	prefix  *string
}

type stuLockInstance struct {
	released    bool
	logc        clog.Instance
	ctx         context.Context
	startedAt   time.Time
	key         string
	redisLock   *redislock.Lock
	cancel      *context.CancelFunc
	etcdSession *concurrency.Session
	etcdMtx     *concurrency.Mutex
}
