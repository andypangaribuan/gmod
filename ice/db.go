/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package ice

import "github.com/andypangaribuan/gmod/model"

type Db interface {
	Postgres(conf model.DbConnection) DbPostgresInstance
	PostgresRW(readConf model.DbConnection, writeConf model.DbConnection) DbPostgresInstance
}

type DbTx interface {
	Commit() error
	Rollback() error
}

type DbPostgresInstance interface {
	DbInstance
}

type DbInstance interface {
	Ping() (string, error)
	PingRead() (string, error)

	Select(out interface{}, query string, args ...interface{}) (*model.DbExecReport, error)
	SelectR2(out interface{}, check func() bool, query string, args ...interface{}) (*model.DbExecReport, error)
	Execute(query string, args ...interface{}) (*model.DbExecReport, error)
	ExecuteRID(query string, args ...interface{}) (*int64, *model.DbExecReport, error)

	TxSelect(tx DbTx, out interface{}, query string, args ...interface{}) (*model.DbExecReport, error)
	TxExecute(tx DbTx, query string, args ...interface{}) (*model.DbExecReport, error)
	TxExecuteRID(tx DbTx, query string, args ...interface{}) (*int64, *model.DbExecReport, error)
}
