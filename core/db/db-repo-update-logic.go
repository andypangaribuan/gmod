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
	"fmt"
	"sort"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mdl"
)

func (slf *stuRepo[T]) update(tx ice.DbTx, builder *stuUpdateBuilder) (*stuReport, error) {
	var (
		report = &stuReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			query:         `UPDATE ::tableName SET`,
		}
		err           error
		execReport    *mdl.DbExecReport
		withUpdatedAt = fm.GetDefault(builder.withAutoUpdatedAt, true)
		setQuery      = ""
		setArgs       = make([]any, 0)
		whereQuery    = ""
		whereArgs     = make([]any, 0)
	)

	if withUpdatedAt {
		setQuery += fm.Ternary(setQuery == "", "", ", ")
		setQuery += "updated_at=?"
		setArgs = append(setArgs, gm.Util.Timenow())
	}

	if builder.setQuery != nil {
		setQuery += fm.Ternary(setQuery == "", "", ", ")
		setQuery += *builder.setQuery
	}

	if builder.setArgs != nil && len(*builder.setArgs) > 0 {
		setArgs = append(setArgs, *builder.setArgs...)
	}

	if builder.setInn != nil && len(*builder.setInn) > 0 {
		var (
			query = ""
			args  = make([]any, 0)
			keys  = make([]string, 0, len(*builder.setInn))
		)

		for key := range *builder.setInn {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			arg := (*builder.setInn)[key]
			if !fm.IsNil(arg) {
				query += fm.Ternary(query == "", "", ", ")
				query += fmt.Sprintf("%v=?", key)
				args = append(args, arg)
			}
		}

		if query != "" {
			setQuery += fm.Ternary(setQuery == "", "", ", ")
			setQuery += query
		}

		if len(args) > 0 {
			setArgs = append(setArgs, args...)
		}
	}

	if builder.whereQuery != nil {
		whereQuery = "\nWHERE " + *builder.whereQuery
	}

	if builder.whereArgs != nil && len(*builder.whereArgs) > 0 {
		whereArgs = append(whereArgs, *builder.whereArgs...)
	}

	report.query += " " + setQuery + whereQuery
	report.args = append(setArgs, whereArgs...)

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
