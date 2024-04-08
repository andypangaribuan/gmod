/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

type Box interface {
	Set(key, subKey string, val any)
	GetValue(key, subKey string) any
	GetSub(key string) map[string]any
}
