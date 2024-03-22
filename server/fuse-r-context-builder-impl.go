/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

func (slf *srFuseContextBuilderR) Build() FuseContextR {
	ctx := &srFuseContextR{
		fiberCtx:     slf.original.fiberCtx,
		endpoint:     slf.original.endpoint,
		isRegulator:  false,
		regulatorCtx: slf.original.regulatorCtx,
	}

	return ctx
}
