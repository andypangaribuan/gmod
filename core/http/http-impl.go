/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package http

import (
	"github.com/andypangaribuan/gmod/ice"
	"github.com/go-resty/resty/v2"
)

func (*stuHttp) Get(url string) ice.HttpBuilder {
	return newHttp(url, "get")
}

func (*stuHttp) Post(url string) ice.HttpBuilder {
	return newHttp(url, "post")
}

func (*stuHttp) Put(url string) ice.HttpBuilder {
	return newHttp(url, "put")
}

func (*stuHttp) Patch(url string) ice.HttpBuilder {
	return newHttp(url, "patch")
}

func (*stuHttp) Delete(url string) ice.HttpBuilder {
	return newHttp(url, "delete")
}

func newHttp(url, method string) ice.HttpBuilder {
	return &stuHttpBuilder{
		url:         url,
		method:      method,
		client:      resty.New(),
		fileReaders: make([]*stuFileReader, 0),
	}
}
