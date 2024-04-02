/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package fm

func Ptr[T any](val T) *T {
	return &val
}

func PtrGetFirst[T any](ls []*T, dval ...*T) *T {
	if len(ls) == 0 {
		if len(dval) > 0 {
			return dval[0]
		}

		return nil
	}

	return ls[0]
}
