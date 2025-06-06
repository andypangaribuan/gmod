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
	"time"
	"unsafe"

	"github.com/json-iterator/go"
)

const (
	ansic       = "ANSIC"
	unixDate    = "UnixDate"
	rubyDate    = "RubyDate"
	rfc822      = "RFC822"
	rfc822Z     = "RFC822Z"
	rfc850      = "RFC850"
	rfc1123     = "RFC1123"
	rfc1123Z    = "RFC1123Z"
	rfc3339     = "RFC3339"
	rfc3339Nano = "RFC3339Nano"
	kitchen     = "Kitchen"
	stamp       = "Stamp"
	stampMilli  = "StampMilli"
	stampMicro  = "StampMicro"
	stampNano   = "StampNano"
)

// time zone alias
const (
	Local = "Local"
	UTC   = "UTC"
)

const (
	tagNameTimeFormat   = "jdt"          // ori: time_format
	tagNameTimeLocation = "jdt_location" // ori: time_location
)

// var configWithCustomTimeFormat = jsoniter.ConfigCompatibleWithStandardLibrary
var configWithCustomTimeFormat = jsoniter.Config{
	EscapeHTML:             false,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

var formatAlias = map[string]string{
	ansic:       time.ANSIC,
	unixDate:    time.UnixDate,
	rubyDate:    time.RubyDate,
	rfc822:      time.RFC822,
	rfc822Z:     time.RFC822Z,
	rfc850:      time.RFC850,
	rfc1123:     time.RFC1123,
	rfc1123Z:    time.RFC1123Z,
	rfc3339:     time.RFC3339,
	rfc3339Nano: time.RFC3339Nano,
	kitchen:     time.Kitchen,
	stamp:       time.Stamp,
	stampMilli:  time.StampMilli,
	stampMicro:  time.StampMicro,
	stampNano:   time.StampNano,
}

var localeAlias = map[string]*time.Location{
	Local: time.Local,
	UTC:   time.UTC,
}

var (
	defaultFormat = time.RFC3339
	defaultLocale = time.Local
)

func init() {
	configWithCustomTimeFormat.RegisterExtension(&customTimeExtension{})
}

func addTimeFormatAlias(alias, format string) {
	formatAlias[alias] = format
}

func addLocaleAlias(alias string, locale *time.Location) {
	localeAlias[alias] = locale
}

func setDefaultTimeFormat(timeFormat string, timeLocation *time.Location) {
	defaultFormat = timeFormat
	defaultLocale = timeLocation
}

type customTimeExtension struct {
	jsoniter.DummyExtension
}

func (extension *customTimeExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, binding := range structDescriptor.Fields {
		var typeErr error
		var isPtr bool
		typeName := binding.Field.Type().String()

		switch typeName {
		case "time.Time":
			isPtr = false

		case "*time.Time":
			isPtr = true

		default:
			continue
		}

		timeFormat := defaultFormat
		formatTag := binding.Field.Tag().Get(tagNameTimeFormat)
		if format, ok := formatAlias[formatTag]; ok {
			timeFormat = format
		} else if formatTag != "" {
			timeFormat = formatTag
		}
		locale := defaultLocale
		if localeTag := binding.Field.Tag().Get(tagNameTimeLocation); localeTag != "" {
			if loc, ok := localeAlias[localeTag]; ok {
				locale = loc
			} else {
				loc, err := time.LoadLocation(localeTag)
				if err != nil {
					typeErr = err
				} else {
					locale = loc
				}
			}
		}

		binding.Encoder = &funcEncoder{fun: func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			if typeErr != nil {
				stream.Error = typeErr
				return
			}

			var tp *time.Time
			if isPtr {
				tpp := (**time.Time)(ptr)
				tp = *(tpp)
			} else {
				tp = (*time.Time)(ptr)
			}

			if tp != nil {
				var str string
				if locale != nil {
					lt := tp.In(locale)
					str = lt.Format(timeFormat)
				} else {
					str = tp.Format(timeFormat)
				}
				stream.WriteString(str)
			} else {
				_, _ = stream.Write([]byte("null"))
			}
		}}
		binding.Decoder = &funcDecoder{fun: func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if typeErr != nil {
				iter.Error = typeErr
				return
			}

			str := iter.ReadString()
			var t *time.Time
			if str != "" {
				var err error
				var tmp time.Time
				if locale != nil {
					tmp, err = time.ParseInLocation(timeFormat, str, locale)
				} else {
					tmp, err = time.Parse(timeFormat, str)
				}
				if err != nil {
					iter.Error = err
					return
				}
				t = &tmp
			} else {
				t = nil
			}

			if isPtr {
				tpp := (**time.Time)(ptr)
				*tpp = t
			} else {
				tp := (*time.Time)(ptr)
				if tp != nil && t != nil {
					*tp = *t
				}
			}
		}}
	}
}

type funcDecoder struct {
	fun jsoniter.DecoderFunc
}

func (decoder *funcDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	decoder.fun(ptr, iter)
}

type funcEncoder struct {
	fun         jsoniter.EncoderFunc
	isEmptyFunc func(ptr unsafe.Pointer) bool
}

func (encoder *funcEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	encoder.fun(ptr, stream)
}

func (encoder *funcEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	if encoder.isEmptyFunc == nil {
		return false
	}
	return encoder.isEmptyFunc(ptr)
}
