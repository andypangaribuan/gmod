/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import "sort"

const (
	OrderedLevelRepo    = 2
	OrderedLevelService = 9
)

func OrderedInit(level int, fn func()) {
	orderedInitLs = append(orderedInitLs, []any{level, fn})
}

func CallOrderedInit() {
	if len(orderedInitLs) > 0 {
		sort.SliceStable(orderedInitLs, func(i, j int) bool {
			return orderedInitLs[i][0].(int) < orderedInitLs[j][0].(int)
		})

		for _, v := range orderedInitLs {
			v[1].(func())()
		}
	}
}
