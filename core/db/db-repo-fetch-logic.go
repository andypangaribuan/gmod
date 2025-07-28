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
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
)

func (slf *stuRepo[T]) fetches(isLimitOne bool, tx ice.DbTx, condition string, args []any) *stuRepoResult[T] {
	var (
		isDirectFullQuery = false
		whereQuery        string
		endQuery          string
		fullQuery         string
		report            = &stuReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			args:          slf.getArgs(args),
			query:         "SELECT * FROM ::tableName",
		}
	)

	condition = strings.TrimSpace(condition)
	query := strings.ToLower(condition)

	if len(query) > 6 && query[:6] == "select" {
		isDirectFullQuery = true
	}

	if len(query) > 4 && query[:4] == "with" {
		isDirectFullQuery = true
	}

	if isDirectFullQuery {
		report.query = condition
	} else {
		whereQuery = slf.getWhereQuery(condition, args)
		endQuery = strings.TrimSpace(slf.getQuery("end-query", args) + fm.Ternary(isLimitOne, " LIMIT 1", ""))
		fullQuery = slf.getQuery("full-query", args)
	}

	if whereQuery != "" {
		report.query += " " + whereQuery
	}

	if endQuery != "" {
		report.query += " " + endQuery
	}

	if fullQuery != "" {
		report.query = fullQuery
	}

	var (
		err        error
		out        []*T
		execReport *mol.DbExecReport
	)

	err = report.transform()
	if err != nil {
		return slf.result(report, err, nil, nil)
	}

	if tx != nil {
		execReport, err = slf.ins.TxSelect(tx, &out, report.query, report.args...)
	} else {
		usingRW := slf.isUsingRW(args)
		if slf.rwFetchWhenNull && !usingRW {
			execReport, err = slf.ins.SelectR2(&out, report.query, report.args, fm.Ptr(func() bool { return len(out) > 0 }))
		} else {
			execReport, err = slf.ins.Select(&out, usingRW, report.query, report.args...)
		}
	}

	report.execReport = execReport
	return slf.result(report, err, nil, out)
}
