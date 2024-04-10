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
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (slf *stuGrpcServerHandler) unaryPanicHandler(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	defer slf.crashHandler(func(r any) {
		err = slf.toPanicError(r)
	})

	return handler(ctx, req)
}

func (slf *stuGrpcServerHandler) streamPanicHandler(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer slf.crashHandler(func(r any) {
		err = slf.toPanicError(r)
	})

	return handler(srv, stream)
}

func (slf *stuGrpcServerHandler) toPanicError(r any) error {
	return status.Errorf(codes.Internal, "panic: %v", r)
}

func (slf *stuGrpcServerHandler) crashHandler(handler func(any)) {
	if r := recover(); r != nil {
		handler(r)
		slf.printPanic(r)
	}
}

func (slf *stuGrpcServerHandler) printPanic(r any) {
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

	message := fmt.Sprintf(strings.TrimSpace(`
## recovered from panic
## detail\n[[%#v]]
## stacktrace:
`), r)

	// log.Printf("## recovered from panic\n")
	// log.Printf("## detail\n[[%#v]]\n", r)
	// log.Printf("## stacktrace:\n")

	startIndex := 0
	if slf.stackTraceSkipLevel > 0 {
		if slf.stackTraceSkipLevel < len(callers) {
			startIndex = slf.stackTraceSkipLevel
		}
	}

	message += "\n"
	for i := startIndex; len(callers) > i; i++ {
		message += callers[i]
		// log.Printf(" %v", callers[i])
	}
	message += "\n\n"

	log.Print(message)
}
