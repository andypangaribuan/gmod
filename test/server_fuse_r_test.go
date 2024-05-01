/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fct"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/server"
	"github.com/andypangaribuan/gmod/test/db/entity"
	"github.com/andypangaribuan/gmod/test/db/repo"
)

// go test -v -run ^TestServerFuseR$
func TestServerFuseR(t *testing.T) {
	server.FuseR(env.AppRestPort, func(router server.RouterR) {
		router.AutoRecover(env.AppAutoRecover)
		router.PrintOnError(env.AppServerPrintOnError)
		router.ErrorHandler(sfrErrorHandler)
		router.Unrouted(sfrUnrouted)

		router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
			"POS: /hi/:firstName-:lastName/:age?": {sfrHi},
		})

		router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-1": {sfrPrivateStatus1},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-2": {sfrPrivateStatus1},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-3": {sfrPrivateStatus1},
			"GET: /private/status-4": {sfrPrivateStatus1, sfrPrivateStatus2},
			"GET: /private/status-5": {sfrPrivateStatus1, sfrPrivateStatus2},
			"GET: /private/status-6": {sfrPrivateStatus1, sfrPrivateStatus2},
		})

		// error or panic
		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-10": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-11": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-12": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-13": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-14": {sfrPrivateStatus1, sfrPrivateStatusPanic, sfrPrivateStatus2},
			"GET: /private/status-15": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusPanic},
		})

		router.Endpoints(sfrRegulator, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /private/status-16": {sfrPrivateStatus1, sfrPrivateStatusErr, sfrPrivateStatus2},
			"GET: /private/status-17": {sfrPrivateStatus1, sfrPrivateStatus2, sfrPrivateStatusErr},
		})

		router.Endpoints(nil, sfrAuth, map[string][]func(server.FuseRContext) any{
			"GET: /fetch-1":          {sfrFetch},
			"POS: /insert-1":         {sfrInsert},
			"POS: /delete-1":         {sfrDelete},
			"POS: /form-1":           {sfrForm},
			"GET: /call-http-1":      {sfrCallHttp1},
			"GET: /call-http-2":      {sfrCallHttp2},
			"GET: /call-http-3":      {sfrCallHttp3},
			"GET: /insert/clog-note": {sfrInsertNote},
		})
	})
}

func sfrUnrouted(ctx server.FuseRContext, method, path, url string) any {
	data := map[string]string{
		"status": "unrouted",
		"method": method,
		"path":   path,
		"url":    url,
	}

	return ctx.R404NotFound(data)
}

func sfrErrorHandler(ctx server.FuseRContext, err error) any {
	// message := fmt.Sprintf("something went wrong: %v\n%+v", err.Error(), err)
	return ctx.R500InternalServerError(err)
}

func sfrRegulator(regulator server.FuseRRegulator) {
	defer regulator.Recover()
	var (
		code    int
		canCall bool
	)

	for {
		canCall = true
		next, handler := regulator.Next()
		if !next {
			break
		}

		if regulator.Endpoint() == "GET: /private/status-4" && regulator.IsHandler(sfrPrivateStatus2) {
			continue
		}

		if regulator.Endpoint() == "GET: /private/status-5" && regulator.IsHandler(sfrPrivateStatus1) {
			canCall = false
			_, meta, _ := regulator.Call(handler, regulator.CallOpt().OverrideHeader(map[string]string{
				"xyz": "Override Header",
			}))
			code = meta.Code
		}

		if canCall {
			_, meta, _ := regulator.Call(handler)
			code = meta.Code
		}

		if code == -1 {
			return
		}

		if code < 200 || code > 299 {
			break
		}
	}
}

func sfrAuth(ctx server.FuseRContext) any {
	ctx.Auth("Halo")
	ctx.UserId("abc")
	ctx.PartnerId("xyz")
	return ctx.R200OK("Andy")
}

func sfrPrivateStatus1(ctx server.FuseRContext) any {
	ctx.SetFiles(map[string]string{
		"file1": "gcs/file1.pdf",
	})

	_, val := ctx.LastResponse()
	return ctx.R200OK(fmt.Sprintf("%v Pangaribuan", val))
}

func sfrPrivateStatus2(ctx server.FuseRContext) any {
	fmt.Printf("private-status-2: header: %v\n", ctx.ReqHeader())
	auth := ctx.Auth().(string)
	_, val := ctx.LastResponse()

	data := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("%v %v", auth, val),
	}

	return ctx.R200OK(data)
}

func sfrPrivateStatusPanic(ctx server.FuseRContext) any {
	auth := ctx.Auth().(int) // panic error
	return ctx.R200OK(auth)
}

func sfrPrivateStatusErr(ctx server.FuseRContext) any {
	return errors.New("test error")
}

func sfrHi(ctx server.FuseRContext) any {
	return ctx.R200OK("ok")
}

