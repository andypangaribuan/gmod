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
	"fmt"
)

func (slf *stuFuseContextR) Regulator() FuseContextRegulatorR {
	if !slf.isRegulator {
		fmt.Printf("fuse server [restful]: forbidden, you're not the regulator")
		return nil
	}

	if slf.regulatorCtx == nil {
		slf.regulatorCtx = &stuFuseContextRegulatorR{
			fuseContext:  slf,
			currentIndex: -1,
		}
	}

	return slf.regulatorCtx
}

func (slf *stuFuseContextR) GetLastResponse() (code int, val any) {
	return slf.lastResponseCode, slf.lastResponseVal
}

func (slf *stuFuseContextR) GetResponse() (code int, val any) {
	slf.regulatorCtx.currentHandlerContext = slf
	return slf.responseCode, slf.responseVal
}

func (slf *stuFuseContextR) SetAuth(obj any) {
	slf.authObj = obj
}

func (slf *stuFuseContextR) Auth() any {
	return slf.authObj
}

func (slf *stuFuseContextR) R200OK(obj any) any {
	slf.setResponse(200, obj)
	return nil
}
