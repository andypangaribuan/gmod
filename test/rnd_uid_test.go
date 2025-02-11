/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

/* spell-checker: disable */
package test

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/stretchr/testify/assert"
)

func printLog(t *testing.T, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	if t != nil {
		t.Log(message)
	}
	fmt.Print(message)
}

// go test -count=5 -v -run "^TestGenerateUid$"
func TestGenerateUid(t *testing.T) {
	startedAt := time.Now()
	defer func() {
		durationMs := time.Since(startedAt).Milliseconds()
		printLog(t, "\n\nduration: %v ms\n", durationMs)
	}()

	uid := UID(t)
	printLog(t, "uid : %v\n", uid)
	assert.Equal(t, 12, len(uid))
}

func TestDecodeUid(t *testing.T) {
	startedAt := time.Now()
	defer func() {
		durationMs := time.Since(startedAt).Milliseconds()
		printLog(t, "\n\nduration: %v ms\n", durationMs)
	}()

	uid := "HZMMHEAALDAw"
	raw, err := decodeUidL3(uid)
	assert.Nil(t, err)

	var (
		year         = raw[0:4]
		month        = raw[4:6]
		day          = raw[6:8]
		hour         = raw[8:10]
		minute       = raw[10:12]
		second       = raw[12:14]
		milliseconds = raw[14:20]
	)

	printLog(t, "%v-%v-%v %v:%v:%v.%v", year, month, day, hour, minute, second, milliseconds)
	assert.Equal(t, "2024-03-28 16:00:01.108160", fmt.Sprintf("%v-%v-%v %v:%v:%v.%v", year, month, day, hour, minute, second, milliseconds))
}

func decodeUidL3(uid string, chunk ...int) (string, error) {
	uid = strings.TrimSpace(uid)
	if uid == "" {
		return "", errors.New("uid-l3: empty value")
	}

	uids, err := uidL3Splitter(uid)
	if err != nil {
		return "", err
	}

	_, _, l3xk := uidL3()
	num := ""
	raw := ""
	chunkSize := *fm.GetFirst(chunk, 5)

	for _, uid := range uids {
		num = l3xk[uid]
		cut := len(num) - chunkSize
		if cut < 0 {
			cut = 0
		}

		raw += num[cut:]
	}

	return raw, nil
}

func uidL3Splitter(uid string) ([]string, error) {
	uid = strings.TrimSpace(uid)

	if len(uid)%3 != 0 {
		return nil, errors.New("uid-l3: invalid uid")
	}

	uids := make([]string, 0)

	for {
		if uid == "" {
			break
		}

		length := len(uid)
		size := fm.Ternary(length >= 3, 3, length)
		uids = append(uids, uid[:size])
		uid = uid[size:]
	}

	return uids, nil
}

func UID(t *testing.T) string {
	// l3: 140.608 item
	l3, l3xn, _ := uidL3()
	length := len(fmt.Sprint(len(l3)))

	checkUnique(t, l3)

	tn := time.Now().Format("2006-01-02 15:04:05.000000")
	printLog(t, "tn: %v\n", tn)

	tn = strings.ReplaceAll(tn, "-", "")
	tn = strings.ReplaceAll(tn, ":", "")
	tn = strings.ReplaceAll(tn, ".", "")
	tn = strings.ReplaceAll(tn, " ", "")

	// 2024-09-30 23:59:59.999999
	// id1:20.240  id2:93.123  id3:59.599  id4:99.999 -> 12 char
	id1 := l3xn[numsAddZero(length, tn[0:5])]
	id2 := l3xn[numsAddZero(length, tn[5:10])]
	id3 := l3xn[numsAddZero(length, tn[10:15])]
	id4 := l3xn[numsAddZero(length, tn[15:20])]

	uid := id1 + id2 + id3 + id4
	return uid
}

func uidL3() ([]string, map[string]string, map[string]string) {
	chars := func(withNumber bool) []string {
		length := 26 * 2
		if withNumber {
			length += 10
		}

		var (
			chars = make([]string, length)
			rn    = rune('0')
			rl    = rune('a')
			ru    = rune('A')
			index = -1
		)

		if withNumber {
			for i := 0; i < 10; i++ {
				index++
				chars[index] = string(rn + int32(i))
			}
		}

		for i := 0; i < 26; i++ {
			index++
			chars[index] = string(ru + int32(i))
		}

		for i := 0; i < 26; i++ {
			index++
			chars[index] = string(rl + int32(i))
		}

		return chars
	}

	uidL3 := func(withNumber bool) []string {
		var (
			chars  = chars(withNumber)
			length = len(chars)
			uids   = make([]string, length*length*length)
			index  = 0
		)

		for _, c1 := range chars {
			for _, c2 := range chars {
				for _, c3 := range chars {
					uids[index] = c1 + c2 + c3
					index++
				}
			}
		}

		return uids
	}

	km := func(uids []string) (map[string]string, map[string]string) {
		length := len(uids)
		lc := len(fmt.Sprint(length))
		mn := make(map[string]string, length)
		mk := make(map[string]string, length)

		for i, k := range uids {
			n := numsAddZero(lc, fmt.Sprint(i))
			mn[n] = k
			mk[k] = n
		}

		return mn, mk
	}

	l3x := uidL3(false) // 140.608 item
	l3xn, l3xk := km(l3x)
	return l3x, l3xn, l3xk
}

func numsAddZero(length int, v string) string {
	for {
		if len(v) >= length {
			break
		}

		v = "0" + v
	}
	return v
}

func checkUnique(t *testing.T, ids []string) {
	mids := make(map[string]any, 0)

	for _, id := range ids {
		if id == "" {
			t.Fatalf("empty id\n")
		}

		_, ok := mids[id]
		if !ok {
			mids[id] = nil
		} else {
			t.Fatalf("have duplicate id: %v\n", id)
		}
	}
}
