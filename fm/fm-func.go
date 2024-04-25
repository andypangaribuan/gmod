/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

func mrf2[A any, B any](key string, arg ...any) (va A, vb B) {
	arr := mainReflection(key, arg...)

	if v, ok := arr[0].(A); ok {
		va = v
	}

	if v, ok := arr[1].(B); ok {
		vb = v
	}

	return
}
