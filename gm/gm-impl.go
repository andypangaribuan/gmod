/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package gm

import "github.com/andypangaribuan/gmod/ice"

func (slf *stuGM) SetBox(box ice.Box) ice.GM {
	Box = box
	return slf
}

func (slf *stuGM) SetClog(clog ice.Clog) ice.GM {
	Clog = clog
	return slf
}

func (slf *stuGM) SetConf(conf ice.Conf) ice.GM {
	Conf = conf
	return slf
}

func (slf *stuGM) SetConv(conv ice.Conv, tm ice.ConvTime) ice.GM {
	Conv = &stuConv{
		iceConv: conv,
		Time:    tm,
	}
	return slf
}

func (slf *stuGM) SetCrypto(crypto ice.Crypto) ice.GM {
	Crypto = crypto
	return slf
}

func (slf *stuGM) SetDb(db ice.Db) ice.GM {
	Db = db
	return slf
}

func (slf *stuGM) SetHttp(http ice.Http) ice.GM {
	Http = http
	return slf
}

func (slf *stuGM) SetJson(json ice.Json) ice.GM {
	Json = json
	return slf
}

func (slf *stuGM) SetNet(net ice.Net) ice.GM {
	Net = net
	return slf
}

func (slf *stuGM) SetTest(test ice.Test) ice.GM {
	Test = test
	return slf
}

func (slf *stuGM) SetUtil(util ice.Util, env ice.UtilEnv) ice.GM {
	Util = &stuUtil{
		iceUtil: util,
		Env:     env,
	}
	return slf
}
