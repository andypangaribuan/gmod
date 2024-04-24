/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import (
	"github.com/andypangaribuan/gmod/gm"
)

func getConfVal[T any](name string) (value T) {
	v, err := gm.Util.ReflectionGet(gm.Conf, name)
	if err == nil {
		if v, ok := v.(T); ok {
			value = v
		}
	}
	return
}

func isTxOnDevMode() bool {
	return txLockEngineAddress == "-"
}
