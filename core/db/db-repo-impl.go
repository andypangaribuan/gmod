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
	"github.com/andypangaribuan/gmod/fm"
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
	models, _, err := slf.fetches(true, nil, condition, args)
	return fm.PtrGetFirst(models), err
}

func (slf *stuRepo[T]) Fetches(clog clog.Instance, condition string, args ...any) ([]*T, error) {
	models, _, err := slf.fetches(false, nil, condition, args)
	return models, err
}

func (slf *stuRepo[T]) VFetches(clog clog.Instance, condition string, args ...any) ([]T, error) {
	models, _, err := slf.vfetches(false, nil, condition, args)
	return models, err
}

func (slf *stuRepo[T]) Insert(clog clog.Instance, e *T) error {
	_, _, err := slf.insert(nil, false, slf.insertColumnFunc(e))
	return err
}

func (slf *stuRepo[T]) InsertRID(clog clog.Instance, e *T) (*int64, error) {
	id, _, err := slf.insert(nil, true, slf.insertColumnFunc(e))
	return id, err
}

func (slf *stuRepo[T]) Update(clog clog.Instance, builder UpdateBuilder) error {
	_, err := slf.update(nil, builder.(*stuUpdateBuilder))
	return err
}

func (slf *stuRepo[T]) Execute(clog clog.Instance, condition string, args ...any) error {
	_, err := slf.execute(nil, condition, args)
	return err
}

func (slf *stuRepo[T]) TxFetch(clog clog.Instance, tx ice.DbTx, condition string, args ...any) (*T, error) {
	models, _, err := slf.fetches(true, tx, condition, args)
	return fm.PtrGetFirst(models), err
}

func (slf *stuRepo[T]) TxFetches(clog clog.Instance, tx ice.DbTx, condition string, args ...any) ([]*T, error) {
	models, _, err := slf.fetches(false, tx, condition, args)
	return models, err
}

func (slf *stuRepo[T]) TxVFetches(clog clog.Instance, tx ice.DbTx, condition string, args ...any) ([]T, error) {
	models, _, err := slf.vfetches(false, tx, condition, args)
	return models, err
}

func (slf *stuRepo[T]) TxInsert(clog clog.Instance, tx ice.DbTx, e *T) error {
	_, _, err := slf.insert(tx, false, slf.insertColumnFunc(e))
	return err
}

func (slf *stuRepo[T]) TxInsertRID(clog clog.Instance, tx ice.DbTx, e *T) (*int64, error) {
	id, _, err := slf.insert(tx, true, slf.insertColumnFunc(e))
	return id, err
}

func (slf *stuRepo[T]) TxBulkInsert(clog clog.Instance, tx ice.DbTx, entities []*T, chunkSize ...int) error {
	_, err := slf.bulkInsert(tx, entities, chunkSize...)
	return err
}

func (slf *stuRepo[T]) TxUpdate(clog clog.Instance, tx ice.DbTx, builder UpdateBuilder) error {
	_, err := slf.update(tx, builder.(*stuUpdateBuilder))
	return err
}

func (slf *stuRepo[T]) TxExecute(clog clog.Instance, tx ice.DbTx, condition string, args ...any) error {
	_, err := slf.execute(tx, condition, args)
	return err
}
