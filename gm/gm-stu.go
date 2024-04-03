/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gm

import "github.com/andypangaribuan/gmod/ice"

type stuGM struct{}

type stuConv struct {
	iceConv
	Time ice.ConvTime
}

type stuUtil struct {
	iceUtil
	Env ice.UtilEnv
}
