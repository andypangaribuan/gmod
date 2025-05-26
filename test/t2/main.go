/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package main

import (
	"fmt"
	"strconv"
)

var (
	queue map[int]any
)

func init() {
	queue = make(map[int]any, 0)
}

func main() {
	for i := range 100 {
		queue[i] = strconv.Itoa(i)
	}

	// qe := queue
	// queue = make(map[int]any, 0)

	qe := copy()

	queue[-1] = "***"
	fmt.Println(queue)
	fmt.Println(len(queue))
	fmt.Println("----")
	fmt.Println(qe)
	fmt.Println(len(qe))
}

func copy() map[int]any {
	cp := queue
	queue = make(map[int]any, 0)
	return cp
}
