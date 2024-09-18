/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func (slf *stuServer) FuseG(grpcPort int, routes func(router RouterG)) {
	slf.fuseGS(grpcPort, routes, 0, nil)
}
