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

	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/require"
)

func TestTxLock(t *testing.T) {
	gm.Test.Start(t, testTxLock)
}

func testTxLock(t *testing.T) {
	// startedAt := time.Now().Add(time.Second * 5)

	lock, err := gm.Lock.Tx("1")
	require.Nil(t, err)
	require.NotNil(t, lock)
}

// func txLockIns()