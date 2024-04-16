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

func (slf *stuFuseRouterR) Group(endpoints map[string][]func(ctx FuseContextR)) {
	for path, handlers := range endpoints {
		slf.Single(path, handlers...)
	}
}

func (slf *stuFuseRouterR) GroupWithAuth(auth func(ctx FuseContextR), endpoints map[string][]func(ctx FuseContextR)) {
	for path, handlers := range endpoints {
		ls := []func(ctx FuseContextR){auth}
		slf.Single(path, append(ls, handlers...)...)
	}
}

func (slf *stuFuseRouterR) Single(path string, handlers ...func(ctx FuseContextR)) {
	ls := make([]func() (bool, func(FuseContextR)), len(handlers))

	for i, handler := range handlers {
		ls[i] = func() (bool, func(FuseContextR)) {
			return false, handler
		}
	}

	slf.rawSingle(path, ls...)
}

func (slf *stuFuseRouterR) rawSingle(endpoint string, controllers ...func() (isRegulator bool, controller func(ctx FuseContextR))) {
	index := strings.Index(endpoint, ":")
	if index == -1 {
		log.Fatalln("fuse server [restful]: endpoint format must be ▶︎ {Method}: {path}")
	}

	method := endpoint[0:index]
	path := strings.TrimSpace(endpoint[index+1:])

	switch method {
	case "GET":
		slf.fiberApp.Get(path, slf.restProcess(endpoint, controllers...))
	case "POS":
		slf.fiberApp.Post(path, slf.restProcess(endpoint, controllers...))
	case "DEL":
		slf.fiberApp.Delete(path, slf.restProcess(endpoint, controllers...))
	case "PUT":
		slf.fiberApp.Put(path, slf.restProcess(endpoint, controllers...))
	case "PAT":
		slf.fiberApp.Patch(path, slf.restProcess(endpoint, controllers...))
	default:
		log.Fatalln("fuse server [restful]: only support method GET, POS, DEL, PUT or PAT")
	}
}

func (slf *stuFuseRouterR) restProcess(endpoint string, controllers ...func() (isRegulator bool, controller func(ctx FuseContextR))) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		var (
			controllerRegulator *func(ctx FuseContextR)
			funcs               = make([]func(ctx FuseContextR), 0)
		)

		for _, fn := range controllers {
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

func (slf *stuFuseRouterR) regulator(fiberCtx *fiber.Ctx, endpoint string, controllerRegulator *func(ctx FuseContextR), controllers ...func(ctx FuseContextR)) {
	regulatorCtx := &stuFuseContextR{
		fiberCtx:    fiberCtx,
		endpoint:    endpoint,
		isRegulator: true,
		controllers: controllers,
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
