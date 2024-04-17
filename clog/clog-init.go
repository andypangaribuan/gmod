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
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

func xinit() {
	mainCLogCallback = func() {
		if val := getConfValue("clogAddress"); val != "" {
			c, err := fm.GrpcClient(val, sclog.NewCLogServiceClient)
			if err == nil {
				client = c
			}
		}

		if val := getConfValue("svcName"); val != "" {
			svcName = val
		}

		if val := getConfValue("svcVersion"); val != "" {
			svcVersion = val
		}
	}
}
