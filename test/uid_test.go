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
	_ "github.com/andypangaribuan/gmod"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/require"

	"testing"
	"time"
)

func TestGenerateL3Uid(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		printLog(t, "\nduration: %v ms\n", durationMs)
	}()

	uid := gm.Util.UID(0)
	printLog(t, "uid: %v\n", uid)
	printLog(t, "len: %v\n", len(uid))
	require.NotEmpty(t, uid)
}

func TestDecodeL3Uid(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		printLog(t, "\nduration: %v ms\n", durationMs)
	}()

	uid := "HZMMIoJnGjknwKY2dXKr"
	rawId, randId, err := gm.Util.DecodeXID(uid)
	require.Nil(t, err)
	require.Equal(t, uid[12:], randId)

	printLog(t, "raw : %v\n", rawId)
	printLog(t, "rand: %v\n", randId)

	require.Equal(t, "2024", rawId[0:4])
}

func TestGenerateUidB52(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		printLog(t, "\nduration: %v ms\n", durationMs)
	}()

	uid := gm.Util.UID52()
	require.NotEmpty(t, uid)

	printLog(t, "uid    : %v\n", uid)
	printLog(t, "len    : %v\n", len(uid))
	timeId, randId, err := gm.Util.DecodeUID52(uid)
	require.Nil(t, err)

	printLog(t, "time-id: %v\n", timeId)
	require.Empty(t, randId)

	printLog(t, "\n")

	uid = gm.Util.UID52(3)
	require.NotEmpty(t, uid)

	printLog(t, "uid    : %v\n", uid)
	printLog(t, "len    : %v\n", len(uid))
	timeId, randId, err = gm.Util.DecodeUID52(uid)
	require.Nil(t, err)

	printLog(t, "time-id: %v\n", timeId)
	printLog(t, "rand-id: %v\n", randId)
	require.NotEmpty(t, randId)
}

func TestGenerateUidB62(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		printLog(t, "\nduration: %v ms\n", durationMs)
	}()

	uid := gm.Util.UID62()
	require.NotEmpty(t, uid)

	printLog(t, "uid    : %v\n", uid)
	printLog(t, "len    : %v\n", len(uid))
	timeId, randId, err := gm.Util.DecodeUID62(uid)
	require.Nil(t, err)

	printLog(t, "time-id: %v\n", timeId)
	require.Empty(t, randId)

	printLog(t, "\n")

	uid = gm.Util.UID62(3)
	require.NotEmpty(t, uid)

	printLog(t, "uid    : %v\n", uid)
	printLog(t, "len    : %v\n", len(uid))
	timeId, randId, err = gm.Util.DecodeUID62(uid)
	require.Nil(t, err)

	printLog(t, "time-id: %v\n", timeId)
	printLog(t, "rand-id: %v\n", randId)
	require.NotEmpty(t, randId)
}
