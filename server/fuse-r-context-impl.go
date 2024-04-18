/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func (slf *stuFuseContextR) LastResponse() (code int, val any) {
	return slf.lastResponseCode, slf.lastResponseVal
}

func (slf *stuFuseContextR) Auth(obj ...any) any {
	if len(obj) > 0 {
		slf.authObj = obj[0]
	}
	return slf.authObj
}

func (slf *stuFuseContextR) R200OK(obj any) error {
	slf.setResponse(200, obj)
	return nil
}
