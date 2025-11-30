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
	"github.com/andypangaribuan/gmod/ice"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func getTxLock(logc clog.Instance, key string, timeout time.Duration, tryFor *time.Duration) (ice.LockInstance, error) {
	if txLockEngineName == "redis" {
		return getTxRedisLock(logc, key, timeout, tryFor)
	}

	if txLockEngineName == "etcd" {
		return getTxEtcdLock(logc, key, timeout, tryFor)
	}

	return &stuLockInstance{
			ctx:       context.Background(),
			startedAt: time.Now(),
			key:       key,
		},
		errors.New("unavailable lock engine")
}

func getTxRedisLock(logc clog.Instance, key string, timeout time.Duration, tryFor *time.Duration) (ice.LockInstance, error) {
	var (
		ins = &stuLockInstance{
			logc:      logc,
			ctx:       context.Background(),
			startedAt: time.Now(),
			key:       key,
		}
	)

	for {
		lock, err := txLockRedisClient.Obtain(ins.ctx, key, timeout, nil)
		if err != nil {
			if tryFor == nil || time.Since(ins.startedAt) > *tryFor {
				err = errors.WithMessage(err, "failed to lock")
				ins.obtainAt = time.Now()
				pushClogReport(logc, key, ins.obtainAt, ins.startedAt, err, "get-lock")
				return ins, err
			}

			time.Sleep(time.Millisecond * 10)
		} else {
			ins.redisLock = lock
			break
		}
	}

	renewCtx, cancel := context.WithCancel(ins.ctx)
	go func() {
		failedToRefresh := 0
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := ins.redisLock.Refresh(ins.ctx, timeout, nil)
				if err == nil {
					failedToRefresh = 0
				} else {
					failedToRefresh++
					if failedToRefresh == 3 {
						return
					}
				}

			case <-renewCtx.Done():
				return
			}
		}
	}()

	ins.cancel = &cancel
	ins.obtainAt = time.Now()

	return ins, nil
}

func getTxEtcdLock(logc clog.Instance, key string, timeout time.Duration, tryFor *time.Duration) (ice.LockInstance, error) {
	var (
		ins = &stuLockInstance{
			logc:      logc,
			ctx:       context.Background(),
			startedAt: time.Now(),
			key:       key,
		}
		mtx *concurrency.Mutex
	)

	ttl := max(int(timeout/time.Second), 3)
	session, err := concurrency.NewSession(txLockEtcdClient, concurrency.WithTTL(ttl))
	if err != nil {
		ins.obtainAt = time.Now()
		pushClogReport(logc, key, ins.obtainAt, ins.startedAt, err, "get-session")
		return nil, err
	}

	for {
		mtx = concurrency.NewMutex(session, "/locks/"+key)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = mtx.Lock(ctx)
		if err == nil {
			break
		}

		if tryFor != nil && time.Since(ins.startedAt) > *tryFor {
			err = errors.WithMessage(err, "failed to lock")
			ins.obtainAt = time.Now()
			pushClogReport(logc, key, ins.obtainAt, ins.startedAt, err, "get-lock")
			return ins, err
		}

		time.Sleep(time.Millisecond * 10)
	}

	ins.etcdSession = session
	ins.etcdMtx = mtx
	ins.obtainAt = time.Now()

	return ins, nil
}
