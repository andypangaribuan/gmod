/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuXdbOptBuilder) RWFetchWhenNull(val ...bool) XdbOptBuilder {
	slf.rwFetchWhenNull = fm.GetFirst(val, true)
	return slf
}
