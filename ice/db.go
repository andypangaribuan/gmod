/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import "github.com/andypangaribuan/gmod/mol"

type Db interface {
	Postgres(conf mol.DbConnection) DbPostgresInstance
	PostgresRW(readConf mol.DbConnection, writeConf mol.DbConnection) DbPostgresInstance
	QueryBuilder() DbQueryBuilder
}

type DbTx interface {
	Commit() error
	Rollback()
}

type DbPostgresInstance interface {
	DbInstance
}

type DbInstance interface {
	Ping() (string, error)
	PingRead() (string, error)
	NewTransaction() (DbTx, error)

	Select(out any, usingRW bool, query string, args ...any) (*mol.DbExecReport, error)
	SelectR2(out any, query string, args []any, check *func() bool) (*mol.DbExecReport, error)
	Execute(query string, args ...any) (*mol.DbExecReport, error)
	ExecuteRID(query string, args ...any) (*int64, *mol.DbExecReport, error)

	TxSelect(tx DbTx, out any, query string, args ...any) (*mol.DbExecReport, error)
	TxExecute(tx DbTx, query string, args ...any) (*mol.DbExecReport, error)
	TxExecuteRID(tx DbTx, query string, args ...any) (*int64, *mol.DbExecReport, error)
}

type DbQueryBuilder interface {
	And(condition string, args ...any) DbQueryBuilder
	AndNotNill(condition string, arg any, validator ...func() (ok bool, args []any)) DbQueryBuilder
	AndNotNill2(condition string, arg1 any, arg2 any, validator ...func() (ok bool, args []any)) DbQueryBuilder

	AndNotNil(condition string, args ...any) DbQueryBuilder

	Build() (where string, args []any)
}
