/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuRepoOptBuilder) WithDeletedAtIsNull(val ...bool) RepoOptBuilder {
	slf.withDeletedAtIsNull = fm.GetFirst(val, true)
	return slf
}

func (slf *stuRepoOptBuilder) RWFetchWhenNull(val ...bool) RepoOptBuilder {
	slf.rwFetchWhenNull = fm.GetFirst(val, true)
	return slf
}
