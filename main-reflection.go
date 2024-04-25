/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package gmod

import "github.com/andypangaribuan/gmod/gm"

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
	}

	return rx()
}
