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
	"strings"

	"github.com/pkg/errors"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/gofiber/fiber/v2"
)

func (slf *stuFuseRRouter) AutoRecover(autoRecover bool) {
	slf.withAutoRecover = autoRecover
}

func (slf *stuFuseRRouter) PrintOnError(printOnError bool) {
	slf.printOnError = printOnError
}

func (slf *stuFuseRRouter) Unrouted(handler func(clog clog.Instance, ctx FuseRContext, method, path, url string) error) {
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
				ctx := &stuFuseRContext{
					fiberCtx:    c,
					endpoint:    endpoint,
					isRegulator: false,
				}

				err = handler(clogNew(), ctx, method, path, url)
			}
		}

		return err
	})
}

func (slf *stuFuseRRouter) ErrorHandler(catcher func(clog clog.Instance, ctx FuseRContext, err error) error) {
	slf.errorHandler = catcher
}

func (slf *stuFuseRRouter) Endpoints(regulator func(clog clog.Instance, regulator FuseRRegulator), auth func(clog.Instance, FuseRContext) error, pathHandlers map[string][]func(clog.Instance, FuseRContext) error) {
	for endpoint, handlers := range pathHandlers {
		var (
			ca = fm.Ternary(auth == nil, 0, 1)
			ls = make([]func(clog.Instance, FuseRContext) error, len(handlers)+ca)
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
