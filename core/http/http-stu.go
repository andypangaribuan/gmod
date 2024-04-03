/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package http

import (
	"io"
	"time"

	"github.com/andypangaribuan/gmod/ice"
	"github.com/go-resty/resty/v2"
)

type stuHttp struct{}

type stuHttpBuilder struct {
	url                string
	method             string
	timeout            *time.Duration
	insecureSkipVerify bool
	retryCondition     *func(resp ice.HttpResponse) bool
	enableTrace        bool
	headers            *map[string]string
	pathParams         *map[string]string
	queryParams        *map[string]string
	formData           *map[string]string
	body               interface{}
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
