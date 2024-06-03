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
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
	"github.com/jmoiron/sqlx"
)

type stuDb struct{}

type stuRepo[T any] struct {
	ins                 ice.DbInstance
	tableName           string
	insertColumn        string
	insertArgSign       string
	withDeletedAtIsNull bool
	rwFetchWhenNull     bool
	insertColumnFunc    func(e *T) []any
}

type stuVDB[T any] struct {
	ins     ice.DbInstance
	dvalSql string
}

type stuConnection struct {
	conf       *mol.DbConnection
	sx         *sqlx.DB
	driverName string
}

type pgInstance struct {
	rw *stuConnection
	ro *stuConnection
}

type pgInstanceTx struct {
	ins        *pgInstance
	tx         *sqlx.Tx
	isCommit   bool
	isRollback bool
	errCommit  error
}

type stuRepoResult[T any] struct {
	entities []*T
	id       *int64
	report   *stuReport
	err      error
}

type stuReport struct {
	tableName     string
	insertColumn  string
	insertArgSign string
	query         string
	args          []any
	execReport    *mol.DbExecReport
}

type stuUnsafe struct {
	query   string
	args    []any
	message string
	trace   string
}

type stuRepoOptBuilder struct {
	withDeletedAtIsNull *bool
	rwFetchWhenNull     *bool
}

type stuFetchOptBuilder struct {
	withDeletedAtIsNull *bool
	endQuery            *string
	fullQuery           *string
}

type stuUpdateBuilder struct {
	withAutoUpdatedAt *bool
	setQuery          *string
	setArgs           *[]any
	setInn            *map[string]any
	whereQuery        *string
	whereArgs         *[]any
}
