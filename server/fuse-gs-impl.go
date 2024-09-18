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
	"net"
	"os"
	"time"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"google.golang.org/grpc"
)

func (slf *stuServer) FuseGS(port int, routes func(router RouterG), ws func(router RouterS)) {
	slf.fuseGS(port, routes, &ws)
}

func (slf *stuServer) fuseGS(port int, routes func(router RouterG), ws *func(router RouterS)) {
	if gm.Net.IsPortUsed(port) {
		fmt.Printf("fuse server [grpc]%v: port %v already in use\n", slf.logSpace, port)
		os.Exit(100)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Printf("fuse server [grpc]%v: failed to listen on port %v\n", slf.logSpace, port)
		os.Exit(100)
	}

	router := &stuFuseGRouter{
		withAutoRecover:     false,
		stackTraceSkipLevel: 3,
	}

	router.fnGetServer = func() *grpc.Server {
		if router.server == nil {
			if !router.withAutoRecover {
				router.server = grpc.NewServer()
			} else {
				handler := &stuGrpcServerHandler{
					stackTraceSkipLevel: router.stackTraceSkipLevel,
				}

				uIntOpt := grpc.UnaryInterceptor(handler.unaryPanicHandler)
				sIntOpt := grpc.StreamInterceptor(handler.streamPanicHandler)
				router.server = grpc.NewServer(uIntOpt, sIntOpt)
			}
		}

		return router.server
	}

	routes(router)

	isListenFailed := false
	go func() {
		tryCount := 0
		maxTry := 30

		for {
			if isListenFailed {
				break
			}

			tryCount++
			if tryCount > maxTry {
				break
			}

			if gm.Net.IsPortUsed(port) {
				fmt.Printf("fuse server [grpc]%v: run at port %v\n", slf.logSpace, port)
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	if ws != nil {
		fuseFiberApp = fiber.New(fiber.Config{
			JSONEncoder:           gm.Json.Marshal,
			JSONDecoder:           gm.Json.Unmarshal,
			DisableStartupMessage: true,
		})

		if router.withAutoRecover {
			fuseFiberApp.Use(recover.New())
		}

		slf.startWebsocket(fuseFiberApp, ws)
	}

	err = router.server.Serve(listener)
	if err != nil {
		isListenFailed = true
		fmt.Printf("fuse server [grpc]%v: failed to listen on port %v\n", slf.logSpace, port)
		os.Exit(100)
	}
}
