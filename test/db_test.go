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
	"testing"

	_ "github.com/andypangaribuan/gmod"

	"github.com/andypangaribuan/gmod/fc"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/test/db/entity"
	"github.com/andypangaribuan/gmod/test/db/repo"
	"github.com/stretchr/testify/require"
)

func TestDbInsert(t *testing.T) {
	gm.Test.Start(t, func(t *testing.T) {
		timenow := gm.Util.Timenow()
		e := &entity.User{
			CreatedAt:  timenow,
			UpdatedAt:  timenow,
			Uid:        gm.Util.UID(),
			Name:       "andy",
			Address:    fm.Ptr("bintaro"),
			Height:     fm.Ptr(10),
			GoldAmount: fm.Ptr(fc.New(100.0)),
		}

		err := repo.User.Insert(e)
		require.Nil(t, err)
	})
}

func TestDbFetches(t *testing.T) {
	gm.Test.Start(t, func(t *testing.T) {
		entities, err := repo.User.Fetches("name=?", "andy")
		require.Nil(t, err)
		require.Greater(t, len(entities), 0)
		gm.Test.Printf(t, "length: %v\n", len(entities))
	})
}

func TestDbDelete(t *testing.T) {
	gm.Test.Start(t, func(t *testing.T) {
		
	})
}
