/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package test

import "github.com/andypangaribuan/gmod/ice"

type stuEnv struct {
	AppName        string
	AppEnv         ice.AppEnv
	AppTimezone    string
	AppRestPort    int
	AppAutoRecover bool
	AppServerPrintOnError bool

	DbHost string
	DbPort int
	DbName string
	DbUser string
	DbPass string
}
