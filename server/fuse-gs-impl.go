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

func (slf *stuServer) FuseGS(grpcPort int, grpcRoutes func(router RouterG), wsPort int, wsRoutes func(router RouterS)) {
	slf.fuseGS(grpcPort, grpcRoutes, wsPort, &wsRoutes)
}

func (slf *stuServer) fuseGS(grpcPort int, grpcRoutes func(router RouterG), wsPort int, wsRoutes *func(router RouterS)) {
	if gm.Net.IsPortUsed(grpcPort) {
		fmt.Printf("fuse server [grpc]     : port %v already in use\n", grpcPort)
		os.Exit(100)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		fmt.Printf("fuse server [grpc]     : failed to listen on port %v\n", grpcPort)
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

	grpcRoutes(router)

	isListenFailed := false
	go func() {
		tryCount := 0
		maxTry := 30

		for !isListenFailed {
			tryCount++
			if tryCount > maxTry {
				break
			}

			if gm.Net.IsPortUsed(grpcPort) {
				fmt.Printf("fuse server [grpc]     : run at port %v\n", grpcPort)
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	go slf.runWS(wsPort, wsRoutes, router.withAutoRecover)

	err = router.server.Serve(listener)
	if err != nil {
		isListenFailed = true
		fmt.Printf("fuse server [grpc]     : failed to listen on port %v\n", grpcPort)
		os.Exit(100)
	}
}

func (slf *stuServer) runWS(wsPort int, wsRoutes *func(router RouterS), withAutoRecover bool) {
	if wsRoutes != nil {
		fuseFiberApp = fiber.New(fiber.Config{
			JSONEncoder:           gm.Json.Marshal,
			JSONDecoder:           gm.Json.Unmarshal,
			DisableStartupMessage: true,
		})

		if withAutoRecover {
			fuseFiberApp.Use(recover.New())
		}

		slf.startWebsocket(fuseFiberApp, wsRoutes)

		isListenFailed := false
		go func() {
			tryCount := 0
			maxTry := 30
			time.Sleep(time.Millisecond * 100)

			for !isListenFailed {
				tryCount++
				if tryCount > maxTry {
					break
				}

				if gm.Net.IsPortUsed(wsPort) {
					fmt.Printf("fuse server [websocket]: run at port %v\n", wsPort)
					break
				}

				time.Sleep(time.Millisecond * 100)
			}
		}()

		err := fuseFiberApp.Listen(fmt.Sprintf(":%v", wsPort))
		if err != nil {
			isListenFailed = true
			fmt.Printf("fuse server [websocket]: failed to listen on port %v\n", wsPort)
			os.Exit(100)
		}
	}
}
