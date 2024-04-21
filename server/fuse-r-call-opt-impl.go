/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func (slf *stuFuseRCallOpt) OverrideHeader(val map[string]string) FuseRCallOpt {
	slf.header = &val
	return slf
}

func (slf *stuFuseRCallOpt) OverrideParam(val map[string]string) FuseRCallOpt {
	slf.param = &val
	return slf
}

func (slf *stuFuseRCallOpt) OverrideQuery(val map[string]string) FuseRCallOpt {
	slf.query = &val
	return slf
}

func (slf *stuFuseRCallOpt) OverrideForm(val map[string][]string) FuseRCallOpt {
	slf.form = &val
	return slf
}
