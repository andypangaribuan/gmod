/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
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