func sfrFetch(ctx server.FuseRContext) any {
	entities, err := repo.User.Fetches(ctx.Clog(), "name=?", "andy")
	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	type response struct {
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  time.Time  `json:"updated_at"`
		DeletedAt  *time.Time `json:"deleted_at"`
		Name       string     `json:"name"`
		Address    *string    `json:"address"`
		Height     *int       `json:"height"`
		GoldAmount *fct.FCT   `json:"gold_amount"`
	}

	users := make([]response, len(entities))
	for i, e := range entities {
		users[i] = response{
			CreatedAt:  e.CreatedAt,
			UpdatedAt:  e.UpdatedAt,
			Name:       e.Name,
			Address:    e.Address,
			Height:     e.Height,
			GoldAmount: e.GoldAmount,
		}

		if i > 0 {
			calVal, err := fct.Calc2(users[0].GoldAmount.Get(fct.Zero), "+", e.GoldAmount.Get(fct.Zero))
			if err != nil {
				return ctx.R500InternalServerError(err)
			}

			users[0].GoldAmount = &calVal
		}
	}

	return ctx.R200OK(users)
}

func sfrInsert(ctx server.FuseRContext) any {
	type stuHeader struct {
		Version string `json:"x-version"`
	}

	type stuRequest struct {
		Name       string   `json:"name"`
		Address    *string  `json:"address"`
		Height     *int     `json:"height"`
		GoldAmount *fct.FCT `json:"gold_amount"`
	}

	var (
		header *stuHeader
		req    *stuRequest
	)

	err := ctx.ReqParser(&header, &req)
	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	timenow := gm.Util.Timenow()
	user := &entity.User{
		CreatedAt:  timenow,
		UpdatedAt:  timenow,
		Uid:        gm.Util.UID(),
		Name:       req.Name,
		Address:    req.Address,
		Height:     req.Height,
		GoldAmount: req.GoldAmount,
	}

	err = repo.User.Insert(ctx.Clog(), user)
	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	return ctx.R200OK("success")
}

func sfrDelete(ctx server.FuseRContext) any {
	query := ctx.ReqQuery()
	if query == nil {
		return ctx.R406NotAcceptable("have no query")
	}

	var name *string
	for k, v := range *query {
		if k == "name" {
			name = &v
		}
	}

	if name == nil {
		return ctx.R400BadRequest("not found: name on query")
	}

	err := repo.User.Delete(ctx.Clog(), "name=?", name)
	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	return ctx.R200OK("ok")
}

func sfrForm(ctx server.FuseRContext) any {
	form := ctx.ReqForm()
	file := ctx.ReqFile()

	if form == nil {
		fmt.Printf("[form] no data\n")
	} else {
		for k, v := range *form {
			fmt.Printf("[form] key: %v, file: %v\n", k, v)
		}
	}

	if file == nil {
		fmt.Printf("[file] no data\n")
	} else {
		for k, v := range *file {
			fmt.Printf("[file] key: %v, total-file: %v\n", k, len(v))
		}
	}

	return ctx.R200OK("ok")
}

func sfrCallHttp1(ctx server.FuseRContext) any {
	url := "http://ipecho.net/plain"

	data, code, err := gm.Http.Get(ctx.Clog(), url).
		SetJsonHeader("1.0", map[string]string{
			"Authorization": "Bearer xyz",
		}).
		Call()

	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	if code != 200 {
		return ctx.R400BadRequest(fmt.Sprintf("code: %v", code))
	}

	return ctx.R200OK(string(data))
}

func sfrCallHttp2(ctx server.FuseRContext) any {
	data, code, err := gm.Http.Get(ctx.Clog(), "http://ipecho.net/plain").
		SetJsonHeader("1.0", map[string]string{
			"Authorization": "Bearer xyz",
		}).
		SetPathParam(map[string]string{
			"userId": "sample@sample.id",
		}).
		SetQueryParam(map[string]string{
			"par1": "ok",
		}).
		SetFormData(map[string]string{
			"data1": "data",
		}).
		SetBody(map[string]string{
			"body1": "body",
		}).
		Call()

	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	if code != 200 {
		return ctx.R400BadRequest(fmt.Sprintf("code: %v", code))
	}

	return ctx.R200OK(string(data))
}

func sfrCallHttp3(ctx server.FuseRContext) any {
	data, code, err := gm.Http.Get(ctx.Clog(), "http://localhost:3321/private/status-2").
		SetJsonHeader("1.0", map[string]string{
			"Authorization": "Bearer xyz",
		}).
		Call()

	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	if code != 200 {
		return ctx.R400BadRequest(fmt.Sprintf("code: %v", code))
	}

	return ctx.R200OK(string(data))
}

func sfrInsertNote(ctx server.FuseRContext) any {
	note := &clog.Note{Data: "halo"}
	err := ctx.Clog().Note(note)
	if err != nil {
		return ctx.R500InternalServerError(err)
	}

	return ctx.R200OK("done")
}
