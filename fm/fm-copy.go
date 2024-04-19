/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import "maps"

func MapCopy[K comparable, V any](original map[K]V) map[K]V {
	dest := make(map[K]V, 0)
	maps.Copy(dest, original)
	return dest
}
