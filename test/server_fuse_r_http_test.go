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
	"fmt"
	"testing"

	"github.com/andypangaribuan/gmod/gm"
)

// go test -v -run ^TestServerFuseRHttp$
func TestServerFuseRHttp(t *testing.T) {
	baseUrl := fmt.Sprintf("http://127.0.0.1:%v", env.AppRestPort)
	url := baseUrl + "/private/status"
	gm.Http.Get(url).Call()
}
