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

	"github.com/andypangaribuan/gmod/gm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GrpcClient[T any](address string, fn func(cc grpc.ClientConnInterface) T) (T, error) {
	var client T

	conn, err := gm.Net.GrpcConnection(address)
	if err != nil {
		return client, err
	}

	return fn(conn), nil
}

func GrpcCall[T any, R any](fn func(ctx context.Context, in *T, opts ...grpc.CallOption) (*R, error), req *T, header ...map[string]string) (*R, error) {
	ctx := context.Background()
	if len(header) > 0 && len(header[0]) > 0 {
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(header[0]))
	}

	return fn(ctx, req)
}
