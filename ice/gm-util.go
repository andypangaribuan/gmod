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

	PanicCatcher(fn func()) (err error)
	ReflectionGet(obj interface{}, fieldName string) (interface{}, error)
	ReflectionSet(obj interface{}, bind map[string]interface{}) error
}
