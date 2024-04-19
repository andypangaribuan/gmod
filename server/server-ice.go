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
	Unrouted(handler func(clog clog.Instance, ctx FuseRContext, method, path, url string) error)

	ErrorHandler(catcher func(clog clog.Instance, ctx FuseRContext, err error) error)
	Endpoints(regulator func(clog clog.Instance, regulator FuseRRegulator), auth func(clog.Instance, FuseRContext) error, pathHandlers map[string][]func(clog.Instance, FuseRContext) error)
}

type RouterG interface {
	AutoRecover(autoRecover bool)
	Server() *grpc.Server
}

type FuseRContext interface {
	LastResponse() (code int, val any)
	Auth(obj ...any) any
	Header() *map[string]string
	Url() string

	R200OK(val any) error
}

type FuseRContextBuilder interface {
	Build() FuseRContext
}

type FuseRRegulator interface {
	Next() (next bool, handler func(clog clog.Instance, ctx FuseRContext) error)
	IsHandler(handler func(clog clog.Instance, ctx FuseRContext) error) bool
	Call(handler func(clog clog.Instance, ctx FuseRContext) error, opt ...FuseRCallOpt) (code int, res any)
	CallOpt() FuseRCallOpt
	Endpoint() string
	Recover()
}

type FuseRCallOpt interface {
	OverrideHeader(header map[string]string) FuseRCallOpt
}
