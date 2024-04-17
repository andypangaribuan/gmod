/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func (slf *stuFuseContextBuilderR) build() *stuFuseContextR {
	ctx := &stuFuseContextR{
		fiberCtx:     slf.original.fiberCtx,
		endpoint:     slf.original.endpoint,
		isRegulator:  false,
		regulatorCtx: slf.original.regulatorCtx,
		authObj:      slf.original.authObj,
	}

	currentHandlerContext := slf.original.regulatorCtx.currentHandlerContext
	if currentHandlerContext != nil {
		ctx.authObj = currentHandlerContext.authObj
		ctx.lastResponseCode = currentHandlerContext.responseCode
		ctx.lastResponseVal = currentHandlerContext.responseVal
	}

	slf.original.regulatorCtx.currentHandlerContext = ctx
	return ctx
}
