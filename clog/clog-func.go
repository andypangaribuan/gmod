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
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func mrf1[A any](key string, arg ...any) (va A) {
	arr := mainReflection(key, arg...)

	if v, ok := arr[0].(A); ok {
		va = v
	}

	return
}

func mrf2[A any, B any](key string, arg ...any) (va A, vb B) {
	arr := mainReflection(key, arg...)

	if v, ok := arr[0].(A); ok {
		va = v
	}

	if v, ok := arr[1].(B); ok {
		vb = v
	}

	return
}

func getConfVal[T any](name string) (value T) {
	val, err := mrf2[any, error]("mrf-conf-val", name)
	if err == nil {
		if v, ok := val.(T); ok {
			value = v
		}
	}
	return
}

func grpcCall[T any, R any](async bool, fn func(ctx context.Context, in *T, opts ...grpc.CallOption) (*R, error), logType string, req *T) (err error) {
	switch {
	case !async:
		_, err = call(fn, req)

	case logType == "":
		go doGrpcCall(fn, req)

	default:
		queueEntry(mrf1[string]("mrf-util-uid"), &stuQueue{
			logType: logType,
			req:     req,
		})
	}

	return
}

func doGrpcCall[T any, R any](fn func(ctx context.Context, in *T, opts ...grpc.CallOption) (*R, error), req *T) {
	startedAt := time.Now()

	for {
		_, err := call(fn, req)
		if err == nil {
			break
		}

		time.Sleep(time.Millisecond * 300)
		if time.Since(startedAt) > retryMaxDuration {
			break
		}
	}
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

	conn, err := mrf2[grpc.ClientConnInterface, error]("mrf-net-grpc-connection", address)
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

func queueEntry(uid string, sq *stuQueue) {
	mx.Lock()
	defer mx.Unlock()

	queue[uid] = sq
}

func queueSize() int {
	mx.RLock()
	defer mx.RUnlock()
	return len(queue)
}

func queueList() []*stuQueue {
	mx.Lock()
	defer mx.Unlock()

	i := -1
	ls := make([]*stuQueue, len(queue))

	for _, v := range queue {
		i++
		ls[i] = v
	}

	queue = make(map[string]*stuQueue, 0)
	return ls
}
