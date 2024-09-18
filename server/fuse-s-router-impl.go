/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "github.com/gofiber/contrib/websocket"

func (slf *stuFuseSRouter) Locals(fn func(sl FuseSLocal)) {
	slf.fnLocals = &fn
}

func (slf *stuFuseSRouter) Register(path string, handler func(ctx FuseSContext)) {
	slf.app.Get(path, websocket.New(func(c *websocket.Conn) {
		ctx := &stuFuseSContext{conn: c}
		handler(ctx)
	}))
}

func (slf *stuFuseSRouter) Run(path string) FuseSRun {
	return slf.run(path)
}
