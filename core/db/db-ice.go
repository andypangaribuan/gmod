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
	Insert(args ...interface{}) error
	InsertRID(args ...interface{}) (*int64, error)
	Update(update *Update) error

	TxFetch(tx ice.DbTx, condition string, args ...interface{}) (*T, error)
	TxFetches(tx ice.DbTx, condition string, args ...interface{}) ([]*T, error)
	TxInsert(tx ice.DbTx, args ...interface{}) error
	TxInsertRID(tx ice.DbTx, args ...interface{}) (*int64, error)
	TxBulkInsert(tx ice.DbTx, entities []*T, args func(e *T) []interface{}, chunkSize ...int) error
	TxUpdate(tx ice.DbTx, update *Update) error
}
