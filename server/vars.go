/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import "github.com/gofiber/fiber/v2"

var (
	serverImpl          server
	fuseFiberApp        *fiber.App
	isFuseRPrintOnError bool
)
