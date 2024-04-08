/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import (
	_ "github.com/andypangaribuan/gmod"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestEnv(t *testing.T) {
	appName := gm.Util.Env.GetString("APP_NAME", "not-found")
	t.Logf("value: %v\n", appName)
	assert.NotEqual(t, "not-found", appName)
}
