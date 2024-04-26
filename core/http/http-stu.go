/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package http

import (
	"io"
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/go-resty/resty/v2"
)

type stuHttp struct{}

type stuHttpBuilder struct {
	clog               clog.Instance
	url                string
	method             string
	timeout            *time.Duration
	insecureSkipVerify bool
	retryCondition     *func(resp ice.HttpResponse, count int) bool
	maxRetry           *int
	retryCount         int
	enableTrace        bool
	headers            *map[string]string
	pathParams         *map[string]string
	queryParams        *map[string]string
	formData           *map[string]string
	body               any
	fileReaders        []*stuFileReader
	files              *map[string]string
}

type stuFileReader struct {
	param    string
	fileName string
	reader   io.Reader
}

type stuRetryCondition struct {
	resp *resty.Response
	err  error
}
