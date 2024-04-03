/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import "google.golang.org/grpc"

func (slf *stuFuseRouterG) AutoRecover(autoRecover bool) {
	slf.withAutoRecover = autoRecover
}

func (slf *stuFuseRouterG) Server() *grpc.Server {
	return slf.fnGetServer()
}
