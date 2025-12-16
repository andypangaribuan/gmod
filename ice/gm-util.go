/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import "time"

type Util interface {
	IsEmailValid(email string, verifyDomain ...bool) bool
	Timenow(timezone ...string) time.Time
	ConcurrentProcess(total, max int, callback func(index int))
	XConcurrentProcess(maxConcurrent int, maxJob int) UtilConcurrentProcess

	UID52(addition ...int) string
	UID62(addition ...int) string
	LiteUID() string
	UID(addition ...int) string
	XID(addition ...int) string
	GetAlphabet(isUpper ...bool) string
	GetNumeric() string
	GetRandom(length int, value string) string
	DecodeUID52(uid string) (timeId *time.Time, randId string, err error)
	DecodeUID62(uid string) (timeId *time.Time, randId string, err error)
	DecodeXID(uid string, addition ...int) (rawId string, randId string, err error)
	ReplaceAll(value *string, replaceValue string, replaceKey ...string) *string

	ReadTextFile(filePath string) ([]string, error)
	LoadEnv(filePath ...string) error
	GetExecDir() (string, error)
	GetExecPathFunc(skip ...int) (string, string)
	SingleExec(fn func())

	PanicCatcher(fn func()) (err error)
	ReflectionGet(obj any, fieldName string) (any, error)
	ReflectionSet(obj any, bind map[string]any) error
	StackTrace(skip ...int) string
	IsAllowedIp(clientIp string, allowedIps ...string) bool
}
