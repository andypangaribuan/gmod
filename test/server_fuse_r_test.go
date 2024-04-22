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

	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/server"
	"github.com/andypangaribuan/gmod/test/db/repo"
)

// go test -v -run ^TestServerFuseR$
func TestServerFuseR(t *testing.T) {
	server.FuseR(env.AppRestPort, func(router server.RouterR) {
		router.AutoRecover(env.AppAutoRecover)
		router.PrintOnError(env.AppServerPrintOnError)
		router.ErrorHandler(sfrErrorHandler)
		router.Unrouted(sfrUnrouted)

		router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
			"POS: /hi/:firstName-:lastName/:age?": {sfrHi},
		})

		router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-1": {sfrPrivateStatus1},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-2": {sfrPrivateStatus1},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-3": {sfrPrivateStatus1},
			"GET: /private/status-4": {sfrPrivateStatus1, sfrPrivateStatus2},
			"GET: /private/status-5": {sfrPrivateStatus1, sfrPrivateStatus2},
			"GET: /private/status-6": {sfrPrivateStatus1, sfrPrivateStatus2},
		})

		// error or panic
		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-10": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-11": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-12": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-13": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-14": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-15": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-16": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-17": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /fetch-1": {sfrFetch},
		})
	})
}

func sfrUnrouted(ctx server.FuseRContext, method, path, url string) any {
	data := map[string]string{
		"status": "unrouted",
		"method": method,
		"path":   path,
		"url":    url,
	}

	return ctx.R404NotFound(data)
}

func sfrErrorHandler(ctx server.FuseRContext, err error) any {
	message := fmt.Sprintf("something went wrong: %v\n%+v", err.Error(), err)
	return ctx.R500InternalServerError(message)
}

func sfrRegulator(regulator server.FuseRRegulator) {
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

		if regulator.Endpoint() == "GET: /private/status-4" && regulator.IsHandler(sfrPrivateStatus2) {
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

func sfrAuth(ctx server.FuseRContext) any {
	ctx.Auth("Halo")
	ctx.UserId("abc")
	ctx.PartnerId("xyz")
	return ctx.R200OK("Andy")
}

func sfrPrivateStatus1(ctx server.FuseRContext) any {
	h := ctx.Header()
	hj, err := gm.Json.Encode(h)
	if err == nil {
		fmt.Println(hj)
	}
	fmt.Printf("private-status-1: header: %v\n", ctx.Header())

	ctx.SetFiles(map[string]string{
		"file1": "gcs/file1.pdf",
	})

	_, val := ctx.LastResponse()
	return ctx.R200OK(fmt.Sprintf("%v Pangaribuan", val))
}

func sfrPrivateStatus2(ctx server.FuseRContext) any {
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

func sfrPrivateStatusPanic(ctx server.FuseRContext) any {
	auth := ctx.Auth().(int) // panic error
	return ctx.R200OK(auth)
}

func sfrPrivateStatusErr(ctx server.FuseRContext) any {
	return errors.New("test error")
}

func sfrHi(ctx server.FuseRContext) any {
	return ctx.R200OK("ok")
}

func sfrFetch(ctx server.FuseRContext) any {
	entities, err := repo.User.Fetches("name=?", "andy")
	if err != nil {
		return ctx.R500InternalServerError(fmt.Sprintf("found some error: %v", err))
	}

	return ctx.R200OK(fmt.Sprintf("total user: %v", len(entities)))
}
