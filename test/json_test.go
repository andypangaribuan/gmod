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
	type stuX struct {
		StartedAt time.Time `json:"started_at"`
	}

	var (
		// jsonStr = `{"started_at": "2024-04-14 10:10:00.000001"}` // error
		// jsonStr = `{"started_at": "2024-04-14 10:10:00.000001 +07:00"}` // ok
		jsonStr = `{"started_at": "2024-04-14 10:10:00 +09:00"}` // ok
		x       *stuX
	)

	err := gm.Json.Decode(jsonStr, &x)
	require.Nil(t, err)

	x1 := stuX{StartedAt: x.StartedAt}
	jsonData, err := gm.Json.Encode(x1)
	require.Nil(t, err)
	printLog(t, "%v\n", jsonData)

	x2 := stuX{StartedAt: gm.Util.Timenow()}
	jsonData, err = gm.Json.Encode(x2)
	require.Nil(t, err)
	printLog(t, "%v\n", jsonData)
}
