/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package json

import (
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/gm"
	jsoniter "github.com/json-iterator/go"
)

var api jsoniter.API

const dvalTimeFormat = "2006-01-02 15:04:05 -07:00"

func init() {
	api = configWithCustomTimeFormat
	jsoniter.RegisterTypeEncoder("[]uint8", &uint8Enc{})

	mainJsonCallback = func() {
		val, err := gm.Util.ReflectionGet(gm.Conf, "timezone")
		if err == nil {
			if v, ok := val.(string); ok {
				v = strings.TrimSpace(v)
				if v != "" {
					loc, _ := time.LoadLocation(v)
					setDefaultTimeFormat(dvalTimeFormat, loc)
				}
			}
		}
	}

	setDefaultTimeFormat(dvalTimeFormat, nil)
	addLocaleAlias("-", nil)

	addTimeFormatAlias("date", "2006-01-02")
	addTimeFormatAlias("time", "15:04:05")
	addTimeFormatAlias("dt", dvalTimeFormat)
	addTimeFormatAlias("dt-ms", "2006-01-02 15:04:05.000 -07:00")
	addTimeFormatAlias("dt-micro", "2006-01-02 15:04:05.000000 -07:00")
}
