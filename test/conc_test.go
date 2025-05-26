/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkConcurrent(b *testing.B) {
	list := makeRandList()
	n := runtime.NumCPU()
	fmt.Printf("cpu: %v\n", n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSumSplit(n, list)
	}
}

func TestConcurrent(t *testing.T) {
	list := makeRandList()
	n := runtime.NumCPU()
	fmt.Printf("cpu: %v\n", n)

	FindSumSplit(n, list)
}

func makeRandList() []int {
	capacity := 1_000_000
	list := make([]int, capacity)
	for i := range list {
		list[i] = rand.Intn(capacity)
	}
	return list
}

func FindSumSplit(n int, list []int) int {
	// divide into non-overlapping segments
	segment := make([][]int, 0, n)
	size := (len(list) + n - 1) / n
	for i := 0; i < len(list); i += size {
		end := i + size
		if end > len(list) {
			end = len(list)
		}
		segment = append(segment, list[i:end])
	}

	// spread the load
	result := make([]int, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i, values := range segment {
		go func(index int, sublist []int) {
			defer wg.Done()
			sum := 0
			for _, number := range sublist {
				sum += number
			}
			result[index] = sum
		}(i, values)
	}
	wg.Wait()

	// total intermediate results
	total := 0
	for _, segmentSum := range result {
		total += segmentSum
	}
	return total
}
