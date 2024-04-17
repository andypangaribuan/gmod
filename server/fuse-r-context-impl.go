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

func (slf *stuFuseContextR) regulator() FuseContextRegulatorR {
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
