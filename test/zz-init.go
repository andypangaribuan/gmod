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
	_ "github.com/andypangaribuan/gmod"
	"github.com/andypangaribuan/gmod/gm"
)

func init() {
	loadEnv()
	loadDb()
	gm.Conf.
		SetTimezone(env.AppTimezone).
		SetCLogAddress(env.ClogAddress, env.AppName, env.AppVersion).
		SetInternalBaseUrls(env.SvcInternalBaseUrls).
		Commit()
}
