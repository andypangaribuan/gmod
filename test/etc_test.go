/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package test

import (
	_ "github.com/andypangaribuan/gmod"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestTimeFormat(t *testing.T) {
	// gm.Conf.SetTimeZone("Asia/Tokyo")
	gm.Conf.SetTimeZone("Asia/Singapore")
	// gm.Conf.SetTimeZone("Asia/Jakarta")

	timenow := gm.Util.Timenow()
	tmFull := gm.Conv.Time.ToStrFull(timenow)
	tmDT := gm.Conv.Time.ToStrDT(timenow)

	t.Logf("full: %v\n", tmFull)
	t.Logf("dt  : %v\n", tmDT)

	assert.NotEmpty(t, tmFull)
	assert.NotEmpty(t, tmDT)
}
