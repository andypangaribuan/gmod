/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	} else {
		return b
	}
}

func TernaryR[T any](condition bool, a T, b func() T) T {
	if condition {
		return a
	} else {
		return b()
	}
}

func TernaryLR[T any](condition bool, a func() T, b func() T) T {
	if condition {
		return a()
	} else {
		return b()
	}
}

func Ternary2LR[A any, B any](condition bool, l, r func() (A, B)) (A, B) {
	if condition {
		return l()
	} else {
		return r()
	}
}
