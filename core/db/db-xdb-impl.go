/*
 * Copyright (c) 2025.
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

func (slf *stuXDB) Select(logc clog.Instance, query string, args ...any) ([]*map[string]any, error) {
	return slf.override(logc, slf.fetches(nil, query, args)).fetches()
}

func (slf *stuXDB) Execute(logc clog.Instance, query string, args ...any) error {
	return slf.override(logc, slf.execute(nil, false, query, args)).err
}

func (slf *stuXDB) ExecuteRID(logc clog.Instance, query string, args ...any) (*int64, error) {
	return slf.override(logc, slf.execute(nil, true, query, args)).execute()
}

func (slf *stuXDB) TxSelect(logc clog.Instance, tx ice.DbTx, query string, args ...any) ([]*map[string]any, error) {
	return slf.override(logc, slf.fetches(tx, query, args)).fetches()
}

func (slf *stuXDB) TxExecute(logc clog.Instance, tx ice.DbTx, query string, args ...any) error {
	return slf.override(logc, slf.execute(tx, false, query, args)).err
}

func (slf *stuXDB) TxExecuteRID(logc clog.Instance, tx ice.DbTx, query string, args ...any) (*int64, error) {
	return slf.override(logc, slf.execute(tx, true, query, args)).execute()
}
