/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mdl"
)

func (*srDb) Postgres(conf mdl.DbConnection) ice.DbPostgresInstance {
	setPgConfDVal(&conf)

	instance := &pgInstance{
		rw: &srConnection{
			conf:       &conf,
			driverName: "postgres",
		},
	}

	return instance
}

func (*srDb) PostgresRW(readConf mdl.DbConnection, writeConf mdl.DbConnection) ice.DbPostgresInstance {
	setPgConfDVal(&readConf)
	setPgConfDVal(&writeConf)

	instance := &pgInstance{
		ro: &srConnection{
			conf:       &readConf,
			driverName: "postgres",
		},
		rw: &srConnection{
			conf:       &writeConf,
			driverName: "postgres",
		},
	}

	return instance
}
