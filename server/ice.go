/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

type server interface {
	FuseR(restPort int, routes func(router RouterR))
}

type RouterR interface {
	AutoRecover(autoRecover bool)
	PrintOnError(printOnError bool)
	Single(path string, handlers ...func(sc Context) error)
	Group(endpoints map[string][]func(sc Context) error)
	Unrouted(handler func(ctx Context, method, path, url string) error)
}

type Context interface {
}
