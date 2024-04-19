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

func (slf *stuFuseRRegulator) Next() (next bool, handler func(clog clog.Instance, ctx FuseRContext) error) {
	slf.currentIndex++
	next = slf.currentIndex < len(slf.original.handlers)
	if next {
		handler = slf.currentHandler()
	}
	return
}

func (slf *stuFuseRRegulator) IsHandler(handler func(clog clog.Instance, ctx FuseRContext) error) bool {
	v1 := runtime.FuncForPC(reflect.ValueOf(slf.currentHandler()).Pointer()).Name()
	v2 := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return v1 == v2
}

func (slf *stuFuseRRegulator) Call(handler func(clog clog.Instance, ctx FuseRContext) error, opt ...FuseRCallOpt) (code int, res any) {
	var (
		builder = &stuFuseRContextBuilder{
			original: slf.original,
		}
		ctx = builder.build()
	)

	for _, v := range opt {
		o, ok := v.(*stuFuseRCallOpt)
		if ok && o != nil {
			if o.header != nil {
				ctx.header = o.header
			}
		}
	}

	defer func() {
		slf.original.authObj = ctx.authObj
		slf.original.userId = ctx.userId
		slf.original.partnerId = ctx.partnerId
	}()

	err := handler(slf.clog, ctx)
	if err != nil && slf.original.errorHandler != nil {
		slf.original.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
		return -1, nil
	}

	return ctx.responseCode, ctx.responseVal
}

func (slf *stuFuseRRegulator) CallOpt() FuseRCallOpt {
	return new(stuFuseRCallOpt)
}

func (slf *stuFuseRRegulator) Endpoint() string {
	return slf.original.val.endpoint
}

func (slf *stuFuseRRegulator) Recover() {
	v := recover()
	if v != nil && slf.original.errorHandler != nil {
		err, ok := v.(error)
		if ok {
			err = errors.WithStack(err)
		} else {
			err = errors.New(fmt.Sprintf("%+v", v))
		}

		slf.original.errorHandler(slf.clog, slf.currentHandlerContext, err)
	}

	err := slf.send()
	if err != nil && slf.original.errorHandler != nil {
		slf.original.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
	}
}
