/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import (
	_ "unsafe"

	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

//go:linkname mainCLogCallback github.com/andypangaribuan/gmod.mainCLogCallback
var mainCLogCallback func()

var (
	client sclog.CLogServiceClient
)
