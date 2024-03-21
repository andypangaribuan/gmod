/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

func (self *srFuseRouterR) AutoRecover(autoRecover bool) {
	self.withAutoRecover = autoRecover
}

func (self *srFuseRouterR) PrintOnError(printOnError bool) {
	self.printOnError = printOnError
}

func (self *srFuseRouterR) Single(path string, handlers ...func(sc Context) error) {

}

func (self *srFuseRouterR) Group(endpoints map[string][]func(sc Context) error) {}

func (self *srFuseRouterR) Unrouted(handler func(ctx Context, method, path, url string) error) {}
