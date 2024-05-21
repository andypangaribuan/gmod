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
	"time"
	_ "unsafe"

	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

//go:linkname mainReflection github.com/andypangaribuan/gmod.mainReflection
var mainReflection func(key string, arg ...any) []any

//go:linkname mainCLogCallback github.com/andypangaribuan/gmod.mainCLogCallback
var mainCLogCallback func()

// clogNew github.com/andypangaribuan/gmod/server.clogNew
// var clogNew func(uid string) Instance

var (
	client           sclog.CLogServiceClient
	svcName          string
	svcVersion       string
	retryMaxDuration = time.Minute * 5
)

// accessed through injection
var (
	//lint:ignore U1000 Ignore unused function temporarily for debugging
	clogSetUserId func(clog Instance, id string) //nolint:golint,unused

	//lint:ignore U1000 Ignore unused function temporarily for debugging
	clogSetPartnerId func(clog Instance, id string) //nolint:golint,unused

	//lint:ignore U1000 Ignore unused function temporarily for debugging
	clogGetId func(clog Instance) (string, *string, *string) //nolint:golint,unused
)
