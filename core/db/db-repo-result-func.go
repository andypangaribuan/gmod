/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuRepoResult[T]) fetch() (*T, error) {
	return fm.PtrGetFirst(slf.entities), slf.err
}

func (slf *stuRepoResult[T]) fetches() ([]*T, error) {
	return slf.entities, slf.err
}

func (slf *stuRepoResult[T]) selectX() ([]map[string]any, error) {
	if slf.rows == nil {
		return []map[string]any{}, slf.err
	}

	return *slf.rows, slf.err
}

func (slf *stuRepoResult[T]) execute() (*int64, error) {
	return slf.id, slf.err
}
