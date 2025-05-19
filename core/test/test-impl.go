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
		timenow := gm.Conv.Time.ToStrDateTime(gm.Util.Timenow())
		switch {
		case durationMs >= oneHour:
			slf.printf(fmt.Sprintf("\n%v duration: %.2f h\n\n\n", timenow, durationMs/oneHour), false)
		case durationMs >= 3*oneMinute:
			slf.printf(fmt.Sprintf("\n%v duration: %.2f m\n\n\n", timenow, durationMs/oneMinute), false)
		case durationMs >= oneSecond:
			slf.printf(fmt.Sprintf("\n%v duration: %.2f s\n\n\n", timenow, durationMs/oneSecond), false)
		default:
			slf.printf(fmt.Sprintf("\n%v duration: %v ms\n\n\n", timenow, int64(durationMs)), false)
		}
	}()

	fn(t)
}

func (slf *stuTest) Printf(t *testing.T, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	slf.printf(message, true)
}

func (*stuTest) printf(message string, usingLog bool) {
	newLines := ""
	for len(message) > 0 && message[:1] == "\n" {
		newLines += "\n"
		message = message[1:]
	}

	if newLines != "" {
		fmt.Print(newLines)
	}

	if usingLog {
		log.Print(message)
	} else {
		fmt.Print(message)
	}
}
