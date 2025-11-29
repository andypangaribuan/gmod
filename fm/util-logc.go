/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import (
	"fmt"
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/pkg/errors"
)

func logcSaveGrpcError(startedAt time.Time, destination string, logc clog.Instance, req any, header map[string]string, code string, err error) {
	if logc != nil {
		execPath, execFunc := gm.Util.GetExecPathFunc(2)
		data := map[string]any{
			"request":    req,
			"error-code": code,
		}

		var (
			errMessage *string
			stackTrace *string
		)

		if err != nil {
			errMessage = Ptr(err.Error())
			stackTrace = Ptr(fmt.Sprintf("%+v", errors.WithStack(err)))
		}

		jsonHeader, headerErr := gm.Json.Encode(header)
		jsonData, dataErr := gm.Json.Encode(data)

		_ = logc.GrpcV1(&clog.GrpcV1{
			Destination: destination,
			Severity:    "error",
			ExecPath:    execPath,
			ExecFunc:    execFunc,
			ReqHeader:   Ternary(headerErr == nil, &jsonHeader, nil),
			Data:        Ternary(dataErr == nil, &jsonData, nil),
			ErrMessage:  errMessage,
			StackTrace:  stackTrace,
			StartedAt:   startedAt,
			FinishedAt:  gm.Util.Timenow(),
		})
	}
}

func logcSaveGrpcSuccess(startedAt time.Time, destination string, logc clog.Instance, req any, header map[string]string, res any) {
	if logc != nil {
		execPath, execFunc := gm.Util.GetExecPathFunc(2)
		data := map[string]any{
			"request":  req,
			"response": res,
		}

		jsonHeader, headerErr := gm.Json.Encode(header)
		jsonData, dataErr := gm.Json.Encode(data)

		_ = logc.GrpcV1(&clog.GrpcV1{
			Destination: destination,
			Severity:    "success",
			ExecPath:    execPath,
			ExecFunc:    execFunc,
			ReqHeader:   Ternary(headerErr == nil, &jsonHeader, nil),
			Data:        Ternary(dataErr == nil, &jsonData, nil),
			StartedAt:   startedAt,
			FinishedAt:  gm.Util.Timenow(),
		})
	}
}
