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
	"fmt"
	"net"
	"net/mail"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unsafe"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
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

func (slf *stuUtil) XConcurrentProcess(maxConcurrent int, maxJob int) ice.UtilConcurrentProcess {
	return slf.xConcurrentProcess(maxConcurrent, maxJob)
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

func (*stuUtil) ReadTextFile(filePath string) ([]string, error) {
	readFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = readFile.Close()
	}()

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

func (*stuUtil) GetExecPathFunc(skip ...int) (string, string) {
	skipFrame := 2
	if len(skip) > 0 && skip[0] > 0 {
		skipFrame += skip[0]
	}

	pc := make([]uintptr, 10)
	n := runtime.Callers(skipFrame, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%v:%v", frame.File, frame.Line), frame.Function
}

func (slf *stuUtil) SingleExec(fn func()) {
	path, fname := slf.GetExecPathFunc(3)
	key := fmt.Sprintf("%v:%v", path, fname)
	slf.initFuncMx.Lock()
	defer slf.initFuncMx.Unlock()

	if slf.initFuncExecuteMap == nil {
		slf.initFuncExecuteMap = make(map[string]any, 0)
	}

	_, ok := slf.initFuncExecuteMap[key]
	if !ok {
		slf.initFuncExecuteMap[key] = nil
		fn()
	}
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

				err = errors.New(fmt.Sprintf("%+v", pv))
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

func (slf *stuUtil) StackTrace(skip ...int) string {
	var (
		buf         = make([]byte, 1024)
		stackTrace  string
		rootRemoved = false
		lines       = make([]string, 0)
		skipFrame   = *fm.GetFirst(skip, 0)
	)

	runtime.Stack(buf, false)
	stackTrace = string(buf)

	ls := strings.Split(stackTrace, "\n")

	for i := 0; i < len(ls); i++ {
		line := ls[i]

		if !rootRemoved && strings.Contains(line, "github.com/andypangaribuan/gmod/core/util.(*stuUtil).StackTrace") {
			rootRemoved = true
			i++
			continue
		}

		if rootRemoved && skipFrame > 0 {
			skipFrame--
			i++
			continue
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func (slf *stuUtil) IsAllowedIp(clientIp string, allowedIps ...string) bool {
	var (
		authorized = false
		clientIp1  = ""
		clientIp2  = ""
		clientIp3  = ""
		clientIp4  = ""
	)

	ls := strings.Split(clientIp, ".")
	for i, v := range ls {
		switch i {
		case 0:
			clientIp1 = v
		case 1:
			clientIp2 = clientIp1 + "." + v
		case 2:
			clientIp3 = clientIp2 + "." + v
		case 3:
			clientIp4 = clientIp3 + "." + v
		}
	}

	for _, ip := range allowedIps {
		if ip == "*" {
			authorized = true
			break
		}

		length := len(strings.Split(ip, "."))
		switch {
		case length == 1 && clientIp1 == ip:
			authorized = true
		case length == 2 && clientIp2 == ip:
			authorized = true
		case length == 3 && clientIp3 == ip:
			authorized = true
		case length == 4 && clientIp4 == ip:
			authorized = true
		}

		if authorized {
			break
		}
	}

	return authorized
}
