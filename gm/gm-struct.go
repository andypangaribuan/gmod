/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gm

import "github.com/andypangaribuan/gmod/ice"

type srGM struct{}

type srConv struct {
	iceConv
	Time ice.ConvTime
}

type srUtil struct {
	iceUtil
	Env ice.UtilEnv
}
