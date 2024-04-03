/*
* Copyright (c) 2024.
* Created by Andy Pangaribuan <https://github.com/apangaribuan>.
* All Rights Reserved.
 */

package util

func init() {
	var (
		util    = new(stuUtil)
		utilEnv = new(stuUtilEnv)
	)

	iceUtil = util
	iceUtilEnv = utilEnv

	xinit()
}
