/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package http

func xinit() {
	internalBaseUrls = make([]string, 0)

	mainHttpCallback = func() {
		if val := getConfVal[string]("svcName"); val != "" {
			svcName = val
		}

		if val := getConfVal[string]("svcVersion"); val != "" {
			svcVersion = val
		}

		if val := getConfVal[[]string]("internalBaseUrls"); val != nil {
			internalBaseUrls = val
		}
	}
}
