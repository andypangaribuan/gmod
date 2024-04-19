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
	"github.com/andypangaribuan/gmod/clog"
	"github.com/gofiber/fiber/v2"
)

var (
	serverImpl   server
	fuseFiberApp *fiber.App
	clogNew      func() clog.Instance
	cip          *stuClientIP
)
