/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"github.com/andypangaribuan/gmod/clog"
	"google.golang.org/grpc"
)

type server interface {
	FuseG(grpcPort int, routes func(router RouterG))
	FuseR(restPort int, routes func(router RouterR))
	FuseGR(grpcPort int, grpcRoutes func(router RouterG), restPort int, restRoutes func(router RouterR))
}

type RouterR interface {
	AutoRecover(autoRecover bool)
	PrintOnError(printOnError bool)
	Unrouted(handler func(clog clog.Instance, ctx FuseContextR, method, path, url string) error)

	ErrorHandler(catcher func(clog clog.Instance, ctx FuseContextR, err error) error)
	Endpoints(regulator func(clog clog.Instance, regulator FuseRegulatorR), auth func(clog.Instance, FuseContextR) error, pathHandlers map[string][]func(clog.Instance, FuseContextR) error)
}

type RouterG interface {
	AutoRecover(autoRecover bool)
	Server() *grpc.Server
}

type FuseContextR interface {
	LastResponse() (code int, val any)
	Auth(obj ...any) any

	R200OK(val any) error
}

type FuseRegulatorR interface {
	Next() (next bool, handler func(clog clog.Instance, ctx FuseContextR) error)
	IsHandler(handler func(clog clog.Instance, ctx FuseContextR) error) bool
	Call(handler func(clog clog.Instance, ctx FuseContextR) error) (code int, res any)
	Endpoint() string
	Recover()
}

type FuseContextBuilderR interface {
	Build() FuseContextR
}
