/*
* Copyright (c) 2024.
* Created by Andy Pangaribuan <https://github.com/apangaribuan>.
* All Rights Reserved.
 */

package ice

type Util interface {
	IsEmailValid(email string, verifyDomain ...bool) bool

	PanicCatcher(fn func()) (err error)
	ReflectionGet(obj interface{}, fieldName string) (interface{}, error)
	ReflectionSet(obj interface{}, bind map[string]interface{}) error
}
