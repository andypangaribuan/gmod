/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func beforeGracefulShutdown(then ...func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigChan

	for _, fn := range then {
		go fn()
	}
}

func gracefulShutdown(waitDuration ...time.Duration) {
	mx.Lock()
	defer mx.Unlock()

	if isGracefulShutdownImpl {
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		<-sigChan

		wait := time.Second * 10
		if len(waitDuration) > 0 {
			wait = waitDuration[0]
		}

		time.Sleep(wait)
		os.Exit(1)
	}()
}
