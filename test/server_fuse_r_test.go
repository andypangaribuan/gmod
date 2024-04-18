/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/server"
)

// go test -v -run ^TestServerFuseR$
func TestServerFuseR(t *testing.T) {
	server.FuseR(env.AppRestPort, func(router server.RouterR) {
		router.AutoRecover(env.AppAutoRecover)
		router.PrintOnError(env.AppServerPrintOnError)
		router.ErrorHandler(sfrErrorHandler)

		router.Endpoints(nil, nil, map[string][]func(clog.Instance, server.FuseRContext) error{
			"GET: /private/status-1": {sfrPrivateStatus1},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(clog.Instance, server.FuseRContext) error{
			"GET: /private/status-2": {sfrPrivateStatus1},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(clog.Instance, server.FuseRContext) error{
			"GET: /private/status-3": {sfrPrivateStatus1},
			"GET: /private/status-4": {sfrPrivateStatus1, sfrPrivateStatus2},
			"GET: /private/status-5": {sfrPrivateStatus1, sfrPrivateStatus2},
			"GET: /private/status-6": {sfrPrivateStatus1, sfrPrivateStatus2},
		})

		// error or panic
		router.Endpoints(nil, sfrAuth, map[string][]func(clog.Instance, server.FuseRContext) error{
			"GET: /private/status-10": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-11": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(clog.Instance, server.FuseRContext) error{
			"GET: /private/status-12": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-13": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(clog.Instance, server.FuseRContext) error{
			"GET: /private/status-14": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-15": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(clog.Instance, server.FuseRContext) error{
			"GET: /private/status-16": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-17": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})
	})
}

func sfrErrorHandler(clog clog.Instance, ctx server.FuseRContext, err error) error {
	message := fmt.Sprintf("something went wrong: %v\n%+v", err.Error(), err)
	return ctx.R200OK(message)
}

func sfrRegulator(clog clog.Instance, regulator server.FuseRRegulator) {
	defer regulator.Recover()
	var (
		code    int
		canCall bool
	)

	for {
		canCall = true
		next, handler := regulator.Next()
		if !next {
			break
		}

		if regulator.Endpoint() == "GET: /private/status-4" && regulator.IsHandler(sfrPrivateStatus1) {
			continue
		}

		if regulator.Endpoint() == "GET: /private/status-5" && regulator.IsHandler(sfrPrivateStatus1) {
			canCall = false
			code, _ = regulator.Call(handler, regulator.CallOpt().OverrideHeader(map[string]string{
				"xyz": "Override Header",
			}))
		}

		if canCall {
			code, _ = regulator.Call(handler)
		}

		if code == -1 {
			return
		}

		if code < 200 || code > 299 {
			break
		}
	}
}

func sfrAuth(clog clog.Instance, ctx server.FuseRContext) error {
	ctx.Auth("Halo")
	return ctx.R200OK("Andy")
}

func sfrPrivateStatus1(clog clog.Instance, ctx server.FuseRContext) error {
	fmt.Printf("private-status-1: header: %v\n", ctx.Header())
	_, val := ctx.LastResponse()
	return ctx.R200OK(fmt.Sprintf("%v Pangaribuan", val))
}

func sfrPrivateStatus2(clog clog.Instance, ctx server.FuseRContext) error {
	fmt.Printf("private-status-2: header: %v\n", ctx.Header())
	auth := ctx.Auth().(string)
	_, val := ctx.LastResponse()

	data := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("%v %v", auth, val),
	}

	return ctx.R200OK(data)
}

func sfrPrivateStatusPanic(clog clog.Instance, ctx server.FuseRContext) error {
	auth := ctx.Auth().(int) // panic error
	return ctx.R200OK(auth)
}

func sfrPrivateStatusErr(clog clog.Instance, ctx server.FuseRContext) error {
	return errors.New("test error")
}
