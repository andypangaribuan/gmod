/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import (
	"slices"
	"sort"
)

func SliceRemove[T any](slice []T, index ...int) []T {
	if len(index) == 0 {
		return slice
	}

	sort.Slice(index, func(i, j int) bool {
		return index[i] > index[j]
	})

	for _, idx := range index {
		slice = slices.Delete(slice, idx, idx+1)
	}

	return slice
}

func UniqueSlice[T comparable](slice []T) []T {
	seen := make(map[T]bool, 0)
	result := make([]T, 0)

	for _, item := range slice {
		if _, ok := seen[item]; !ok {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
