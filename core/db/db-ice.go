/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import (
	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/ice"
)

type Repo[T any] interface {
	SetInsert(columns string, fn func(e *T) []any)

	Fetch(clog clog.Instance, condition string, args ...any) (*T, error)
	Fetches(clog clog.Instance, condition string, args ...any) ([]*T, error)
	Insert(clog clog.Instance, e *T) error
	InsertRID(clog clog.Instance, e *T) (*int64, error)
	Update(clog clog.Instance, builder UpdateBuilder) error
	Delete(clog clog.Instance, condition string, args ...any) error

	TxFetch(clog clog.Instance, tx ice.DbTx, condition string, args ...any) (*T, error)
	TxFetches(clog clog.Instance, tx ice.DbTx, condition string, args ...any) ([]*T, error)
	TxInsert(clog clog.Instance, tx ice.DbTx, e *T) error
	TxInsertRID(clog clog.Instance, tx ice.DbTx, e *T) (*int64, error)
	TxBulkInsert(clog clog.Instance, tx ice.DbTx, entities []*T, chunkSize ...int) error
	TxUpdate(clog clog.Instance, tx ice.DbTx, builder UpdateBuilder) error
	TxDelete(clog clog.Instance, tx ice.DbTx, condition string, args ...any) error
}

type VDB[T any] interface {
	Sql(sqlName string) string

	Fetch(clog clog.Instance, sqlName string, args ...any) (*T, error)
	Fetches(clog clog.Instance, sqlName string, args ...any) ([]*T, error)

	TxFetch(clog clog.Instance, tx ice.DbTx, sqlName string, args ...any) (*T, error)
	TxFetches(clog clog.Instance, tx ice.DbTx, sqlName string, args ...any) ([]*T, error)
}

type RepoOptBuilder interface {
	WithDeletedAtIsNull(val ...bool) RepoOptBuilder
	RWFetchWhenNull(val ...bool) RepoOptBuilder
}

type FetchOptBuilder interface {
	WithDeletedAtIsNull(val ...bool) FetchOptBuilder
	EndQuery(query string) FetchOptBuilder
	FullQuery(query string) FetchOptBuilder
	FullQueryFormatter(formatter func(query string) string) FetchOptBuilder
}

type UpdateBuilder interface {
	AutoUpdatedAt(val ...bool) UpdateBuilder
	Set(query string, args ...any) UpdateBuilder
	SetIfNotNil(keyVal map[string]any) UpdateBuilder
	Where(query string, args ...any) UpdateBuilder
}
