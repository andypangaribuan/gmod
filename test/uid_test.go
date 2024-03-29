/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package test

import (
	"fmt"

	_ "github.com/andypangaribuan/gmod"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/assert"

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
	assert.NotEmpty(t, uid)
}

func TestDecodeL3Uid(t *testing.T) {
	startedTime := time.Now()
	defer func() {
		durationMs := time.Since(startedTime).Milliseconds()
		printLog(t, "\nduration: %v ms\n", durationMs)
	}()

	uid := "HZMMGyAIUOcnwk6TuB5o"
	rawId, randId, err := gm.Util.DecodeUID(uid)
	assert.Nil(t, err)
	assert.Equal(t, uid[12:], randId)

	var (
		year         = rawId[0:4]
		month        = rawId[4:6]
		day          = rawId[6:8]
		hour         = rawId[8:10]
		minute       = rawId[10:12]
		second       = rawId[12:14]
		milliseconds = rawId[14:20]
	)

	tn := fmt.Sprintf("%v-%v-%v %v:%v:%v.%v", year, month, day, hour, minute, second, milliseconds)
	printLog(t, "raw : %v\n", tn)
	printLog(t, "rand: %v\n", randId)
	
	assert.Equal(t, "2024-03-28 10:00:43.639351", tn)
}
