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
	SetBox(box Box) GM
	SetCli(cli Cli) GM
	SetConf(conf Conf) GM
	SetConv(conv Conv, tm ConvTime) GM
	SetCrypto(crypto Crypto) GM
	SetDb(db Db) GM
	SetHttp(http Http) GM
	SetJson(json Json) GM
	SetLock(lock Lock) GM
	SetNet(net Net) GM
	SetTest(test Test) GM
	SetUse(use Use, useGcs UseGcs) GM
	SetUtil(util Util, env UtilEnv) GM
}
