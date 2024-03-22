/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package fm

func GetDefault[T any](val *T, defaultValue T) T {
	if val == nil {
		return defaultValue
	}

	return *val
}
