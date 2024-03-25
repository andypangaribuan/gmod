/*
* Copyright (c) 2024.
* Created by Andy Pangaribuan <https://github.com/apangaribuan>.
* All Rights Reserved.
 */

package ice

type UtilEnv interface {
	GetAppEnv(key string) AppEnv
	GetString(key string, dval ...string) string
	GetInt(key string, dval ...int) int
	GetInt32(key string, dval ...int32) int32
	GetInt64(key string, dval ...int64) int64
	GetBool(key string, dval ...bool) bool
}
