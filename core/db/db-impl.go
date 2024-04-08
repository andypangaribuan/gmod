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
)

func (*stuDb) Postgres(conf mol.DbConnection) ice.DbPostgresInstance {
	setPgConfDVal(&conf)

	instance := &pgInstance{
		rw: &stuConnection{
			conf:       &conf,
			driverName: "postgres",
		},
	}

	return instance
}

func (*stuDb) PostgresRW(readConf mol.DbConnection, writeConf mol.DbConnection) ice.DbPostgresInstance {
	setPgConfDVal(&readConf)
	setPgConfDVal(&writeConf)

	instance := &pgInstance{
		ro: &stuConnection{
			conf:       &readConf,
			driverName: "postgres",
		},
		rw: &stuConnection{
			conf:       &writeConf,
			driverName: "postgres",
		},
	}

	return instance
}
