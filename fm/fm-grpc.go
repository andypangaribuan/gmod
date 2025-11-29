/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import (
	"context"
	"strings"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GrpcClient[T any](address string, fn func(cc grpc.ClientConnInterface) T) (T, error) {
	var client T

	conn, err := mrf2[grpc.ClientConnInterface, error]("mrf-net-grpc-connection", address)
	if err != nil {
		return client, err
	}

	return fn(conn), nil
}

func GrpcCall[T any, R any](clog clog.Instance, fn func(ctx context.Context, in *T, opts ...grpc.CallOption) (*R, error), req *T, header ...map[string]string) (*R, error) {
	var (
		ctx        = context.Background()
		metaHeader = make(map[string]string, 0)
	)

	if clog != nil {
		uid, userId, partnerId := clogGetId(clog)
		if uid != "" {
			metaHeader["gmod-uid"] = uid
		}

		if userId != nil {
			metaHeader["gmod-user-id"] = *userId
		}

		if partnerId != nil {
			metaHeader["gmod-partner-id"] = *partnerId
		}
	}

	svcName, err := mrf2[string, error]("mrf-conf-val", "svcName")
	if err == nil {
		metaHeader["gmod-from-svc-name"] = svcName
	}

	svcVersion, err := mrf2[string, error]("mrf-conf-val", "svcVersion")
	if err == nil {
		metaHeader["gmod-from-svc-version"] = svcVersion
	}

	if len(header) > 0 && len(header[0]) > 0 {
		for k, v := range header[0] {
			metaHeader[k] = v
		}
	}

	if len(metaHeader) > 0 {
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(metaHeader))
	}

	return fn(ctx, req)
}

func GrpcHeader(ctx context.Context) map[string]string {
	header := make(map[string]string, 0)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for key, val := range md {
			if len(val) > 0 {
				header[key] = val[0]
			}
		}
	}

	return header
}

func GrpcClientIp(ctx context.Context) (clientIp string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		vals := md.Get("cf-connecting-ip")
		for _, v := range vals {
			v = strings.TrimSpace(v)
			if len(v) > 0 {
				return v
			}
		}

		vals = md.Get("x-original-forwarded-for")
		for _, v := range vals {
			v = strings.TrimSpace(v)
			if len(v) > 0 {
				return v
			}
		}

		vals = md.Get(":authority")
		for _, v := range vals {
			v = strings.ToLower(strings.TrimSpace(v))
			if strings.Contains(v, ":") {
				addr := strings.Split(v, ":")[0]
				if addr == "localhost" || addr == "127.0.0.1" {
					return "127.0.0.1"
				}
			}
		}
	}

	p, ok := peer.FromContext(ctx)
	if ok && p != nil {
		addr := p.Addr.String()
		if addr == "localhost" {
			return "127.0.0.1"
		}
		return addr
	}

	return
}

func GrpcExec[RQ any, RS any](destination string, ctx context.Context, req *RQ, exec func(send ice.GrpcSender[RS], logc clog.Instance, req *RQ) (*RS, error), onError func(code *wrapperspb.StringValue, message *wrapperspb.StringValue) *RS) (*RS, error) {
	var (
		header = GrpcHeader(ctx)
		logc   = clog.New(header)
		send   = grpcSender(ctx, destination, logc, req, header, onError)
	)

	return exec(send, logc, req)
}

func GrpcExec2[RQ any, RS any](send ice.GrpcSender[RS], exec func(send ice.GrpcSender[RS], logc clog.Instance, req *RQ) (*RS, error)) (*RS, error) {
	sender := send.(*stuGrpcSender[RQ, RS])
	return exec(send, sender.logc, sender.req)
}

func GrpcSender[RQ any, RS any](destination string, ctx context.Context, req *RQ, onError func(code *wrapperspb.StringValue, message *wrapperspb.StringValue) *RS) (logc clog.Instance, send ice.GrpcSender[RS]) {
	header := GrpcHeader(ctx)
	logc = clog.New(header)
	send = grpcSender(ctx, destination, logc, req, header, onError)

	return logc, send
}

func grpcSender[RQ any, RS any](ctx context.Context, destination string, logc clog.Instance, req *RQ, header map[string]string, getErrResponse func(code *wrapperspb.StringValue, message *wrapperspb.StringValue) *RS) ice.GrpcSender[RS] {
	return &stuGrpcSender[RQ, RS]{
		startedAt:      gm.Util.Timenow(),
		ctx:            ctx,
		destination:    destination,
		logc:           logc,
		header:         header,
		req:            req,
		getErrResponse: getErrResponse,
	}
}

func (slf *stuGrpcSender[RQ, RS]) Context() context.Context {
	return slf.ctx
}

func (slf *stuGrpcSender[RQ, RS]) Header() map[string]string {
	return slf.header
}

func (slf *stuGrpcSender[RQ, RS]) Error(code string, err error) (*RS, error) {
	var (
		errCode    = PbwString(&code)
		errMessage = TernaryR(err == nil, nil, func() *wrapperspb.StringValue { return PbwString(Ptr(err.Error())) })
	)

	logcSaveGrpcError(slf.startedAt, slf.destination, slf.logc, slf.req, slf.header, code, err)
	return slf.getErrResponse(errCode, errMessage), nil
}

func (slf *stuGrpcSender[RQ, RS]) Success(result *RS) (*RS, error) {
	logcSaveGrpcSuccess(slf.startedAt, slf.destination, slf.logc, slf.req, slf.header, result)
	return result, nil
}
