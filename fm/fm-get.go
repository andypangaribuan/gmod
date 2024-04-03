/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package fm

func GetDefault[T any](val *T, dval T) T {
	if IsNil(val) {
		return dval
	}

	return *val
}

func GetFirst[T any](ls []T, dval ...T) *T {
	if len(ls) == 0 {
		if len(dval) > 0 {
			return &dval[0]
		}

		return nil
	}

	return &ls[0]
}
