/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import "github.com/pkg/errors"

func (slf *stuLockInstance) isHaveLock() error {
	if slf.lock != nil {
		return nil
	}

	return errors.New("you don't have any lock")
}