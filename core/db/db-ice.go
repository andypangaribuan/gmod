/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

type Repo[T any] interface {
	SetInsertColumn(columns string)
	Fetch(condition string, args ...interface{}) (*T, error)
}
