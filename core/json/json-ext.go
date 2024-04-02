/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package json

import (
	jsoniter "github.com/json-iterator/go"
)

var api jsoniter.API

func init() {
	api = configWithCustomTimeFormat

	setDefaultTimeFormat("2006-01-02 15:04:05.000000", nil)
	addLocaleAlias("-", nil)

	addTimeFormatAlias("date", "2006-01-02")
	addTimeFormatAlias("time", "15:04:05")
	addTimeFormatAlias("full", "2006-01-02 15:04:05")
	addTimeFormatAlias("full-millis", "2006-01-02 15:04:05.000")
	addTimeFormatAlias("full-micros", "2006-01-02 15:04:05.000000")
}
