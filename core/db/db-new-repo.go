/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/ice"

func NewRepo[T any](db ice.DbInstance, tableName string, opt ...RepoOptBuilder) Repo[T] {
	stu := &stuRepo[T]{
		ins:                 db,
		tableName:           tableName,
		withDeletedAtIsNull: true,
		rwFetchWhenNull:     true,
	}

	for _, val := range opt {
		v, ok := val.(*stuRepoOptBuilder)
		if ok && v != nil {
			if v.withDeletedAtIsNull != nil {
				stu.withDeletedAtIsNull = *v.withDeletedAtIsNull
			}

			if v.rwFetchWhenNull != nil {
				stu.rwFetchWhenNull = *v.rwFetchWhenNull
			}
		}
	}

	return stu
}
