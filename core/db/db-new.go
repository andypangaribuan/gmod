/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

func RepoOpt() RepoOptBuilder {
	return new(stuRepoOptBuilder)
}

func FetchOpt() FetchOptBuilder {
	return new(stuFetchOptBuilder)
}

func Update() UpdateBuilder {
	return new(stuUpdateBuilder)
}
