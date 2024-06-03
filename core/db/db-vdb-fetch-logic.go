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

func (slf *stuVDB[T]) fetches(isLimitOne bool, tx ice.DbTx, sqlName string, args []any) *stuRepoResult[T] {
	var (
		endQuery  = strings.TrimSpace(slf.getQuery("end-query", args) + fm.Ternary(isLimitOne, " LIMIT 1", ""))
		fullQuery = slf.getQuery("full-query", args)
		report    = &stuReport{
			args:  slf.getArgs(args),
			query: slf.dvalSql[sqlName],
		}
	)

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
		execReport, err = slf.ins.Select(&out, report.query, report.args...)
	}

	report.execReport = execReport
	return slf.result(report, err, nil, out)
}
