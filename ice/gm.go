/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

type GM interface {
	SetBox(box Box)
	SetConf(conf Conf)
	SetConv(conv Conv, tm ConvTime)
	SetCrypto(crypto Crypto)
	SetDb(db Db)
	SetHttp(http Http)
	SetJson(json Json)
	SetUtil(util Util, env UtilEnv)
}
