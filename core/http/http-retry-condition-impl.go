/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package http

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
