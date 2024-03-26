/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"strings"

	"github.com/andypangaribuan/gmod/fm"
)

func (slf *srRepo[T]) SetInsertColumn(columns string) {
	columns = strings.TrimSpace(columns)
	for {
		columns = strings.ReplaceAll(columns, " ", "")
		columns = strings.ReplaceAll(columns, "\n", "")
		if !strings.Contains(columns, " ") && !strings.Contains(columns, "\n") {
			break
		}
	}

	slf.insertColumn = slf.formatInsertColumnArgs(columns)
	slf.insertArgSign = slf.formatInsertColumnArgs(slf.generateArgSign(columns))
}

func (slf *srRepo[T]) Fetch(condition string, args ...interface{}) (*T, error) {
	var (
		whereQuery = slf.getWhereQuery(condition, args...)
		endQuery   = strings.TrimSpace(slf.getEndQuery(args...) + " LIMIT 1")
	)

	models, _, err := slf.fetches(nil, whereQuery, endQuery, args)
	return fm.GetFirst(models), err
}

func (slf *srRepo[T]) Fetches(condition string, args ...interface{}) ([]T, error) {
	var (
		whereQuery = slf.getWhereQuery(condition, args...)
		endQuery   = slf.getEndQuery(args...)
	)

	models, _, err := slf.fetches(nil, whereQuery, endQuery, args)
	return models, err
}
