/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

func Val[T any](condition bool, fn func() T) *T {
	if !condition {
		return nil
	}

	val := fn()
	return &val
}
