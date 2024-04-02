/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package http

import (
	"io"

	"github.com/go-resty/resty/v2"
)

type stuHttp struct{}

type stuHttpBuilder struct {
	url         string
	method      string
	client      *resty.Client
	enableTrace bool
	headers     *map[string]string
	pathParams  *map[string]string
	queryParams *map[string]string
	formData    *map[string]string
	body        interface{}
	fileReaders []*stuFileReader
	files       *map[string]string
}

type stuFileReader struct {
	param    string
	fileName string
	reader   io.Reader
}
