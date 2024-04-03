/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mdl"
)

func (slf *srRepo[T]) update(tx ice.DbTx, update *Update) (*srReport, error) {
	update.Set = strings.TrimSpace(update.Set)
	update.Where = strings.TrimSpace(update.Where)

	var (
		report = &srReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			query:         `UPDATE ::tableName SET`,
		}
		err          error
		execReport   *mdl.DbExecReport
		withUpdateAt = fm.GetDefault(update.WithUpdateAt, true)
	)

	if withUpdateAt && !strings.Contains(update.Set, "updated_at") {
		update.Set += ", updated_at=?"
		update.SetArgs = append(update.SetArgs, gm.Util.Timenow())
	}

	report.query += " " + update.Set
	report.query += fm.Ternary(update.Where == "", "", "\nWHERE "+update.Where)
	report.args = append(update.SetArgs, update.WhereArgs...)

	err = report.transform()
	if err != nil {
		return report, err
	}

	if tx != nil {
		execReport, err = slf.ins.TxExecute(tx, report.query, report.args...)
	} else {
		execReport, err = slf.ins.Execute(report.query, report.args...)
	}

	report.execReport = execReport
	return report, err
}
