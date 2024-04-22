/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

func (slf *stuRepoFuncOpt) SetSkipLevel(level int) RepoFuncOpt {
	slf.skipLevel = &level
	return slf
}

func (slf *stuRepoFuncOpt) SetChunkSize(size int) RepoFuncOpt {
	slf.chunkSize = &size
	return slf
}
