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

//go:linkname clogNew github.com/andypangaribuan/gmod/server.clogNew
var clogNew func() Instance

var (
	client     sclog.CLogServiceClient
	svcName    string
	svcVersion string
)

// through injection
var (
	clogSetUserId    func(clog Instance, id string)
	clogSetPartnerId func(clog Instance, id string)
)
