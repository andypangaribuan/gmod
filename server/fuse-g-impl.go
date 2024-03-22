/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/andypangaribuan/gmod/gm"
	"google.golang.org/grpc"
)

func (*srServer) FuseG(grpcPort int, routes func(router RouterG)) {
	if gm.Net.IsPortUsed(grpcPort) {
		fmt.Printf("fuse server [grpc]: port %v already in use\n", grpcPort)
		os.Exit(100)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		fmt.Printf("fuse server [grpc]: failed to listen on port %v\n", grpcPort)
		os.Exit(100)
	}

	router := &srFuseRouterG{
		withAutoRecover:     false,
		stackTraceSkipLevel: 3,
	}

	router.fnGetServer = func() *grpc.Server {
		if router.server == nil {
			if !router.withAutoRecover {
				router.server = grpc.NewServer()
			} else {
				handler := &srGrpcServerHandler{
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

			if gm.Net.IsPortUsed(grpcPort) {
				fmt.Printf("fuse server [grpc]: run at port %v\n", grpcPort)
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	err = router.server.Serve(listener)
	if err != nil {
		isListenFailed = true
		fmt.Printf("fuse server [grpc]: failed to listen on port %v\n", grpcPort)
		os.Exit(100)
	}
}