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
	"fmt"
	"testing"

	"github.com/andypangaribuan/gmod/server"
)

// go test -v -run ^TestServerFuseR$
func TestServerFuseR(t *testing.T) {
	server.FuseR(env.AppRestPort, func(router server.RouterR) {
		router.AutoRecover(env.AppAutoRecover)
		router.PrintOnError(env.AppServerPrintOnError)

		router.Endpoints(nil, nil, map[string][]func(ctx server.FuseContextR) any{
			"GET: /private/status-1": {serverFuseRAuth, serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
			"GET: /private/status-2": {serverFuseRAuth, serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
		})

		router.Endpoints(nil, serverFuseRAuth, map[string][]func(ctx server.FuseContextR) any{
			"GET: /private/status-3": {serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
			"GET: /private/status-4": {serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
		})

		router.Endpoints(serverFuseRRegulator, nil, map[string][]func(ctx server.FuseContextR) any{
			"GET: /private/status-5": {serverFuseRAuth, serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
		})

		router.Endpoints(serverFuseRRegulator, serverFuseRAuth, map[string][]func(ctx server.FuseContextR) any{
			"GET: /private/status-6": {serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
		})
	})
}

func serverFuseRRegulator(ctx server.FuseContextR) any {
	regulator := ctx.Regulator()

	for {
		next, handler := regulator.Next()
		if !next {
			break
		}

		if regulator.IsHandler(serverRuseRPrivateStatus1) {
			continue
		}

		builder := regulator.ContextBuilder()
		ctx := builder.Build()
		handler()(ctx)

		code, _ := ctx.GetResponse()
		if code < 200 || code >= 300 {
			break
		}
	}

	return regulator.Send()
}

func serverFuseRAuth(ctx server.FuseContextR) any {
	ctx.SetAuth("Halo")
	return ctx.R200OK("Andy")
}

func serverRuseRPrivateStatus1(ctx server.FuseContextR) any {
	_, val := ctx.GetLastResponse()
	return ctx.R200OK(fmt.Sprintf("%v Pangaribuan", val))
}

func serverFuseRPrivateStatus2(ctx server.FuseContextR) any {
	auth := ctx.Auth().(string)
	_, val := ctx.GetLastResponse()

	data := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("%v %v", auth, val),
	}

	return ctx.R200OK(data)
}
