/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package gmod

import (
	"log"

	_ "github.com/andypangaribuan/gmod/clog"
	_ "github.com/andypangaribuan/gmod/gm"

	_ "github.com/andypangaribuan/gmod/core/box"
	_ "github.com/andypangaribuan/gmod/core/cli"
	_ "github.com/andypangaribuan/gmod/core/conf"
	_ "github.com/andypangaribuan/gmod/core/conv"
	_ "github.com/andypangaribuan/gmod/core/crypto"
	_ "github.com/andypangaribuan/gmod/core/db"
	_ "github.com/andypangaribuan/gmod/core/http"
	_ "github.com/andypangaribuan/gmod/core/json"
	_ "github.com/andypangaribuan/gmod/core/lock"
	_ "github.com/andypangaribuan/gmod/core/net"
	_ "github.com/andypangaribuan/gmod/core/test"
	_ "github.com/andypangaribuan/gmod/core/use"
	_ "github.com/andypangaribuan/gmod/core/util"

	"github.com/KimMachineGun/automemlimit/memlimit"
	"github.com/andypangaribuan/gmod/ice"
	"go.uber.org/automaxprocs/maxprocs"
)

var (
	iceGM ice.GM

	iceBox      ice.Box
	iceCli      ice.Cli
	iceConf     ice.Conf
	iceConv     ice.Conv
	iceConvTime ice.ConvTime
	iceCrypto   ice.Crypto
	iceDb       ice.Db
	iceHttp     ice.Http
	iceJson     ice.Json
	iceLock     ice.Lock
	iceNet      ice.Net
	iceTest     ice.Test
	iceUse      ice.Use
	iceUseGcs   ice.UseGcs
	iceUtil     ice.Util
	iceUtilEnv  ice.UtilEnv
)

// accessed through injection
var (
	//lint:ignore U1000 Ignore unused function temporarily for debugging
	mainReflection func(key string, arg ...any) []any //nolint:golint,unused

	//lint:ignore U1000 Ignore unused function temporarily for debugging
	mainConfCommit func() //nolint:golint,unused
)

var (
	mainHttpCallback func()
	mainJsonCallback func()
	mainLockCallback func()
	mainUtilCallback func()
	mainCLogCallback func()
)

const memRation float64 = 0.8

func init() {
	_, _ = maxprocs.Set()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	_, _ = memlimit.SetGoMemLimitWithOpts(
		memlimit.WithRatio(memRation),
		memlimit.WithProvider(memlimit.FromCgroup),
	)
	_, _ = memlimit.SetGoMemLimitWithOpts(
		memlimit.WithRatio(memRation),
		memlimit.WithProvider(
			memlimit.ApplyFallback(
				memlimit.FromCgroup,
				memlimit.FromSystem,
			),
		),
	)
	_, _ = memlimit.SetGoMemLimit(memRation)
	_, _ = memlimit.SetGoMemLimitWithProvider(memlimit.Limit(1024*1024), memRation)
	_, _ = memlimit.SetGoMemLimitWithProvider(memlimit.FromCgroup, memRation)
	_, _ = memlimit.SetGoMemLimitWithProvider(memlimit.FromCgroupV1, memRation)
	_, _ = memlimit.SetGoMemLimitWithProvider(memlimit.FromCgroupHybrid, memRation)
	_, _ = memlimit.SetGoMemLimitWithProvider(memlimit.FromCgroupV2, memRation)

	mainReflection = reflection

	iceGM.
		SetBox(iceBox).
		SetCli(iceCli).
		SetConf(iceConf).
		SetConv(iceConv, iceConvTime).
		SetCrypto(iceCrypto).
		SetDb(iceDb).
		SetHttp(iceHttp).
		SetJson(iceJson).
		SetLock(iceLock).
		SetNet(iceNet).
		SetTest(iceTest).
		SetUse(iceUse, iceUseGcs).
		SetUtil(iceUtil, iceUtilEnv)

	mainConfCommit = func() {
		mainHttpCallback()
		mainJsonCallback()
		mainUtilCallback()
		mainLockCallback()
		mainCLogCallback()
	}
}
