/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mdl"
)

func (slf *srRepo[T]) fetches(isFetch bool, tx ice.DbTx, condition string, args []interface{}) ([]*T, *srReport, error) {
	var (
		whereQuery = slf.getWhereQuery(condition, args)
		endQuery   = strings.TrimSpace(slf.getEndQuery(args) + fm.Ternary(isFetch, " LIMIT 1", ""))
		report     = &srReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			args:          slf.getArgs(args),
			query:         "SELECT * FROM ::tableName",
		}
	)

	if whereQuery != "" {
		report.query += " " + whereQuery
	}

	if endQuery != "" {
		report.query += " " + endQuery
	}

	var (
		err        error
		out        []*T
		execReport *mdl.DbExecReport
	)

	err = report.transform()
	if err != nil {
		return nil, report, err
	}

	if tx != nil {
		execReport, err = slf.ins.TxSelect(tx, &out, report.query, report.args...)
	} else {
		if slf.rwFetchWhenNull {
			execReport, err = slf.ins.SelectR2(&out, report.query, report.args, fm.Ptr(func() bool { return len(out) > 0 }))
		} else {
			execReport, err = slf.ins.Select(&out, report.query, report.args...)
		}
	}

	report.execReport = execReport
	return out, report, err
}
