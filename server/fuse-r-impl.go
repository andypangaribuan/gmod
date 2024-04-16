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

func (*stuServer) FuseR(restPort int, routes func(router RouterR)) {
	if gm.Net.IsPortUsed(restPort) {
		fmt.Printf("fuse server [restful]: port %v already in use\n", restPort)
		os.Exit(100)
	}

	fuseFiberApp = fiber.New(fiber.Config{
		JSONEncoder:           gm.Json.Marshal,
		JSONDecoder:           gm.Json.UnMarshal,
		DisableStartupMessage: true,
	})

	router := &stuFuseRouterR{
		fiberApp:        fuseFiberApp,
		withAutoRecover: false,
		printOnError:    true,
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

		for {
			if isListenFailed {
				break
			}

			tryCount++
			if tryCount > maxTry {
				break
			}

			if gm.Net.IsPortUsed(restPort) {
				fmt.Printf("fuse server [restful]: run at port %v\n", restPort)
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	err := fuseFiberApp.Listen(fmt.Sprintf(":%v", restPort))
	if err != nil {
		isListenFailed = true
		fmt.Printf("fuse server [restful]: failed to listen on port %v\n", restPort)
		os.Exit(100)
	}
}
