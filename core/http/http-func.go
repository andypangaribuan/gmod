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
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
)

func newHttp(url, method string) ice.HttpBuilder {
	return &stuHttpBuilder{
		url:         url,
		method:      method,
		fileReaders: make([]*stuFileReader, 0),
	}
}

func getConfVal[T any](name string) (value T) {
	v, err := gm.Util.ReflectionGet(gm.Conf, name)
	if err == nil {
		if v, ok := v.(T); ok {
			value = v
		}
	}
	return
}

func (*stuHttp) isInternalSvc(url string) bool {
	var (
		urlLength = len(url)
		length    = 0
	)

	for _, base := range internalBaseUrls {
		length = len(base)
		if urlLength >= len(base) {
			if url[:length] == base {
				return true
			}
		}
	}
	return false
}
