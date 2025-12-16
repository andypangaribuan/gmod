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
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/fm"
)

func l3uidGenerate() ([]string, map[string]string, map[string]string) {
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
			for i := range 10 {
				index++
				chars[index] = string(rn + int32(i))
			}
		}

		for i := range 26 {
			index++
			chars[index] = string(ru + int32(i))
		}

		for i := range 26 {
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
			n := l3uidAddZero(lc, fmt.Sprint(i))
			mn[n] = k
			mk[k] = n
		}

		return mn, mk
	}

	l3x := uidL3(false) // 140.608 item
	l3xn, l3xk := km(l3x)
	return l3x, l3xn, l3xk
}

func l3uidAddZero(length int, v string) string {
	for len(v) < length {
		v = "0" + v
	}

	return v
}

func l3uidCheckUnique(ids []string) {
	mids := make(map[string]any, 0)

	for _, id := range ids {
		if id == "" {
			log.Fatalf("l3-uid: found empty id")
		}

		_, ok := mids[id]
		if !ok {
			mids[id] = nil
		} else {
			log.Fatalf(`l3-uid: have duplicate id "%v"`, id)
		}
	}
}

func l3uidSplitter(uid string) ([]string, error) {
	uid = strings.TrimSpace(uid)
	if len(uid)%3 != 0 {
		return nil, errors.New("l3-uid: invalid uid")
	}

	uids := make([]string, 0)
	for uid != "" {
		length := len(uid)
		size := fm.Ternary(length >= 3, 3, length)
		uids = append(uids, uid[:size])
		uid = uid[size:]
	}

	return uids, nil
}

func getRandom(max int) []int {
	xRandMx.Lock()
	defer xRandMx.Unlock()
	return xRand.Perm(max)
}

func encodeToBase52(n uint64, length int) string {
	out := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		out[i] = base52Alphabet[n%base52]
		n /= base52
	}
	return string(out)
}

func encodeTimestampBase52(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	micro := t.Nanosecond() / 1000

	m := int(month) - 1
	d := day - 1

	var index uint64
	index = uint64(year)
	index = index*12 + uint64(m)
	index = index*31 + uint64(d)
	index = index*24 + uint64(hour)
	index = index*60 + uint64(minute)
	index = index*60 + uint64(second)
	index = index*1_000_000 + uint64(micro)

	return encodeToBase52(index, 11)
}

func decodeFromBase52(s string) uint64 {
	var n uint64
	for i := 0; i < len(s); i++ {
		ch := s[i]
		pos := uint64(0)
		for j := 0; j < base52; j++ {
			if base52Alphabet[j] == ch {
				pos = uint64(j)
				break
			}
		}
		n = n*base52 + pos
	}
	return n
}

func decodeTimestampBase52(code string) time.Time {
	n := decodeFromBase52(code)

	micro := int(n % 1_000_000)
	n /= 1_000_000

	second := int(n % 60)
	n /= 60

	minute := int(n % 60)
	n /= 60

	hour := int(n % 24)
	n /= 24

	d := int(n % 31)
	n /= 31

	m := int(n % 12)
	n /= 12

	year := int(n)

	month := time.Month(m + 1)
	day := d + 1

	return time.Date(year, month, day, hour, minute, second, micro*1000, time.UTC)
}

func encodeToBase62(n uint64, length int) string {
	out := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		out[i] = base62Alphabet[n%base62]
		n /= base62
	}
	return string(out)
}

func encodeTimestampBase62(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	micro := t.Nanosecond() / 1000

	m := int(month) - 1
	d := day - 1

	var index uint64
	index = uint64(year)
	index = index*12 + uint64(m)
	index = index*31 + uint64(d)
	index = index*24 + uint64(hour)
	index = index*60 + uint64(minute)
	index = index*60 + uint64(second)
	index = index*1_000_000 + uint64(micro)

	return encodeToBase62(index, 10)
}

func decodeFromBase62(s string) uint64 {
	var n uint64
	for i := 0; i < len(s); i++ {
		ch := s[i]
		pos := uint64(0)
		for j := 0; j < base62; j++ {
			if base62Alphabet[j] == ch {
				pos = uint64(j)
				break
			}
		}
		n = n*base62 + pos
	}
	return n
}

func decodeTimestampBase62(code string) time.Time {
	n := decodeFromBase62(code)

	micro := int(n % 1_000_000)
	n /= 1_000_000

	second := int(n % 60)
	n /= 60

	minute := int(n % 60)
	n /= 60

	hour := int(n % 24)
	n /= 24

	d := int(n % 31)
	n /= 31

	m := int(n % 12)
	n /= 12

	year := int(n)

	month := time.Month(m + 1)
	day := d + 1

	return time.Date(year, month, day, hour, minute, second, micro*1000, time.UTC)
}
