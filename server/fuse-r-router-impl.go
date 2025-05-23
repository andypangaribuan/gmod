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
	"strings"

	"github.com/pkg/errors"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/gofiber/fiber/v2"
)

func (slf *stuFuseRRouter) AutoRecover(autoRecover bool) {
	slf.withAutoRecover = autoRecover
}

func (slf *stuFuseRRouter) PrintOnError(printOnError bool) {
	slf.printOnError = printOnError
}

func (slf *stuFuseRRouter) Unrouted(handler func(ctx FuseRContext, method, path, url string) any) {
	slf.fiberApp.Use(func(fcx *fiber.Ctx) error {
		err := fcx.Next()

		var fe *fiber.Error
		if errors.As(err, &fe) && fe.Code == 404 {
			slf.execute(fcx, "unrouted", nil, handler)
			return nil
		}

		return err
	})
}

func (slf *stuFuseRRouter) ErrorHandler(catcher func(ctx FuseRContext, err error) any) {
	slf.errorHandler = catcher
}

func (slf *stuFuseRRouter) NoLog(paths []string) {
	for _, path := range paths {
		path = strings.ToLower(path)
		path = strings.ReplaceAll(path, " ", "")
		slf.noLogPaths[path] = path
	}
}

func (slf *stuFuseRRouter) Endpoints(regulator func(regulator FuseRRegulator), auth func(FuseRContext) any, pathHandlers map[string][]func(FuseRContext) any) {
	for endpoint, handlers := range pathHandlers {
		var (
			ca = fm.Ternary(auth == nil, 0, 1)
			ls = make([]func(FuseRContext) any, len(handlers)+ca)
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

func (slf *stuFuseRRouter) Static(endpointPaths map[string]string, config ...fiber.Static) {
	for endpoint, path := range endpointPaths {
		slf.static(endpoint, path, config...)
	}
}
