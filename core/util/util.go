/*
* Copyright (c) 2024.
* Created by Andy Pangaribuan <https://github.com/apangaribuan>.
* All Rights Reserved.
 */

package util

func init() {
	var (
		util    = new(srUtil)
		utilEnv = new(srUtilEnv)
	)

	iceUtil = util
	iceUtilEnv = utilEnv

	xinit()
}
