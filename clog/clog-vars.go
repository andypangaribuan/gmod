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
	"google.golang.org/grpc"
)

//go:linkname mainCLogCallback github.com/andypangaribuan/gmod.mainCLogCallback
var mainCLogCallback func()

//go:linkname mainCLogUtilReflectionGetConf github.com/andypangaribuan/gmod.mainCLogUtilReflectionGetConf
var mainCLogUtilReflectionGetConf func(fieldName string) (any, error)

//go:linkname mainCLogNetGrpcConnection github.com/andypangaribuan/gmod.mainCLogNetGrpcConnection
var mainCLogNetGrpcConnection func(address string) (grpc.ClientConnInterface, error)

//go:linkname mainCLogUtilUid github.com/andypangaribuan/gmod.mainCLogUtilUid
var mainCLogUtilUid func() (string)

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
