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
	"errors"

	"github.com/andypangaribuan/gmod/gm"
)

func JsonCast[T any](value any, links ...string) (*T, error) {
	if len(links) == 0 {
		switch val := value.(type) {
		case T:
			return &val, nil

		case *T:
			return val, nil

		case []byte:
			var out *map[string]any
			err := gm.Json.UnMarshal(val, &out)
			if err != nil {
				return nil, err
			}

			if out == nil {
				return nil, errors.New("unable to unmarshal the data")
			}

			return JsonCast[T](*out)

		default:
			return nil, errors.New("cannot cast the value")
		}
	}

	switch val := value.(type) {
	case []byte:
		var out *map[string]any
		err := gm.Json.UnMarshal(val, &out)
		if err != nil {
			return nil, err
		}

		if out == nil {
			return nil, errors.New("unable to unmarshal the data")
		}

		if len(links) == 0 {
			return JsonCast[T](*out)
		}

		v, ok := (*out)[links[0]]
		if !ok {
			return nil, errors.New("doesn't have child " + links[0])
		}

		return JsonCast[T](v, links[1:]...)

	case map[string]any:
		v, ok := val[links[0]]
		if !ok {
			return nil, errors.New("doesn't have child " + links[0])
		}

		return JsonCast[T](v, links[1:]...)

	case *map[string]any:
		if val == nil {
			return nil, errors.New("cannot cast the value, value is nil")
		}

		v, ok := (*val)[links[0]]
		if !ok {
			return nil, errors.New("doesn't have child " + links[0])
		}

		return JsonCast[T](v, links[1:]...)
	}

	return nil, errors.New("cannot cast the value")
}
