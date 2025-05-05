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

func (slf *stuVDB[T]) Sql(sqlName string) string {
	return slf.dvalSql[sqlName]
}

func (slf *stuVDB[T]) Fetch(clog clog.Instance, sqlName string, args ...any) (*T, error) {
	return slf.override(clog, slf.fetches(true, nil, sqlName, args, false)).fetch()
}

func (slf *stuVDB[T]) Fetches(clog clog.Instance, sqlName string, args ...any) ([]*T, error) {
	return slf.override(clog, slf.fetches(false, nil, sqlName, args, false)).fetches()
}

func (slf *stuVDB[T]) Select(clog clog.Instance, sqlName string, args ...any) ([]map[string]any, error) {
	return slf.override(clog, slf.fetches(false, nil, sqlName, args, true)).selectX()
}

func (slf *stuVDB[T]) TxFetch(clog clog.Instance, tx ice.DbTx, sqlName string, args ...any) (*T, error) {
	return slf.override(clog, slf.fetches(true, tx, sqlName, args, false)).fetch()
}

func (slf *stuVDB[T]) TxFetches(clog clog.Instance, tx ice.DbTx, sqlName string, args ...any) ([]*T, error) {
	return slf.override(clog, slf.fetches(false, tx, sqlName, args, false)).fetches()
}
