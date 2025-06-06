/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import (
	"io"
	"time"

	"github.com/andypangaribuan/gmod/clog"
)

type Http interface {
	Get(clog clog.Instance, url string) HttpBuilder
	Post(clog clog.Instance, url string) HttpBuilder
	Put(clog clog.Instance, url string) HttpBuilder
	Patch(clog clog.Instance, url string) HttpBuilder
	Delete(clog clog.Instance, url string) HttpBuilder
}

type HttpBuilder interface {
	SetTimeout(duration time.Duration) HttpBuilder

	// disable security check (https)
	InsecureSkipVerify(skip ...bool) HttpBuilder

	// Example:
	//	SetRetryCondition(func(resp ice.HttpResponse, count int) bool {
	//		return resp.IsConnectionReset() ||
	//			resp.IsTimeout() ||
	//			resp.Code() == http.StatusTooManyRequests
	//	})
	SetRetryCondition(condition func(resp HttpResponse, count int) bool) HttpBuilder
	SetMaxRetry(max int) HttpBuilder

	EnableTrace(enable ...bool) HttpBuilder
	SetHeader(args map[string]string) HttpBuilder

	// Example:
	//  SetJsonHeader("1.0", map[string]string{
	// 		"Authorization": "Bearer xyz",
	//  })
	// string "1.0" it will convert to "X-Version",
	// then map[string]string will added to header value
	SetJsonHeader(opt ...any) HttpBuilder

	// Examples:
	//
	//	Get(".../users/{userId}/{subAccountId}/details").
	//	SetPathParams(map[string]any{
	//		"userId": "sample@sample.com",
	//		"subAccountId": "100002",
	//	})
	SetPathParam(args map[string]string) HttpBuilder

	SetQueryParam(args map[string]string) HttpBuilder
	SetFormData(args map[string]string) HttpBuilder

	// Examples:
	//
	//	SetBody(User{
	//		Username: "jeeva@myjeeva.com",
	//		Password: "welcome2resty",
	//	})
	//
	//	SetBody(map[string]any{
	//		"username": "jeeva@myjeeva.com",
	//		"password": "welcome2resty",
	//		"address": &Address{
	//			City: "My City",
	//			ZipCode: 00000,
	//		},
	//	})
	//
	//	SetBody(`{
	//		"username": "jeeva@getrightcare.com",
	//		"password": "admin"
	//	}`)
	//
	//	SetBody([]byte("This is my raw request, sent as-is"))
	//
	SetBody(value any) HttpBuilder

	// Examples:
	// profileImgBytes, _ := os.ReadFile("/andy/test-img.png")
	// notesBytes, _ := os.ReadFile("/andy/text-file.txt")
	//
	//	AddFileReader("profile_img", "my-profile-img.png", bytes.NewReader(profileImgBytes)).
	//	AddFileReader("notes", "user-notes.txt", bytes.NewReader(notesBytes))
	AddFileReader(param, fileName string, reader io.Reader) HttpBuilder

	// Examples:
	//
	//	SetFiles(map[string]string{
	//		"file1": "/andy/invoice.pdf",
	//		"file2": "/andy/detail.pdf",
	//		"file3": "/andy/summary.pdf",
	//	})
	SetFiles(files map[string]string) HttpBuilder

	Call() (data []byte, code int, err error)
}

type HttpResponse interface {
	Body() []byte
	IsError() bool
	IsSuccess() bool
	Error() error
	Code() int

	IsTimeout() bool
	IsConnectionReset() bool
}
