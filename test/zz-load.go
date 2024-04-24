/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/mol"
	"github.com/andypangaribuan/gmod/test/db/repo"
)

func loadEnv() {
	dirPath := getDirPath()
	gm.Util.LoadEnv(dirPath + "/.env")

	env = &stuEnv{
		AppName:               gm.Util.Env.GetString("APP_NAME"),
		AppVersion:            gm.Util.Env.GetString("APP_VERSION", "0.0.0"),
		AppEnv:                gm.Util.Env.GetAppEnv("APP_ENV"),
		AppTimezone:           gm.Util.Env.GetString("APP_TIMEZONE"),
		AppRestPort:           gm.Util.Env.GetInt("APP_REST_PORT"),
		AppAutoRecover:        gm.Util.Env.GetBool("APP_AUTO_RECOVER"),
		AppServerPrintOnError: gm.Util.Env.GetBool("APP_SERVER_PRINT_ON_ERROR"),

		ClogAddress:         gm.Util.Env.GetString("CLOG_ADDRESS"),
		SvcInternalBaseUrls: gm.Util.Env.GetStringSlice("SVC_INTERNAL_BASE_URLS", "|", []string{}),
		
		TxLockEngineAddress: gm.Util.Env.GetString("TX_LOCK_ENGINE_ADDRESS"),

		DbHost: gm.Util.Env.GetString("DB_HOST"),
		DbPort: gm.Util.Env.GetInt("DB_PORT"),
		DbName: gm.Util.Env.GetString("DB_NAME"),
		DbUser: gm.Util.Env.GetString("DB_USER"),
		DbPass: gm.Util.Env.GetString("DB_PASS"),
	}
}

func loadDb() {
	conn := mol.DbConnection{
		AppName:  env.AppName,
		Host:     env.DbHost,
		Port:     env.DbPort,
		Name:     env.DbName,
		Username: env.DbUser,
		Password: env.DbPass,
	}

	dbi = gm.Db.PostgresRW(conn, conn)
	repo.Init(dbi)
	fm.CallOrderedInit()
}
