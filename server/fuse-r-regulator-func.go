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
	"github.com/andypangaribuan/gmod/fm"
)

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
