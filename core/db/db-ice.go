/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/ice"

type Repo[T any] interface {
	SetInsertColumn(columns string)

	Fetch(condition string, args ...interface{}) (*T, error)
	Fetches(condition string, args ...interface{}) ([]*T, error)
	VFetches(condition string, args ...interface{}) ([]T, error)
	Insert(args ...interface{}) error
	InsertRID(args ...interface{}) (*int64, error)
	Update(builder UpdateBuilder) error

	TxFetch(tx ice.DbTx, condition string, args ...interface{}) (*T, error)
	TxFetches(tx ice.DbTx, condition string, args ...interface{}) ([]*T, error)
	TxVFetches(tx ice.DbTx, condition string, args ...interface{}) ([]T, error)
	TxInsert(tx ice.DbTx, args ...interface{}) error
	TxInsertRID(tx ice.DbTx, args ...interface{}) (*int64, error)
	TxBulkInsert(tx ice.DbTx, entities []*T, args func(e *T) []interface{}, chunkSize ...int) error
	TxUpdate(tx ice.DbTx, builder UpdateBuilder) error
}

type RepoOptBuilder interface {
	WithDeletedAtIsNull(val ...bool) RepoOptBuilder
	RWFetchWhenNull(val ...bool) RepoOptBuilder
}

type FetchOptBuilder interface {
	WithDeletedAtIsNull(val ...bool) FetchOptBuilder
	EndQuery(query string) FetchOptBuilder
}

type UpdateBuilder interface {
	AutoUpdatedAt(val ...bool) UpdateBuilder
	Set(query string, args ...interface{}) UpdateBuilder
	SetIfNotNil(keyVal map[string]interface{}) UpdateBuilder
	Where(query string, args ...interface{}) UpdateBuilder
}
