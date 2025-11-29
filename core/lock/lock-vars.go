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
	"time"
	_ "unsafe"

	"github.com/andypangaribuan/gmod/ice"
	"github.com/bsm/redislock"
	etcdclientv3 "go.etcd.io/etcd/client/v3"
)

//go:linkname iceLock github.com/andypangaribuan/gmod.iceLock
var iceLock ice.Lock

//go:linkname mainLockCallback github.com/andypangaribuan/gmod.mainLockCallback
var mainLockCallback func()

var (
	dvalTxTimeout       time.Duration
	dvalTxTryFor        *time.Duration
	txLockRedisClient   *redislock.Client
	txLockEtcdClient    *etcdclientv3.Client
	txLockEngineAddress string
	txLockEngineName    string
)
