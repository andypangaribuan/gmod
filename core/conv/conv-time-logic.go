/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package conv

import (
	"strings"
	"time"
)

var timeReplacer = [][]string{
	{"yyyy", "2006"},
	{"yy", "06"},

	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},

	{"dd", "02"},
	{"HH", "15"},
	{"mm", "04"},
	{"ss", "05"},

	{"SSSSSS", "999999"},
	{"SSSSS", "99999"},
	{"SSSS", "9999"},
	{"SSS", "999"},
	{"SS", "99"},
	{"S", "9"},

	{"TZ", "-07:00"},
}

func (slf *srConvTime) toStr(val time.Time, layout string) string {
	return val.Format(slf.replace(layout))
}

func (slf *srConvTime) replace(layout string) string {
	for _, arr := range timeReplacer {
		layout = strings.Replace(layout, arr[0], arr[1], -1)
	}
	return layout
}
