/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package conf

type stuConf struct {
	zxEnvName           string
	timezone            string   // accessed from reflection
	clogAddress         string   // accessed from reflection
	svcName             string   // accessed from reflection
	svcVersion          string   // accessed from reflection
	internalBaseUrls    []string // accessed from reflection
	txLockEngineAddress string   // accessed from reflection
}
