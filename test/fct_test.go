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
	// v1, err := fct.New("a")
	// require.NotNil(t, err)
	// require.Nil(t, v1)

	// v2 := fct.UnsafeNew("10000")
	// require.NotNil(t, v2)

	v3 := fct.UnsafeNew("10000.12345678901234")
	require.NotNil(t, v3)
}