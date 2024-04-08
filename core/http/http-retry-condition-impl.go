/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package http

import "strings"

func (slf *stuRetryCondition) Body() []byte {
	return slf.resp.Body()
}

func (slf *stuRetryCondition) IsError() bool {
	return slf.resp.IsError()
}

func (slf *stuRetryCondition) IsSuccess() bool {
	return slf.resp.IsSuccess()
}

func (slf *stuRetryCondition) Error() error {
	return slf.err
}

func (slf *stuRetryCondition) Code() int {
	return slf.resp.StatusCode()
}

func (slf *stuRetryCondition) IsTimeout() bool {
	var (
		keyA1     = "dial tcp"
		keyA2     = "i/o timeout"
		keyB1     = "net/http"
		keyB2     = "handshake timeout"
		isTimeout = false
	)

	if slf.err != nil {
		msg := strings.ToLower(slf.err.Error())
		isTimeout = (strings.Contains(msg, keyA1) && len(msg) > len(keyA2) && msg[len(msg)-len(keyA2):] == keyA2) ||
			(strings.Contains(msg, keyB1) && len(msg) > len(keyB2) && msg[len(msg)-len(keyB2):] == keyB2)
	}

	return isTimeout
}

func (slf *stuRetryCondition) IsConnectionReset() bool {
	var (
		key     = "connection reset by peer"
		isReset = false
	)

	if slf.err != nil {
		msg := strings.ToLower(slf.err.Error())
		isReset = len(msg) > len(key) && msg[len(msg)-len(key):] == key
	}

	return isReset
}
