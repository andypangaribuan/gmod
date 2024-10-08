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
	"net/http"
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/fasthttp/websocket"
)

func WsStream(logc clog.Instance, addr string, header *map[string]string, connectedMessage string, callback func(logc clog.Instance, message string), async ...bool) {
	if *fm.GetFirst(async, true) {
		go wsStream(logc, addr, header, connectedMessage, callback)
	} else {
		wsStream(logc, addr, header, connectedMessage, callback)
	}
}

func wsStream(logc clog.Instance, addr string, header *map[string]string, connectedMessage string, callback func(logc clog.Instance, message string)) {
	var (
		err       error
		conn      *websocket.Conn
		reqHeader = make(http.Header)
		dialer    = websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 5 * time.Second,
		}
	)

	if header != nil {
		for k, v := range *header {
			reqHeader.Add(k, v)
		}
	}

	doConnection := func() {
		for {
			conn, _, err = dialer.Dial(addr, reqHeader)
			if err == nil {
				fmt.Println(connectedMessage)
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}

	doConnection()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			_ = conn.Close()
			doConnection()
			continue
		}

		message := string(msg)
		if msgType == websocket.TextMessage && message != "" {
			callback(logc, message)
		}
	}
}
