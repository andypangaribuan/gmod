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
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/andypangaribuan/gmod/gm"
)

func (slf *stuTest) Start(t *testing.T, fn func(t *testing.T)) {
	var (
		startedTime = time.Now()
		oneSecond   = float64(1000)
		oneMinute   = float64(1000 * 60)
		oneHour     = float64(1000 * 60 * 60)
	)

	defer func() {
		durationMs := float64(time.Since(startedTime).Milliseconds())
		timenow := gm.Conv.Time.ToStrDT(gm.Util.Timenow())
		switch {
		case durationMs >= oneHour:
			slf.printf(t, fmt.Sprintf("\n\n%v duration: %.2f h\n", timenow, durationMs/oneHour), false)
		case durationMs >= 3*oneMinute:
			slf.printf(t, fmt.Sprintf("\n\n%v duration: %.2f m\n", timenow, durationMs/oneMinute), false)
		case durationMs >= oneSecond:
			slf.printf(t, fmt.Sprintf("\n\n%v duration: %.2f s\n", timenow, durationMs/oneSecond), false)
		default:
			slf.printf(t, fmt.Sprintf("\n\n%v duration: %v ms\n", timenow, int64(durationMs)), false)
		}
	}()

	fn(t)
}

func (slf *stuTest) Printf(t *testing.T, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	slf.printf(t, message, true)
}

func (*stuTest) printf(t *testing.T, message string, usingLog bool) {
	if t != nil {
		t.Logf(message)
	}

	if usingLog {
		log.Print(message)
	} else {
		fmt.Print(message)
	}
}
