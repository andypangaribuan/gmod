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
	"strings"

	"github.com/pkg/errors"
)

func (slf *stuFuseRRegulator) Next() (next bool, handler func(ctx FuseRContext) any) {
	slf.currentIndex++
	next = slf.currentIndex < len(slf.mcx.handlers)
	if next {
		handler = slf.currentHandler()
	}

	return
}

func (slf *stuFuseRRegulator) IsHandler(handler func(ctx FuseRContext) any) bool {
	v1 := runtime.FuncForPC(reflect.ValueOf(slf.currentHandler()).Pointer()).Name()
	v2 := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return v1 == v2
}

func (slf *stuFuseRRegulator) Call(handler func(ctx FuseRContext) any, opt ...FuseRCallOpt) (code int, res any) {
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

	handler(ctx)
	slf.mcx.responseCode = ctx.responseCode
	slf.mcx.responseVal = ctx.responseVal

	return ctx.responseCode, ctx.responseVal
}

func (slf *stuFuseRRegulator) CallOpt() FuseRCallOpt {
	return new(stuFuseRCallOpt)
}

func (slf *stuFuseRRegulator) Endpoint() string {
	return slf.mcx.val.endpoint
}

func (slf *stuFuseRRegulator) Recover() {
	v := recover()

	if v != nil {
		err, ok := v.(error)
		if ok {
			err = errors.WithStack(err)
		} else {
			err = errors.New(fmt.Sprintf("%+v", v))
		}

		errMessage := err.Error()
		stackTrace := fmt.Sprintf("%+v", err)
		idx := strings.Index(stackTrace, errMessage)
		if idx == 0 {
			stackTrace = strings.Replace(stackTrace, errMessage, "", 1)
			stackTrace = strings.TrimSpace(stackTrace)
		}

		slf.mcx.errMessage = &errMessage
		slf.mcx.stackTrace = &stackTrace

		if slf.mcx.errorHandler != nil {
			slf.mcx.errorHandler(slf.currentHandlerContext, err)
		} else {
			slf.currentHandlerContext.R500InternalServerError("We apologize and are fixing the problem. Please try again at a later stage.")
		}
	}

	slf.send()
}
