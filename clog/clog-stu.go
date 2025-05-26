/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

type stuInstance struct {
	uid              string
	userId           *string
	partnerId        *string
	svcParentName    *string
	svcParentVersion *string
}

type stuQueue struct {
	logType string
	req     any
}
