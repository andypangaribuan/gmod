/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "google.golang.org/grpc"

type server interface {
	FuseG(grpcPort int, routes func(router RouterG))
	FuseR(restPort int, routes func(router RouterR))
	FuseGR(grpcPort int, grpcRoutes func(router RouterG), restPort int, restRoutes func(router RouterR))
}

type RouterR interface {
	AutoRecover(autoRecover bool)
	PrintOnError(printOnError bool)
	Unrouted(controller func(ctx FuseContextR, method, path, url string))
	Group(endpoints map[string][]func() (isRegulator bool, controller func(ctx FuseContextR)))
	Single(path string, controllers ...func() (isRegulator bool, controller func(ctx FuseContextR)))
}

type RouterG interface {
	AutoRecover(autoRecover bool)
	Server() *grpc.Server
}

type FuseContextR interface {
	Regulator() FuseContextRegulatorR
	GetResponse() (code int, obj any)
	SetAuth(val any)
}

type FuseContextRegulatorR interface {
	Next() (canNext bool, ctrl func() func(ctx FuseContextR))
	ContextBuilder() FuseContextBuilderR
	Send()
}

type FuseContextBuilderR interface {
	Build() FuseContextR
}
