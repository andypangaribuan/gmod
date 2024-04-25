/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fct

func (slf *FCT) Float64(dval ...FCT) (float64, error) {
	return slf.float64(dval...)
}

func (slf *FCT) PtrFloat64(dval ...FCT) (*float64, error) {
	return slf.ptrFloat64(dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafeFloat64(dval ...FCT) float64 {
	return slf.unsafeFloat64(dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafePtrFloat64(dval ...FCT) *float64 {
	return slf.unsafePtrFloat64(dval...)
}

func (slf *FCT) Round(places int, dval ...FCT) (FCT, error) {
	return slf.round(places, dval...)
}

func (slf *FCT) PtrRound(places int, dval ...FCT) (*FCT, error) {
	return slf.ptrRound(places, dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafeRound(places int, dval ...FCT) FCT {
	return slf.unsafeRound(places, dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafePtrRound(places int, dval ...FCT) *FCT {
	return slf.unsafePtrRound(places, dval...)
}

func (slf *FCT) PtrFloor(places int, dval ...FCT) (*FCT, error) {
	return slf.ptrFloor(places, dval...)
}

func (slf *FCT) Floor(places int, dval ...FCT) (FCT, error) {
	return slf.floor(places, dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafeFloor(places int, dval ...FCT) FCT {
	return slf.unsafeFloor(places, dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafePtrFloor(places int, dval ...FCT) *FCT {
	return slf.unsafePtrFloor(places, dval...)
}

func (slf *FCT) Ceil(places int, dval ...FCT) (FCT, error) {
	return slf.ceil(places, dval...)
}

func (slf *FCT) PtrCeil(places int, dval ...FCT) (*FCT, error) {
	return slf.ptrCeil(places, dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafeCeil(places int, dval ...FCT) FCT {
	return slf.unsafeCeil(places, dval...)
}

// be careful, unsafe perform panic if theres an error
func (slf *FCT) UnsafePtrCeil(places int, dval ...FCT) *FCT {
	return slf.unsafePtrCeil(places, dval...)
}
