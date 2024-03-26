/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"time"

	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/model"
	"github.com/jmoiron/sqlx"
)

type srDb struct{}

type srRepo[T any] struct {
	ins                 ice.DbInstance
	tableName           string
	insertColumn        string
	insertArgSign       string
	withDeletedAtIsNull bool
	rwFetchWhenNull     bool
}

type srConnection struct {
	conf       *model.DbConnection
	sx         *sqlx.DB
	driverName string
}

type pgInstance struct {
	rw *srConnection
	ro *srConnection
}

type pgInstanceTx struct {
	ins        *pgInstance
	tx         *sqlx.Tx
	isCommit   bool
	isRollback bool
	errCommit  error
}

type srReport struct {
	tableName     string
	insertColumn  string
	insertArgSign string
	query         string
	args          []interface{}
	execReport    *model.DbExecReport
}

type srReportHost struct {
	host       string
	startedAt  time.Time
	finishedAt time.Time
	durationMs int64
}

type srUnsafe struct {
	query   string
	args    []interface{}
	message string
	trace   string
}
