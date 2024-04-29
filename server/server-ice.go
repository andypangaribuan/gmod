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
	"mime/multipart"

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

	ReqHeader() *map[string]string
	ReqParam() *map[string]string
	ReqQuery() *map[string]string
	ReqForm() *map[string][]string
	ReqFile() *map[string][]*multipart.FileHeader

	GetHeader(key string, dval ...string) *string

	ReqParser(header any, body any) error
	ReqParserPQF(param any, query any, form any) error

	LastResponse() (val any, meta ResponseMeta)

	R200OK(val any, opt ...ResponseOpt) any
	R201Created(val any, opt ...ResponseOpt) any
	R202Accepted(val any, opt ...ResponseOpt) any
	R204NoContent(val any, opt ...ResponseOpt) any

	R301MovedPermanently(val any, opt ...ResponseOpt) any
	R307TemporaryRedirect(val any, opt ...ResponseOpt) any
	R308PermanentRedirect(val any, opt ...ResponseOpt) any

	R400BadRequest(val any, opt ...ResponseOpt) any
	R401Unauthorized(val any, opt ...ResponseOpt) any
	R403Forbidden(val any, opt ...ResponseOpt) any
	R404NotFound(val any, opt ...ResponseOpt) any
	R406NotAcceptable(val any, opt ...ResponseOpt) any
	R412PreconditionFailed(val any, opt ...ResponseOpt) any
	R418Teapot(val any, opt ...ResponseOpt) any
	R428PreconditionRequired(val any, opt ...ResponseOpt) any

	R500InternalServerError(val any, opt ...ResponseOpt) any
	R503ServiceUnavailable(val any, opt ...ResponseOpt) any
}

type FuseRContextBuilder interface {
	Build() FuseRContext
}

type FuseRRegulator interface {
	Next() (next bool, handler func(ctx FuseRContext) any)
	IsHandler(handler func(ctx FuseRContext) any) bool
	Call(handler func(ctx FuseRContext) any, opt ...FuseRCallOpt) (res any, meta ResponseMeta, raw bool)
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
