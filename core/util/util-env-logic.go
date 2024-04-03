/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

import (
	"log"
	"os"
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
)

var (
	zxEnv     map[string]string
	zxEnvName *string
)

func init() {
	zxEnv = make(map[string]string, 0)
}

func getFromZxEnv(key string) string {
	if len(zxEnv) > 0 {
		return zxEnv[key]
	}

	if zxEnvName == nil {
		val, err := iceUtil.ReflectionGet(gm.Conf, "zxEnvName")
		if err != nil {
			zxEnvName = fm.Ptr("")
		} else {
			val, _ := val.(string)
			zxEnvName = &val
		}
	}

	if *zxEnvName == "" {
		return ""
	}

	value := os.Getenv(*zxEnvName)
	if value == "" {
		return ""
	}

	lines := strings.Split(value, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		eqIdx := strings.Index(line, "=")

		if line == "" || line[:1] == "#" || eqIdx < 1 {
			continue
		}

		key := strings.TrimSpace(line[:eqIdx])
		val := strings.TrimSpace(line[eqIdx+1:])

		if key != "" && val != "" {
			zxEnv[key] = val
		}
	}

	return zxEnv[key]
}

func getEnv[T any](key string, dval ...T) (string, *T) {
	value := getFromZxEnv(key)
	if value == "" {
		value = strings.TrimSpace(os.Getenv(key))
	}

	switch {
	case value == "" && len(dval) > 0:
		return value, &dval[0]
	case value == "":
		log.Fatalf(`env key "%v" doesn't exists`, key)
	}

	return value, nil
}

func (*stuUtilEnv) invalid(key string, sval string, typ string, err ...error) {
	if len(err) == 0 || err[0] != nil {
		log.Fatalf(`env value "%v", from key env key "%v" is not a valid %v value`, sval, key, typ)
	}
}
