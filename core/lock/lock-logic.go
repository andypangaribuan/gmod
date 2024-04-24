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
	"context"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

func getLock(key string, timeout time.Duration) {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer client.Close()

	var (
		locker = redislock.New(client)
		ctx    = context.Background()
	)

	locker.Obtain(ctx, key, timeout, nil)
}
