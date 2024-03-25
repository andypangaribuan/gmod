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
}
