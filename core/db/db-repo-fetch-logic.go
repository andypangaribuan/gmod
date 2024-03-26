/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/model"
)

func (slf *srRepo[T]) fetches(tx ice.DbTx, whereQuery, endQuery string, args []interface{}) ([]T, *srReport, error) {
	report := &srReport{
		tableName:     slf.tableName,
		insertColumn:  slf.insertColumn,
		insertArgSign: slf.insertArgSign,
		query:         "SELECT * FROM ::tableName",
		args:          args,
	}

	if whereQuery != "" {
		report.query += " " + whereQuery
	}

	if endQuery != "" {
		report.query += " " + endQuery
	}

	var (
		err        error
		out        []T
		execReport *model.DbExecReport
	)

	err = report.transform()
	if err != nil {
		return nil, report, err
	}

	if tx != nil {
		execReport, err = slf.ins.TxSelect(tx, &out, report.query, report.args...)
		report.execReport = execReport
		return out, report, err
	}

	if slf.rwFetchWhenNull {
		check := func() bool {
			return out != nil && len(out) > 0
		}
		execReport, err = slf.ins.SelectR2(&out, check, report.query, report.args...)
	} else {
		execReport, err = slf.ins.Select(&out, report.query, report.args...)
	}

	return out, report, err
}
