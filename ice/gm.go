/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package ice

type GM interface {
	SetConf(conf Conf)
	SetConv(conv Conv, tm ConvTime)
	SetDb(db Db)
	SetHttp(http Http)
	SetJson(json Json)
	SetNet(net Net)
	SetUtil(util Util, env UtilEnv)
}
