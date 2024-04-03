/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mdl"
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
}

type stuConnection struct {
	conf       *mdl.DbConnection
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

type stuReport struct {
	tableName     string
	insertColumn  string
	insertArgSign string
	query         string
	args          []interface{}
	execReport    *mdl.DbExecReport
}

type stuUnsafe struct {
	query   string
	args    []interface{}
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
}

type stuUpdateBuilder struct {
	withAutoUpdatedAt *bool
	setQuery          *string
	setArgs           *[]interface{}
	setInn            *map[string]interface{}
	whereQuery        *string
	whereArgs         *[]interface{}
}
