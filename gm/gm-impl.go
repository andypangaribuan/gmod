/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gm

import "github.com/andypangaribuan/gmod/ice"

func (*srGM) SetConf(conf ice.Conf) {
	Conf = conf
}

func (*srGM) SetConv(conv ice.Conv, tm ice.ConvTime) {
	Conv = &srConv{
		iceConv: conv,
		Time:    tm,
	}
}

func (*srGM) SetDb(db ice.Db) {
	Db = db
}

func (*srGM) SetNet(net ice.Net) {
	Net = net
}

func (*srGM) SetJson(json ice.Json) {
	Json = json
}

func (*srGM) SetUtil(util ice.Util, env ice.UtilEnv) {
	Util = &srUtil{
		iceUtil: util,
		Env:     env,
	}
}
