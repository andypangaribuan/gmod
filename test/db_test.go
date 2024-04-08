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
	"fmt"
	"testing"
	"time"

	_ "github.com/andypangaribuan/gmod"

	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/mol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	dbi       ice.DbInstance
	repoTUid1 *stuTUid1Repo
)

func initDb() {
	conn := mol.DbConnection{
		AppName:  gm.Util.Env.GetString("APP_NAME"),
		Host:     gm.Util.Env.GetString("DBW_HOST"),
		Port:     gm.Util.Env.GetInt("DBW_PORT"),
		Name:     gm.Util.Env.GetString("DBW_NAME"),
		Username: gm.Util.Env.GetString("DBW_USER"),
		Password: gm.Util.Env.GetString("DBW_PASS"),
	}

	dbi = gm.Db.PostgresRW(conn, conn)
	repoTUid1 = newTUid1Repo(dbi)
}

func TestDbFetches(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		fmt.Printf("\nduration: %v ms\n", durationMs)
	}()

	initDb()

	// models, err := repoTUid1.Fetches("", db.FetchOpt{EndQuery: fm.Ptr("ORDER BY uid ASC"), WithDeletedAtIsNull: fm.Ptr(false)})
	// models, err := repoTUid1.Fetches("", db.FetchOpt{EndQuery: fm.Ptr("ORDER BY uid ASC"), WithDeletedAtIsNull: fm.Ptr(false)})
	models, err := repoTUid1.Fetches("", db.FetchOpt().WithDeletedAtIsNull(false).EndQuery("ORDER BY uid ASC"))
	assert.Nil(t, err)

	l3, _, _ := uidL3()
	// require.Equal(t, len(models), 0)
	require.Equal(t, len(models), len(l3))

	for i := range l3 {
		require.Equal(t, l3[i], models[i].Uid)
		require.Equal(t, int64(i+1), models[i].Id)
	}
}

func TestDbInsert(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		fmt.Printf("\nduration: %v ms\n", durationMs)
	}()

	initDb()

	model := &stuTUid1Model{
		Uid: "AA9",
	}

	err := repoTUid1.Insert(model)
	assert.Nil(t, err)
}

func TestUidL3Insert(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		fmt.Printf("\nduration: %v ms\n", durationMs)
	}()

	initDb()

	l3, _, _ := uidL3()

	tx, err := dbi.NewTransaction()
	defer tx.Rollback()
	assert.Nil(t, err)

	models := make([]*stuTUid1Model, len(l3))

	for i, uid := range l3 {
		models[i] = &stuTUid1Model{
			Uid: uid,
		}
	}

	err = repoTUid1.TxBulkInsert(tx, models)
	assert.Nil(t, err)

	err = tx.Commit()
	assert.Nil(t, err)
}
