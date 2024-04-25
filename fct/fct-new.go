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

func UNew(val any, dval ...FCT) FCT {
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

func NNew(val any, dval ...FCT) *FCT {
	v, _ := New(val, dval...)
	return v
}
