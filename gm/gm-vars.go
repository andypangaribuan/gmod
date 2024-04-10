/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package gm

import (
	"github.com/andypangaribuan/gmod/ice"

	_ "unsafe"
)

//go:linkname iceGM github.com/andypangaribuan/gmod.iceGM
var iceGM ice.GM

var initExecuted = false

var (
	Box    ice.Box
	Conf   ice.Conf
	Conv   *stuConv
	Crypto ice.Crypto
	Db     ice.Db
	Http   ice.Http
	Json   ice.Json
	Net    ice.Net
	Test   ice.Test
	Util   *stuUtil
)
