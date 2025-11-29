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
	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/pkg/errors"
)

func (slf *stuLock) NewOpt() ice.LockOpt {
	return new(stuLockOpt)
}

func (slf *stuLock) Tx(logc clog.Instance, key string, opt ...ice.LockOpt) (ice.LockInstance, error) {
	if isTxOnDevMode() {
		return new(stuLockInstance), nil
	}

	if (txLockRedisClient == nil && txLockEtcdClient == nil) ||
		(txLockEngineName == "redis" && txLockRedisClient == nil) ||
		(txLockEngineName == "etcd" && txLockEtcdClient == nil) {
		return new(stuLockInstance), errors.New("doesn't have lock engine, please set from gm.Conf.SetTxLockEngine")
	}

	var (
		timeout = dvalTxTimeout
		tryFor  = dvalTxTryFor
		prefix  string
	)

	for _, o := range opt {
		if v, ok := o.(*stuLockOpt); ok {
			if v.timeout != nil {
				timeout = *v.timeout
			}

			if v.tryFor != nil {
				tryFor = v.tryFor
			}

			if v.prefix != nil {
				prefix = *v.prefix
			}
		}
	}

	return getTxLock(logc, prefix+key, timeout, tryFor)
}
