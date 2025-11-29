/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"time"

	_ "github.com/andypangaribuan/gmod"
	"github.com/andypangaribuan/gmod/gm"
)

func init() {
	loadEnv()
	loadDb()
	gm.Conf.
		SetTimezone(env.AppTimezone).
		SetClogAddress(env.ClogAddress, env.AppName, env.AppVersion).
		SetClogRetryMaxDuration(time.Minute*5).
		SetInternalBaseUrls(env.SvcInternalBaseUrls).
		SetTxLockEngine(env.TxLockEngine, env.TxLockEngineAddress, env.TxLockTimeout, &env.TxLockTryFor).
		Commit()
}
