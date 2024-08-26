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
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type stuServer struct {
	logSpace string
}

type stuGrpcServerHandler struct {
	stackTraceSkipLevel int
}

type stuCronRouter struct {
	everyItems     []*stuCronEveryItem
	everyNItems    []*stuCronNEveryItem
	everyDayItems  []*stuCronEveryDayItem
	everyDayNItems []*stuCronNEveryDayItem
}

type stuCronEveryItem struct {
	uid            string
	duration       string
	fn             func()
	startUpDelayed *time.Duration
	allowParallel  bool
}

type stuCronNEveryItem struct {
	uid            string
	duration       string
	fns            []func()
	startUpDelayed *time.Duration
	allowParallel  bool
}

type stuCronEveryDayItem struct {
	uid           string
	at            string
	fn            func()
	allowParallel bool
}

type stuCronNEveryDayItem struct {
	uid           string
	at            string
	fns           []func()
	allowParallel bool
}

type stuFuseRRouter struct {
	fiberApp        *fiber.App
	withAutoRecover bool
	printOnError    bool
	errorHandler    func(FuseRContext, error) any
	noLogPaths      map[string]any
}

type stuFuseGRouter struct {
	server              *grpc.Server
	withAutoRecover     bool
	stackTraceSkipLevel int
	fnGetServer         func() *grpc.Server
}

type stuFuseSRouter struct {
	app      *fiber.App
	locals   map[string]string
	fnLocals *func(sl FuseSLocal)
}

type stuFuseSLocal struct {
	router *stuFuseSRouter
}

type stuFuseSContext struct {
	conn *websocket.Conn
}

type stuFuseRMainContext struct {
	startedAt    time.Time
	fcx          *fiber.Ctx
	clog         clog.Instance
	handlers     []func(FuseRContext) any
	errorHandler func(FuseRContext, error) any

	clientIP  string
	authObj   any
	userId    any
	partnerId any
	files     *map[string]string

	val          *stuFuseRVal
	responseVal  any
	responseMeta ResponseMeta
	responseRaw  bool
	execPath     string
	execFunc     string
	errMessage   *string
	stackTrace   *string
}

type stuFuseRContext struct {
	mcx *stuFuseRMainContext

	header     *map[string]string
	param      *map[string]string
	queries    *map[string]string
	form       *map[string][]string
	file       *map[string][]*multipart.FileHeader
	bodyParser func(out any) error

	responseVal  any
	responseMeta ResponseMeta
	responseRaw  bool
}

type stuFuseRVal struct {
	endpoint string
	url      string
	clientIP string
	unrouted bool

	bodyParser func(out any) error

	header  *map[string]string
	param   *map[string]string
	queries *map[string]string
	form    *map[string][]string
	file    *map[string][]*multipart.FileHeader

	fromSvcName    *string
	fromSvcVersion *string
	reqVersion     *string
	reqSource      *string
	reqHeader      *string
	reqParam       *string
	reqQuery       *string
	reqForm        *string
	reqBody        *string
}

type stuFuseRRegulator struct {
	mcx                   *stuFuseRMainContext
	currentIndex          int
	currentHandlerContext *stuFuseRContext
}

type stuFuseRCallOpt struct {
	header *map[string]string
	param  *map[string]string
	query  *map[string]string
	form   *map[string][]string
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
