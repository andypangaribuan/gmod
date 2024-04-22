/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "strings"

func (*stuRepo[T]) generateArgSign(column string) (argSign string) {
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

func (*stuRepo[T]) formatInsertColumnArgs(val string) string {
	var (
		maxOneLine = 5
		fourSpace  = "    "
		formatted  = ""
		count      = 0
		ls         = strings.Split(val, ",")
	)

	for i, v := range ls {
		if count >= maxOneLine {
			count = 0
		}
		count++

		if count == 1 {
			if i > 0 {
				formatted += "\n"
			}
			formatted += fourSpace
		}
		formatted += strings.TrimSpace(v)

		if i < len(ls)-1 {
			formatted += ","
			if count < maxOneLine {
				formatted += " "
			}
		}
	}

	return formatted
}

func (slf *stuRepo[T]) getWhereQuery(condition string, args []any) string {
	condition = strings.TrimSpace(condition)
	var (
		whereQuery          = ""
		withDeletedAtIsNull = slf.isWithDeletedAtIsNull(args)
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

func (slf *stuRepo[T]) isWithDeletedAtIsNull(args []any) bool {
	isWith := slf.withDeletedAtIsNull

	for _, arg := range args {
		switch val := arg.(type) {
		case FetchOptBuilder:
			v, ok := val.(*stuFetchOptBuilder)
			if ok && v != nil && v.withDeletedAtIsNull != nil {
				isWith = *v.withDeletedAtIsNull
			}
		}
	}

	return isWith
}

func (slf *stuRepo[T]) getQuery(typ string, args []any) string {
	query := ""

	for _, arg := range args {
		switch val := arg.(type) {
		case FetchOptBuilder:
			v, ok := val.(*stuFetchOptBuilder)
			if ok && v != nil {
				switch typ {
				case "end-query":
					if v.endQuery != nil {
						if query != "" {
							query += " "
						}
						query += strings.TrimSpace(*v.endQuery)
					}

				case "full-query":
					if v.fullQuery != nil {
						if query != "" {
							query += " "
						}
						query += strings.TrimSpace(*v.fullQuery)
					}
				}
			}
		}
	}

	return strings.TrimSpace(query)
}

func (slf *stuRepo[T]) getArgs(args []any) []any {
	filtered := make([]any, 0)

	for _, arg := range args {
		switch arg.(type) {
		case FetchOptBuilder, *stuRepoFuncOpt, RepoFuncOpt:
			continue
		default:
			filtered = append(filtered, arg)
		}
	}

	return filtered
}
