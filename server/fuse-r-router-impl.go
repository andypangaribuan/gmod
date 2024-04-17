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
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/gofiber/fiber/v2"
)

func (slf *stuFuseRouterR) AutoRecover(autoRecover bool) {
	slf.withAutoRecover = autoRecover
}

func (slf *stuFuseRouterR) PrintOnError(printOnError bool) {
	slf.printOnError = printOnError
}

func (slf *stuFuseRouterR) Unrouted(controller func(ctx FuseContextR, method, path, url string)) {
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

				controller(ctx, method, path, url)
			}
		}

		return err
	})
}

func (slf *stuFuseRouterR) Endpoints(regulator func(ctx FuseContextR), auth func(ctx FuseContextR), pathHandlers map[string][]func(ctx FuseContextR)) {
	for endpoint, handlers := range pathHandlers {
		var (
			cr = fm.Ternary(regulator == nil, 0, 1)
			ca = fm.Ternary(auth == nil, 0, 1)
			ls = make([]func() (bool, func(FuseContextR)), len(handlers)+cr+ca)
		)

		if regulator != nil {
			ls[0] = func() (bool, func(FuseContextR)) {
				return true, regulator
			}
		}

		if auth != nil {
			ls[cr] = func() (bool, func(FuseContextR)) {
				return false, auth
			}
		}

		for i, handler := range handlers {
			ls[i+cr+ca] = func() (bool, func(FuseContextR)) {
				return false, handler
			}
		}

		slf.register(endpoint, ls...)
	}
}

func (slf *stuFuseRouterR) register(endpoint string, handlers ...func() (isRegulator bool, controller func(ctx FuseContextR))) {
	index := strings.Index(endpoint, ":")
	if index == -1 {
		log.Fatalln("fuse server [restful]: endpoint format must be ▶︎ {Method}: {path}")
	}

	method := endpoint[0:index]
	path := strings.TrimSpace(endpoint[index+1:])

	switch method {
	case "GET":
		slf.fiberApp.Get(path, slf.restProcess(endpoint, handlers...))
	case "POS":
		slf.fiberApp.Post(path, slf.restProcess(endpoint, handlers...))
	case "DEL":
		slf.fiberApp.Delete(path, slf.restProcess(endpoint, handlers...))
	case "PUT":
		slf.fiberApp.Put(path, slf.restProcess(endpoint, handlers...))
	case "PAT":
		slf.fiberApp.Patch(path, slf.restProcess(endpoint, handlers...))
	default:
		log.Fatalln("fuse server [restful]: only support method GET, POS, DEL, PUT or PAT")
	}
}

func (slf *stuFuseRouterR) restProcess(endpoint string, handlers ...func() (isRegulator bool, controller func(ctx FuseContextR))) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		var (
			controllerRegulator *func(ctx FuseContextR)
			funcs               = make([]func(ctx FuseContextR), 0)
		)

		for _, fn := range handlers {
			isRegulator, controller := fn()
			if isRegulator && controllerRegulator == nil {
				controllerRegulator = &controller
			} else {
				funcs = append(funcs, controller)
			}
		}

		slf.regulator(fiberCtx, endpoint, controllerRegulator, funcs...)
		return nil
	}
}

func (slf *stuFuseRouterR) regulator(fiberCtx *fiber.Ctx, endpoint string, controllerRegulator *func(ctx FuseContextR), handlers ...func(ctx FuseContextR)) {
	regulatorCtx := &stuFuseContextR{
		fiberCtx:    fiberCtx,
		endpoint:    endpoint,
		isRegulator: true,
		handlers:    handlers,
	}

	if controllerRegulator != nil {
		controller := *controllerRegulator
		controller(regulatorCtx)
	} else {
		slf.defaultRegulatorController(regulatorCtx)
	}
}

func (slf *stuFuseRouterR) defaultRegulatorController(ctx FuseContextR) {
	regulator := ctx.Regulator()

	for {
		canNext, handler := regulator.Next()
		if !canNext {
			break
		}

		builder := regulator.ContextBuilder()
		ctx := builder.Build()
		handler()(ctx)

		code, _ := ctx.GetResponse()
		if code < 200 || code >= 300 {
			break
		}
	}

	regulator.Send()
}
