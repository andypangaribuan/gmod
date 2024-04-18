/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "fmt"

func (slf *stuFuseContextR) regulator() FuseRegulatorR {
	if !slf.isRegulator {
		fmt.Printf("fuse server [restful]: forbidden, you're not the regulator")
		return nil
	}

	if slf.regulatorCtx == nil {
		slf.regulatorCtx = &stuFuseRegulatorR{
			clog:         slf.clog,
			fuseContext:  slf,
			currentIndex: -1,
		}
	}

	return slf.regulatorCtx
}

func (slf *stuFuseContextR) setResponse(code int, val any) {
	slf.responseCode = code
	slf.responseVal = val
}
