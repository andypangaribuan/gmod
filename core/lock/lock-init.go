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
	"log"
	"strings"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	etcdclientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func xinit() {
	mainLockCallback = func() {
		dvalTxTryFor = getConfVal[*time.Duration]("txLockDvalTryFor")
		dvalTxTimeout = getConfVal[time.Duration]("txLockDvalTimeout")
		engineAddress := getConfVal[string]("txLockEngineAddress")
		txLockEngineName = getConfVal[string]("txLockEngineName")
		if txLockEngineName == "" {
			txLockEngineName = "redis"
		}

		if engineAddress == "" || engineAddress == "-" {
			return
		}

		if txLockEngineName == "redis" {
			txLockEngineAddress = engineAddress
			client := redis.NewClient(&redis.Options{
				Network: "tcp",
				Addr:    engineAddress,
			})

			txLockRedisClient = redislock.New(client)
		}

		if txLockEngineName == "etcd" {
			ls := strings.Split(engineAddress, ",")
			urls := make([]string, 0)
			for _, v := range ls {
				v = strings.TrimSpace(v)
				if v != "" {
					urls = append(urls, v)
				}
			}

			client, err := etcdclientv3.New(etcdclientv3.Config{
				Endpoints:   urls,
				DialTimeout: 3 * time.Second,
				Logger:      zap.NewNop(),
			})

			if err != nil {
				log.Printf("etcd lock create client is error, error: %+v\n", err)
				return
			}

			txLockEtcdClient = client
		}
	}
}
