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

	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
)

func (slf *stuRepo[T]) delete(tx ice.DbTx, condition string, args []any) *stuRepoResult[T] {
	var (
		whereQuery = slf.getWhereQuery(condition, args)
		endQuery   = strings.TrimSpace(slf.getQuery("end-query", args))
		fullQuery  = slf.getQuery("full-query", args)
		report     = &stuReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			args:          slf.getArgs(args),
			query:         "DELETE FROM ::tableName",
		}
	)

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
		execReport *mol.DbExecReport
	)

	err = report.transform()
	if err != nil {
		return slf.result(report, err, nil, nil)
	}

	if tx == nil {
		execReport, err = slf.ins.Execute(report.query, report.args...)
	} else {
		execReport, err = slf.ins.TxExecute(tx, report.query, report.args...)
	}

	report.execReport = execReport
	return slf.result(report, err, nil, nil)
}
