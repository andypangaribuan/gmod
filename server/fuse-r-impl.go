/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package server

import (
	"fmt"
	"log"
	"time"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (*srServer) FuseR(restPort int, routes func(router RouterR)) {
	if gm.Net.IsPortUsed(restPort) {
		panic("rest-port already in use")
	}

	fuseFiberApp = fiber.New(fiber.Config{
		JSONEncoder: gm.Json.Marshal,
		JSONDecoder: gm.Json.UnMarshal,
		DisableStartupMessage: true,
	})

	router := &srFuseRouterR{
		fiberApp:        fuseFiberApp,
		withAutoRecover: false,
		printOnError:    true,
	}

	routes(router)

	isFuseRPrintOnError = router.printOnError
	if router.withAutoRecover {
		fuseFiberApp.Use(recover.New())
	}

	go func ()  {
		tryCount := 0
		maxTry := 5
		time.Sleep(time.Millisecond * 100)

		for {
			tryCount++
			if tryCount > maxTry {
				break
			}

			if gm.Net.IsPortUsed(restPort) {
				fmt.Printf("fuse server: restful run at port %v\n", restPort)
				break
			}

			time.Sleep(time.Second)
		}
	}()

	log.Fatal(fuseFiberApp.Listen(fmt.Sprintf(":%v", restPort)))
}
