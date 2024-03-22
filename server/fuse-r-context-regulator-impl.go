/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import "github.com/andypangaribuan/gmod/fm"

func (slf *srFuseContextRegulatorR) Next() (canNext bool, ctrl func() func(ctx FuseContextR)) {
	slf.currentIndex++
	return slf.currentIndex < len(slf.fuseContext.controllers), slf.getCtrl
}

func (slf *srFuseContextRegulatorR) getCtrl() func(ctx FuseContextR) {
	return slf.fuseContext.controllers[slf.currentIndex]
}

func (slf *srFuseContextRegulatorR) ContextBuilder() FuseContextBuilderR {
	builder := &srFuseContextBuilderR{
		original: slf.fuseContext,
	}

	return builder
}

func (slf *srFuseContextRegulatorR) Send() {
	ctx := slf.fuseContext.fiberCtx.Status(slf.currentControllerContext.responseCode)

	switch val := slf.currentControllerContext.responseObj.(type) {
	case string:
		ctx.SendString(val)
	case *string:
		ctx.SendString(fm.GetDefault(val, ""))
	case []byte:
		ctx.Send(val)
	case *[]byte:
		ctx.Send(fm.GetDefault(val, []byte{}))
	case interface{}:
		ctx.JSON(val)
	default:
		ctx.SendString("invalid response object")
	}
}
