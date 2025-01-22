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
	"strings"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/ice"
)

func (slf *stuRepo[T]) SetInsert(columns string, fn func(e *T) []any) {
	columns = strings.TrimSpace(columns)
	for {
		columns = strings.ReplaceAll(columns, " ", "")
		columns = strings.ReplaceAll(columns, "\n", "")
		if !strings.Contains(columns, " ") && !strings.Contains(columns, "\n") {
			break
		}
	}

	slf.insertColumn = slf.formatInsertColumnArgs(columns)
	slf.insertArgSign = slf.formatInsertColumnArgs(slf.generateArgSign(columns))
	slf.insertColumnFunc = fn
}

func (slf *stuRepo[T]) Fetch(clog clog.Instance, condition string, args ...any) (*T, error) {
	return slf.override(clog, slf.fetches(true, nil, condition, args)).fetch()
}

func (slf *stuRepo[T]) Fetches(clog clog.Instance, condition string, args ...any) ([]*T, error) {
	return slf.override(clog, slf.fetches(false, nil, condition, args)).fetches()
}

func (slf *stuRepo[T]) Insert(clog clog.Instance, e *T) error {
	return slf.override(clog, slf.insert(nil, false, slf.insertColumnFunc(e))).err
}

func (slf *stuRepo[T]) InsertRID(clog clog.Instance, e *T) (*int64, error) {
	return slf.override(clog, slf.insert(nil, true, slf.insertColumnFunc(e))).execute()
}

func (slf *stuRepo[T]) Update(clog clog.Instance, builder UpdateBuilder) error {
	return slf.override(clog, slf.update(nil, builder.(*stuUpdateBuilder))).err
}

func (slf *stuRepo[T]) Delete(clog clog.Instance, condition string, args ...any) error {
	return slf.override(clog, slf.delete(nil, condition, args)).err
}

func (slf *stuRepo[T]) TxFetch(clog clog.Instance, tx ice.DbTx, condition string, args ...any) (*T, error) {
	return slf.override(clog, slf.fetches(true, tx, condition, args)).fetch()
}

func (slf *stuRepo[T]) TxFetches(clog clog.Instance, tx ice.DbTx, condition string, args ...any) ([]*T, error) {
	return slf.override(clog, slf.fetches(false, tx, condition, args)).fetches()
}

func (slf *stuRepo[T]) TxInsert(clog clog.Instance, tx ice.DbTx, e *T) error {
	return slf.override(clog, slf.insert(tx, false, slf.insertColumnFunc(e))).err
}

func (slf *stuRepo[T]) TxInsertRID(clog clog.Instance, tx ice.DbTx, e *T) (*int64, error) {
	return slf.override(clog, slf.insert(tx, true, slf.insertColumnFunc(e))).execute()
}

func (slf *stuRepo[T]) TxBulkInsert(clog clog.Instance, tx ice.DbTx, entities []*T, chunkSize ...int) error {
	return slf.override(clog, slf.bulkInsert(tx, entities, chunkSize...)).err
}

func (slf *stuRepo[T]) TxUpdate(clog clog.Instance, tx ice.DbTx, builder UpdateBuilder) error {
	return slf.override(clog, slf.update(tx, builder.(*stuUpdateBuilder))).err
}

func (slf *stuRepo[T]) TxDelete(clog clog.Instance, tx ice.DbTx, condition string, args ...any) error {
	return slf.override(clog, slf.delete(tx, condition, args)).err
}
