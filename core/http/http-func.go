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
	"bytes"
	"encoding/json"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
)

func newHttp(clog clog.Instance, url, method string) ice.HttpBuilder {
	return &stuHttpBuilder{
		clog:        clog,
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

func getJsonIndent(args *map[string]string) *string {
	if args != nil && len(*args) > 0 {
		data, err := gm.Json.Marshal(*args)
		if err == nil {
			var out bytes.Buffer
			err = json.Indent(&out, data, "", "  ")
			if err == nil {
				return fm.Ptr(out.String())
			}
		}
	}

	return nil
}
