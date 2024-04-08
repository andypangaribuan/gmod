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

	uid := gm.Util.UID()
	printLog(t, "uid: %v\n", uid)
	require.NotEmpty(t, uid)
}

func TestDecodeL3Uid(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		printLog(t, "\nduration: %v ms\n", durationMs)
	}()

	uid := "HZMMIoJnGjknwKY2dXKr"
	rawId, randId, err := gm.Util.DecodeUID(uid)
	require.Nil(t, err)
	require.Equal(t, uid[12:], randId)

	printLog(t, "raw : %v\n", rawId)
	printLog(t, "rand: %v\n", randId)

	require.Equal(t, "2024", rawId[0:4])
}
