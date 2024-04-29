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

func (slf *stuFuseRRegulator) Call(handler func(ctx FuseRContext) any, opt ...FuseRCallOpt) (res any, meta ResponseMeta, raw bool) {
	return slf.call(handler, nil, opt...)
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

		var (
			errMessage = err.Error()
			stackTrace = fmt.Sprintf("%+v", err)
			idx        = strings.Index(stackTrace, errMessage)
		)
		if idx == 0 {
			stackTrace = strings.Replace(stackTrace, errMessage, "", 1)
			stackTrace = strings.TrimSpace(stackTrace)
		}

		if slf.mcx.errorHandler != nil {
			slf.mcx.errorHandler(slf.currentHandlerContext, err)
		} else {
			slf.currentHandlerContext.R500InternalServerError(err)
		}

		slf.mcx.errMessage = &errMessage
		slf.mcx.stackTrace = &stackTrace
		slf.mcx.responseVal = slf.currentHandlerContext.responseVal
		slf.mcx.responseMeta = slf.currentHandlerContext.responseMeta
		slf.mcx.responseRaw = slf.currentHandlerContext.responseRaw
	}

	_ = slf.send()
}
