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

	"github.com/gofiber/contrib/websocket"
)

func (slf *stuFuseSContext) ReadMessage() (message string, err error) {
	msgType, msg, err := slf.conn.ReadMessage()
	if err != nil {
		return "", err
	}

	if msgType == websocket.TextMessage {
		return string(msg), nil
	}

	return "", nil
}

func (slf *stuFuseSContext) WriteMessage(message string) error {
	return slf.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (slf *stuFuseSContext) GetLocal(key string) string {
	val := slf.conn.Locals(key)
	if val == nil {
		return ""
	}
	switch v := val.(type) {
	case string:
		return v
	case *string:
		if v == nil {
			return ""
		}
		return *v

	default:
		return fmt.Sprintf("%v", val)
	}
}

func (slf *stuFuseSContext) GetParam(key string) string {
	return slf.conn.Params(key)
}

func (slf *stuFuseSContext) GetQuery(key string) string {
	return slf.conn.Query(key)
}

func (slf *stuFuseSContext) Close() {
	_ = slf.conn.WriteMessage(websocket.CloseMessage, []byte{})
	_ = slf.conn.Close()
}
