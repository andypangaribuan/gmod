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
	Unrouted(handler func(ctx FuseRContext, method, path, url string) any)

	ErrorHandler(catcher func(ctx FuseRContext, err error) any)
	Endpoints(regulator func(regulator FuseRRegulator), auth func(FuseRContext) any, pathHandlers map[string][]func(FuseRContext) any)
}

type RouterG interface {
	AutoRecover(autoRecover bool)
	Server() *grpc.Server
}

type FuseRContext interface {
	Clog() clog.Instance
	Auth(obj ...any) any
	UserId(id ...any) any
	PartnerId(id ...any) any
	SetFiles(files map[string]string)

	Url() string
	Header() *map[string]string

	LastResponse() (code int, val any)

	R200OK(val any) any
	R201Created(val any) any
	R202Accepted(val any) any
	R204NoContent(val any) any

	R301MovedPermanently(val any) any
	R307TemporaryRedirect(val any) any
	R308PermanentRedirect(val any) any

	R400BadRequest(val any) any
	R401Unauthorized(val any) any
	R403Forbidden(val any) any
	R404NotFound(val any) any
	R406NotAcceptable(val any) any
	R412PreconditionFailed(val any) any
	R418Teapot(val any) any
	R428PreconditionRequired(val any) any

	R500InternalServerError(val any) any
	R503ServiceUnavailable(val any) any
}

type FuseRContextBuilder interface {
	Build() FuseRContext
}

type FuseRRegulator interface {
	Next() (next bool, handler func(ctx FuseRContext) any)
	IsHandler(handler func(ctx FuseRContext) any) bool
	Call(handler func(ctx FuseRContext) any, opt ...FuseRCallOpt) (code int, res any)
	CallOpt() FuseRCallOpt
	Endpoint() string
	Recover()
}

type FuseRCallOpt interface {
	OverrideHeader(val map[string]string) FuseRCallOpt
	OverrideParam(val map[string]string) FuseRCallOpt
	OverrideQuery(val map[string]string) FuseRCallOpt
	OverrideForm(val map[string][]string) FuseRCallOpt
}
