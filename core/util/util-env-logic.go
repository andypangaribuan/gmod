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

func loadZxEnv() {
	if len(zxEnv) > 0 {
		return
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
		return
	}

	value := os.Getenv(*zxEnvName)
	if value == "" {
		return
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
}

func getFromZxEnv(key string) string {
	loadZxEnv()
	val, ok := zxEnv[key]
	if !ok {
		return ""
	}

	return val
}

func getEnv[T any](key string, dval ...T) (string, *T) {
	value := getEnvVal(key)

	switch {
	case value == "" && len(dval) > 0:
		return value, &dval[0]
	case value == "":
		log.Fatalf(`env key "%v" doesn't exists`, key)
	}

	return value, nil
}

func getEnvVal(key string) string {
	value := getFromZxEnv(key)
	if value == "" {
		value = strings.TrimSpace(os.Getenv(key))
	}

	switch {
	case strings.Contains(value, " #> "):
		ls := strings.Split(value, " #> ")
		value = strings.TrimSpace(ls[0])

	case strings.Contains(value, " ## "):
		ls := strings.Split(value, " ## ")
		value = strings.TrimSpace(ls[0])

	case strings.Contains(value, " ### "):
		ls := strings.Split(value, " ### ")
		value = strings.TrimSpace(ls[0])
	}

	return value
}

func (*stuUtilEnv) invalid(key string, sval string, typ string, err ...error) {
	if len(err) == 0 || err[0] != nil {
		log.Fatalf(`env value "%v", from key env key "%v" is not a valid %v value`, sval, key, typ)
	}
}

func getKeysByPrefix(prefix string) []string {
	prefixLength := len(prefix)
	ls := make([]string, 0)

	loadZxEnv()
	if len(zxEnv) > 0 {
		for key := range zxEnv {
			if len(key) >= prefixLength && key[:prefixLength] == prefix {
				ls = append(ls, key)
			}
		}
	} else {
		envVars := os.Environ()
		for _, key := range envVars {
			key = strings.Split(key, "=")[0]
			if len(key) >= prefixLength && key[:prefixLength] == prefix {
				ls = append(ls, key)
			}
		}
	}

	return ls
}
