/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fc

import (
	"errors"
	"log"
	"reflect"
	"runtime/debug"

	"github.com/andypangaribuan/gmod/fm"
)

func Compare(v1 any, operation string, v2 any) bool {
	v, err := SCompare(v1, operation, v2)
	if err != nil {
		debug.PrintStack()
		log.Panicf("error: %+v\nv1: %v, op: %v, v2: %v\n", err, v1, operation, v2)
	}

	return v
}

func SCompare(v1 any, operation string, v2 any) (bool, error) {
	if !fm.IfHaveIn(operation, "==", "!=", "<", "<=", ">=", ">") {
		return false, errors.New("fc.SCompare: invalid operation")
	}

	if fm.IsNil(v1) {
		return false, errors.New("v1 cannot nil")
	}

	if fm.IsNil(v2) {
		return false, errors.New("v2 cannot nil")
	}

	if rv := reflect.ValueOf(v1); rv.Kind() == reflect.Ptr {
		v1 = rv.Elem().Interface()
	}

	if rv := reflect.ValueOf(v2); rv.Kind() == reflect.Ptr {
		v2 = rv.Elem().Interface()
	}

	var (
		fv1 FCT
		fv2 FCT
	)

	switch v := v1.(type) {
	case FCT:
		fv1 = v

	default:
		fv1 = New(v)
	}

	switch v := v2.(type) {
	case FCT:
		fv2 = v
	default:
		fv2 = New(v)
	}

	switch operation {
	case "==":
		return fv1.vd.Equal(fv2.vd), nil

	case "!=":
		return !fv1.vd.Equal(fv2.vd), nil

	case "<":
		return fv1.vd.LessThan(fv2.vd), nil

	case "<=":
		return fv1.vd.LessThanOrEqual(fv2.vd), nil

	case ">=":
		return fv1.vd.GreaterThanOrEqual(fv2.vd), nil

	case ">":
		return fv1.vd.GreaterThan(fv2.vd), nil
	}

	return false, errors.New("unknown error")
}
