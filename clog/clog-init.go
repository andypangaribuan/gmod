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
	mainCLogCallback = func() {
		if val := getConfValue("clogAddress"); val != "" {
			c, err := createClient(val, sclog.NewCLogServiceClient)
			if err != nil {
				go connect(val)
			} else {
				client = c
				fmt.Println("connected to central log")
			}
		}

		if val := getConfValue("svcName"); val != "" {
			svcName = val
		}

		if val := getConfValue("svcVersion"); val != "" {
			svcVersion = val
		}
	}

	clogNew = func() Instance {
		if client == nil {
			return nil
		}

		return &stuInstance{
			uid: mainCLogUtilUid(),
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
