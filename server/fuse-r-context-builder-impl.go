/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

func (slf *srFuseContextBuilderR) Build() FuseContextR {
	ctx := &srFuseContextR{
		fiberCtx:     slf.original.fiberCtx,
		endpoint:     slf.original.endpoint,
		isRegulator:  false,
		regulatorCtx: slf.original.regulatorCtx,
		authObj:      slf.original.authObj,
	}

	currentControllerContext := slf.original.regulatorCtx.currentControllerContext
	if currentControllerContext != nil {
		ctx.authObj = currentControllerContext.authObj
	}

	return ctx
}
