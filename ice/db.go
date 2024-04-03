/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package ice

import "github.com/andypangaribuan/gmod/mdl"

type Db interface {
	Postgres(conf mdl.DbConnection) DbPostgresInstance
	PostgresRW(readConf mdl.DbConnection, writeConf mdl.DbConnection) DbPostgresInstance
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
	NewTransaction() (DbTx, error)

	Select(out interface{}, query string, args ...interface{}) (*mdl.DbExecReport, error)
	SelectR2(out interface{}, query string, args []interface{}, check *func() bool) (*mdl.DbExecReport, error)
	Execute(query string, args ...interface{}) (*mdl.DbExecReport, error)
	ExecuteRID(query string, args ...interface{}) (*int64, *mdl.DbExecReport, error)

	TxSelect(tx DbTx, out interface{}, query string, args ...interface{}) (*mdl.DbExecReport, error)
	TxExecute(tx DbTx, query string, args ...interface{}) (*mdl.DbExecReport, error)
	TxExecuteRID(tx DbTx, query string, args ...interface{}) (*int64, *mdl.DbExecReport, error)
}
