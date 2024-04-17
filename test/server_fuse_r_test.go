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

	"github.com/andypangaribuan/gmod/server"
)

// go test -v -run ^TestServerFuseR$
func TestServerFuseR(t *testing.T) {
	server.FuseR(env.AppRestPort, func(router server.RouterR) {
		router.AutoRecover(env.AppAutoRecover)
		router.PrintOnError(env.AppServerPrintOnError)
		router.ErrorHandler(sfrErrorHandler)

		router.Endpoints(nil, nil, map[string][]func(ctx server.FuseContextR) error{
			"GET: /private/status-1": {sfrPrivateStatus1},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(ctx server.FuseContextR) error{
			"GET: /private/status-2": {sfrPrivateStatus1},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(ctx server.FuseContextR) error{
			"GET: /private/status-3": {sfrPrivateStatus1},
			"GET: /private/status-4": {sfrPrivateStatus1, sfrPrivateStatus2},
			"GET: /private/status-5": {sfrPrivateStatus1, sfrPrivateStatus2},
		})

		// error or panic
		router.Endpoints(nil, sfrAuth, map[string][]func(ctx server.FuseContextR) error{
			"GET: /private/status-10": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-11": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(ctx server.FuseContextR) error{
			"GET: /private/status-12": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-13": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(ctx server.FuseContextR) error{
			"GET: /private/status-14": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-15": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(ctx server.FuseContextR) error{
			"GET: /private/status-16": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-17": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})
	})
}

func sfrErrorHandler(ctx server.FuseContextR, err error) error {
	message := fmt.Sprintf("something went wrong: %+v", err)
	return ctx.R200OK(message)
}

func sfrRegulator(regulator server.FuseContextRegulatorR) {
	defer regulator.Recover()

	for {
		next, handler := regulator.Next()
		if !next {
			break
		}

		if regulator.Endpoint() == "GET: /private/status-4" && regulator.IsHandler(sfrPrivateStatus1) {
			continue
		}

		code, _ := regulator.Call(handler)
		if code == -1 {
			return
		}

		if code < 200 || code >= 300 {
			break
		}
	}
}

func sfrAuth(ctx server.FuseContextR) error {
	ctx.Auth("Halo")
	return ctx.R200OK("Andy")
}

func sfrPrivateStatus1(ctx server.FuseContextR) error {
	_, val := ctx.LastResponse()
	return ctx.R200OK(fmt.Sprintf("%v Pangaribuan", val))
}

func sfrPrivateStatus2(ctx server.FuseContextR) error {
	auth := ctx.Auth().(string)
	_, val := ctx.LastResponse()

	data := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("%v %v", auth, val),
	}

	return ctx.R200OK(data)
}

func sfrPrivateStatusPanic(ctx server.FuseContextR) error {
	auth := ctx.Auth().(int) // panic error
	return ctx.R200OK(auth)
}

func sfrPrivateStatusErr(ctx server.FuseContextR) error {
	return errors.New("test error")
}
