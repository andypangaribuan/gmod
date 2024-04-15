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

func (slf *stuRepo[T]) Fetch(condition string, args ...any) (*T, error) {
	models, _, err := slf.fetches(true, nil, condition, args)
	return fm.PtrGetFirst(models), err
}

func (slf *stuRepo[T]) Fetches(condition string, args ...any) ([]*T, error) {
	models, _, err := slf.fetches(false, nil, condition, args)
	return models, err
}

func (slf *stuRepo[T]) VFetches(condition string, args ...any) ([]T, error) {
	models, _, err := slf.vfetches(false, nil, condition, args)
	return models, err
}

func (slf *stuRepo[T]) Insert(e *T) error {
	_, _, err := slf.insert(nil, false, slf.insertColumnFunc(e))
	return err
}

func (slf *stuRepo[T]) InsertRID(args ...any) (*int64, error) {
	id, _, err := slf.insert(nil, true, args)
	return id, err
}

func (slf *stuRepo[T]) Update(builder UpdateBuilder) error {
	_, err := slf.update(nil, builder.(*stuUpdateBuilder))
	return err
}

func (slf *stuRepo[T]) TxFetch(tx ice.DbTx, condition string, args ...any) (*T, error) {
	models, _, err := slf.fetches(true, tx, condition, args)
	return fm.PtrGetFirst(models), err
}

func (slf *stuRepo[T]) TxFetches(tx ice.DbTx, condition string, args ...any) ([]*T, error) {
	models, _, err := slf.fetches(false, tx, condition, args)
	return models, err
}

func (slf *stuRepo[T]) TxVFetches(tx ice.DbTx, condition string, args ...any) ([]T, error) {
	models, _, err := slf.vfetches(false, tx, condition, args)
	return models, err
}

func (slf *stuRepo[T]) TxInsert(tx ice.DbTx, args ...any) error {
	_, _, err := slf.insert(tx, false, args)
	return err
}

func (slf *stuRepo[T]) TxInsertRID(tx ice.DbTx, args ...any) (*int64, error) {
	id, _, err := slf.insert(tx, true, args)
	return id, err
}

func (slf *stuRepo[T]) TxBulkInsert(tx ice.DbTx, entities []*T, args func(e *T) []any, chunkSize ...int) error {
	_, err := slf.bulkInsert(tx, entities, args, chunkSize...)
	return err
}

func (slf *stuRepo[T]) TxUpdate(tx ice.DbTx, builder UpdateBuilder) error {
	_, err := slf.update(tx, builder.(*stuUpdateBuilder))
	return err
}
