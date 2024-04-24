/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import "github.com/andypangaribuan/gmod/ice"

func (slf *stuLock) NewOpt() ice.LockOpt {
	return new(stuLockOpt)
}

func (slf *stuLock) Tx(id string, opt ...ice.LockOpt) (code int, err error) {
	return -1, nil
}
