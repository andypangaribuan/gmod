/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func (*srUtil) IsEmailValid(email string, verifyDomain ...bool) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	if len(verifyDomain) > 0 && verifyDomain[0] {
		parts := strings.Split(email, "@")
		mx, err := net.LookupMX(parts[1])
		if err != nil || len(mx) == 0 {
			return false
		}
	}

	return true
}

func (slf *srUtil) Timenow(timezone ...string) time.Time {
	location := slf.getTimeLocation(timezone...)
	return time.Now().In(location)
}

func (*srUtil) PanicCatcher(fn func()) (err error) {
	defer func() {
		pv := recover()
		if pv != nil {
			if v, ok := pv.(error); ok {
				err = v
			} else {
				rv := reflect.ValueOf(pv)
				if rv.Kind() == reflect.Ptr {
					pv = rv.Elem()
				}

				msg := fmt.Sprintf("%+v", pv)
				err = errors.New(msg)
			}
		}
	}()
	fn()
	return
}

func (*srUtil) ReflectionGet(obj interface{}, fieldName string) (interface{}, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		return nil, errors.New("obj must be a pointer")
	}

	val = val.Elem()
	typ := val.Type()

	if val.Kind() != reflect.Struct && val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	if val.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			rs := typ.Field(i)
			rf := val.Field(i)

			if rs.Name == fieldName {
				if rf.IsValid() {
					if rs.IsExported() {
						return rf.Interface(), nil
					}

					value := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem().Interface()
					return value, nil
				}
			}
		}
	}

	return nil, errors.New("not found")
}

func (slf *srUtil) ReflectionSet(obj interface{}, bind map[string]interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		return errors.New("obj must be a pointer")
	}

	val = val.Elem()
	typ := val.Type()

	if val.Kind() != reflect.Struct && val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	if val.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			rs := typ.Field(i)
			rf := val.Field(i)
			fieldName := rs.Name

			if bindVal, ok := bind[fieldName]; ok {
				if !rf.IsValid() {
					return fmt.Errorf("invalid field: %v", fieldName)
				}

				if rs.IsExported() {
					err := slf.reflectionSet(rs, rf, bindVal)
					if err != nil {
						return err
					}
				}

				reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).
					Elem().
					Set(reflect.ValueOf(bindVal))
			}
		}
	}

	return nil
}
