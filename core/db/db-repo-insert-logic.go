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
	"errors"
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
)

func (slf *stuRepo[T]) insert(tx ice.DbTx, rid bool, args []any) *stuRepoResult[T] { // (*int64, *stuReport, error) {
	var (
		report = &stuReport{
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
		execReport *mol.DbExecReport
	)

	if rid {
		report.query += `
RETURNING id`
	}

	err = report.transform()
	if err != nil {
		return slf.result(report, err, id, nil)
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
	return slf.result(report, err, id, nil)
}

func (slf *stuRepo[T]) bulkInsert(tx ice.DbTx, entities []*T, chunkSize ...int) *stuRepoResult[T] { // (*stuReport, error) {
	if tx == nil {
		return slf.result(nil, errors.New("db: bulk insert only available via transaction"), nil, nil)
	}

	var (
		report = &stuReport{
			tableName:     slf.tableName,
			insertColumn:  slf.insertColumn,
			insertArgSign: slf.insertArgSign,
			query: `INSERT INTO ::tableName (
::insertColumn
) VALUES `,
			execReport: &mol.DbExecReport{
				StartedAt: gm.Util.Timenow(),
			},
		}
		err             error
		insertChunkSize = *fm.GetFirst(chunkSize, 1000)
		partSize        = make([]int, 0)
		partArgs        = make([][]any, 0)
		part            = make([]any, 0)
		count           = 0
	)

	defer func() {
		report.execReport.FinishedAt = gm.Util.Timenow()
		report.execReport.DurationMs = report.execReport.FinishedAt.Sub(report.execReport.StartedAt).Milliseconds()
	}()

	for _, e := range entities {
		count++
		ar := slf.insertColumnFunc(e)
		if len(ar) == 0 {
			break
		}

		part = append(part, ar...)
		if count == insertChunkSize {
			count = 0
			partSize = append(partSize, len(part))
			partArgs = append(partArgs, part)
			part = make([]any, 0)
		}
	}

	if len(part) > 0 {
		partSize = append(partSize, len(part))
		partArgs = append(partArgs, part)
	}

	if len(partSize) == 0 {
		return slf.result(nil, errors.New("db: no data to process"), nil, nil)
	}

	valFormat := `(
::insertArgSign
)`

	for i, args := range partArgs {
		size := partSize[i] / strings.Count(report.insertArgSign, "?")
		query := strings.ReplaceAll(report.query, "::tableName", report.tableName)
		query = strings.ReplaceAll(query, "::insertColumn", report.insertColumn)
		valQuery := strings.ReplaceAll(valFormat, "::insertArgSign", report.insertArgSign)

		for i := range size {
			query += fm.Ternary(i == 0, valQuery, ", "+valQuery)
		}

		_, err = slf.ins.TxExecute(tx, query, args...)
		if err != nil {
			return slf.result(report, err, nil, nil)
		}
	}

	return slf.result(report, nil, nil, nil)
}
