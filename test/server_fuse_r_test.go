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

		router.Endpoints(map[string][]func(ctx server.FuseContextR){
			"GET: /private/status-1": {serverFuseRAuth, serverFuseRPrivateStatus},
			"GET: /private/status-2": {serverFuseRAuth, serverFuseRPrivateStatus},
		})

		router.EndpointsWithAuth(serverFuseRAuth, map[string][]func(ctx server.FuseContextR){
			"GET: /private/status-3": {serverFuseRPrivateStatus},
			"GET: /private/status-4": {serverFuseRPrivateStatus},
		})
	})
}

func serverFuseRAuth(ctx server.FuseContextR) {
	ctx.SetAuth("Andy")
	ctx.R200OK("Pangaribuan")
}

func serverFuseRPrivateStatus(ctx server.FuseContextR) {
	auth := ctx.Auth().(string)
	_, val := ctx.GetResponse()

	data := struct {
		From string `json:"from"`
	}{
		From: fmt.Sprintf("%v %v", auth, val),
	}

	ctx.R200OK(data)
}
