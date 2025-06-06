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
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func setPgConfDVal(conf *mol.DbConnection) {
	if conf.AppName == "" {
		log.Fatalf("db connection must have application name")
	}

	var (
		dvalHost           = gm.Util.Env.GetString("DB_HOST", "127.0.0.1")
		dvalPort           = gm.Util.Env.GetInt("DB_PORT", 5432)
		dvalScheme         = gm.Util.Env.GetString("DB_SCHEME", "public")
		dvalTimezone       = gm.Util.Env.GetString("APP_TIMEZONE", "Asia/Jakarta")
		dvalMaxLifeTimeIns = gm.Util.Env.GetInt("DB_MAX_LIFE_TIME_INS", 300)
		dvalMaxIdleTimeIns = gm.Util.Env.GetInt("DB_MAX_IDLE_TIME_INS", 30)
		dvalMaxIdle        = gm.Util.Env.GetInt("DB_MAX_IDLE", 2)
		dvalMaxOpen        = gm.Util.Env.GetInt("DB_MAX_OPEN", 10)

		dvalUnsafeCompatibility = fm.Ptr(gm.Util.Env.GetBool("DB_UNSAFE_COMPATIBILITY", false))
		dvalPrintUnsafeError    = fm.Ptr(gm.Util.Env.GetBool("DB_PRINT_UNSAFE_ERROR", true))
		dvalAutoRebind          = fm.Ptr(gm.Util.Env.GetBool("DB_AUTO_REBIND", true))
		dvalPrintSql            = fm.Ptr(gm.Util.Env.GetBool("DB_PRINT_SQL", true))
	)

	conf.Host = fm.Ternary(conf.Host != "", conf.Host, dvalHost)
	conf.Port = fm.Ternary(conf.Port != 0, conf.Port, dvalPort)
	conf.Scheme = fm.Ternary(conf.Scheme != "", conf.Scheme, dvalScheme)
	conf.Timezone = fm.Ternary(conf.Timezone != "", conf.Timezone, dvalTimezone)
	conf.MaxLifeTimeIns = fm.Ternary(conf.MaxLifeTimeIns != 0, conf.MaxLifeTimeIns, dvalMaxLifeTimeIns)
	conf.MaxIdleTimeIns = fm.Ternary(conf.MaxIdleTimeIns != 0, conf.MaxIdleTimeIns, dvalMaxIdleTimeIns)
	conf.MaxIdle = fm.Ternary(conf.MaxIdle != 0, conf.MaxIdle, dvalMaxIdle)
	conf.MaxOpen = fm.Ternary(conf.MaxOpen != 0, conf.MaxOpen, dvalMaxOpen)

	conf.UnsafeCompatibility = fm.Ternary(conf.UnsafeCompatibility != nil, conf.UnsafeCompatibility, dvalUnsafeCompatibility)
	conf.PrintUnsafeError = fm.Ternary(conf.PrintUnsafeError != nil, conf.PrintUnsafeError, dvalPrintUnsafeError)
	conf.AutoRebind = fm.Ternary(conf.AutoRebind != nil, conf.AutoRebind, dvalAutoRebind)
	conf.PrintSql = fm.Ternary(conf.PrintSql != nil, conf.PrintSql, dvalPrintSql)
}

func getPgConnectionString(conf *mol.DbConnection) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s application_name=%s timezone=%s sslmode=disable",
		conf.Host, conf.Port, conf.Username, conf.Password, conf.Name, conf.Scheme, conf.AppName, conf.Timezone)
}

func (slf *stuConnection) createConnection() error {
	connStr := getPgConnectionString(slf.conf)
	instance, err := sqlx.Connect(slf.driverName, connStr)
	if err == nil {
		instance.SetConnMaxLifetime(time.Second * time.Duration(slf.conf.MaxLifeTimeIns))
		instance.SetConnMaxIdleTime(time.Second * time.Duration(slf.conf.MaxIdleTimeIns))
		instance.SetMaxIdleConns(slf.conf.MaxIdle)
		instance.SetMaxOpenConns(slf.conf.MaxOpen)
		err = instance.Ping()
	}

	if err == nil {
		slf.sx = instance
	}

	return err
}

func (slf *stuConnection) normalizeQueryArgs(query string, args []any) (string, []any) {
	if len(args) == 1 {
		switch val := args[0].(type) {
		case []any:
			args = val
		}
	}

	if slf.driverName == "postgres" && *slf.conf.AutoRebind && len(args) > 0 {
		query = sqlx.Rebind(sqlx.DOLLAR, query)
	}
	return query, args
}

func (slf *stuConnection) execute(tx ice.DbTx, query string, args ...any) (sql.Result, error) {
	query, args = slf.normalizeQueryArgs(query, args)

	if tx != nil {
		switch v := tx.(type) {
		case *pgInstanceTx:
			stmt, err := v.tx.Prepare(query)
			if err != nil {
				return nil, err
			}
			defer func() {
				_ = stmt.Close()
			}()

			res, err := stmt.Exec(args...)
			return res, err
		}
	}

	stmt, err := slf.sx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	res, err := stmt.Exec(args...)
	return res, err
}

func (slf *stuConnection) executeRID(tx ice.DbTx, query string, args ...any) (*int64, string, error) {
	query, args = slf.normalizeQueryArgs(query, args)

	if tx != nil {
		switch v := tx.(type) {
		case *pgInstanceTx:
			stmt, err := v.tx.Prepare(query)
			if err != nil {
				return nil, slf.conf.Host, err
			}
			defer func() {
				_ = stmt.Close()
			}()

			var id *int64
			err = stmt.QueryRow(args...).Scan(&id)
			if err != nil {
				return nil, slf.conf.Host, err
			}

			return id, slf.conf.Host, err
		}
	}

	stmt, err := slf.sx.Prepare(query)
	if err != nil {
		return nil, slf.conf.Host, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	var id *int64
	err = stmt.QueryRow(args...).Scan(&id)
	if err != nil {
		return nil, slf.conf.Host, err
	}

	return id, slf.conf.Host, err
}

func (slf *stuConnection) printSql(startTime time.Time, query string, args []any) {
	if *slf.conf.PrintSql {
		durationMs := gm.Util.Timenow().Sub(startTime).Milliseconds()
		sqlSpaceNL := fm.Ternary(strings.Contains(query, "\n"), "\n", " ")
		if len(args) == 0 {
			log.Printf("\nHOST: %v\nSQL :%v%v\nDUR : %v ms\n---\n\n", slf.conf.Host, sqlSpaceNL, query, durationMs)
		} else {
			var argsVal any = args
			if v, err := gm.Json.Encode(args); err == nil {
				argsVal = v
			}
			log.Printf("\nHOST: %v\nSQL :%v%v\nARGS: %v\nDUR : %v ms\n---\n\n", slf.conf.Host, sqlSpaceNL, query, argsVal, durationMs)
		}
	}
}
