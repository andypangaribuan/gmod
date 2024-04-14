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
	jsoniter "github.com/json-iterator/go"
)

var api jsoniter.API

func init() {
	api = configWithCustomTimeFormat

	setDefaultTimeFormat("2006-01-02 15:04:05.000000 -07:00", nil)
	addLocaleAlias("-", nil)

	addTimeFormatAlias("date", "2006-01-02")
	addTimeFormatAlias("time", "15:04:05")
	addTimeFormatAlias("dt", "2006-01-02 15:04:05 -07:00")
	addTimeFormatAlias("dt-ms", "2006-01-02 15:04:05.000 -07:00")
	addTimeFormatAlias("dt-micro", "2006-01-02 15:04:05.000000 -07:00")
}
