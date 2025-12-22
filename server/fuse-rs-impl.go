/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"fmt"
	"os"
	"time"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (slf *stuServer) FuseRS(restPort int, routes func(router RouterR), ws func(router RouterS)) {
	slf.fuseRS(restPort, routes, &ws)
}

func (slf *stuServer) fuseRS(restPort int, routes func(router RouterR), ws *func(router RouterS)) {
	if gm.Net.IsPortUsed(restPort) {
		fmt.Printf("fuse server [restful]  : port %v already in use\n", restPort)
		os.Exit(100)
	}

	fuseFiberApp = fiber.New(fiber.Config{
		JSONEncoder:           gm.Json.Marshal,
		JSONDecoder:           gm.Json.Unmarshal,
		DisableStartupMessage: true,
	})

	router := &stuFuseRRouter{
		fiberApp:        fuseFiberApp,
		withAutoRecover: false,
		printOnError:    true,
		noLogPaths:      make(map[string]any, 0),
	}

	routes(router)

	if router.withAutoRecover {
		fuseFiberApp.Use(recover.New())
	}

	isListenFailed := false
	go func() {
		tryCount := 0
		maxTry := 30
		time.Sleep(time.Millisecond * 100)

		for !isListenFailed {
			tryCount++
			if tryCount > maxTry {
				break
			}

			if gm.Net.IsPortUsed(restPort) {
				fmt.Printf("fuse server [restful]  : run at port %v\n", restPort)
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	slf.startWebsocket(fuseFiberApp, ws)
	go gracefulShutdown()

	err := fuseFiberApp.Listen(fmt.Sprintf(":%v", restPort))
	if err != nil {
		isListenFailed = true
		fmt.Printf("fuse server [restful]  : failed to listen on port %v\n", restPort)
		os.Exit(100)
	}
}
