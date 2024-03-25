/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gmod

import (
	"log"

	_ "github.com/andypangaribuan/gmod/gm"

	_ "github.com/andypangaribuan/gmod/core/conf"
	_ "github.com/andypangaribuan/gmod/core/db"
	_ "github.com/andypangaribuan/gmod/core/json"
	_ "github.com/andypangaribuan/gmod/core/net"
	_ "github.com/andypangaribuan/gmod/core/util"

	"github.com/andypangaribuan/gmod/ice"
)

var (
	iceGM      ice.GM
	iceConf    ice.Conf
	iceDb      ice.Db
	iceJson    ice.Json
	iceNet     ice.Net
	iceUtil    ice.Util
	iceUtilEnv ice.UtilEnv
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	iceGM.SetConf(iceConf)
	iceGM.SetDb(iceDb)
	iceGM.SetJson(iceJson)
	iceGM.SetNet(iceNet)
	iceGM.SetUtil(iceUtil, iceUtilEnv)
}
