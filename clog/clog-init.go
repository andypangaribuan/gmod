/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import (
	"fmt"
	"time"

	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

func xinit() {
	go pusher()

	mainCLogCallback = func() {
		if val := getConfVal[string]("clogAddress"); val != "" {
			c, err := createClient(val, sclog.NewCLogServiceClient)
			if err != nil {
				go connect(val)
			} else {
				client = c
				fmt.Println("connected to central log")
			}
		}

		if val := getConfVal[string]("svcName"); val != "" {
			svcName = val
		}

		if val := getConfVal[string]("svcVersion"); val != "" {
			svcVersion = val
		}

		if val := getConfVal[*time.Duration]("clogRetryMaxDuration"); val != nil {
			retryMaxDuration = *val
		}

		if val := getConfVal[int]("clogMaxConcurrentPusher"); val > 0 {
			maxConcurrentPusher = val
		}
	}

	clogSetUserId = func(clog Instance, id string) {
		if clog != nil {
			c, ok := clog.(*stuInstance)
			if ok && c != nil {
				c.userId = &id
			}
		}
	}

	clogSetPartnerId = func(clog Instance, id string) {
		if clog != nil {
			c, ok := clog.(*stuInstance)
			if ok && c != nil {
				c.partnerId = &id
			}
		}
	}

	clogGetId = func(clog Instance) (string, *string, *string) {
		if clog != nil {
			v, ok := clog.(*stuInstance)
			if ok {
				return v.uid, v.userId, v.partnerId
			}
		}

		return "", nil, nil
	}
}

func connect(address string) {
	for {
		time.Sleep(time.Millisecond * 300)
		c, err := createClient(address, sclog.NewCLogServiceClient)
		if err == nil {
			client = c
			fmt.Println("connected to central log")
			break
		}
	}
}

func pusher() {
	var xcp any

	for {
		size := queueSize()
		if client == nil || size == 0 {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		concurrentPusher := maxConcurrentPusher
		if concurrentPusher == 0 {
			concurrentPusher = 100
		}

		maxJob := concurrentPusher * 10
		if xcp == nil {
			xcp = mrf1[any]("mrf-util-x-concurrent-process", concurrentPusher, maxJob)
		}

		startAt := time.Now()
		msg := "\n\n\n\n\n"
		msg += fmt.Sprintf("before clog queue size: %v\n", size)

		logs := queueList()
		chunk := chunkSlice(logs, maxJob)

		msg += fmt.Sprintf("after clog queue size : %v\n", queueSize())
		msg += fmt.Sprintf("copy clog logs size   : %v\n", len(logs))
		if maxConcurrentPusher > 0 {
			fmt.Printf("%v\n\n\n\n\n", msg)
		}

		for _, logs := range chunk {
			mainReflection("mrf-util-x-concurrent-process-run", xcp, len(logs), func(index int) {
				sq := logs[index]
				switch sq.logType {
				case "NoteV1":
					req := sq.req.(*sclog.RequestNoteV1)
					doGrpcCall(client.NoteV1, req)

				case "DbqV1":
					req := sq.req.(*sclog.RequestDbqV1)
					doGrpcCall(client.DbqV1, req)

				case "HttpCallV1":
					req := sq.req.(*sclog.RequestHttpCallV1)
					doGrpcCall(client.HttpCallV1, req)

				case "ServicePieceV1":
					req := sq.req.(*sclog.RequestServicePieceV1)
					doGrpcCall(client.ServicePieceV1, req)

				case "ServiceV1":
					req := sq.req.(*sclog.RequestServiceV1)
					doGrpcCall(client.ServiceV1, req)

				case "GrpcV1":
					req := sq.req.(*sclog.RequestGrpcV1)
					doGrpcCall(client.GrpcV1, req)
				}
			})
		}

		msg += fmt.Sprintf("max concurrent pusher : %v\n", maxConcurrentPusher)
		msg += fmt.Sprintf("pusher duration       : %v\n", time.Since(startAt))
		if maxConcurrentPusher > 0 {
			fmt.Printf("%v\n\n\n\n\n", msg)
		}
	}
}
