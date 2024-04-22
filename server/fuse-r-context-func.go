/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "github.com/andypangaribuan/gmod/gm"

func (slf *stuFuseRContext) setExecPathFunc() {
	slf.mcx.execPath, slf.mcx.execFunc = gm.Util.GetExecPathFunc(3)
}

func (slf *stuFuseRContext) setResponse(code int, val any) any {
	slf.setExecPathFunc()
	slf.responseCode = code
	slf.responseVal = val
	return nil
}

func (slf *stuFuseRContext) pushUserIdToClog() {
	if slf.mcx.clog != nil {
		id := slf.mcx.getUserId()
		if id != nil {
			clogSetUserId(slf.mcx.clog, *id)
		}
	}
}

func (slf *stuFuseRContext) pushPartnerIdToClog() {
	if slf.mcx.clog != nil {
		id := slf.mcx.getPartnerId()
		if id != nil {
			clogSetPartnerId(slf.mcx.clog, *id)
		}
	}
}
