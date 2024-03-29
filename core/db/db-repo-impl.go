/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/ice"
)

func (slf *srRepo[T]) SetInsertColumn(columns string) {
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
}

func (slf *srRepo[T]) Fetch(condition string, args ...interface{}) (*T, error) {
	models, _, err := slf.fetches(true, nil, condition, args)
	return fm.PtrGetFirst(models), err
}

func (slf *srRepo[T]) Fetches(condition string, args ...interface{}) ([]*T, error) {
	models, _, err := slf.fetches(false, nil, condition, args)
	return models, err
}

func (slf *srRepo[T]) Insert(args ...interface{}) error {
	_, _, err := slf.insert(nil, false, args)
	return err
}

func (slf *srRepo[T]) InsertRID(args ...interface{}) (*int64, error) {
	id, _, err := slf.insert(nil, true, args)
	return id, err
}

func (slf *srRepo[T]) TxFetch(tx ice.DbTx, condition string, args ...interface{}) (*T, error) {
	models, _, err := slf.fetches(true, tx, condition, args)
	return fm.PtrGetFirst(models), err
}

func (slf *srRepo[T]) TxFetches(tx ice.DbTx, condition string, args ...interface{}) ([]*T, error) {
	models, _, err := slf.fetches(false, tx, condition, args)
	return models, err
}

func (slf *srRepo[T]) TxInsert(tx ice.DbTx, args ...interface{}) error {
	_, _, err := slf.insert(tx, false, args)
	return err
}

func (slf *srRepo[T]) TxInsertRID(tx ice.DbTx, args ...interface{}) (*int64, error) {
	id, _, err := slf.insert(tx, true, args)
	return id, err
}

func (slf *srRepo[T]) TxBulkInsert(tx ice.DbTx, entities []*T, args func(e *T) []interface{}, chunkSize ...int) error {
	_, err := slf.bulkInsert(tx, entities, args, chunkSize...)
	return err
}
