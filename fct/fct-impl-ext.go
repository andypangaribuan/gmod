/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fct

func (slf *FCT) unsafeFloat64(dval ...FCT) float64 {
	val, err := slf.float64(dval...)
	slf.panic(err)
	return val
}

func (slf *FCT) ptrFloat64(dval ...FCT) (*float64, error) {
	val, err := slf.float64(dval...)
	if err != nil {
		return nil, err
	}

	return &val, nil
}

func (slf *FCT) unsafePtrFloat64(dval ...FCT) *float64 {
	val := slf.unsafeFloat64(dval...)
	return &val
}

func (slf *FCT) round(places int, dval ...FCT) (FCT, error) {
	val, err := slf.ptrRound(places, dval...)
	if err != nil {
		return emptyFCT, err
	}

	return *val, nil
}

func (slf *FCT) unsafeRound(places int, dval ...FCT) FCT {
	val, err := slf.ptrRound(places, dval...)
	slf.panic(err)
	return *val
}

func (slf *FCT) unsafePtrRound(places int, dval ...FCT) *FCT {
	val, err := slf.ptrRound(places, dval...)
	slf.panic(err)
	return val
}

func (slf *FCT) floor(places int, dval ...FCT) (FCT, error) {
	val, err := slf.ptrFloor(places, dval...)
	if err != nil {
		return emptyFCT, err
	}

	return *val, nil
}

func (slf *FCT) unsafeFloor(places int, dval ...FCT) FCT {
	val, err := slf.ptrFloor(places, dval...)
	slf.panic(err)
	return *val
}

func (slf *FCT) unsafePtrFloor(places int, dval ...FCT) *FCT {
	val, err := slf.ptrFloor(places, dval...)
	slf.panic(err)
	return val
}

func (slf *FCT) ceil(places int, dval ...FCT) (FCT, error) {
	val, err := slf.ptrCeil(places, dval...)
	if err != nil {
		return emptyFCT, err
	}

	return *val, nil
}

func (slf *FCT) unsafeCeil(places int, dval ...FCT) FCT {
	val, err := slf.ptrCeil(places, dval...)
	slf.panic(err)
	return *val
}

func (slf *FCT) unsafePtrCeil(places int, dval ...FCT) *FCT {
	val, err := slf.ptrCeil(places, dval...)
	slf.panic(err)
	return val
}