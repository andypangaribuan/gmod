/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/ice"

func NewRepo[T any](db ice.DbInstance, tableName string, opt ...RepoOpt) Repo[T] {
	sr := &srRepo[T]{
		ins:                 db,
		tableName:           tableName,
		withDeletedAtIsNull: true,
		rwFetchWhenNull:     true,
	}

	for _, val := range opt {
		if val.WithDeletedAtIsNull != nil {
			sr.withDeletedAtIsNull = *val.WithDeletedAtIsNull
		}
		if val.RWFetchWhenNull != nil {
			sr.rwFetchWhenNull = *val.RWFetchWhenNull
		}
	}

	return sr
}
