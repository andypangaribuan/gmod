/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "time"

func (slf *stuServer) FuseGR(grpcPort int, grpcRoutes func(router RouterG), restPort int, restRoutes func(router RouterR)) {
	go slf.FuseG(grpcPort, grpcRoutes)
	time.Sleep(time.Millisecond * 100)
	slf.FuseR(restPort, restRoutes)
}
