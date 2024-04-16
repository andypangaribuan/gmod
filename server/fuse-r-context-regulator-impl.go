/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuFuseContextRegulatorR) Next() (canNext bool, ctrl func() func(ctx FuseContextR)) {
	slf.currentIndex++
	return slf.currentIndex < len(slf.fuseContext.controllers), slf.getCtrl
}

func (slf *stuFuseContextRegulatorR) getCtrl() func(ctx FuseContextR) {
	return slf.fuseContext.controllers[slf.currentIndex]
}

func (slf *stuFuseContextRegulatorR) ContextBuilder() FuseContextBuilderR {
	builder := &stuFuseContextBuilderR{
		original: slf.fuseContext,
	}

	return builder
}

func (slf *stuFuseContextRegulatorR) Send() {
	ctx := slf.fuseContext.fiberCtx.Status(slf.currentControllerContext.responseCode)

	switch val := slf.currentControllerContext.responseVal.(type) {
	case string:
		ctx.SendString(val)
	case *string:
		ctx.SendString(fm.GetDefault(val, ""))
	case []byte:
		ctx.Send(val)
	case *[]byte:
		ctx.Send(fm.GetDefault(val, []byte{}))
	case any:
		ctx.JSON(val)
	default:
		ctx.SendString("invalid response object")
	}
}
