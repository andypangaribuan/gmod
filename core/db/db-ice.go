/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/ice"

type Repo[T any] interface {
	SetInsert(columns string, fn func(e *T) []any)

	Fetch(condition string, args ...any) (*T, error)
	Fetches(condition string, args ...any) ([]*T, error)
	VFetches(condition string, args ...any) ([]T, error)
	Insert(e *T) error
	InsertRID(e *T) (*int64, error)
	Update(builder UpdateBuilder) error
	Execute(condition string, args ...any) error

	TxFetch(tx ice.DbTx, condition string, args ...any) (*T, error)
	TxFetches(tx ice.DbTx, condition string, args ...any) ([]*T, error)
	TxVFetches(tx ice.DbTx, condition string, args ...any) ([]T, error)
	TxInsert(tx ice.DbTx, e *T) error
	TxInsertRID(tx ice.DbTx, e *T) (*int64, error)
	TxBulkInsert(tx ice.DbTx, entities []*T, chunkSize ...int) error
	TxUpdate(tx ice.DbTx, builder UpdateBuilder) error
	TxExecute(tx ice.DbTx, condition string, args ...any) error
}

type RepoOptBuilder interface {
	WithDeletedAtIsNull(val ...bool) RepoOptBuilder
	RWFetchWhenNull(val ...bool) RepoOptBuilder
}

type FetchOptBuilder interface {
	WithDeletedAtIsNull(val ...bool) FetchOptBuilder
	EndQuery(query string) FetchOptBuilder
	FullQuery(query string) FetchOptBuilder
}

type UpdateBuilder interface {
	AutoUpdatedAt(val ...bool) UpdateBuilder
	Set(query string, args ...any) UpdateBuilder
	SetIfNotNil(keyVal map[string]any) UpdateBuilder
	Where(query string, args ...any) UpdateBuilder
}
