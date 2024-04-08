/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package util

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/joho/godotenv"
)

func (*stuUtil) IsEmailValid(email string, verifyDomain ...bool) bool {
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

func (slf *stuUtil) Timenow(timezone ...string) time.Time {
	location := slf.getTimeLocation(timezone...)
	return time.Now().In(location)
}

func (slf *stuUtil) ConcurrentProcess(total, max int, fn func(index int)) {
	slf.concurrentProcess(total, max, fn)
}

func (slf *stuUtil) LiteUID() string {
	return slf.UID(3)
}

func (slf *stuUtil) UID(addition ...int) string {
	length := *fm.GetFirst(addition, 8)
	length = fm.Ternary(length < 0, 0, length)

	randId := ""
	if length > 0 {
		randId = slf.GetRandom(length, numeric+alphabetLower+alphabetUpper)
	}

	tn := time.Now().UTC().Format("2006-01-02 15:04:05.000000")
	tn = *slf.ReplaceAll(&tn, "", "-", " ", ":", ".")

	// 2024-09-30 23:59:59.999999
	// id1:20.240  id2:93.123  id3:59.599  id4:99.999 -> 12 char
	id1 := l3uidN[l3uidAddZero(l3uidLength, tn[0:5])]
	id2 := l3uidN[l3uidAddZero(l3uidLength, tn[5:10])]
	id3 := l3uidN[l3uidAddZero(l3uidLength, tn[10:15])]
	id4 := l3uidN[l3uidAddZero(l3uidLength, tn[15:20])]

	return id1 + id2 + id3 + id4 + randId
}

func (slf *stuUtil) GetAlphabet(isUpper ...bool) string {
	upper := *fm.GetFirst(isUpper, false)
	return fm.Ternary(upper, alphabetUpper, alphabetLower)
}

func (slf *stuUtil) GetNumeric() string {
	return numeric
}

func (slf *stuUtil) GetRandom(length int, value string) string {
	if length < 1 || length >= 100000 || len(value) == 0 {
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

func (slf *stuUtil) DecodeUID(uid string, addition ...int) (rawId string, randId string, err error) {
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

	var (
		year         = raw[0:4]
		month        = raw[4:6]
		day          = raw[6:8]
		hour         = raw[8:10]
		minute       = raw[10:12]
		second       = raw[12:14]
		milliseconds = raw[14:20]
	)

	rawId = fmt.Sprintf("%v-%v-%v %v:%v:%v.%v", year, month, day, hour, minute, second, milliseconds)
	return rawId, randId, nil
}

func (slf *stuUtil) ReplaceAll(value *string, replaceValue string, replaceKey ...string) *string {
	if value == nil {
		return value
	}

	newValue := *value
	for _, key := range replaceKey {
		newValue = strings.ReplaceAll(newValue, key, replaceValue)
	}

	return &newValue
}

func (slf *stuUtil) IsPortUsed(port int, host ...string) bool {
	var (
		targetHost = "127.0.0.1"
		timeout    = time.Second * 3
	)

	if len(host) > 0 && host[0] != "" {
		targetHost = host[0]
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetHost, strconv.Itoa(port)), timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}

func (*stuUtil) ReadTextFile(filePath string) ([]string, error) {
	readFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}

func (*stuUtil) LoadEnv(filePath ...string) error {
	return godotenv.Load(filePath...)
}

func (*stuUtil) GetExecDir() (string, error) {
	exec, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(exec), nil
}

func (*stuUtil) PanicCatcher(fn func()) (err error) {
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

func (*stuUtil) ReflectionGet(obj any, fieldName string) (any, error) {
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

func (slf *stuUtil) ReflectionSet(obj any, bind map[string]any) error {
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
