/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fct

import "log"

func New(val any, dval ...FCT) (*FCT, error) {
	deci, err := convert(val)
	if err != nil {
		dv := dvalFCT(dval...)
		if dv != nil {
			return dv, nil
		}

		return nil, err
	}

	return create(*deci), nil
}

// be careful, unsafe perform panic if theres an error
func UnsafeNew(val any, dval ...FCT) FCT {
	deci, err := convert(val)
	if err != nil {
		dv := dvalFCT(dval...)
		if dv != nil {
			return *dv
		}

		log.Panic(err)
	}

	return *create(*deci)
}

// be careful, unsafe perform panic if theres an error
func UnsafePtrNew(val any, dval ...FCT) *FCT {
	v := UnsafeNew(val, dval...)
	return &v
}
