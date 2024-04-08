/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
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
