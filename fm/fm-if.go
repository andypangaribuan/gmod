/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

func IfHaveIn[T comparable](val T, in ...T) bool {
	for _, v := range in {
		if val == v {
			return true
		}
	}

	return false
}
