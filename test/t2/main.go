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

	chunk := chunkSlice(qe, 1000)

	fmt.Println(queue)
	fmt.Println(len(queue))
	fmt.Println("----")
	fmt.Println(qe)
	fmt.Println(len(qe))
	fmt.Println("----")
	fmt.Println(len(chunk))
	fmt.Println(chunk[0])
	fmt.Println(chunk[1])
}

func copy() []int {
	i := -1
	ls := make([]int, len(queue))
	for key := range queue {
		i++
		ls[i] = key
	}

	// cp := queue
	queue = make(map[int]any, 0)
	return ls
}

func chunkSlice[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := min(i+chunkSize, len(slice))
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
