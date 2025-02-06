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
		switch v := val.(type) {
		case string:
			continue

		case *string:
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
