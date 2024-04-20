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
	}

	return slf.mcx.userId
}

func (slf *stuFuseRContext) PartnerId(id ...any) any {
	if len(id) > 0 {
		slf.mcx.partnerId = id[0]
	}

	return slf.mcx.partnerId
}

func (slf *stuFuseRContext) Header() *map[string]string {
	return slf.header
}

func (slf *stuFuseRContext) Url() string {
	return slf.mcx.val.url
}

func (slf *stuFuseRContext) LastResponse() (code int, val any) {
	return slf.mcx.responseCode, slf.mcx.responseVal
}

func (slf *stuFuseRContext) R200OK(val any) any {
	return slf.setResponse(200, val)
}

func (slf *stuFuseRContext) R500InternalServerError(val any) any {
	return slf.setResponse(500, val)
}
