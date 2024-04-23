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
	"github.com/andypangaribuan/gmod/ice"
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

func (slf *stuHttp) GetJsonHeader(url string, opt ...any) map[string]string {
	var (
		reqVer *string
		add    = make(map[string]string, 0)
		header = map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		}
	)

	if len(opt) > 0 {
		for _, o := range opt {
			switch v := o.(type) {
			case string:
				reqVer = &v
			case *string:
				reqVer = v

			case map[string]string:
				add = v
			case *map[string]string:
				if v != nil {
					add = *v
				}
			}
		}
	}

	if slf.isInternalSvc(url) {
		header["X-From-SvcName"] = svcName
		header["X-From-SvcVersion"] = svcVersion
		if reqVer != nil {
			header["X-Version"] = *reqVer
		}
	}

	for k, v := range add {
		header[k] = v
	}

	return header
}
