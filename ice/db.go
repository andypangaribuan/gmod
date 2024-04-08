/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
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

	Select(out any, query string, args ...any) (*mdl.DbExecReport, error)
	SelectR2(out any, query string, args []any, check *func() bool) (*mdl.DbExecReport, error)
	Execute(query string, args ...any) (*mdl.DbExecReport, error)
	ExecuteRID(query string, args ...any) (*int64, *mdl.DbExecReport, error)

	TxSelect(tx DbTx, out any, query string, args ...any) (*mdl.DbExecReport, error)
	TxExecute(tx DbTx, query string, args ...any) (*mdl.DbExecReport, error)
	TxExecuteRID(tx DbTx, query string, args ...any) (*int64, *mdl.DbExecReport, error)
}
