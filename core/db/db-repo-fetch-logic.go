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

func (slf *stuRepo[T]) vfetches(isFetch bool, tx ice.DbTx, condition string, args []any) ([]T, *stuReport, error) {
	entities, report, err := slf.fetches(isFetch, tx, condition, args)
	ls := make([]T, len(entities))

	for i, e := range entities {
		ls[i] = *e
	}

	return ls, report, err
}

func (slf *stuRepo[T]) fetches(isFetch bool, tx ice.DbTx, condition string, args []any) ([]*T, *stuReport, error) {
	var (
		whereQuery = slf.getWhereQuery(condition, args)
		endQuery   = strings.TrimSpace(slf.getEndQuery(args) + fm.Ternary(isFetch, " LIMIT 1", ""))
		report     = &stuReport{
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
		execReport *mol.DbExecReport
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
