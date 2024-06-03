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

	"github.com/andypangaribuan/gmod/clog"
)

func (slf *stuVDB[T]) getFunc(typ string, args []any) any {
	for _, arg := range args {
		switch val := arg.(type) {
		case FetchOptBuilder:
			v, ok := val.(*stuFetchOptBuilder)
			if ok && v != nil {
				switch typ {
				case "full-query-formatter":
					if v.fullQueryFormatter != nil {
						return *v.fullQueryFormatter
					}
				}
			}
		}
	}

	return nil
}

func (slf *stuVDB[T]) getQuery(typ string, args []any) string {
	query := ""

	for _, arg := range args {
		switch val := arg.(type) {
		case FetchOptBuilder:
			v, ok := val.(*stuFetchOptBuilder)
			if ok && v != nil {
				switch typ {
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

func (slf *stuVDB[T]) getArgs(args []any) []any {
	filtered := make([]any, 0)

	for _, arg := range args {
		switch arg.(type) {
		case FetchOptBuilder:
			continue
		default:
			filtered = append(filtered, arg)
		}
	}

	return filtered
}

func (slf *stuVDB[T]) result(report *stuReport, err error, id *int64, entities []*T) *stuRepoResult[T] {
	return &stuRepoResult[T]{
		report:   report,
		err:      err,
		id:       id,
		entities: entities,
	}
}

func (slf *stuVDB[T]) override(clog clog.Instance, res *stuRepoResult[T]) *stuRepoResult[T] {
	pushClogReport(clog, res.report, res.err, 4)
	return res
}
