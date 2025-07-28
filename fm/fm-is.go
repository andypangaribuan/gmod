/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import (
	"fmt"
	"reflect"
	"time"

	"github.com/andypangaribuan/gmod/fct"
)

func IsNil(val any) bool {
	if val == nil {
		return true
	}

	switch v := val.(type) {
	case *bool:
		if v == nil {
			return true
		}

	case *string:
		if v == nil {
			return true
		}

	case *int:
		if v == nil {
			return true
		}

	case *int8:
		if v == nil {
			return true
		}

	case *int16:
		if v == nil {
			return true
		}

	case *int32:
		if v == nil {
			return true
		}

	case *int64:
		if v == nil {
			return true
		}

	case *uint:
		if v == nil {
			return true
		}

	case *uint8:
		if v == nil {
			return true
		}

	case *uint16:
		if v == nil {
			return true
		}

	case *uint32:
		if v == nil {
			return true
		}

	case *uint64:
		if v == nil {
			return true
		}

	case *float32:
		if v == nil {
			return true
		}

	case *float64:
		if v == nil {
			return true
		}

	case *complex64:
		if v == nil {
			return true
		}

	case *complex128:
		if v == nil {
			return true
		}

	case *time.Time:
		if v == nil {
			return true
		}

	case *fct.FCT:
		if v == nil {
			return true
		}

	default:
		if fmt.Sprintf("%+v", val) == "<nil>" {
			return true
		}
	}

	v := reflect.ValueOf(val)
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return false
}
