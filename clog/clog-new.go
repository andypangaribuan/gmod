/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import "strings"

func New(args ...interface{}) Instance {
	var (
		uid              string
		userId           *string
		partnerId        *string
		svcParentName    *string
		svcParentVersion *string
	)

	for _, arg := range args {
		switch val := arg.(type) {
		case string:
			uid = strings.TrimSpace(val)

		case *string:
			if val != nil {
				uid = strings.TrimSpace(*val)
			}

		case map[string]string:
			for k, v := range val {
				switch k {
				case "gmod-uid":
					uid = strings.TrimSpace(v)

				case "gmod-user-id":
					v = strings.TrimSpace(v)
					userId = &v

				case "gmod-partner-id":
					v = strings.TrimSpace(v)
					partnerId = &v

				case "gmod-from-svc-name":
					v = strings.TrimSpace(v)
					svcParentName = &v

				case "gmod-from-svc-version":
					v = strings.TrimSpace(v)
					svcParentVersion = &v
				}
			}
		}
	}

	if uid == "" {
		uid = mrf1[string]("mrf-util-uid")
	}

	return &stuInstance{
		uid:              uid,
		userId:           userId,
		partnerId:        partnerId,
		svcParentName:    svcParentName,
		svcParentVersion: svcParentVersion,
	}
}
