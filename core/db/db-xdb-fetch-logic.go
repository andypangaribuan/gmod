/*
 * Copyright (c) 2025.
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

func (slf *stuXDB) fetches(tx ice.DbTx, query string, args []any) *stuRepoResult[map[string]any] {
	var (
		report = &stuReport{
			args:  slf.getArgs(args),
			query: query,
		}
	)

	var (
		err        error
		out        []*map[string]any
		execReport *mol.DbExecReport
	)

	err = report.transform()
	if err != nil {
		return slf.result(report, err, nil, nil)
	}

	if tx != nil {
		execReport, err = slf.ins.TxSelect(tx, &out, report.query, report.args...)
	} else {
		usingRW := false
		execReport, err = slf.ins.Select(&out, usingRW, report.query, report.args...)
	}

	report.execReport = execReport
	return slf.result(report, err, nil, out)
}

func (slf *stuXDB) execute(tx ice.DbTx, rid bool, query string, args []any) *stuRepoResult[map[string]any] {
	var (
		report = &stuReport{
			args:  slf.getArgs(args),
			query: query,
		}
	)

	var (
		err        error
		id         *int64
		execReport *mol.DbExecReport
	)

	err = report.transform()
	if err != nil {
		return slf.result(report, err, nil, nil)
	}

	if tx != nil {
		if rid {
			id, execReport, err = slf.ins.TxExecuteRID(tx, query, args...)
		} else {
			execReport, err = slf.ins.TxExecute(tx, query, args...)
		}
	} else {
		if rid {
			id, execReport, err = slf.ins.ExecuteRID(query, args...)
		} else {
			execReport, err = slf.ins.Execute(query, args...)
		}
	}

	report.execReport = execReport
	return slf.result(report, err, id, nil)
}
