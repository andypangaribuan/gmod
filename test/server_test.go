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

func TestServerGRPC(t *testing.T) {
	server.FuseG(11011, func(router server.RouterG) {
		router.AutoRecover(false)
		router.Server()
		fmt.Printf("grpc routes\n")
	})
}
