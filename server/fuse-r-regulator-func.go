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
	"fmt"
	"net/http"
	"strings"

	"github.com/andypangaribuan/gmod/fm"
)

func (slf *stuFuseRRegulator) call(handler func(ctx FuseRContext) any, unrouted func(ctx FuseRContext, method, path, url string) any, opt ...FuseRCallOpt) (res any, meta ResponseMeta, raw bool) {
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

		slf.mcx.val.endpoint = fmt.Sprintf("%v: %v", method, path)
		unrouted(ctx, method, path, slf.mcx.val.url)
	}

	slf.mcx.responseVal = ctx.responseVal
	slf.mcx.responseMeta = ctx.responseMeta
	slf.mcx.responseRaw = ctx.responseRaw

	return ctx.responseVal, ctx.responseMeta, ctx.responseRaw
}

func (slf *stuFuseRRegulator) buildContext() *stuFuseRContext {
	ctx := &stuFuseRContext{
		mcx: slf.mcx,

		header:     slf.mcx.val.header,
		param:      slf.mcx.val.param,
		queries:    slf.mcx.val.queries,
		form:       slf.mcx.val.form,
		file:       slf.mcx.val.file,
		bodyParser: slf.mcx.val.bodyParser,
	}

	slf.currentHandlerContext = ctx
	return ctx
}

func (slf *stuFuseRRegulator) currentHandler() func(FuseRContext) any {
	return slf.mcx.handlers[slf.currentIndex]
}

func (slf *stuFuseRRegulator) send() error {
	var (
		ctx    = slf.mcx.fcx.Status(slf.currentHandlerContext.responseMeta.Code)
		resVal = slf.currentHandlerContext.responseVal
	)

	if !slf.currentHandlerContext.responseRaw {
		slf.mcx.responseVal = Response{
			Meta: slf.currentHandlerContext.responseMeta,
			Data: resVal,
		}

		return ctx.JSON(slf.mcx.responseVal)
	}

	switch slf.currentHandlerContext.responseMeta.Code {
	case http.StatusMovedPermanently, http.StatusTemporaryRedirect, http.StatusPermanentRedirect:
		switch val := resVal.(type) {
		case map[string]string:
			url, ok := val["redirect"]
			if ok {
				return ctx.Redirect(url, slf.currentHandlerContext.responseMeta.Code)
			}
		}
	}

	switch val := resVal.(type) {
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
