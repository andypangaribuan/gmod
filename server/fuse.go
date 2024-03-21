/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

func FuseR(restPort int, routes func(router RouterR)) {
	serverImpl.FuseR(restPort, routes)
}
