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
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	}

	return
}
