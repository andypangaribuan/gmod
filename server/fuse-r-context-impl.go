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
	"os"
)

func (slf *stuFuseContextR) Regulator() FuseContextRegulatorR {
	if !slf.isRegulator {
		fmt.Printf("fuse server [restful]: you cannot call this function because you're not the regulator")
		os.Exit(100)
	}

	if slf.regulatorCtx == nil {
		slf.regulatorCtx = &stuFuseContextRegulatorR{
			fuseContext:  slf,
			currentIndex: -1,
		}
	}

	return slf.regulatorCtx
}

func (slf *stuFuseContextR) GetResponse() (code int, obj any) {
	slf.regulatorCtx.currentControllerContext = slf
	return slf.responseCode, slf.responseObj
}

func (slf *stuFuseContextR) SetAuth(obj any) {
	slf.authObj = obj
}
