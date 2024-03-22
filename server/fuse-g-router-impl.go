/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import "google.golang.org/grpc"

func (slf *srFuseRouterG) AutoRecover(autoRecover bool) {
	slf.withAutoRecover = autoRecover
}

func (slf *srFuseRouterG) Server() *grpc.Server {
	return slf.fnGetServer()
}
