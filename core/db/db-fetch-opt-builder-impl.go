/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
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

func (slf *stuFetchOptBuilder) FullQuery(query string) FetchOptBuilder {
	slf.fullQuery = &query
	return slf
}

func (slf *stuFetchOptBuilder) FullQueryFormatter(formatter func(query string) string) FetchOptBuilder {
	slf.fullQueryFormatter = &formatter
	return slf
}

func (slf *stuFetchOptBuilder) UsingRW(val ...bool) FetchOptBuilder {
	slf.usingRW = fm.GetFirst(val, true)
	return slf
}

func (slf *stuFetchOptBuilder) Out(ref any) FetchOptBuilder {
	slf.out = ref
	return slf
}
