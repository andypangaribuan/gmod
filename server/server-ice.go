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
	Unrouted(handler func(ctx FuseContextR, method, path, url string) error)

	ErrorHandler(catcher func(ctx FuseContextR, err error) error)
	Endpoints(regulator func(regulator FuseContextRegulatorR), auth func(ctx FuseContextR) error, pathHandlers map[string][]func(ctx FuseContextR) error)
}

type RouterG interface {
	AutoRecover(autoRecover bool)
	Server() *grpc.Server
}

type FuseContextR interface {
	GetLastResponse() (code int, val any)
	GetResponse() (code int, val any)
	SetAuth(val any)
	Auth() any

	R200OK(val any) error
}

type FuseContextRegulatorR interface {
	Next() (next bool, getHandler func(ctx FuseContextR) error)
	IsHandler(handler func(ctx FuseContextR) error) bool
	Call(handler func(ctx FuseContextR) error) (code int, res any, err error)
	OnError(err error) bool
	Recover()
	Send() error
}

type FuseContextBuilderR interface {
	Build() FuseContextR
}
