/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import "time"

type Conf interface {
	SetZxEnvName(name string) Conf
	SetTimezone(timezone string) Conf
	SetCLogAddress(address string, svcName string, svcVersion string) Conf
	SetInternalBaseUrls(urls []string) Conf
	SetTxLockEngineAddress(address string, dvalTimeout time.Duration, dvalTryFor *time.Duration) Conf
	Commit()
}
