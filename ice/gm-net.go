/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import (
	"github.com/andypangaribuan/gmod/mol"
	"google.golang.org/grpc"
)

type Net interface {
	IsPortUsed(port int, host ...string) bool
	GrpcConnection(address string, opt ...mol.NetOpt) (grpc.ClientConnInterface, error)
}
