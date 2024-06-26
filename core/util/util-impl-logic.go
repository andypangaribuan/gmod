/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package util

import (
	"fmt"
	"reflect"
	"time"
	"unicode"
	"unsafe"

	"github.com/andypangaribuan/gmod/gm"
)

func (slf *stuUtil) dvalTimezone() string {
	if !isGetDvalTimezone {
		isGetDvalTimezone = true
		val, err := slf.ReflectionGet(gm.Conf, "timezone")
		if err == nil {
			if v, ok := val.(string); ok {
				dvalTimezone = v
			}
		}
	}

	return dvalTimezone
}

func (slf *stuUtil) getTimeLocation(timezone ...string) *time.Location {
	zone := ""
	if len(timezone) > 0 {
		zone = timezone[0]
	} else {
		zone = slf.dvalTimezone()
	}

	if zone == "" {
		return time.UTC
	}

	loc, ok := timezones[zone]
	if ok {
		return loc
	}

	timezoneLocking.Lock()
	defer timezoneLocking.Unlock()

	loc, ok = timezones[zone]
	if ok {
		return loc
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		timezones[zone] = time.UTC
	} else {
		timezones[zone] = loc
	}

	return timezones[zone]
}

func (slf *stuUtil) reflectionSet(sf reflect.StructField, rv reflect.Value, obj any) (err error) {
	switch rv.CanSet() {
	case true:
		err = slf.reflectionPublicSet(sf, rv, obj)
	case false:
		err = slf.reflectionPrivateSet(sf, rv, obj)
	}
	return
}

func (slf *stuUtil) reflectionPublicSet(rs reflect.StructField, rv reflect.Value, obj any) error {
	err := slf.PanicCatcher(func() {
		rv.Set(reflect.ValueOf(obj))
	})
	return slf.reflectionSetError(rs.Name, err)
}

func (slf *stuUtil) reflectionPrivateSet(rs reflect.StructField, rv reflect.Value, obj any) error {
	var first rune
	for _, c := range rs.Name {
		first = c
		break
	}

	if unicode.IsUpper(first) {
		return fmt.Errorf("cannot set the field: %v", rs.Name)
	}

	ptr := unsafe.Pointer(rv.UnsafeAddr())
	newRV := reflect.NewAt(rv.Type(), ptr)
	val := newRV.Elem()
	err := slf.PanicCatcher(func() {
		val.Set(reflect.ValueOf(obj))
	})
	return slf.reflectionSetError(rs.Name, err)
}

func (*stuUtil) reflectionSetError(fieldName string, err error) error {
	if err != nil {
		return fmt.Errorf("%v\nfield name: %v", err, fieldName)
	}
	return err
}
