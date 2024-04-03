/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"strconv"
	"strings"

	"github.com/andypangaribuan/gmod/ice"
)

func (slf *stuUtilEnv) GetAppEnv(key string) ice.AppEnv {
	return &stuUtilEnvAppEnv{
		val: slf.GetString(key),
	}
}

func (*stuUtilEnv) GetString(key string, dval ...string) string {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	return sval
}

func (slf *stuUtilEnv) GetInt(key string, dval ...int) int {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	value, err := strconv.Atoi(sval)
	slf.invalid(key, sval, "int", err)
	return value
}

func (slf *stuUtilEnv) GetInt32(key string, dval ...int32) int32 {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	value, err := strconv.ParseInt(sval, 10, 32)
	slf.invalid(key, sval, "int32", err)
	return int32(value)
}

func (slf *stuUtilEnv) GetInt64(key string, dval ...int64) int64 {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	value, err := strconv.ParseInt(sval, 10, 64)
	slf.invalid(key, sval, "int64", err)
	return value
}

func (slf *stuUtilEnv) GetBool(key string, dval ...bool) bool {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	switch strings.ToLower(sval) {
	case "1", "true":
		return true
	case "0", "false":
		return false
	}

	slf.invalid(key, sval, "boolean")
	return false
}
