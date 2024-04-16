/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func (slf *stuFuseContextR) setResponse(code int, val any) {
	slf.responseCode = code
	slf.responseVal = val
}
