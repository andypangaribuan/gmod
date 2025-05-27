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
	"math/rand"
	"strconv"
	"time"
)

var (
	queue map[int]any
)

func init() {
	queue = make(map[int]any, 0)
}

func main() {
	fmt.Println(Shuffle(3, "12345"))
	fmt.Println(Shuffle(4, "12345"))
	fmt.Println(Shuffle(5, "12345"))
	fmt.Println(Shuffle(6, "12345"))
	fmt.Println(Shuffle(7, "12345"))

	for i := range 100 {
		queue[i] = strconv.Itoa(i)
	}

	// qe := queue
	// queue = make(map[int]any, 0)

	qe := copy()
	queue[-1] = "***"

	chunk := chunkSlice(qe, 10)

	fmt.Println(queue)
	fmt.Println(len(queue))
	fmt.Println("----")
	fmt.Println(qe)
	fmt.Println(len(qe))
	fmt.Println("----")
	fmt.Println(len(chunk))
	fmt.Println(chunk[0])
	fmt.Println(chunk[1])

	xRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	mySlice := []int{10, 20, 30}
	randomIndex := xRand.Intn(len(mySlice)) - 2

	fmt.Println("Random Index:", randomIndex)
	fmt.Println("Value at index:", mySlice[randomIndex]) // Potential panic here
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

var rd = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(length int, value string) string {
	var (
		// rd    = rand.New(rand.NewSource(time.Now().Unix()))
		res   = ""
		max   = len(value)
		count = 0
	)

	for count < length {
		perm := rd.Perm(max)
		for _, randIndex := range perm {
			res += value[randIndex : randIndex+1]
			count++
			if count == length {
				break
			}
		}
	}

	// perm := rd.Perm(max)
	// for i, randIndex := range perm {
	// 	if i == length {
	// 		break
	// 	}

	// 	res += value[randIndex : randIndex+1]
	// }

	return res
}
