/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import (
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func DirectPbwGet[T any](obj any, dval ...T) *T {
	objVal, _ := PbwGet(obj, dval...)
	return objVal
}

func PbwGet[T any](obj any, dval ...T) (*T, bool) {
	var objVal any

	switch v := obj.(type) {
	case *wrapperspb.BoolValue:
		if v != nil {
			objVal = v.Value
		}

	case *wrapperspb.StringValue:
		if v != nil {
			objVal = v.Value
		}

	case *wrapperspb.Int32Value:
		if v != nil {
			objVal = v.Value
			if val, ok := objVal.(T); !ok {
				switch fmt.Sprintf("%T", val) {
				case "int":
					objVal = int(v.Value)
				case "int64":
					objVal = int64(v.Value)
				}
			}
		}

	case *wrapperspb.Int64Value:
		if v != nil {
			objVal = v.Value
			if val, ok := objVal.(T); !ok {
				switch fmt.Sprintf("%T", val) {
				case "int":
					objVal = int(v.Value)
				case "int32":
					objVal = int32(v.Value)
				}
			}
		}

	case *wrapperspb.FloatValue:
		if v != nil {
			objVal = v.Value
			if val, ok := objVal.(T); !ok {
				switch fmt.Sprintf("%T", val) {
				case "float64":
					objVal = float64(v.Value)
				}
			}
		}

	case *wrapperspb.DoubleValue:
		if v != nil {
			objVal = v.Value
			if val, ok := objVal.(T); !ok {
				switch fmt.Sprintf("%T", val) {
				case "float32":
					objVal = float32(v.Value)
				}
			}
		}
	}

	switch {
	case objVal == nil && len(dval) > 0:
		return &dval[0], true
	case objVal == nil:
		return nil, false
	}

	val, ok := objVal.(T)
	return &val, ok
}
