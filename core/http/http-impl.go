/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package http

import (
	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/ice"
)

func (*stuHttp) Get(clog clog.Instance, url string) ice.HttpBuilder {
	return newHttp(clog, url, "get")
}

func (*stuHttp) Post(clog clog.Instance, url string) ice.HttpBuilder {
	return newHttp(clog, url, "post")
}

func (*stuHttp) Put(clog clog.Instance, url string) ice.HttpBuilder {
	return newHttp(clog, url, "put")
}

func (*stuHttp) Patch(clog clog.Instance, url string) ice.HttpBuilder {
	return newHttp(clog, url, "patch")
}

func (*stuHttp) Delete(clog clog.Instance, url string) ice.HttpBuilder {
	return newHttp(clog, url, "delete")
}
