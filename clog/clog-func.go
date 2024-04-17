/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import "github.com/andypangaribuan/gmod/gm"

func New() Instance {
	if client == nil {
		return nil
	}

	return &stuInstance{
		uid: gm.Util.UID(),
	}
}
