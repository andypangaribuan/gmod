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
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/go-resty/resty/v2"
)

func (slf *stuHttpBuilder) SetTimeout(duration time.Duration) ice.HttpBuilder {
	slf.timeout = &duration
	return slf
}

func (slf *stuHttpBuilder) InsecureSkipVerify(skip ...bool) ice.HttpBuilder {
	slf.insecureSkipVerify = *fm.GetFirst(skip, false)
	return slf
}

func (slf *stuHttpBuilder) SetRetryCondition(condition func(resp ice.HttpResponse, count int) bool) ice.HttpBuilder {
	slf.retryCondition = &condition
	return slf
}

func (slf *stuHttpBuilder) SetMaxRetry(max int) ice.HttpBuilder {
	max = fm.Ternary(max < 0, -1, max)
	slf.maxRetry = &max
	return slf
}

func (slf *stuHttpBuilder) EnableTrace(enable ...bool) ice.HttpBuilder {
	slf.enableTrace = *fm.GetFirst(enable, true)
	return slf
}

func (slf *stuHttpBuilder) SetHeader(args map[string]string) ice.HttpBuilder {
	header := slf.setInternalHeader(args)
	slf.headers = &header
	return slf
}

func (slf *stuHttpBuilder) SetJsonHeader(opt ...any) ice.HttpBuilder {
	header := slf.getJsonHeader(opt...)
	slf.headers = &header
	return slf
}

func (slf *stuHttpBuilder) SetQueryParam(args map[string]string) ice.HttpBuilder {
	slf.queryParams = &args
	return slf
}

func (slf *stuHttpBuilder) SetPathParam(args map[string]string) ice.HttpBuilder {
	slf.pathParams = &args
	return slf
}

func (slf *stuHttpBuilder) SetFormData(args map[string]string) ice.HttpBuilder {
	slf.formData = &args
	return slf
}

func (slf *stuHttpBuilder) SetBody(value any) ice.HttpBuilder {
	slf.body = value
	return slf
}

func (slf *stuHttpBuilder) AddFileReader(param, fileName string, reader io.Reader) ice.HttpBuilder {
	slf.fileReaders = append(slf.fileReaders, &stuFileReader{
		param:    param,
		fileName: fileName,
		reader:   reader,
	})
	return slf
}

func (slf *stuHttpBuilder) SetFiles(files map[string]string) ice.HttpBuilder {
	slf.files = &files
	return slf
}

