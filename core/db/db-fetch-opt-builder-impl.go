/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuFetchOptBuilder) WithDeletedAtIsNull(val ...bool) FetchOptBuilder {
	slf.withDeletedAtIsNull = fm.GetFirst(val, true)
	return slf
}

func (slf *stuFetchOptBuilder) EndQuery(query string) FetchOptBuilder {
	slf.endQuery = &query
	return slf
}
