/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package fm

func GetDefault[T any](val *T, dval T) T {
	if val == nil {
		return dval
	}

	return *val
}
