/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package gmod

import (
	"log"

	_ "github.com/andypangaribuan/gmod/gm"

	"github.com/andypangaribuan/gmod/core/json"
	"github.com/andypangaribuan/gmod/core/net"
	"github.com/andypangaribuan/gmod/ice"
)

var iceGM ice.GM

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	iceGM.SetJson(json.Create())
	iceGM.SetNet(net.Create())
}
