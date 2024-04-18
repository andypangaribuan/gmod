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
	"google.golang.org/grpc"
)

func (slf *stuServer) FuseG(grpcPort int, routes func(router RouterG)) {
	if gm.Net.IsPortUsed(grpcPort) {
		fmt.Printf("fuse server [grpc]%v: port %v already in use\n", slf.logSpace, grpcPort)
		os.Exit(100)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		fmt.Printf("fuse server [grpc]%v: failed to listen on port %v\n", slf.logSpace, grpcPort)
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

			if gm.Net.IsPortUsed(grpcPort) {
				fmt.Printf("fuse server [grpc]%v: run at port %v\n", slf.logSpace, grpcPort)
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	err = router.server.Serve(listener)
	if err != nil {
		isListenFailed = true
		fmt.Printf("fuse server [grpc]%v: failed to listen on port %v\n", slf.logSpace, grpcPort)
		os.Exit(100)
	}
}
