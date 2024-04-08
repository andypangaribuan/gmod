/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/ice"
)

type stuTUid1Model struct {
	Id  int64  `db:"id"`
	Uid string `db:"uid"`
}

type stuTUid1Repo struct {
	repo db.Repo[stuTUid1Model]
}

func newTUid1Repo(dbi ice.DbInstance) *stuTUid1Repo {
	sr := stuTUid1Repo{}
	sr.repo = db.NewRepo[stuTUid1Model](dbi, "temp_uid1", db.RepoOpt().WithDeletedAtIsNull(false))

	sr.repo.SetInsertColumn(`uid`)

	return &sr
}

func (slf *stuTUid1Repo) getInsertColumn(e *stuTUid1Model) []any {
	return []any{
		e.Uid,
	}
}

func (slf *stuTUid1Repo) Fetches(condition string, args ...any) ([]*stuTUid1Model, error) {
	return slf.repo.Fetches(condition, args...)
}

func (slf *stuTUid1Repo) Insert(m *stuTUid1Model) error {
	return slf.repo.Insert(m.Uid)
}

func (slf *stuTUid1Repo) TxInsert(tx ice.DbTx, m *stuTUid1Model) error {
	return slf.repo.TxInsert(tx, m.Uid)
}

func (slf *stuTUid1Repo) TxBulkInsert(tx ice.DbTx, models []*stuTUid1Model) error {
	return slf.repo.TxBulkInsert(tx, models, slf.getInsertColumn, 1000)
}
