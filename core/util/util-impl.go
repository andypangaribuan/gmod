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

	"github.com/andypangaribuan/gmod/fm"
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

func (slf *srUtil) SmallUID() string {
	return slf.UID(0)
}

func (slf *srUtil) LiteUID() string {
	return slf.UID(3)
}

func (slf *srUtil) UID(addition ...int) string {
	length := *fm.GetFirst(addition, 8)
	length = fm.Ternary(length < 0, 0, length)

	randId := ""
	if length > 0 {
		randId = slf.GetRandom(length, numeric+alphabetLower+alphabetUpper)
	}

	tn := time.Now().UTC().Format("2006-01-02 15:04:05.000000")
	tn = strings.ReplaceAll(tn, "-", "")
	tn = strings.ReplaceAll(tn, ":", "")
	tn = strings.ReplaceAll(tn, ".", "")
	tn = strings.ReplaceAll(tn, " ", "")

	// 2024-09-30 23:59:59.999999
	// id1:20.240  id2:93.123  id3:59.599  id4:99.999 -> 12 char
	id1 := l3uidN[l3uidAddZero(l3uidLength, tn[0:5])]
	id2 := l3uidN[l3uidAddZero(l3uidLength, tn[5:10])]
	id3 := l3uidN[l3uidAddZero(l3uidLength, tn[10:15])]
	id4 := l3uidN[l3uidAddZero(l3uidLength, tn[15:20])]

	return id1 + id2 + id3 + id4 + randId
}

func (slf *srUtil) DecodeUID(uid string, addition ...int) (rawId string, randId string, err error) {
	length := *fm.GetFirst(addition, 8)
	length = fm.Ternary(length < 0, 0, length)

	if length > 0 && len(uid) > 12 {
		randId = uid[12:]
		uid = uid[:12]
	}

	uids, err := l3uidSplitter(uid)
	if err != nil {
		return "", "", err
	}

	raw := ""
	num := ""
	chunkSize := 5

	for _, uid := range uids {
		num = l3uidK[uid]
		cut := len(num) - chunkSize
		if cut < 0 {
			cut = 0
		}

		raw += num[cut:]
	}

	return raw, randId, nil
}

func (slf *srUtil) GetAlphabet(isUpper ...bool) string {
	upper := *fm.GetFirst(isUpper, false)
	return fm.Ternary(upper, alphabetUpper, alphabetLower)
}

func (slf *srUtil) GetNumeric() string {
	return numeric
}

func (slf *srUtil) GetRandom(length int, value string) string {
	if length < 1 {
		return ""
	}

	var (
		res   = ""
		count = -1
		max   = len(value)
		min   = 1
		rin   int
	)

	for {
		count++
		if count == length {
			break
		}

		rin = xRand.Intn(max) + min
		res += value[rin-1 : rin]
	}

	return res
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
