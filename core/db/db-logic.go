/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/model"
	"github.com/jmoiron/sqlx"
)

func setPgConfDVal(conf *model.DbConnection) {
	if conf.AppName == "" {
		log.Fatalf("db connection must have application name")
	}

	conf.Host = fm.Ternary(conf.Host != "", conf.Host, "127.0.0.1")
	conf.Port = fm.Ternary(conf.Port != 0, conf.Port, 5432)
	conf.Scheme = fm.Ternary(conf.Scheme != "", conf.Scheme, "public")
	conf.MaxLifeTimeIns = fm.Ternary(conf.MaxLifeTimeIns != 0, conf.MaxLifeTimeIns, 300)
	conf.MaxIdleTimeIns = fm.Ternary(conf.MaxIdleTimeIns != 0, conf.MaxIdleTimeIns, 30)
	conf.MaxIdle = fm.Ternary(conf.MaxIdle != 0, conf.MaxIdle, 2)
	conf.MaxOpen = fm.Ternary(conf.MaxOpen != 0, conf.MaxOpen, 10)
}

func getPgConnectionString(conf *model.DbConnection) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s application_name=%s sslmode=disable", conf.Host, conf.Port, conf.Username, conf.Password, conf.Name, conf.Scheme, conf.AppName)
}

func createConnection(conn *srConnection) (*sqlx.DB, error) {
	connStr := getPgConnectionString(conn.conf)
	instance, err := sqlx.Connect(conn.driverName, connStr)
	if err == nil {
		instance.SetConnMaxLifetime(time.Second * time.Duration(conn.conf.MaxLifeTimeIns))
		instance.SetConnMaxIdleTime(time.Second * time.Duration(conn.conf.MaxIdleTimeIns))
		instance.SetMaxIdleConns(conn.conf.MaxIdle)
		instance.SetMaxOpenConns(conn.conf.MaxOpen)
		err = instance.Ping()
	}

	return instance, err
}

func normalizeSqlQueryParams(conn *srConnection, query string, args []interface{}) (string, []interface{}) {
	if len(args) == 1 {
		switch val := args[0].(type) {
		case []interface{}:
			args = val
		}
	}

	if conn.driverName == "postgres" && conn.conf.AutoRebind && len(args) > 0 {
		query = sqlx.Rebind(sqlx.DOLLAR, query)
	}
	return query, args
}

func execute(conn *srConnection, tx ice.DbTx, query string, args ...interface{}) (sql.Result, error) {
	query, args = normalizeSqlQueryParams(conn, query, args)

	if tx != nil {
		switch v := tx.(type) {
		case *pqInstanceTx:
			stmt, err := v.tx.Prepare(query)
			if err != nil {
				return nil, err
			}
			defer stmt.Close()

			res, err := stmt.Exec(args...)
			return res, err
		}
	}

	stmt, err := conn.sx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	return res, err
}

func executeRID(conn *srConnection, tx ice.DbTx, query string, args ...interface{}) (*int64, string, error) {
	query, args = normalizeSqlQueryParams(conn, query, args)

	if tx != nil {
		switch v := tx.(type) {
		case *pqInstanceTx:
			stmt, err := v.tx.Prepare(query)
			if err != nil {
				return nil, conn.conf.Host, err
			}
			defer stmt.Close()

			var id *int64
			err = stmt.QueryRow(args...).Scan(&id)
			if err != nil {
				return nil, conn.conf.Host, err
			}

			return id, conn.conf.Host, err
		}
	}

	stmt, err := conn.sx.Prepare(query)
	if err != nil {
		return nil, conn.conf.Host, err
	}
	defer stmt.Close()

	var id *int64
	err = stmt.QueryRow(args...).Scan(&id)
	if err != nil {
		return nil, conn.conf.Host, err
	}

	return id, conn.conf.Host, err
}

func printSql(conn *srConnection, startTime time.Time, query string, args ...interface{}) {
	if conn.conf.PrintSql {
		durationMs := gm.Util.Timenow().Sub(startTime).Milliseconds()
		if len(args) == 0 {
			log.Printf("\nHOST: \"%v\"\nSQL: \"%v\"\nDUR: %vms", conn.conf.Host, query, durationMs)
		} else {
			log.Printf("\nHOST: \"%v\"\nSQL: \"%v\"\nARGS: %v\nDUR: %vms", conn.conf.Host, query, args, durationMs)
		}
	}
}
