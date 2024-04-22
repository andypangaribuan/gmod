/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package repo

import (
	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/core/db"
)

type xrepo[T any] struct {
	repo db.Repo[T]
}

func (slf *xrepo[T]) Insert(clog clog.Instance, e *T) error {
	return slf.repo.Insert(clog, e)
}

func (slf *xrepo[T]) Fetches(clog clog.Instance, condition string, args ...any) ([]*T, error) {
	return slf.repo.Fetches(clog, condition, args...)
}

func (slf *xrepo[T]) Execute(clog clog.Instance, condition string, args ...any) error {
	return slf.repo.Execute(clog, condition, args...)
}
