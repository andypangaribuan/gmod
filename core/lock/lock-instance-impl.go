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
)

func (slf *stuLockInstance) Release() {
	if isTxOnDevMode() {
		return
	}

	if txLockEngineName == "redis" {
		defer slf.clean()
		if slf.released || slf.redisLock == nil {
			return
		}

		if slf.cancel != nil {
			cancel := *slf.cancel
			cancel()
		}

		err := slf.redisLock.Release(slf.ctx)
		pushClogReport(slf.logc, slf.key, slf.obtainAt, slf.startedAt, err, "release-lock")
	}

	if txLockEngineName == "etcd" {
		defer slf.clean()
		if slf.released || slf.etcdSession == nil || slf.etcdMtx == nil {
			return
		}

		if slf.cancel != nil {
			cancel := *slf.cancel
			cancel()
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		err := slf.etcdMtx.Unlock(ctx)
		if err != nil {
			pushClogReport(slf.logc, slf.key, slf.obtainAt, slf.startedAt, err, "release-lock:mutex-unlock")
			return
		}

		err = slf.etcdSession.Close()
		pushClogReport(slf.logc, slf.key, slf.obtainAt, slf.startedAt, err, "release-lock:session-close")
	}
}

func (slf *stuLockInstance) IsLocked() (bool, error) {
	if isTxOnDevMode() {
		return true, nil
	}

	err := slf.isHaveLock()
	if err != nil {
		return false, err
	}

	ttl, err := slf.redisLock.TTL(slf.ctx)
	if err != nil {
		return false, err
	}

	return ttl > 0, nil
}

func (slf *stuLockInstance) Extend(duration time.Duration) error {
	if isTxOnDevMode() {
		return nil
	}

	err := slf.isHaveLock()
	if err != nil {
		return err
	}

	if txLockEngineName == "redis" {
		err := slf.redisLock.Refresh(slf.ctx, duration, nil)
		pushClogReport(slf.logc, slf.key, slf.obtainAt, slf.startedAt, err, "extend-lock")
		return err
	}

	return nil
}

func (slf *stuLockInstance) clean() {
	slf.released = true
	slf.logc = nil
	slf.ctx = nil
	slf.redisLock = nil
	slf.cancel = nil
	slf.etcdSession = nil
	slf.etcdMtx = nil
}
