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
	"github.com/andypangaribuan/gmod/fm"
)

func (slf *stuFuseRegulatorR) currentHandler() func(clog.Instance, FuseContextR) error {
	return slf.fuseContext.handlers[slf.currentIndex]
}

func (slf *stuFuseRegulatorR) send() error {
	ctx := slf.fuseContext.fiberCtx.Status(slf.currentHandlerContext.responseCode)

	switch val := slf.currentHandlerContext.responseVal.(type) {
	case string:
		return ctx.SendString(val)
	case *string:
		return ctx.SendString(fm.GetDefault(val, ""))
	case []byte:
		return ctx.Send(val)
	case *[]byte:
		return ctx.Send(fm.GetDefault(val, []byte{}))
	case any:
		return ctx.JSON(val)
	default:
		return ctx.SendString("invalid response object")
	}
}
