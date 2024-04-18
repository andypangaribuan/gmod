/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import (
	"strings"

	"github.com/andypangaribuan/gmod/gm"
)

func getConfValue(name string) (value string) {
	val, err := gm.Util.ReflectionGet(gm.Conf, name)
	if err == nil {
		if v, ok := val.(string); ok {
			value = strings.TrimSpace(v)
		}
	}
	return
}
