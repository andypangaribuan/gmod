/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func getConfValue(name string) (value string) {
	val, err := mainCLogUtilReflectionGetConf(name)
	if err == nil {
		if v, ok := val.(string); ok {
			value = strings.TrimSpace(v)
		}
	}
	return
}

func grpcCall[T any, R any](async bool, fn func(ctx context.Context, in *T, opts ...grpc.CallOption) (*R, error), req *T, header ...map[string]string) (err error) {
	if !async {
		_, err = call(fn, req, header...)
	} else {
		go func() {
			for {
				_, err = call(fn, req, header...)
				if err == nil {
					break
				}
			}
		}()
	}

	return
}

func pbwString(val *string) *wrapperspb.StringValue {
	if val == nil {
		return nil
	}

	return &wrapperspb.StringValue{Value: *val}
}

func pbwInt32(val *int) *wrapperspb.Int32Value {
	if val == nil {
		return nil
	}

	return &wrapperspb.Int32Value{Value: int32(*val)}
}

func call[T any, R any](fn func(ctx context.Context, in *T, opts ...grpc.CallOption) (*R, error), req *T, header ...map[string]string) (*R, error) {
	ctx := context.Background()
	if len(header) > 0 && len(header[0]) > 0 {
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(header[0]))
	}

	return fn(ctx, req)
}

func createClient[T any](address string, fn func(cc grpc.ClientConnInterface) T) (T, error) {
	var client T

	conn, err := mainCLogNetGrpcConnection(address)
	if err != nil {
		return client, err
	}

	return fn(conn), nil
}

func getFirst[T any](ls []T, dval ...T) *T {
	if len(ls) == 0 {
		if len(dval) > 0 {
			return &dval[0]
		}

		return nil
	}

	return &ls[0]
}

func timeToStrFull(val time.Time) string {
	return val.Format("2006-01-02 15:04:05.999999 -07:00")
}
