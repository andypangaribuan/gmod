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

	"github.com/andypangaribuan/gmod/ice"
)

func NewVDB[T any](db ice.DbInstance, dvalSql map[string]string) VDB[T] {
	for k, v := range dvalSql {
		dvalSql[k] = strings.TrimSpace(v)
	}

	stu := &stuVDB[T]{
		ins:     db,
		dvalSql: dvalSql,
	}

	return stu
}
