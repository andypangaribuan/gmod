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
	mainHttpCallback = func() {
		if val := getConfValue("svcName"); val != "" {
			svcName = val
		}

		if val := getConfValue("svcVersion"); val != "" {
			svcVersion = val
		}
	}
}
