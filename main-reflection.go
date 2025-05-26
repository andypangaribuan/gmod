/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package gmod

import (
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
)

func rx(val ...any) []any {
	return val
}

func reflection(key string, arg ...any) []any {
	switch key {
	case "mrf-ice-net":
		return rx(iceNet)

	case "mrf-conf-val":
		return rx(iceUtil.ReflectionGet(gm.Conf, arg[0].(string)))

	case "mrf-net-grpc-connection":
		return rx(iceNet.GrpcConnection(arg[0].(string)))

	case "mrf-util-uid":
		return rx(iceUtil.UID())

	case "mrf-util-get-exec-path-func":
		return rx(iceUtil.GetExecPathFunc(arg[0].(int)))

	case "mrf-util-concurrent-process":
		iceUtil.ConcurrentProcess(arg[0].(int), arg[1].(int), arg[2].(func(index int)))

	case "mrf-util-x-concurrent-process":
		return rx(iceUtil.XConcurrentProcess(arg[0].(int), arg[1].(int)))

	case "mrf-util-x-concurrent-process-run":
		ucp := arg[0].(ice.UtilConcurrentProcess)
		ucp.Run(arg[1].(int), arg[2].(func(index int)))

	case "mrf-util-x-concurrent-process-prune":
		ucp := arg[0].(ice.UtilConcurrentProcess)
		ucp.Prune()
	}

	return rx()
}
