/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/clog"

func (slf *stuXDB) getArgs(args []any) []any {
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

func (slf *stuXDB) result(report *stuReport, err error, id *int64, entities []*map[string]any) *stuRepoResult[map[string]any] {
	return &stuRepoResult[map[string]any]{
		report:   report,
		err:      err,
		id:       id,
		entities: entities,
	}
}

func (slf *stuXDB) override(logc clog.Instance, res *stuRepoResult[map[string]any]) *stuRepoResult[map[string]any] {
	pushClogReport(logc, res.report, res.err, 4)
	return res
}
