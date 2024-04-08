/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package http

import "github.com/go-resty/resty/v2"

func (slf *stuHttpBuilder) callRetryCondition(resp *resty.Response, err error) bool {
	if slf.retryCondition == nil {
		return false
	}

	if slf.maxRetry != nil {
		if *slf.maxRetry == 0 {
			return false
		}

		if *slf.maxRetry == slf.retryCount {
			return false
		}
	}

	stu := &stuRetryCondition{
		resp: resp,
		err:  err,
	}

	retry := (*slf.retryCondition)(stu, slf.retryCount)
	if retry {
		slf.retryCount++
	}

	return retry
}
