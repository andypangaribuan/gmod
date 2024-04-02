/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gm

import "github.com/andypangaribuan/gmod/ice"

func (*stuGM) SetConf(conf ice.Conf) {
	Conf = conf
}

func (*stuGM) SetConv(conv ice.Conv, tm ice.ConvTime) {
	Conv = &stuConv{
		iceConv: conv,
		Time:    tm,
	}
}

func (*stuGM) SetDb(db ice.Db) {
	Db = db
}

func (*stuGM) SetHttp(http ice.Http) {
	Http = http
}

func (*stuGM) SetNet(net ice.Net) {
	Net = net
}

func (*stuGM) SetJson(json ice.Json) {
	Json = json
}

func (*stuGM) SetUtil(util ice.Util, env ice.UtilEnv) {
	Util = &stuUtil{
		iceUtil: util,
		Env:     env,
	}
}
