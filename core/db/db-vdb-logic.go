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

func (slf *stuVDB[T]) override(clog clog.Instance, res *stuRepoResult[T]) *stuRepoResult[T] {
	pushClogReport(clog, res.report, res.err, 4)
	return res
}
