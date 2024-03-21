/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gm

import "github.com/andypangaribuan/gmod/ice"

func (*srGM) SetNet(net ice.Net) {
	Net = &srNet{
		net,
	}
}

func (*srGM) SetJson(json ice.Json) {
	Json = &srJson{
		json,
	}
}
