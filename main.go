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

	_ "github.com/andypangaribuan/gmod/gm"

	_ "github.com/andypangaribuan/gmod/core/box"
	_ "github.com/andypangaribuan/gmod/core/conf"
	_ "github.com/andypangaribuan/gmod/core/conv"
	_ "github.com/andypangaribuan/gmod/core/crypto"
	_ "github.com/andypangaribuan/gmod/core/db"
	_ "github.com/andypangaribuan/gmod/core/http"
	_ "github.com/andypangaribuan/gmod/core/json"
	_ "github.com/andypangaribuan/gmod/core/util"

	"github.com/andypangaribuan/gmod/ice"
	"go.uber.org/automaxprocs/maxprocs"
)

var (
	iceGM ice.GM

	iceBox      ice.Box
	iceConf     ice.Conf
	iceConv     ice.Conv
	iceConvTime ice.ConvTime
	iceCrypto   ice.Crypto
	iceDb       ice.Db
	iceHttp     ice.Http
	iceJson     ice.Json
	iceUtil     ice.Util
	iceUtilEnv  ice.UtilEnv
)

func init() {
	maxprocs.Set()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	iceGM.SetBox(iceBox)
	iceGM.SetConf(iceConf)
	iceGM.SetConv(iceConv, iceConvTime)
	iceGM.SetCrypto(iceCrypto)
	iceGM.SetDb(iceDb)
	iceGM.SetHttp(iceHttp)
	iceGM.SetJson(iceJson)
	iceGM.SetUtil(iceUtil, iceUtilEnv)
}
