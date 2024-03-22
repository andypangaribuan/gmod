/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type srServer struct{}

type srGrpcServerHandler struct {
	stackTraceSkipLevel int
}

type srFuseRouterR struct {
	fiberApp        *fiber.App
	withAutoRecover bool
	printOnError    bool
}

type srFuseRouterG struct {
	server              *grpc.Server
	withAutoRecover     bool
	stackTraceSkipLevel int
	fnGetServer         func() *grpc.Server
}

type srFuseContextR struct {
	fiberCtx     *fiber.Ctx
	endpoint     string
	isRegulator  bool
	regulatorCtx *srFuseContextRegulatorR

	controllers  []func(ctx FuseContextR)
	responseCode int
	responseObj  interface{}
}

type srFuseContextRegulatorR struct {
	fuseContext              *srFuseContextR
	currentIndex             int
	currentControllerContext *srFuseContextR
}

type srFuseContextBuilderR struct {
	original *srFuseContextR
}
