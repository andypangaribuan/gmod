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

func (slf *stuUpdateBuilder) AutoUpdatedAt(val ...bool) UpdateBuilder {
	slf.withAutoUpdatedAt = fm.GetFirst(val, true)
	return slf
}

func (slf *stuUpdateBuilder) Set(query string, args ...any) UpdateBuilder {
	slf.setQuery = &query
	slf.setArgs = &args
	return slf
}

func (slf *stuUpdateBuilder) SetIfNotNil(keyVal map[string]any) UpdateBuilder {
	slf.setInn = &keyVal
	return slf
}

func (slf *stuUpdateBuilder) Where(query string, args ...any) UpdateBuilder {
	slf.whereQuery = &query
	slf.whereArgs = &args
	return slf
}
