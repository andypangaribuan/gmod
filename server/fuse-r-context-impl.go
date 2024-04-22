/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "github.com/andypangaribuan/gmod/clog"

func (slf *stuFuseRContext) Clog() clog.Instance {
	return slf.mcx.clog
}

func (slf *stuFuseRContext) Auth(obj ...any) any {
	if len(obj) > 0 {
		slf.mcx.authObj = obj[0]
	}

	return slf.mcx.authObj
}

func (slf *stuFuseRContext) UserId(id ...any) any {
	if len(id) > 0 {
		slf.mcx.userId = id[0]
		slf.pushUserIdToClog()
	}

	return slf.mcx.userId
}

func (slf *stuFuseRContext) PartnerId(id ...any) any {
	if len(id) > 0 {
		slf.mcx.partnerId = id[0]
		slf.pushPartnerIdToClog()
	}

	return slf.mcx.partnerId
}

func (slf *stuFuseRContext) SetFiles(files map[string]string) {
	if len(files) > 0 {
		slf.mcx.files = &files
	}
}

func (slf *stuFuseRContext) Header() *map[string]string {
	return slf.header
}

func (slf *stuFuseRContext) LastResponse() (code int, val any) {
	return slf.mcx.responseCode, slf.mcx.responseVal
}

func (slf *stuFuseRContext) R200OK(val any) any {
	return slf.setResponse(200, val)
}

func (slf *stuFuseRContext) R201Created(val any) any {
	return slf.setResponse(201, val)
}

func (slf *stuFuseRContext) R202Accepted(val any) any {
	return slf.setResponse(202, val)
}

func (slf *stuFuseRContext) R204NoContent(val any) any {
	return slf.setResponse(204, val)
}

func (slf *stuFuseRContext) R301MovedPermanently(val any) any {
	return slf.setResponse(301, val)
}

func (slf *stuFuseRContext) R307TemporaryRedirect(val any) any {
	return slf.setResponse(307, val)
}

func (slf *stuFuseRContext) R308PermanentRedirect(val any) any {
	return slf.setResponse(308, val)
}

func (slf *stuFuseRContext) R400BadRequest(val any) any {
	return slf.setResponse(400, val)
}

func (slf *stuFuseRContext) R401Unauthorized(val any) any {
	return slf.setResponse(401, val)
}

func (slf *stuFuseRContext) R403Forbidden(val any) any {
	return slf.setResponse(403, val)
}

func (slf *stuFuseRContext) R404NotFound(val any) any {
	return slf.setResponse(404, val)
}

func (slf *stuFuseRContext) R406NotAcceptable(val any) any {
	return slf.setResponse(406, val)
}

func (slf *stuFuseRContext) R412PreconditionFailed(val any) any {
	return slf.setResponse(412, val)
}

func (slf *stuFuseRContext) R418Teapot(val any) any {
	return slf.setResponse(418, val)
}

func (slf *stuFuseRContext) R428PreconditionRequired(val any) any {
	return slf.setResponse(428, val)
}

func (slf *stuFuseRContext) R500InternalServerError(val any) any {
	return slf.setResponse(500, val)
}

func (slf *stuFuseRContext) R503ServiceUnavailable(val any) any {
	return slf.setResponse(503, val)
}
