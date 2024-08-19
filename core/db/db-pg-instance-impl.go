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
	"errors"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
)

func (slf *pgInstance) crw() (*stuConnection, error) {
	connWriteLocking.Lock()
	defer connWriteLocking.Unlock()

	if slf.rw.sx == nil {
		err := slf.rw.createConnection()
		if err != nil {
			return nil, err
		}
	}

	return slf.rw, nil
}

func (slf *pgInstance) cro() (*stuConnection, error) {
	connReadLocking.Lock()
	defer connReadLocking.Unlock()

	if slf.ro == nil {
		return nil, errors.New("database configuration doesn't have read only connection")
	}

	if slf.ro.sx == nil {
		err := slf.ro.createConnection()
		if err != nil {
			return nil, err
		}
	}

	return slf.ro, nil
}

func (slf *pgInstance) Ping() (string, error) {
	conn, err := slf.crw()
	if err != nil {
		return conn.conf.Host, err
	}

	return conn.conf.Host, conn.sx.Ping()
}

func (slf *pgInstance) PingRead() (string, error) {
	conn, err := slf.cro()
	if err != nil {
		return conn.conf.Host, err
	}

	return conn.conf.Host, conn.sx.Ping()
}

func (slf *pgInstance) NewTransaction() (ice.DbTx, error) {
	conn, err := slf.crw()
	if err != nil {
		return nil, err
	}

	tx, err := conn.sx.Beginx()
	if err != nil {
		return nil, err
	}

	insx := &pgInstanceTx{
		ins: slf,
		tx:  tx,
	}

	return insx, err
}

func (slf *pgInstance) Select(out any, usingRW bool, query string, args ...any) (*mol.DbExecReport, error) {
	report := &mol.DbExecReport{
		StartedAt: gm.Util.Timenow(),
		Hosts:     make([]*mol.DbExecReportHost, 0),
	}
	defer updateReport(report)

	var (
		conn *stuConnection
		err  error
	)

	reportHost := &mol.DbExecReportHost{StartedAt: report.StartedAt}
	report.Hosts = append(report.Hosts, reportHost)

	if slf.ro != nil && !usingRW {
		conn, err = slf.cro()
	} else {
		conn, err = slf.crw()
	}

	if err != nil {
		updateReportHost(conn, reportHost)
		return report, err
	}

	query, args = conn.normalizeQueryArgs(query, args)
	report.Query = query
	report.Args = args

	err = slf.execSelect(conn, reportHost, nil, out, query, args)
	conn.printSql(reportHost.StartedAt, query, args)

	return report, err
}

func (slf *pgInstance) SelectR2(out any, query string, args []any, check *func() bool) (*mol.DbExecReport, error) {
	report := &mol.DbExecReport{
		StartedAt: gm.Util.Timenow(),
		Hosts:     make([]*mol.DbExecReportHost, 0),
	}
	defer updateReport(report)

	var (
		conn *stuConnection
		err  error
	)

	reportHost := &mol.DbExecReportHost{StartedAt: report.StartedAt}
	report.Hosts = append(report.Hosts, reportHost)

	if slf.ro != nil && check != nil {
		conn, err = slf.cro()
	} else {
		conn, err = slf.crw()
	}

	if err != nil {
		updateReportHost(conn, reportHost)
		return report, err
	}

	query, args = conn.normalizeQueryArgs(query, args)
	report.Query = query
	report.Args = args

	err = slf.execSelect(conn, reportHost, nil, out, query, args)
	conn.printSql(reportHost.StartedAt, query, args)

	if err != nil {
		return report, err
	}

	if slf.ro != nil && check != nil {
		c := *check
		if !c() {
			reportHost := &mol.DbExecReportHost{StartedAt: gm.Util.Timenow()}
			report.Hosts = append(report.Hosts, reportHost)

			conn, err = slf.crw()
			if err != nil {
				updateReportHost(conn, reportHost)
				return report, err
			}

			err = slf.execSelect(conn, reportHost, nil, out, query, args)
		}
	}

	return report, err
}

func (slf *pgInstance) Execute(query string, args ...any) (*mol.DbExecReport, error) {
	_, report, err := slf.execute(false, nil, query, args...)
	return report, err
}

func (slf *pgInstance) ExecuteRID(query string, args ...any) (*int64, *mol.DbExecReport, error) {
	return slf.execute(true, nil, query, args...)
}

func (slf *pgInstance) TxSelect(tx ice.DbTx, out any, query string, args ...any) (*mol.DbExecReport, error) {
	report := &mol.DbExecReport{
		StartedAt: gm.Util.Timenow(),
		Hosts:     make([]*mol.DbExecReportHost, 0),
	}
	defer updateReport(report)

	if tx == nil {
		return report, errors.New("db: tx is nil")
	}

	switch v := tx.(type) {
	case *pgInstanceTx:
		var (
			conn *stuConnection
			err  error
		)

		reportHost := &mol.DbExecReportHost{StartedAt: report.StartedAt}
		report.Hosts = append(report.Hosts, reportHost)

		conn, err = slf.crw()
		if err != nil {
			updateReportHost(conn, reportHost)
			return report, err
		}

		query, args = conn.normalizeQueryArgs(query, args)
		report.Query = query
		report.Args = args

		err = slf.execSelect(conn, reportHost, v, &out, query, args)
		conn.printSql(reportHost.StartedAt, query, args)

		return report, err
	}

	return report, errors.New("db: unknown tx transaction type")
}

func (slf *pgInstance) TxExecute(tx ice.DbTx, query string, args ...any) (*mol.DbExecReport, error) {
	_, report, err := slf.execute(false, tx, query, args...)
	return report, err
}

func (slf *pgInstance) TxExecuteRID(tx ice.DbTx, query string, args ...any) (*int64, *mol.DbExecReport, error) {
	return slf.execute(true, tx, query, args...)
}
