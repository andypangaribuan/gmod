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
	"strings"

	"github.com/andypangaribuan/gmod/fm"
)

func (slf *stuFuseRRegulator) call(handler func(ctx FuseRContext) any, unrouted func(ctx FuseRContext, method, path, url string) any, opt ...FuseRCallOpt) (code int, res any) {
	ctx := slf.buildContext()

	// override request from opt call
	for _, v := range opt {
		o, ok := v.(*stuFuseRCallOpt)
		if ok && o != nil {
			if o.header != nil {
				ctx.header = o.header
			}

			if o.param != nil {
				ctx.param = o.param
			}

			if o.query != nil {
				ctx.queries = o.query
			}

			if o.form != nil {
				ctx.form = o.form
			}
		}
	}

	if handler != nil {
		handler(ctx)
	}

	if unrouted != nil {
		var (
			fcx    = slf.mcx.fcx
			method = ""
			path   = fcx.Path()
		)

		if fcx.Route() != nil {
			method = strings.ToLower(fcx.Route().Method)
		}

		unrouted(ctx, method, path, slf.mcx.val.url)
	}

	slf.mcx.responseCode = ctx.responseCode
	slf.mcx.responseVal = ctx.responseVal

	return ctx.responseCode, ctx.responseVal
}

func (slf *stuFuseRRegulator) buildContext() *stuFuseRContext {
	ctx := &stuFuseRContext{
		mcx: slf.mcx,

		header:  slf.mcx.val.header,
		param:   slf.mcx.val.param,
		queries: slf.mcx.val.queries,
		form:    slf.mcx.val.form,
		file:    slf.mcx.val.file,
	}

	slf.currentHandlerContext = ctx
	return ctx
}

func (slf *stuFuseRRegulator) currentHandler() func(FuseRContext) any {
	return slf.mcx.handlers[slf.currentIndex]
}

func (slf *stuFuseRRegulator) send() error {
	ctx := slf.mcx.fcx.Status(slf.currentHandlerContext.responseCode)

	switch val := slf.currentHandlerContext.responseVal.(type) {
	case string:
		return ctx.SendString(val)
	case *string:
		return ctx.SendString(fm.GetDefault(val, ""))
	case []byte:
		return ctx.Send(val)
	case *[]byte:
		return ctx.Send(fm.GetDefault(val, []byte{}))
	case any:
		return ctx.JSON(val)
	default:
		return ctx.SendString("invalid response object")
	}
}
