/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func FuseR(restPort int, routes func(router RouterR)) {
	serverImpl.FuseR(restPort, routes)
}
func FuseRS(restPort int, routes func(router RouterR), ws func(router RouterS)) {
	serverImpl.FuseRS(restPort, routes, ws)
}

func FuseG(grpcPort int, routes func(router RouterG)) {
	serverImpl.FuseG(grpcPort, routes)
}

func FuseGR(grpcPort int, grpcRoutes func(router RouterG), restPort int, restRoutes func(router RouterR)) {
	serverImpl.FuseGR(grpcPort, grpcRoutes, restPort, restRoutes)
}
