/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gm

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceGM github.com/andypangaribuan/gmod.iceGM
var iceGM ice.GM

var (
	Conf   ice.Conf
	Conv   *stuConv
	Crypto ice.Crypto
	Db     ice.Db
	Http   ice.Http
	Json   ice.Json
	Util   *stuUtil
)
