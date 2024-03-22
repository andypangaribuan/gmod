/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import "time"

func (slf *srServer) FuseGR(grpcPort int, grpcRoutes func(router RouterG), restPort int, restRoutes func(router RouterR)) {
	go slf.FuseG(grpcPort, grpcRoutes)
	time.Sleep(time.Millisecond * 100)
	slf.FuseR(restPort, restRoutes)
}
