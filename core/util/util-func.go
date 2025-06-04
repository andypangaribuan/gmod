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
