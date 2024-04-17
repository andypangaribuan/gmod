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

	"github.com/pkg/errors"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/gofiber/fiber/v2"
)

func (slf *stuFuseRouterR) AutoRecover(autoRecover bool) {
	slf.withAutoRecover = autoRecover
}

func (slf *stuFuseRouterR) PrintOnError(printOnError bool) {
	slf.printOnError = printOnError
}

func (slf *stuFuseRouterR) Unrouted(handler func(ctx FuseContextR, method, path, url string) error) {
	slf.fiberApp.Use(func(c *fiber.Ctx) error {
		err := c.Next()

		var fe *fiber.Error
		if errors.As(err, &fe) && fe.Code == 404 {
			var (
				url      = c.OriginalURL()
				path     = c.Path()
				method   = ""
				endpoint = ""
			)

			if c.Route() != nil {
				method = strings.ToUpper(c.Route().Method)

				m := method
				if len(m) > 3 {
					m = m[:3]
				}
				endpoint = fmt.Sprintf("%v: %v", m, path)
			}

			if fe.Message == fmt.Sprintf("Cannot %v %v", method, path) {
				ctx := &stuFuseContextR{
					fiberCtx:    c,
					endpoint:    endpoint,
					isRegulator: false,
				}

				err = handler(ctx, method, path, url)
			}
		}

		return err
	})
}

func (slf *stuFuseRouterR) ErrorHandler(catcher func(ctx FuseContextR, err error) error) {
	slf.errorHandler = catcher
}

func (slf *stuFuseRouterR) Endpoints(regulator func(regulator FuseContextRegulatorR), auth func(ctx FuseContextR) error, pathHandlers map[string][]func(ctx FuseContextR) error) {
	for endpoint, handlers := range pathHandlers {
		var (
			ca = fm.Ternary(auth == nil, 0, 1)
			ls = make([]func(FuseContextR) error, len(handlers)+ca)
		)

		if auth != nil {
			ls[0] = auth
		}

		for i, handler := range handlers {
			ls[i+ca] = handler
		}

		slf.register(endpoint, regulator, ls...)
	}
}

func (slf *stuFuseRouterR) register(endpoint string, regulator func(FuseContextRegulatorR), handlers ...func(ctx FuseContextR) error) {
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

func (slf *stuFuseRouterR) restProcess(endpoint string, regulator func(FuseContextRegulatorR), handlers ...func(ctx FuseContextR) error) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		slf.execRegulator(fiberCtx, endpoint, regulator, handlers...)
		return nil
	}
}

func (slf *stuFuseRouterR) execRegulator(fiberCtx *fiber.Ctx, endpoint string, regulator func(FuseContextRegulatorR), handlers ...func(ctx FuseContextR) error) {
	regulatorCtx := &stuFuseContextR{
		fiberCtx:    fiberCtx,
		endpoint:    endpoint,
		isRegulator: true,
		handlers:    handlers,

		errorHandler: slf.errorHandler,
	}

	if regulator != nil {
		regulator(regulatorCtx.regulator())
	} else {
		slf.defaultHandlerRegulator(regulatorCtx.regulator())
	}
}

func (slf *stuFuseRouterR) defaultHandlerRegulator(regulator FuseContextRegulatorR) {
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
