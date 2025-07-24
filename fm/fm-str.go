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
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/fct"
)

func TrimSpace(val *string) *string {
	if val == nil {
		return nil
	}

	v := strings.TrimSpace(*val)
	return &v
}

func StringEqual(val *string, compare string) bool {
	if val == nil {
		return false
	}

	return *val == compare
}

func MultiTrimSpace(vals ...any) {
	for _, val := range vals {
		switch v := val.(type) {
		case *string:
			if v != nil {
				*v = strings.TrimSpace(*v)
			}

		case **string:
			if v != nil {
				x := *v
				if x != nil {
					**v = strings.TrimSpace(*x)
				}
			}
		}
	}
}

func MultiToLower(vals ...any) {
	for _, val := range vals {
		switch v := val.(type) {
		case *string:
			if v != nil {
				*v = strings.ToLower(*v)
			}

		case **string:
			if v != nil {
				x := *v
				if x != nil {
					**v = strings.ToLower(*x)
				}
			}
		}
	}
}

func FindEmptyString(vals map[string]any) string {
	for k, val := range vals {
		switch v := val.(type) {
		case string:
			if v == "" {
				return k
			}

		case *string:
			if v != nil && *v == "" {
				return k
			}
		}
	}

	return ""
}

func FindNil(keyVals map[string]any) string {
	for key, val := range keyVals {
		if val == nil {
			return key
		}

		switch v := val.(type) {
		case *string:
			if v == nil {
				return key
			}

		case *int:
			if v == nil {
				return key
			}

		case *int16:
			if v == nil {
				return key
			}

		case *int32:
			if v == nil {
				return key
			}

		case *int64:
			if v == nil {
				return key
			}

		case *float32:
			if v == nil {
				return key
			}

		case *float64:
			if v == nil {
				return key
			}

		case *bool:
			if v == nil {
				return key
			}

		case *time.Time:
			if v == nil {
				return key
			}

		case *fct.FCT:
			if v == nil {
				return key
			}

		default:
			if fmt.Sprintf("%+v", val) == "<nil>" {
				return key
			}
		}
	}

	return ""
}

func FindNilOrEmptyString(keyVals map[string]any) string {
	keyNil := FindNil(keyVals)
	if keyNil != "" {
		return keyNil
	}

	return FindEmptyString(keyVals)
}

func SetTrimSpaceNilIfEmptyString(vals ...**string) {
	for _, val := range vals {
		if val != nil && *val != nil {
			nv := strings.TrimSpace(**val)
			if nv == "" {
				*val = nil
			} else {
				*val = &nv
			}
		}
	}
}
