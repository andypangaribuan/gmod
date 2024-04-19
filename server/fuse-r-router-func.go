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
	"github.com/andypangaribuan/gmod/gm"
	"github.com/gofiber/fiber/v2"
)

func (slf *stuFuseRRouter) register(endpoint string, regulator func(clog.Instance, FuseRRegulator), handlers ...func(clog.Instance, FuseRContext) error) {
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

func (slf *stuFuseRRouter) restProcess(endpoint string, regulator func(clog.Instance, FuseRRegulator), handlers ...func(clog.Instance, FuseRContext) error) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		slf.execute(fiberCtx, endpoint, regulator, handlers...)
		return nil
	}
}

func (slf *stuFuseRRouter) execute(fiberCtx *fiber.Ctx, endpoint string, regulator func(clog.Instance, FuseRRegulator), handlers ...func(clog.Instance, FuseRContext) error) {
	var (
		startedAt = gm.Util.Timenow()
		original  = &stuFuseRContext{
			fiberCtx:     fiberCtx,
			clog:         clogNew(),
			isRegulator:  true,
			handlers:     handlers,
			errorHandler: slf.errorHandler,
			val: &stuFuseRVal{
				endpoint: endpoint,
				url:      fiberCtx.Request().URI().String(),
				clientIP: cip.getClientIP(fiberCtx),
			},
		}

		reqHeader  = make(map[string]string, 0)
		reqParam   = fiberCtx.AllParams()
		reqQueries = fiberCtx.Queries()
	)

	for key, ls := range fiberCtx.GetReqHeaders() {
		key = strings.TrimSpace(strings.ToLower(key))
		val := ""
		if len(ls) > 0 {
			val = strings.TrimSpace(ls[0])
		}

		if val != "" {
			reqHeader[key] = val

			switch key {
			case "x-from-svcname":
				original.val.fromSvcName = &val

			case "x-from-svcversion":
				original.val.fromSvcVersion = &val

			case "x-version":
				original.val.reqVersion = &val
			}
		}
	}

	if len(reqHeader) > 0 {
		original.header = &reqHeader
		jons, err := gm.Json.Encode(reqHeader)
		if err == nil {
			original.val.reqHeader = &jons
		}
	}

	if len(reqParam) > 0 {
		original.param = &reqParam
		jons, err := gm.Json.Encode(reqParam)
		if err == nil {
			original.val.reqParam = &jons
		}
	}

	if len(reqQueries) > 0 {
		original.queries = &reqQueries
		jons, err := gm.Json.Encode(reqQueries)
		if err == nil {
			original.val.reqQuery = &jons
		}
	}

	part, err := fiberCtx.MultipartForm()
	if err == nil && part != nil {
		original.file = &part.File
		original.form = &part.Value
		form := part.Value

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
			jons, err := gm.Json.Encode(original.form)
			if err == nil {
				original.val.reqForm = &jons
			}
		}
	}

	if original.clog != nil {
		mol := &clog.ServicePieceV1{
			SvcParentName:    original.val.fromSvcName,
			SvcParentVersion: original.val.fromSvcVersion,
			Endpoint:         original.val.endpoint,
			Url:              original.val.url,
			ReqVersion:       original.val.reqVersion,
			ReqHeader:        original.val.reqHeader,
			ReqParam:         original.val.reqParam,
			ReqQuery:         original.val.reqQuery,
			ReqForm:          original.val.reqForm,
			ClientIp:         original.val.clientIP,
			StartedAt:        startedAt,
		}

		original.clog.ServicePieceV1(mol)
	}

	if regulator != nil {
		regulator(original.clog, original.regulator())
	} else {
		slf.defaultHandlerRegulator(original.regulator())
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
