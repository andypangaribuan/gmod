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
	// next = slf.currentIndex < len(slf.original.handlers)
	next = slf.currentIndex < len(slf.mcx.handlers)
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
	// var (
	// 	builder = &stuFuseRContextBuilder{
	// 		original: slf.original,
	// 	}
	// 	ctx = builder.build()
	// )

	ctx := slf.buildContext()

	// override request from opt call
	for _, v := range opt {
		o, ok := v.(*stuFuseRCallOpt)
		if ok && o != nil {
			if o.header != nil {
				ctx.header = o.header
			}

			if o.param != nil {
				ctx.param = o.param
			}

			if o.query != nil {
				ctx.queries = o.query
			}

			if o.form != nil {
				ctx.form = o.form
			}
		}
	}

	// defer func() {
	// 	// slf.original.authObj = fm.Ternary(ctx.isSetAuthObj, ctx.authObj, slf.original.authObj)
	// 	// slf.original.userId = fm.Ternary(ctx.isSetUserId, ctx.userId, slf.original.userId)
	// 	// slf.original.partnerId = fm.Ternary(ctx.isSetPartnerId, ctx.partnerId, slf.original.partnerId)
	// 	slf.mcx.authObj = fm.Ternary(ctx.isSetAuthObj, ctx.authObj, slf.mcx.authObj)
	// 	slf.mcx.userId = fm.Ternary(ctx.isSetUserId, ctx.userId, slf.mcx.userId)
	// 	slf.mcx.partnerId = fm.Ternary(ctx.isSetPartnerId, ctx.partnerId, slf.mcx.partnerId)
	// }()

	err := handler(slf.clog, ctx)
	defer func() {
		slf.mcx.responseCode = ctx.responseCode
		slf.mcx.responseVal = ctx.responseVal
	}()
	// if err != nil && slf.original.errorHandler != nil {
	// 	slf.original.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
	// 	return -1, nil
	// }
	if err != nil && slf.mcx.errorHandler != nil {
		slf.mcx.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
		return -1, nil
	}

	return ctx.responseCode, ctx.responseVal
}

func (slf *stuFuseRRegulator) CallOpt() FuseRCallOpt {
	return new(stuFuseRCallOpt)
}

func (slf *stuFuseRRegulator) Endpoint() string {
	// return slf.original.val.endpoint
	return slf.mcx.val.endpoint
}

func (slf *stuFuseRRegulator) Recover() {
	v := recover()
	// if v != nil && slf.original.errorHandler != nil {
	// 	err, ok := v.(error)
	// 	if ok {
	// 		err = errors.WithStack(err)
	// 	} else {
	// 		err = errors.New(fmt.Sprintf("%+v", v))
	// 	}

	// 	slf.original.errorHandler(slf.clog, slf.currentHandlerContext, err)
	// }

	// err := slf.send()
	// if err != nil && slf.original.errorHandler != nil {
	// 	slf.original.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
	// }

	if v != nil && slf.mcx.errorHandler != nil {
		err, ok := v.(error)
		if ok {
			err = errors.WithStack(err)
		} else {
			err = errors.New(fmt.Sprintf("%+v", v))
		}

		slf.mcx.errorHandler(slf.clog, slf.currentHandlerContext, err)
	}

	err := slf.send()
	if err != nil && slf.mcx.errorHandler != nil {
		slf.mcx.errorHandler(slf.clog, slf.currentHandlerContext, errors.WithStack(err))
	}
}
