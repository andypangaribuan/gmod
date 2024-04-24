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

	"github.com/pkg/errors"

	"github.com/andypangaribuan/gmod/ice"
)

func getLock(key string, timeout time.Duration, tryFor *time.Duration) (ice.LockInstance, error) {
	if txLockEngine == nil {
		return nil, errors.New("tx lock engine address is empty, please set from gm.Conf.SetTxLockEngineAddress")
	}

	startedAt := time.Now()
	ins := &stuLockInstance{
		ctx: context.Background(),
	}

	for {
		lock, err := txLockEngine.Obtain(ins.ctx, key, timeout, nil)
		if err != nil {
			if tryFor != nil && time.Since(startedAt) > *tryFor {
				return nil, err
			}

			time.Sleep(time.Millisecond * 10)
		} else {
			ins.lock = lock
			break
		}
	}

	return ins, nil
}
