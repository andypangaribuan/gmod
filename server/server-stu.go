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
	"mime/multipart"
	"net"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type stuServer struct {
	logSpace string
}

type stuGrpcServerHandler struct {
	stackTraceSkipLevel int
}

type stuFuseRRouter struct {
	fiberApp        *fiber.App
	withAutoRecover bool
	printOnError    bool
	errorHandler    func(clog.Instance, FuseRContext, error) error
}

type stuFuseGRouter struct {
	server              *grpc.Server
	withAutoRecover     bool
	stackTraceSkipLevel int
	fnGetServer         func() *grpc.Server
}

type stuFuseRContext struct {
	fiberCtx     *fiber.Ctx
	clog         clog.Instance
	isRegulator  bool
	regulatorCtx *stuFuseRRegulator
	authObj      any

	errorHandler func(clog.Instance, FuseRContext, error) error

	handlers         []func(clog.Instance, FuseRContext) error
	lastResponseCode int
	lastResponseVal  any
	responseCode     int
	responseVal      any

	header  *map[string]string
	param   *map[string]string
	queries *map[string]string
	form    *map[string][]string
	file    *map[string][]*multipart.FileHeader
	val     *stuFuseRVal
}

type stuFuseRVal struct {
	endpoint string
	url      string
	clientIP string

	fromSvcName    *string
	fromSvcVersion *string
	reqVersion     *string
	reqHeader      *string
	reqParam       *string
	reqQuery       *string
	reqForm        *string
}

type stuFuseRContextBuilder struct {
	original *stuFuseRContext
}

type stuFuseRRegulator struct {
	clog                  clog.Instance
	fuseContext           *stuFuseRContext
	currentIndex          int
	currentHandlerContext *stuFuseRContext
}

type stuFuseRCallOpt struct {
	header *map[string]string
}

type stuClientIP struct {
	cidrs                       []*net.IPNet
	xOriginalForwardedForHeader string
	xForwardedForHeader         string
	xForwardedHeader            string
	forwardedForHeader          string
	forwardedHeader             string
	xClientIPHeader             string
	xRealIPHeader               string
	cfConnectingIPHeader        string
	fastlyClientIPHeader        string
	trueClientIPHeader          string
}
