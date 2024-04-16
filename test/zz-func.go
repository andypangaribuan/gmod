/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	"runtime"
	"strings"
)

func getDirPath() string {
	_, filename, _, _ := runtime.Caller(0)
	idx := strings.LastIndex(filename, "/")
	if idx > -1 {
		return filename[:idx]
	}

	return filename
}