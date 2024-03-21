/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import "github.com/gofiber/fiber/v2"

type srServer struct{}

type srFuseRouterR struct {
	fiberApp        *fiber.App
	withAutoRecover bool
	printOnError    bool
}
