/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuUpdateBuilder) AutoUpdatedAt(val ...bool) UpdateBuilder {
	slf.withAutoUpdatedAt = fm.GetFirst(val, true)
	return slf
}

func (slf *stuUpdateBuilder) Set(query string, args ...interface{}) UpdateBuilder {
	slf.setQuery = &query
	slf.setArgs = &args
	return slf
}

func (slf *stuUpdateBuilder) SetIfNotNil(keyVal map[string]interface{}) UpdateBuilder {
	slf.setInn = &keyVal
	return slf
}

func (slf *stuUpdateBuilder) Where(query string, args ...interface{}) UpdateBuilder {
	slf.whereQuery = &query
	slf.whereArgs = &args
	return slf
}
