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
	"reflect"
	"runtime"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/pkg/errors"
)

func (slf *stuFuseContextRegulatorR) Next() (next bool, handler func(ctx FuseContextR) error) {
	slf.currentIndex++
	next = slf.currentIndex < len(slf.fuseContext.handlers)
	if next {
		handler = slf.fuseContext.handlers[slf.currentIndex]
	}
	return
}

func (slf *stuFuseContextRegulatorR) getHandler() func(ctx FuseContextR) error {
	return slf.fuseContext.handlers[slf.currentIndex]
}

func (slf *stuFuseContextRegulatorR) IsHandler(handler func(ctx FuseContextR) error) bool {
	v1 := runtime.FuncForPC(reflect.ValueOf(slf.getHandler()).Pointer()).Name()
	v2 := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return v1 == v2
}

func (slf *stuFuseContextRegulatorR) Call(handler func(ctx FuseContextR) error) (code int, res any, err error) {
	var (
		builder = &stuFuseContextBuilderR{
			original: slf.fuseContext,
		}
		ctx = builder.build()
	)

	err = handler(ctx)
	return ctx.responseCode, ctx.responseVal, err
}

func (slf *stuFuseContextRegulatorR) OnError(err error) bool {
	if err != nil && slf.fuseContext.errorHandler != nil {
		slf.fuseContext.errorHandler(slf.currentHandlerContext, errors.WithStack(err))
	}

	return err != nil
}

func (slf *stuFuseContextRegulatorR) Endpoint() string {
	return slf.fuseContext.endpoint
}

func (slf *stuFuseContextRegulatorR) Recover() {
	v := recover()
	if v != nil && slf.fuseContext.errorHandler != nil {
		err, ok := v.(error)
		if ok {
			err = errors.WithStack(err)
		} else {
			err = errors.New(fmt.Sprintf("%+v", v))
		}

		slf.fuseContext.errorHandler(slf.currentHandlerContext, err)
	}

	err := slf.send()
	if slf.fuseContext.errorHandler != nil {
		slf.fuseContext.errorHandler(slf.currentHandlerContext, errors.WithStack(err))
	}
}

func (slf *stuFuseContextRegulatorR) send() error {
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
