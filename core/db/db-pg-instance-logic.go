/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"fmt"
	"log"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/model"
)

func updateReport(report *model.DbExecReport) {
	report.FinishedAt = gm.Util.Timenow()
	report.DurationMs = report.FinishedAt.Sub(report.StartedAt).Milliseconds()
}

func updateReportHost(conn *srConnection, reportHost *model.DbExecReportHost) {
	if conn != nil {
		reportHost.Host = conn.conf.Host
	}
	reportHost.FinishedAt = gm.Util.Timenow()
	reportHost.DurationMs = reportHost.FinishedAt.Sub(reportHost.StartedAt).Milliseconds()
}

func (slf *pgInstance) execute(rid bool, tx ice.DbTx, query string, args ...interface{}) (*int64, *model.DbExecReport, error) {
	report := &model.DbExecReport{
		StartedAt: gm.Util.Timenow(),
		Hosts:     make([]*model.DbExecReportHost, 0),
	}
	defer updateReport(report)

	var (
		conn *srConnection
		err  error
		id   *int64
	)

	reportHost := &model.DbExecReportHost{StartedAt: report.StartedAt}
	report.Hosts = append(report.Hosts, reportHost)
	defer updateReportHost(conn, reportHost)

	conn, err = slf.crw()
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

func (slf *pgInstance) execSelect(conn *srConnection, reportHost *model.DbExecReportHost, insTx *pgInstanceTx, out interface{}, query string, args []interface{}) (err error) {
	defer updateReportHost(conn, reportHost)

	if insTx != nil {
		err = insTx.tx.Select(out, query, args...)
	} else {
		err = conn.sx.Select(out, query, args...)
	}

	if err != nil && *conn.conf.UnsafeCompatibility {
		unsafe := &srUnsafe{
			query:   query,
			args:    args,
			message: err.Error(),
			trace:   fmt.Sprintf("%+v", err),
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

func (slf *pgInstance) onUnsafe(unsafe *srUnsafe) {
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
			gm.Conv.Time.ToStrDT(gm.Util.Timenow()),
			unsafe.message,
			unsafe.query,
			unsafe.args,
			unsafe.trace,
		)
	}
}
