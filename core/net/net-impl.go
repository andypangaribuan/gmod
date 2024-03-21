/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package net

import (
	"net"
	"strconv"
	"time"
)

func (*srNet) IsPortUsed(port int, host ...string) bool {
	var (
		targetHost = "127.0.0.1"
		timeout    = time.Second * 3
	)

	if len(host) > 0 && host[0] != "" {
		targetHost = host[0]
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetHost, strconv.Itoa(port)), timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}
