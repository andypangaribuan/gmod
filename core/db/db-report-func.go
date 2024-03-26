/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

func (slf *srReport) transform() error {
	slf.query = strings.ReplaceAll(slf.query, "::tableName", slf.tableName)
	slf.query = strings.ReplaceAll(slf.query, "::insertColumn", slf.insertColumn)
	slf.query = strings.ReplaceAll(slf.query, "::insertArgSign", slf.insertArgSign)

	if strings.Contains(strings.ToUpper(slf.query), " IN ") {
		query, args, err := sqlx.In(slf.query, slf.args...)
		if err != nil {
			return err
		}

		slf.query = query
		slf.args = args
	}

	return nil
}
