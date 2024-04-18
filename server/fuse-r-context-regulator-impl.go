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

	"github.com/andypangaribuan/gmod/clog"
	"github.com/pkg/errors"
)

func (slf *stuFuseRegulatorR) Next() (next bool, handler func(clog clog.Instance, ctx FuseContextR) error) {
	slf.currentIndex++
	next = slf.currentIndex < len(slf.fuseContext.handlers)
	if next {
		handler = slf.currentHandler()
	}
	return
}

func (slf *stuFuseRegulatorR) IsHandler(handler func(clog clog.Instance, ctx FuseContextR) error) bool {
	v1 := runtime.FuncForPC(reflect.ValueOf(slf.currentHandler()).Pointer()).Name()
	v2 := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return v1 == v2
}

func (slf *stuFuseRegulatorR) Call(handler func(clog clog.Instance, ctx FuseContextR) error) (code int, res any) {
	var (
		builder = &stuFuseContextBuilderR{
			original: slf.fuseContext,
		}
		ctx = builder.build()
	)

	err := handler(slf.clog, ctx)
	if err != nil && slf.fuseContext.errorHandler != nil {
		slf.fuseContext.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
		return -1, nil
	}

	return ctx.responseCode, ctx.responseVal
}

func (slf *stuFuseRegulatorR) Endpoint() string {
	return slf.fuseContext.endpoint
}

func (slf *stuFuseRegulatorR) Recover() {
	v := recover()
	if v != nil && slf.fuseContext.errorHandler != nil {
		err, ok := v.(error)
		if ok {
			err = errors.WithStack(err)
		} else {
			err = errors.New(fmt.Sprintf("%+v", v))
		}

		slf.fuseContext.errorHandler(slf.clog, slf.currentHandlerContext, err)
	}

	err := slf.send()
	if err != nil && slf.fuseContext.errorHandler != nil {
		slf.fuseContext.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
	}
}
