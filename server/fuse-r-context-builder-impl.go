/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

// func (slf *stuFuseRContextBuilder) build() *stuFuseRContext {
// 	ctx := &stuFuseRContext{
// 		fiberCtx:     slf.original.fiberCtx,
// 		isRegulator:  false,
// 		regulatorCtx: slf.original.regulatorCtx,

// 		header:  slf.original.header,
// 		param:   slf.original.param,
// 		queries: slf.original.queries,
// 		form:    slf.original.form,
// 		file:    slf.original.file,

// 		val: slf.original.val,
// 	}

// 	currentHandlerContext := slf.original.regulatorCtx.currentHandlerContext
// 	if currentHandlerContext != nil {
// 		// ctx.authObj = currentHandlerContext.authObj
// 		// ctx.userId = currentHandlerContext.userId
// 		// ctx.partnerId = currentHandlerContext.partnerId
// 		ctx.lastResponseCode = currentHandlerContext.responseCode
// 		ctx.lastResponseVal = currentHandlerContext.responseVal

// 		// slf.original.authObj = currentHandlerContext.authObj
// 		// slf.original.userId = currentHandlerContext.userId
// 		// slf.original.partnerId = currentHandlerContext.partnerId
// 	}

// 	slf.original.regulatorCtx.currentHandlerContext = ctx
// 	return ctx
// }
