/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"github.com/andypangaribuan/gmod/model"
	"github.com/jmoiron/sqlx"
)

type srDb struct{}

type srConnection struct {
	conf       *model.DbConnection
	sx         *sqlx.DB
	driverName string
}

type pgInstance struct {
	rw *srConnection
	ro *srConnection
}

type pqInstanceTx struct {
	ins        *pgInstance
	tx         *sqlx.Tx
	isCommit   bool
	isRollback bool
	errCommit  error
}