func (slf *stuHttpBuilder) Call() (data []byte, code int, err error) {
	var (
		startedAt = gm.Util.Timenow()
		client    = resty.New()
		resp      *resty.Response

		joinedFiles = make(map[string][]string, 0)
		severity    = "unknown"
		reqHeader   *string
		reqParam    *string
		reqQuery    *string
		reqForm     *string
		reqFiles    *string
		reqBody     *string
		resData     *string
		errMessage  *string
		stackTrace  *string
	)

	if slf.timeout != nil {
		client.SetTimeout(*slf.timeout)
	}

	if slf.insecureSkipVerify {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: slf.insecureSkipVerify})
	}

	if slf.retryCondition != nil {
		client.AddRetryCondition(func(resp *resty.Response, err error) bool {
			return slf.callRetryCondition(resp, err)
		})
	}

	req := client.R()

	if slf.enableTrace {
		req.EnableTrace()
	}

	if slf.headers != nil && len(*slf.headers) > 0 {
		req.SetHeaders(*slf.headers)

		if slf.clog != nil {
			jons, err := gm.Json.Encode(slf.headers)
			if err == nil {
				reqHeader = &jons
			}
		}
	}

	if slf.pathParams != nil && len(*slf.pathParams) > 0 {
		req.SetPathParams(*slf.pathParams)

		if slf.clog != nil {
			jons, err := gm.Json.Encode(slf.pathParams)
			if err == nil {
				reqParam = &jons
			}
		}
	}

	if slf.queryParams != nil && len(*slf.queryParams) > 0 {
		req.SetQueryParams(*slf.queryParams)

		if slf.clog != nil {
			jons, err := gm.Json.Encode(slf.queryParams)
			if err == nil {
				reqQuery = &jons
			}
		}
	}

	if slf.formData != nil && len(*slf.formData) > 0 {
		req.SetFormData(*slf.formData)

		if slf.clog != nil {
			jons, err := gm.Json.Encode(slf.formData)
			if err == nil {
				reqForm = &jons
			}
		}
	}

	if slf.body != nil {
		req.SetBody(slf.body)

		if slf.clog != nil {
			jons, err := gm.Json.Encode(slf.body)
			if err == nil {
				reqBody = &jons
			}
		}
	}

	if len(slf.fileReaders) > 0 {
		for _, f := range slf.fileReaders {
			req.SetFileReader(f.param, f.fileName, f.reader)

			ls, ok := joinedFiles[f.param]
			if ok {
				joinedFiles[f.param] = append(ls, f.fileName)
			} else {
				joinedFiles[f.param] = []string{f.fileName}
			}
		}
	}

	if slf.files != nil && len(*slf.files) > 0 {
		req.SetFiles(*slf.files)
		for k, v := range *slf.files {
			ls, ok := joinedFiles[k]
			if ok {
				joinedFiles[k] = append(ls, v)
			} else {
				joinedFiles[k] = []string{v}
			}
		}
	}

	if slf.clog != nil && len(joinedFiles) > 0 {
		jons, err := gm.Json.Encode(joinedFiles)
		if err == nil {
			reqFiles = &jons
		}
	}

	for {
		switch slf.method {
		case "get":
			resp, err = req.Get(slf.url)

		case "post":
			resp, err = req.Post(slf.url)

		case "put":
			resp, err = req.Put(slf.url)

		case "patch":
			resp, err = req.Patch(slf.url)

		case "delete":
			resp, err = req.Delete(slf.url)
		}

		if resp != nil {
			code := strconv.Itoa(resp.StatusCode())
			if code[:1] != "2" {
				if !slf.callRetryCondition(resp, err) {
					break
				} else {
					continue
				}
			}
		}

		if err == nil {
			break
		}

		if !slf.callRetryCondition(resp, err) {
			break
		}
	}

	var (
		body       = resp.Body()
		statusCode = resp.StatusCode()
	)

	if slf.enableTrace {
		ti := resp.Request.TraceInfo()
		msg := `
!! Response Info
Url         : %v
Method      : %v
Status Code : %v
Status      : %v
Proto       : %v
Duration    : %v
Received At : %v
Error       : %+v %v
Response    : %v



!! Trace Info
DNSLookup     : %v
ConnTime      : %v
TCPConnTime   : %v
TLSHandshake  : %v
ServerTime    : %v
ResponseTime  : %v
TotalTime     : %v
IsConnReused  : %v
IsConnWasIdle : %v
ConnIdleTime  : %v
RequestAttempt: %v
RemoteAddr    : %v
`

		var (
			response = fmt.Sprintf("%v", resp)
			ext      = ""
			out      bytes.Buffer
		)

		if err := json.Indent(&out, body, "", "  "); err == nil {
			response = "\n" + out.String()
		}

		if val := getJsonIndent(slf.headers); val != nil {
			ext += fmt.Sprintf("\nHeaders     :\n%v", *val)
		}

		if val := getJsonIndent(slf.pathParams); val != nil {
			ext += fmt.Sprintf("\nPath Params :\n%v", *val)
		}

		if val := getJsonIndent(slf.queryParams); val != nil {
			ext += fmt.Sprintf("\nQuery Params:\n%v", *val)
		}

		if val := getJsonIndent(slf.formData); val != nil {
			ext += fmt.Sprintf("\nForm Data   :\n%v", *val)
		}

		if val := getJsonIndent(slf.files); val != nil {
			ext += fmt.Sprintf("\nFiles       :\n%v", *val)
		}

		if slf.body != nil {
			data, err := gm.Json.Marshal(slf.body)
			if err == nil {
				var out bytes.Buffer
				err = json.Indent(&out, data, "", "  ")
				if err == nil {
					ext += fmt.Sprintf("\nBody        :\n%v", out.String())
				}
			}
		}

		msg = fmt.Sprintf(msg,
			slf.url,
			slf.method,
			resp.StatusCode(),
			resp.Status(),
			resp.Proto(),
			resp.Time(),
			gm.Conv.Time.ToStrFull(resp.ReceivedAt()),
			err, ext,
			response,

			ti.DNSLookup,
			ti.ConnTime,
			ti.TCPConnTime,
			ti.TLSHandshake,
			ti.ServerTime,
			ti.ResponseTime,
			ti.TotalTime,
			ti.IsConnReused,
			ti.IsConnWasIdle,
			ti.ConnIdleTime,
			ti.RequestAttempt,
			ti.RemoteAddr.String(),
		)

		fmt.Println(msg)
	}

	if slf.clog != nil {
		if err != nil {
			errMessage = fm.Ptr(err.Error())
			stackTrace = fm.Ptr(fmt.Sprintf("%+v", err))
		}

		switch {
		case statusCode >= 100 && statusCode <= 199:
			severity = "server"

		case statusCode >= 200 && statusCode <= 299:
			severity = "success"

		case statusCode >= 300 && statusCode <= 399:
			severity = "server"

		case statusCode >= 400 && statusCode <= 499:
			severity = "warning"

		case statusCode >= 500 && statusCode <= 599:
			severity = "error"
		}

		if len(body) > 0 {
			resData = fm.Ptr(string(body))
		}
	}

	if slf.clog != nil {
		mol := &clog.HttpCallV1{
			Url:        slf.url,
			Severity:   severity,
			ReqHeader:  reqHeader,
			ReqParam:   reqParam,
			ReqQuery:   reqQuery,
			ReqForm:    reqForm,
			ReqFiles:   reqFiles,
			ReqBody:    reqBody,
			ResData:    resData,
			ResCode:    statusCode,
			ErrMessage: errMessage,
			StackTrace: stackTrace,
			StartedAt:  startedAt,
			FinishedAt: gm.Util.Timenow(),
		}

		slf.clog.HttpCallV1(mol)
	}

	return body, statusCode, err
}
