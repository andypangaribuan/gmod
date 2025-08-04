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
	"encoding/base64"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/fct"
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

func (slf *stuUtilEnv) GetFloat32(key string, dval ...float32) float32 {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	value, err := strconv.ParseFloat(sval, 32)
	slf.invalid(key, sval, "float32", err)
	return float32(value)
}

func (slf *stuUtilEnv) GetFloat64(key string, dval ...float64) float64 {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	value, err := strconv.ParseFloat(sval, 64)
	slf.invalid(key, sval, "float64", err)
	return value
}

func (slf *stuUtilEnv) GetFCT(key string, dval ...fct.FCT) fct.FCT {
	sval, val := getEnv(key, dval...)
	if val != nil {
		return *val
	}

	value, err := fct.New(sval)
	slf.invalid(key, sval, "fct.FCT", err)
	return *value
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

func (slf *stuUtilEnv) GetStringSlice(key string, separator string, dval ...[]string) []string {
	value := getEnvVal(key)

	switch {
	case value == "" && len(dval) > 0:
		return dval[0]
	case value == "":
		log.Fatalf(`env key "%v" doesn't exists`, key)
	}

	ls := strings.Split(value, separator)
	for i, v := range ls {
		ls[i] = strings.TrimSpace(v)
	}

	return ls
}

func (slf *stuUtilEnv) GetDurationMs(key string, dval ...time.Duration) time.Duration {
	val := getEnvVal(key)
	if val == "" {
		if len(dval) > 0 {
			return dval[0]
		}

		log.Fatalf(`env key "%v" doesn't exists`, key)
	}

	value, err := strconv.Atoi(val)
	slf.invalid(key, val, "duration-ms (int)", err)

	return time.Millisecond * time.Duration(value)
}

func (slf *stuUtilEnv) GetBase64(key string) ice.UtilEnvBase64 {
	val := getEnvVal(key)
	if val == "" {
		log.Fatalf(`env key "%v" doesn't exists`, key)
	}

	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		log.Fatalf("base64 decode failed, error: %+v\n", err)
	}

	return &stuEnvBase64{
		key:  key,
		data: data,
	}
}

func (slf *stuUtilEnv) GetKeysByPrefix(prefix string) []string {
	return getKeysByPrefix(prefix)
}
