/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import "strings"

func (*srRepo[T]) generateArgSign(column string) (argSign string) {
	if column != "" {
		ls := strings.Split(column, ",")
		for i := range ls {
			if i > 0 {
				argSign += ","
			}
			argSign += " ?"
		}

		argSign = strings.TrimSpace(argSign)
	}

	return
}

func (*srRepo[T]) formatInsertColumnArgs(val string) string {
	var (
		maxOneLine = 5
		fourSpace  = "    "
		formatted  = ""
		count      = 0
		ls         = strings.Split(val, ",")
	)

	for i, v := range ls {
		if count > maxOneLine {
			count = 0
		}
		count++

		if count == 1 {
			formatted += fourSpace
		}
		formatted += strings.TrimSpace(v)

		if i < len(ls)-1 {
			formatted += ","
			if count <= maxOneLine {
				formatted += " "
			}
		}
	}

	return formatted
}

func (slf *srRepo[T]) getWhereQuery(condition string, args ...interface{}) string {
	condition = strings.TrimSpace(condition)
	var (
		whereQuery          = ""
		withDeletedAtIsNull = slf.isWithDeletedAtIsNull(args...)
	)

	if withDeletedAtIsNull {
		whereQuery = "WHERE deleted_at IS NULL"
		if condition != "" {
			whereQuery += " AND " + condition
		}
	} else if condition != "" {
		whereQuery = "WHERE " + condition
	}

	return whereQuery
}

func (slf *srRepo[T]) isWithDeletedAtIsNull(args ...interface{}) bool {
	isWith := slf.withDeletedAtIsNull

	for _, arg := range args {
		switch v := arg.(type) {
		case FetchOpt:
			if v.WithDeletedAtIsNull != nil {
				isWith = *v.WithDeletedAtIsNull
			}

		case *FetchOpt:
			if v != nil && v.WithDeletedAtIsNull != nil {
				isWith = *v.WithDeletedAtIsNull
			}
		}
	}

	return isWith
}

func (slf *srRepo[T]) getEndQuery(args ...interface{}) string {
	endQuery := ""

	for _, arg := range args {
		switch v := arg.(type) {
		case FetchOpt:
			if v.EndQuery != nil {
				if endQuery != "" {
					endQuery += " "
				}
				endQuery += strings.TrimSpace(*v.EndQuery)
			}

		case *FetchOpt:
			if v != nil && v.EndQuery != nil {
				if endQuery != "" {
					endQuery += " "
				}
				endQuery += strings.TrimSpace(*v.EndQuery)
			}
		}
	}

	return strings.TrimSpace(endQuery)
}

func (slf *srRepo[T]) getArgs(args ...interface{}) []interface{} {
	filtered := make([]interface{}, 0)

	for _, arg := range args {
		switch arg.(type) {
		case FetchOpt, *FetchOpt:
			continue
		default:
			filtered = append(filtered, arg)
		}
	}

	return filtered
}
