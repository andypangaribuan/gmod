/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func (slf *stuFuseContextBuilderR) Build() FuseContextR {
	ctx := &stuFuseContextR{
		fiberCtx:     slf.original.fiberCtx,
		endpoint:     slf.original.endpoint,
		isRegulator:  false,
		regulatorCtx: slf.original.regulatorCtx,
		authObj:      slf.original.authObj,
	}

	if slf.original.regulatorCtx.currentControllerContext != nil {
		ctx.lastResponseCode = slf.original.regulatorCtx.currentControllerContext.responseCode
		ctx.lastResponseVal = slf.original.regulatorCtx.currentControllerContext.responseVal
	}

	currentControllerContext := slf.original.regulatorCtx.currentControllerContext
	if currentControllerContext != nil {
		ctx.authObj = currentControllerContext.authObj
	}

	return ctx
}
