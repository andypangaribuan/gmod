/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"fmt"
	"reflect"
	"unicode"
	"unsafe"
)

func (slf *srUtil) reflectionSet(sf reflect.StructField, rv reflect.Value, obj interface{}) (err error) {
	switch rv.CanSet() {
	case true:
		err = slf.reflectionPublicSet(sf, rv, obj)
	case false:
		err = slf.reflectionPrivateSet(sf, rv, obj)
	}
	return
}

func (slf *srUtil) reflectionPublicSet(rs reflect.StructField, rv reflect.Value, obj interface{}) error {
	err := slf.PanicCatcher(func() {
		rv.Set(reflect.ValueOf(obj))
	})
	return slf.reflectionSetError(rs.Name, err)
}

func (slf *srUtil) reflectionPrivateSet(rs reflect.StructField, rv reflect.Value, obj interface{}) error {
	var first rune
	for _, c := range rs.Name {
		first = c
		break
	}

	if unicode.IsUpper(first) {
		return fmt.Errorf("cannot set the field: %v", rs.Name)
	}

	ptr := unsafe.Pointer(rv.UnsafeAddr())
	newRV := reflect.NewAt(rv.Type(), ptr)
	val := newRV.Elem()
	err := slf.PanicCatcher(func() {
		val.Set(reflect.ValueOf(obj))
	})
	return slf.reflectionSetError(rs.Name, err)
}

func (*srUtil) reflectionSetError(fieldName string, err error) error {
	if err != nil {
		return fmt.Errorf("%v\nfield name: %v", err, fieldName)
	}
	return err
}
