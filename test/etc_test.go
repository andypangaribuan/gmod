/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package test

import (
	"testing"
	"time"

	_ "github.com/andypangaribuan/gmod"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/assert"
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

func TestRandom(t *testing.T) {
	var (
		loop          = 10
		sleepDuration = time.Microsecond * 10
		alphabet      = "abcde"
	)

	doTest := func(title, alphabet string, length int, withSleep bool) {
		temp := ""
		random := ""
		for i := 0; i < loop; i++ {
			random = gm.Util.GetRandom(length, alphabet)
			if i != 0 {
				temp += ", "
			}
			temp += random
			if withSleep {
				time.Sleep(sleepDuration)
			}
		}
		printLog(t, "%v%v\n", title, temp)
	}

	doTest("simp-1: ", alphabet, 3, false)
	doTest("simp-2: ", alphabet, 3, true)

	alphabet = gm.Util.GetAlphabet(true) + gm.Util.GetAlphabet() + gm.Util.GetNumeric()
	doTest("full-1: ", alphabet, 3, false)
	doTest("full-2: ", alphabet, 3, true)
}
