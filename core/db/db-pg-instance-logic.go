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
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/andypangaribuan/gmod/fct"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func updateReport(report *mol.DbExecReport) {
	report.FinishedAt = gm.Util.Timenow()
	report.DurationMs = report.FinishedAt.Sub(report.StartedAt).Milliseconds()
}

func updateReportHost(conn *stuConnection, reportHost *mol.DbExecReportHost) {
	if conn != nil {
		reportHost.Host = conn.conf.Host
		reportHost.Name = conn.conf.Name
		reportHost.Scheme = conn.conf.Scheme
	}

	reportHost.FinishedAt = gm.Util.Timenow()
	reportHost.DurationMs = reportHost.FinishedAt.Sub(reportHost.StartedAt).Milliseconds()
}

func (slf *pgInstance) execute(rid bool, tx ice.DbTx, query string, args ...any) (*int64, *mol.DbExecReport, error) {
	report := &mol.DbExecReport{
		StartedAt: gm.Util.Timenow(),
		Hosts:     make([]*mol.DbExecReportHost, 0),
	}
	defer updateReport(report)

	var (
		conn *stuConnection
		err  error
		id   *int64
	)

	reportHost := &mol.DbExecReportHost{StartedAt: report.StartedAt}
	report.Hosts = append(report.Hosts, reportHost)

	conn, err = slf.crw()
	defer updateReportHost(conn, reportHost)

	if err != nil {
		return nil, report, err
	}

	query, args = conn.normalizeQueryArgs(query, args)
	report.Query = query
	report.Args = args

	if rid {
		id, _, err = conn.executeRID(tx, query, args...)
	} else {
		_, err = conn.execute(tx, query, args...)
	}

	conn.printSql(reportHost.StartedAt, query, args)
	return id, report, err
}

func (slf *pgInstance) execSelect(conn *stuConnection, reportHost *mol.DbExecReportHost, insTx *pgInstanceTx, out any, query string, args []any) (err error) {
	defer updateReportHost(conn, reportHost)

	if _, ok := out.(*[]map[string]any); ok {
		var (
			rows *sqlx.Rows
			err  error
			arr  = make([]map[string]any, 0)
		)
		if insTx != nil {
			rows, err = insTx.tx.Queryx(query, args...)
		} else {
			rows, err = conn.sx.Queryx(query, args...)
		}

		if err != nil {
			return err
		}

		for rows.Next() {
			kvs := make(map[string]any, 0)

			err := rows.MapScan(kvs)
			if err != nil {
				_ = rows.Close()
				return err
			}

			for key, val := range kvs {
				switch v := val.(type) {
				case []uint8:
					fv, err := fct.New(string(v))
					if err != nil {
						_ = rows.Close()
						return err
					}

					kvs[key] = fv
				}
			}

			arr = append(arr, kvs)
		}

		elem := reflect.ValueOf(out).Elem()
		elem.Set(reflect.ValueOf(arr))

		return err
	}

	if insTx != nil {
		err = insTx.tx.Select(out, query, args...)
	} else {
		err = conn.sx.Select(out, query, args...)
	}

	if err != nil && *conn.conf.UnsafeCompatibility {
		unsafe := &stuUnsafe{
			query:   query,
			args:    args,
			message: err.Error(),
			trace:   fmt.Sprintf("%+v", err),
		}

		if unsafe.message == unsafe.trace || len(strings.Split(unsafe.trace, "\n")) < 3 {
			unsafe.trace = fmt.Sprintf("%+v", errors.WithStack(err))
		}

		slf.onUnsafe(unsafe)

		if insTx != nil {
			err = insTx.tx.Unsafe().Select(out, query, args...)
		} else {
			err = conn.sx.Unsafe().Select(out, query, args...)
		}
	}

	return
}

func (slf *pgInstance) onUnsafe(unsafe *stuUnsafe) {
	printUnsafeLog := false
	if unsafe != nil {
		if slf.ro != nil {
			printUnsafeLog = *slf.ro.conf.PrintUnsafeError
		} else if slf.rw != nil {
			printUnsafeLog = *slf.rw.conf.PrintUnsafeError
		}
	}

	if printUnsafeLog {
		log.Printf("[%v] db unsafe select\nmessage: %v\nquery: %v\nargs: %v\ntrace: %v\n",
			gm.Conv.Time.ToStrDateTime(gm.Util.Timenow()),
			unsafe.message,
			unsafe.query,
			unsafe.args,
			unsafe.trace,
		)
	}
}
