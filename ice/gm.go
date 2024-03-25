/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package ice

type GM interface {
	SetConf(conf Conf)
	SetDb(db Db)
	SetJson(json Json)
	SetNet(net Net)
	SetUtil(util Util, env UtilEnv)
}
