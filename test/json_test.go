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
	"time"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/require"
)

func TestJson(t *testing.T) {
	stu := struct {
		StartedAt time.Time `json:"started_at"`
	}{
		StartedAt: gm.Util.Timenow(),
	}

	jsonData, err := gm.Json.Encode(stu)
	require.Nil(t, err)

	printLog(t, "%v\n", jsonData)
}
