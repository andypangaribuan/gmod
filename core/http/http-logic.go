/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package http

import (
	"bytes"
	"encoding/json"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
)

func getJsonIndent(args *map[string]string) *string {
	if args != nil && len(*args) > 0 {
		data, err := gm.Json.Marshal(*args)
		if err == nil {
			var out bytes.Buffer
			err = json.Indent(&out, data, "", "  ")
			if err == nil {
				return fm.Ptr(out.String())
			}
		}
	}

	return nil
}
