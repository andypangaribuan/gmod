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

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"google.golang.org/grpc"
)

func getConfValue(name string) (value string) {
	val, err := gm.Util.ReflectionGet(gm.Conf, name)
	if err == nil {
		if v, ok := val.(string); ok {
			value = strings.TrimSpace(v)
		}
	}
	return
}

func grpcCall[T any, R any](async bool, fn func(ctx context.Context, in *T, opts ...grpc.CallOption) (*R, error), req *T, header ...map[string]string) (err error) {
	if !async {
		_, err = fm.GrpcCall(fn, req, header...)
	} else {
		go func ()  {
			for {
				_, err = fm.GrpcCall(fn, req, header...)
				if err == nil {
					break
				}
			}
		}()
	}

	return
}
