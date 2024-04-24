/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import "time"

func (slf *stuLockInstance) Release() {
	slf.lock.Release(slf.ctx)
}

func (slf *stuLockInstance) IsLocked() (bool, error) {
	ttl, err := slf.lock.TTL(slf.ctx)
	if err != nil {
		return false, err
	}

	return ttl > 0, nil
}

func (slf *stuLockInstance) Extend(duration time.Duration) error {
	return slf.lock.Refresh(slf.ctx, duration, nil)
}
