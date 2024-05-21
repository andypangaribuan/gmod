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
	_ "unsafe"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/gofiber/fiber/v2"
)

//go:linkname clogSetUserId github.com/andypangaribuan/gmod/clog.clogSetUserId
var clogSetUserId func(clog clog.Instance, id string)

//go:linkname clogSetPartnerId github.com/andypangaribuan/gmod/clog.clogSetPartnerId
var clogSetPartnerId func(clog clog.Instance, id string)

var (
	serverImpl   server
	fuseFiberApp *fiber.App
	// clogNew      func(uid string) clog.Instance // set by reflection from clog package
	cip *stuClientIP
)

const messageInternalServerError = "We apologize and are fixing the problem. Please try again at a later stage."
