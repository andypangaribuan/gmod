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
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (*stuServer) startWebsocket(fuseFiberApp *fiber.App, ws *func(router RouterS)) {
	if ws != nil {
		router := &stuFuseSRouter{
			app:    fuseFiberApp,
			locals: make(map[string]string, 0),
		}

		fuseFiberApp.Use("/ws", func(c *fiber.Ctx) error {
			if websocket.IsWebSocketUpgrade(c) {
				c.Locals("allowed", true)

				for k, v := range router.locals {
					c.Locals(k, string(c.Request().Header.Peek(v)))
				}

				return c.Next()
			}

			return fiber.ErrUpgradeRequired
		})

		(*ws)(router)
		if router.fnLocals != nil {
			sl := &stuFuseSLocal{router: router}
			(*router.fnLocals)(sl)
		}
	}
}
