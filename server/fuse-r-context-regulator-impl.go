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
	"reflect"
	"runtime"

	"github.com/andypangaribuan/gmod/fm"
)

func (slf *stuFuseContextRegulatorR) Next() (next bool, getHandler func() func(ctx FuseContextR) error) {
	slf.currentIndex++
	return slf.currentIndex < len(slf.fuseContext.handlers), slf.getHandler
}

func (slf *stuFuseContextRegulatorR) getHandler() func(ctx FuseContextR) error {
	return slf.fuseContext.handlers[slf.currentIndex]
}

func (slf *stuFuseContextRegulatorR) IsHandler(handler func(ctx FuseContextR) error) bool {
	v1 := runtime.FuncForPC(reflect.ValueOf(slf.getHandler()).Pointer()).Name()
	v2 := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return v1 == v2
}

func (slf *stuFuseContextRegulatorR) ContextBuilder() FuseContextBuilderR {
	builder := &stuFuseContextBuilderR{
		original: slf.fuseContext,
	}

	return builder
}

func (slf *stuFuseContextRegulatorR) Send() error {
	ctx := slf.fuseContext.fiberCtx.Status(slf.currentHandlerContext.responseCode)

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
