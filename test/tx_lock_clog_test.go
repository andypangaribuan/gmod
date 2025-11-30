/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"testing"
	"time"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/server"
)

type srTestTxLockClog struct{}

func TestTxLockClog(t *testing.T) {
	sr := new(srTestTxLockClog)

	server.FuseR(env.AppRestPort, func(router server.RouterR) {
		router.AutoRecover(env.AppAutoRecover)
		router.PrintOnError(env.AppServerPrintOnError)
		router.ErrorHandler(sr.errorHandler)
		router.Unrouted(sr.unrouted)

		router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
			"GET: /lock": {sr.lock},
		})
	})
}

func (slf *srTestTxLockClog) errorHandler(ctx server.FuseRContext, err error) any {
	return ctx.R500InternalServerError(err)
}

func (slf *srTestTxLockClog) unrouted(ctx server.FuseRContext, method, path, url string) any {
	data := map[string]string{
		"status": "unrouted",
		"method": method,
		"path":   path,
		"url":    url,
	}

	return ctx.R404NotFound(data)
}

func (slf *srTestTxLockClog) lock(ctx server.FuseRContext) any {
	logc := ctx.Clog()

	lock, err := gm.Lock.Tx(logc, "OK")
	defer lock.Release()

	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	time.Sleep(time.Second * 3)
	return ctx.R200OK("ok")
}
