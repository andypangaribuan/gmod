/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import (
	"fmt"
	"os"
)

func (slf *srFuseContextR) Regulator() FuseContextRegulatorR {
	if !slf.isRegulator {
		fmt.Printf("fuse server [restful]: you cannot call this function because you're not the regulator")
		os.Exit(100)
	}

	if slf.regulatorCtx == nil {
		slf.regulatorCtx = &srFuseContextRegulatorR{
			fuseContext:  slf,
			currentIndex: -1,
		}
	}

	return slf.regulatorCtx
}

func (slf *srFuseContextR) GetResponse() (code int, obj interface{}) {
	slf.regulatorCtx.currentControllerContext = slf
	return slf.responseCode, slf.responseObj
}
