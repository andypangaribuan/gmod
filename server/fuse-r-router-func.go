/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"fmt"
	"log"
	"strings"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/gofiber/fiber/v2"
)

func (slf *stuFuseRRouter) register(endpoint string, regulator func(FuseRRegulator), handlers ...func(FuseRContext) any) {
	index := strings.Index(endpoint, ":")
	if index == -1 {
		log.Fatalln("fuse server [restful]: endpoint format must be ▶︎ {Method}: {path}")
	}

	method := endpoint[0:index]
	path := strings.TrimSpace(endpoint[index+1:])

	switch method {
	case "GET":
		slf.fiberApp.Get(path, slf.restProcess(endpoint, regulator, handlers...))
	case "POS":
		slf.fiberApp.Post(path, slf.restProcess(endpoint, regulator, handlers...))
	case "DEL":
		slf.fiberApp.Delete(path, slf.restProcess(endpoint, regulator, handlers...))
	case "PUT":
		slf.fiberApp.Put(path, slf.restProcess(endpoint, regulator, handlers...))
	case "PAT":
		slf.fiberApp.Patch(path, slf.restProcess(endpoint, regulator, handlers...))
	default:
		log.Fatalln("fuse server [restful]: only support method GET, POS, DEL, PUT or PAT")
	}
}

func (slf *stuFuseRRouter) restProcess(endpoint string, regulator func(FuseRRegulator), handlers ...func(FuseRContext) any) func(*fiber.Ctx) error {
	return func(fcx *fiber.Ctx) error {
		slf.execute(fcx, endpoint, regulator, handlers...)
		return nil
	}
}

func (slf *stuFuseRRouter) execute(fcx *fiber.Ctx, endpoint string, regulator func(FuseRRegulator), handlers ...func(FuseRContext) any) {
	var (
		mcx = &stuFuseRMainContext{
			startedAt:    gm.Util.Timenow(),
			fcx:          fcx,
			clog:         clogNew(),
			handlers:     handlers,
			errorHandler: slf.errorHandler,
			val: &stuFuseRVal{
				endpoint: endpoint,
				url:      fcx.Request().URI().String(),
				clientIP: cip.getClientIP(fcx),
			},
		}

		contentType string
		reqHeader   = make(map[string]string, 0)
		reqParam    = fcx.AllParams()
		reqQueries  = fcx.Queries()
		reqBody     = fcx.Request().Body()
	)

	for key, ls := range fcx.GetReqHeaders() {
		key = strings.TrimSpace(strings.ToLower(key))
		val := ""
		if len(ls) > 0 {
			val = strings.TrimSpace(ls[0])
		}

		if val != "" {
			reqHeader[key] = val

			switch key {
			case "content-type":
				ls := strings.Split(val, ";")
				contentType = ls[0]

			case "x-from-svcname":
				mcx.val.fromSvcName = &val

			case "x-from-svcversion":
				mcx.val.fromSvcVersion = &val

			case "x-version":
				mcx.val.reqVersion = &val
			}
		}
	}

	if len(reqHeader) > 0 {
		mcx.val.header = &reqHeader
		jons, err := gm.Json.Encode(reqHeader)
		if err == nil {
			mcx.val.reqHeader = &jons
		}
	}

	if len(reqParam) > 0 {
		mcx.val.param = &reqParam
		jons, err := gm.Json.Encode(reqParam)
		if err == nil {
			mcx.val.reqParam = &jons
		}
	}

	if len(reqQueries) > 0 {
		mcx.val.queries = &reqQueries
		jons, err := gm.Json.Encode(reqQueries)
		if err == nil {
			mcx.val.reqQuery = &jons
		}
	}

	part, err := fcx.MultipartForm()
	if err == nil && part != nil {
		mcx.val.file = &part.File
		mcx.val.form = &part.Value
		form := fm.MapCopy(part.Value)

		for key, ls := range part.File {
			arr := make([]string, 0)
			for _, header := range ls {
				arr = append(arr, header.Filename)
			}

			if len(arr) > 0 {
				form[fmt.Sprintf("file-header: %v", key)] = arr
			}
		}

		if len(form) > 0 {
			jons, err := gm.Json.Encode(form)
			if err == nil {
				mcx.val.reqForm = &jons
			}
		}
	}

	if len(reqBody) > 0 {
		mcx.val.body = reqBody
		if contentType == "application/json" {
			mcx.val.reqBody = fm.Ptr(string(reqBody))
		}
	}

	if mcx.clog != nil {
		mol := &clog.ServicePieceV1{
			SvcParentName:    mcx.val.fromSvcName,
			SvcParentVersion: mcx.val.fromSvcVersion,
			Endpoint:         mcx.val.endpoint,
			Url:              mcx.val.url,
			ReqVersion:       mcx.val.reqVersion,
			ReqHeader:        mcx.val.reqHeader,
			ReqParam:         mcx.val.reqParam,
			ReqQuery:         mcx.val.reqQuery,
			ReqForm:          mcx.val.reqForm,
			ReqBody:          mcx.val.reqBody,
			ClientIp:         mcx.val.clientIP,
			StartedAt:        mcx.startedAt,
		}

		mcx.clog.ServicePieceV1(mol)
	}

	defer func() {
		if mcx.clog != nil {
			var (
				reqFiles *string
				resData  *string
			)

			if mcx.files != nil {
				jons, err := gm.Json.Encode(mcx.files)
				if err == nil {
					reqFiles = &jons
				}
			}

			if mcx.responseVal != nil {
				jons, err := gm.Json.Encode(mcx.responseVal)
				if err == nil {
					resData = &jons
				}
			}

			mol := &clog.ServiceV1{
				UserId:           mcx.getUserId(),
				PartnerId:        mcx.getPartnerId(),
				SvcParentName:    mcx.val.fromSvcName,
				SvcParentVersion: mcx.val.fromSvcVersion,
				Endpoint:         mcx.val.endpoint,
				Url:              mcx.val.url,
				Severity:         mcx.severity(),
				ExecPath:         mcx.execPath,
				ExecFunc:         mcx.execFunc,
				ReqVersion:       mcx.val.reqVersion,
				ReqHeader:        mcx.val.reqHeader,
				ReqParam:         mcx.val.reqParam,
				ReqQuery:         mcx.val.reqQuery,
				ReqForm:          mcx.val.reqForm,
				ReqFiles:         reqFiles,
				ReqBody:          mcx.val.reqBody,
				ResData:          resData,
				ResCode:          mcx.responseCode,
				ErrMessage:       mcx.errMessage,
				StackTrace:       mcx.stackTrace,
				ClientIp:         mcx.val.clientIP,
				StartedAt:        mcx.startedAt,
				FinishedAt:       gm.Util.Timenow(),
			}

			mcx.clog.ServiceV1(mol)
		}
	}()

	if regulator != nil {
		regulator(mcx.regulator())
	} else {
		slf.defaultHandlerRegulator(mcx.regulator())
	}
}

func (slf *stuFuseRRouter) defaultHandlerRegulator(regulator FuseRRegulator) {
	defer regulator.Recover()

	for {
		next, handler := regulator.Next()
		if !next {
			break
		}

		code, _ := regulator.Call(handler)
		if code == -1 {
			return
		}

		if code < 200 || code > 299 {
			break
		}
	}
}
