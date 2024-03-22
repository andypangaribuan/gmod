/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (slf *srGrpcServerHandler) unaryPanicHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer slf.crashHandler(func(r interface{}) {
		err = slf.toPanicError(r)
	})

	return handler(ctx, req)
}

func (slf *srGrpcServerHandler) streamPanicHandler(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer slf.crashHandler(func(r interface{}) {
		err = slf.toPanicError(r)
	})

	return handler(srv, stream)
}

func (slf *srGrpcServerHandler) toPanicError(r interface{}) error {
	return status.Errorf(codes.Internal, "panic: %v", r)
}

func (slf *srGrpcServerHandler) crashHandler(handler func(interface{})) {
	if r := recover(); r != nil {
		handler(r)
		slf.printPanic(r)
	}
}

func (slf *srGrpcServerHandler) printPanic(r interface{}) {
	var callers []string

	for i := 0; true; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		callers = append(callers, fmt.Sprintf("%d: %v:%v\n", i, file, line))
	}

	if len(callers) == 0 {
		return
	}

	log.Printf("## recovered from panic\n")
	log.Printf("## detail\n[[%#v]]\n", r)
	log.Printf("## stacktrace:\n")

	startIndex := 0
	if slf.stackTraceSkipLevel > 0 {
		if slf.stackTraceSkipLevel < len(callers) {
			startIndex = slf.stackTraceSkipLevel
		}
	}

	for i := startIndex; len(callers) > i; i++ {
		log.Printf(" %v", callers[i])
	}

	fmt.Println()
}
