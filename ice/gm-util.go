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
	ConcurrentProcess(total, max int, fn func(index int))

	LiteUID() string
	UID(addition ...int) string
	GetAlphabet(isUpper ...bool) string
	GetNumeric() string
	GetRandom(length int, value string) string
	DecodeUID(uid string, addition ...int) (rawId string, randId string, err error)
	ReplaceAll(value *string, replaceValue string, replaceKey ...string) *string

	ReadTextFile(filePath string) ([]string, error)
	LoadEnv(filePath ...string) error
	GetExecDir() (string, error)

	PanicCatcher(fn func()) (err error)
	ReflectionGet(obj any, fieldName string) (any, error)
	ReflectionSet(obj any, bind map[string]any) error
}
