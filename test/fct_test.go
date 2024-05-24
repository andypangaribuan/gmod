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

	"github.com/andypangaribuan/gmod/fct"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/require"
)

func TestFCT(t *testing.T) {
	gm.Test.Start(t, testFCT)
}

func testFCT(t *testing.T) {
	v1, err := fct.New("a")
	require.NotNil(t, err)
	require.Nil(t, v1)

	v2 := fct.UnsafeNew("10000")
	require.NotNil(t, v2)
	require.Equal(t, "10000", v2.UnsafeToString())

	var (
		val    = "10000"
		points = "12345678901234"
		amount = val + "." + points
	)

	v3 := fct.UnsafeNew(amount)
	require.NotNil(t, v3)
	require.Equal(t, amount, v3.UnsafeToString())

	type model struct {
		F1 fct.FCT  `json:"f1"`
		F2 *fct.FCT `json:"f2"`
		F3 *fct.FCT `json:"f3"`
	}

	m1 := &model{
		F1: v3,
		F2: &v3,
		F3: nil,
	}

	jons, err := gm.Json.Encode(m1)
	require.Nil(t, err)
	require.Contains(t, jons, amount)

	var m2 *model
	err = gm.Json.Decode(jons, &m2)
	require.Nil(t, err)
	require.NotNil(t, m2)

	if fct.UnsafeCompare(v3, "!=", m2.F1) {
		require.FailNow(t, "not equal")
	}
}
