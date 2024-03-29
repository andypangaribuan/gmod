/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package test

import (
	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/ice"
)

type srTUid1Model struct {
	Id  int64  `db:"id"`
	Uid string `db:"uid"`
}

type srTUid1Repo struct {
	repo db.Repo[srTUid1Model]
}

func newTUid1Repo(dbi ice.DbInstance) *srTUid1Repo {
	sr := srTUid1Repo{}
	sr.repo = db.NewRepo[srTUid1Model](dbi, "temp_uid1", db.RepoOpt{WithDeletedAtIsNull: fm.Ptr(false)})

	sr.repo.SetInsertColumn(`uid`)

	return &sr
}

func (slf *srTUid1Repo) getInsertColumn(e *srTUid1Model) []interface{} {
	return []interface{}{
		e.Uid,
	}
}

func (slf *srTUid1Repo) Fetches(condition string, args ...interface{}) ([]*srTUid1Model, error) {
	return slf.repo.Fetches(condition, args...)
}

func (slf *srTUid1Repo) Insert(m *srTUid1Model) error {
	return slf.repo.Insert(m.Uid)
}

func (slf *srTUid1Repo) TxInsert(tx ice.DbTx, m *srTUid1Model) error {
	return slf.repo.TxInsert(tx, m.Uid)
}

func (slf *srTUid1Repo) TxBulkInsert(tx ice.DbTx, models []*srTUid1Model) error {
	return slf.repo.TxBulkInsert(tx, models, slf.getInsertColumn, 1000)
}
