/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuFuseSLocal) Set(key string, header ...string) {
	slf.router.locals[key] = fm.TernaryR(len(header) == 0, key, func() string { return header[0] })
}
