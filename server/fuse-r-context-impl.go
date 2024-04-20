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
	"github.com/andypangaribuan/gmod/gm"
)

func (slf *stuFuseRContext) LastResponse() (code int, val any) {
	return slf.lastResponseCode, slf.lastResponseVal
}

func (slf *stuFuseRContext) Auth(obj ...any) any {
	if len(obj) > 0 {
		slf.authObj = obj[0]
		slf.isSetAuthObj = true
	}

	return slf.authObj
}

func (slf *stuFuseRContext) UserId(id ...any) any {
	if len(id) > 0 {
		slf.userId = id[0]
		slf.isSetUserId = true
	}

	return slf.userId
}

func (slf *stuFuseRContext) PartnerId(id ...any) any {
	if len(id) > 0 {
		slf.partnerId = id[0]
		slf.isSetPartnerId = true
	}

	return slf.partnerId
}

func (slf *stuFuseRContext) Header() *map[string]string {
	return slf.header
}

func (slf *stuFuseRContext) Url() string {
	return slf.val.url
}

func (slf *stuFuseRContext) setExecPathFunc() {
	slf.execPath, slf.execFunc = gm.Util.GetExecPathFunc(2)
}

func (slf *stuFuseRContext) R200OK(obj any) error {
	slf.setResponse(200, obj)
	slf.setExecPathFunc()
	return nil
}
