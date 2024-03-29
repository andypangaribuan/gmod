/*
* Copyright (c) 2024.
* Created by Andy Pangaribuan <https://github.com/apangaribuan>.
* All Rights Reserved.
 */

package ice

import "time"

type Util interface {
	IsEmailValid(email string, verifyDomain ...bool) bool
	Timenow(timezone ...string) time.Time

	SmallUID() string
	LiteUID() string
	UID(addition ...int) string
	GetAlphabet(isUpper ...bool) string
	GetNumeric() string
	GetRandom(length int, value string) string
	DecodeUID(uid string, addition ...int) (rawId string, randId string, err error)

	PanicCatcher(fn func()) (err error)
	ReflectionGet(obj interface{}, fieldName string) (interface{}, error)
	ReflectionSet(obj interface{}, bind map[string]interface{}) error
}
