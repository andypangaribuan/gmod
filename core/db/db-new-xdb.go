/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/ice"

func NewXDB(db ice.DbInstance, opt ...XdbOptBuilder) XDB {
	stu := &stuXDB{
		ins:             db,
		rwFetchWhenNull: true,
	}

	for _, val := range opt {
		v, ok := val.(*stuXdbOptBuilder)
		if ok && v != nil {
			if v.rwFetchWhenNull != nil {
				stu.rwFetchWhenNull = *v.rwFetchWhenNull
			}
		}
	}

	return stu
}
