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
