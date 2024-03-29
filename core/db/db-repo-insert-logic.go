/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"errors"
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/model"
)

func (slf *srRepo[T]) insert(tx ice.DbTx, rid bool, args []interface{}) (*int64, *srReport, error) {
	var (
		report = &srReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			args:          args,
			query: `INSERT INTO ::tableName (
::insertColumn
) VALUES (
::insertArgSign
)`,
		}
		err        error
		id         *int64
		execReport *model.DbExecReport
	)

	err = report.transform()
	if err != nil {
		return id, report, err
	}

	if tx != nil {
		if rid {
			id, execReport, err = slf.ins.TxExecuteRID(tx, report.query, report.args...)
		} else {
			execReport, err = slf.ins.TxExecute(tx, report.query, report.args...)
		}
	} else {
		if rid {
			id, execReport, err = slf.ins.ExecuteRID(report.query, report.args...)
		} else {
			execReport, err = slf.ins.Execute(report.query, report.args...)
		}
	}

	report.execReport = execReport
	return id, report, err
}

func (slf *srRepo[T]) bulkInsert(tx ice.DbTx, entities []*T, args func(e *T) []interface{}, chunkSize ...int) (*srReport, error) {
	if tx == nil {
		return nil, errors.New("db: bulk insert only available via transaction")
	}

	var (
		report = &srReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			query: `INSERT INTO ::tableName (
::insertColumn
) VALUES `,
			execReport: &model.DbExecReport{
				StartedAt: gm.Util.Timenow(),
			},
		}
		err             error
		insertChunkSize = *fm.GetFirst(chunkSize, 100)
		partSize        = make([]int, 0)
		partArgs        = make([][]interface{}, 0)
		part            = make([]interface{}, 0)
		count           = 0
	)

	defer func() {
		report.execReport.FinishedAt = gm.Util.Timenow()
		report.execReport.DurationMs = report.execReport.FinishedAt.Sub(report.execReport.StartedAt).Milliseconds()
	}()

	for _, e := range entities {
		count++
		ar := args(e)
		if len(ar) == 0 {
			break
		}

		part = append(part, ar...)
		if count == insertChunkSize {
			count = 0
			partSize = append(partSize, len(part))
			partArgs = append(partArgs, part)
			part = make([]interface{}, 0)
		}
	}

	if len(part) > 0 {
		partSize = append(partSize, len(part))
		partArgs = append(partArgs, part)
	}

	if len(partSize) == 0 {
		return nil, errors.New("db: no data to process")
	}

	valFormat := `(
::insertArgSign
)`

	for i, args := range partArgs {
		size := partSize[i]
		query := strings.ReplaceAll(report.query, "::tableName", report.tableName)
		query = strings.ReplaceAll(query, "::insertColumn", report.insertColumn)
		valQuery := strings.ReplaceAll(valFormat, "::insertArgSign", report.insertArgSign)

		for i := 0; i < size; i++ {
			query += fm.Ternary(i == 0, valQuery, ", "+valQuery)
		}

		_, err = slf.ins.TxExecute(tx, query, args...)
		if err != nil {
			return report, err
		}
	}

	return report, nil
}
