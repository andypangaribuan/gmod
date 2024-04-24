/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import (
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

func xinit() {
	mainLockCallback = func() {
		if val := getConfVal[string]("txLockEngineAddress"); val != "" {
			txLockEngineAddress = val
			client := redis.NewClient(&redis.Options{
				Network: "tcp",
				Addr:    val,
			})
			// defer client.Close()

			txLockEngine = redislock.New(client)
		}

		dvalTxTryFor = getConfVal[*time.Duration]("txLockDvalTryFor")
		dvalTxTimeout = getConfVal[time.Duration]("txLockDvalTimeout")
	}
}
