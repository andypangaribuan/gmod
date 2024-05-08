/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import (
	"time"

	"github.com/andypangaribuan/gmod/fct"
)

type UtilEnv interface {
	GetAppEnv(key string) AppEnv
	GetString(key string, dval ...string) string
	GetInt(key string, dval ...int) int
	GetInt32(key string, dval ...int32) int32
	GetInt64(key string, dval ...int64) int64
	GetFloat32(key string, dval ...float32) float32
	GetFloat64(key string, dval ...float64) float64
	GetFCT(key string, dval ...fct.FCT) fct.FCT
	GetBool(key string, dval ...bool) bool

	GetStringSlice(key string, separator string, dval ...[]string) []string
	GetDurationMs(key string, dval ...time.Duration) time.Duration
}
