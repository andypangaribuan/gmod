/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/clog"

func (slf *stuVDB[T]) Fetch(clog clog.Instance, args ...any) (*T, error) {
	return slf.override(clog, slf.fetches(true, nil, condition, args)).fetch()
}
