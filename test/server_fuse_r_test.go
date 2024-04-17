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

		router.Endpoints(nil, nil, map[string][]func(ctx server.FuseContextR){
			"GET: /private/status-1": {serverFuseRAuth, serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
			"GET: /private/status-2": {serverFuseRAuth, serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
		})

		router.Endpoints(nil, serverFuseRAuth, map[string][]func(ctx server.FuseContextR){
			"GET: /private/status-3": {serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
			"GET: /private/status-4": {serverRuseRPrivateStatus1, serverFuseRPrivateStatus2},
		})
	})
}

func serverFuseRAuth(ctx server.FuseContextR) {
	ctx.SetAuth("Halo")
	ctx.R200OK("Andy")
}

func serverRuseRPrivateStatus1(ctx server.FuseContextR) {
	_, val := ctx.GetLastResponse()
	ctx.R200OK(fmt.Sprintf("%v Pangaribuan", val))
}

func serverFuseRPrivateStatus2(ctx server.FuseContextR) {
	auth := ctx.Auth().(string)
	_, val := ctx.GetLastResponse()

	data := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("%v %v", auth, val),
	}

	ctx.R200OK(data)
}
