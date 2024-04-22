/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package repo

import "github.com/andypangaribuan/gmod/core/db"

type xrepo[T any] struct {
	repo db.Repo[T]
}

func (slf *xrepo[T]) Insert(e *T) error {
	return slf.repo.Insert(e)
}

func (slf *xrepo[T]) Fetches(condition string, args ...any) ([]*T, error) {
	return slf.repo.Fetches(condition, args...)
}

func (slf *xrepo[T]) Execute(condition string, args ...any) error {
	return slf.repo.Execute(condition, args...)
}
