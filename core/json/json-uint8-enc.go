/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package json

import (
	"strconv"
	"strings"
	"unsafe"

	"github.com/andypangaribuan/gmod/fct"
	jsoniter "github.com/json-iterator/go"
)

type uint8Enc struct{}

func (ue *uint8Enc) IsEmpty(ptr unsafe.Pointer) bool {
	data := *((*[]uint8)(ptr))
	return len(data) == 0
}
func (ue *uint8Enc) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	const hexTable = "0123456789abcdef"
	var (
		data = *((*[]uint8)(ptr))
		sb   strings.Builder
	)

	for _, v := range data {
		if strconv.IsPrint(rune(v)) {
			sb.WriteByte(v)
		} else {
			sb.WriteString(`\x`)
			sb.WriteByte(hexTable[v>>4])
			sb.WriteByte(hexTable[v&0x0f])
		}
	}

	sval := sb.String()
	val, err := fct.New(sval)
	if err == nil {
		v, err := val.ToString()
		if err == nil {
			stream.WriteRaw(v)
			return
		}
	}

	stream.WriteString(sval)
}
