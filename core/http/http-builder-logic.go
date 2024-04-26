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

func (slf *stuHttpBuilder) isInternalUrl() bool {
	var (
		urlLength = len(slf.url)
		length    int
	)

	for _, base := range internalBaseUrls {
		length = len(base)
		if urlLength >= len(base) {
			if slf.url[:length] == base {
				return true
			}
		}
	}
	return false
}

func (slf *stuHttpBuilder) getJsonHeader(opt ...any) map[string]string {
	var (
		add    = make(map[string]string, 0)
		header = map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		}
	)

	if len(opt) > 0 {
		for _, o := range opt {
			switch v := o.(type) {
			case map[string]string:
				add = v
			case *map[string]string:
				if v != nil {
					add = *v
				}
			}
		}
	}

	header = slf.setInternalHeader(header, opt...)

	for k, v := range add {
		header[k] = v
	}

	return header
}

func (slf *stuHttpBuilder) setInternalHeader(header map[string]string, opt ...any) map[string]string {
	var reqVer *string

	if len(opt) > 0 {
		for _, o := range opt {
			switch v := o.(type) {
			case string:
				reqVer = &v
			case *string:
				reqVer = v
			}
		}
	}

	if slf.isInternalUrl() {
		header["X-From-SvcName"] = svcName
		header["X-From-SvcVersion"] = svcVersion

		if reqVer != nil {
			header["X-Version"] = *reqVer
		}

		if slf.clog != nil {
			uid, userId, partnerId := clogGetId(slf.clog)
			if uid != "" {
				header["X-Clog-Uid"] = uid
			}

			if userId != nil {
				header["X-Clog-UserId"] = *userId
			}

			if partnerId != nil {
				header["X-Clog-PartnerId"] = *partnerId
			}
		}
	}

	return header
}

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
