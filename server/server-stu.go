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
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type stuServer struct {
	logSpace string
}

type stuGrpcServerHandler struct {
	stackTraceSkipLevel int
}

type stuFuseRouterR struct {
	fiberApp        *fiber.App
	withAutoRecover bool
	printOnError    bool
}

type stuFuseRouterG struct {
	server              *grpc.Server
	withAutoRecover     bool
	stackTraceSkipLevel int
	fnGetServer         func() *grpc.Server
}

type stuFuseContextR struct {
	fiberCtx     *fiber.Ctx
	endpoint     string
	isRegulator  bool
	regulatorCtx *stuFuseContextRegulatorR
	authObj      any

	handlers         []func(ctx FuseContextR) error
	lastResponseCode int
	lastResponseVal  any
	responseCode     int
	responseVal      any
}

type stuFuseContextRegulatorR struct {
	fuseContext           *stuFuseContextR
	currentIndex          int
	currentHandlerContext *stuFuseContextR
}

type stuFuseContextBuilderR struct {
	original *stuFuseContextR
}
