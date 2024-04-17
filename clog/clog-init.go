/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import (
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

func xinit() {
	mainCLogCallback = func() {
		val, err := gm.Util.ReflectionGet(gm.Conf, "clogAddress")
		if err == nil {
			if v, ok := val.(string); ok {
				v = strings.TrimSpace(v)
				if v != "" {
					c, err := fm.GrpcClient(v, sclog.NewCLogServiceClient)
					if err == nil {
						client = c
					}
				}
			}
		}
	}
}
