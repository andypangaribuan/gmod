/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package conf

import (
	"time"

	"github.com/andypangaribuan/gmod/ice"
)

func (slf *stuConf) SetZxEnvName(name string) ice.Conf {
	slf.zxEnvName = name
	return slf
}

func (slf *stuConf) SetTimezone(timezone string) ice.Conf {
	slf.timezone = timezone
	return slf
}

func (slf *stuConf) SetClogAddress(address string, svcName string, svcVersion string) ice.Conf {
	if address != "" && address != "-" {
		slf.clogAddress = address
	}
	slf.svcName = svcName
	slf.svcVersion = svcVersion
	return slf
}

func (slf *stuConf) SetClogRetryMaxDuration(duration time.Duration) ice.Conf {
	slf.clogRetryMaxDuration = &duration
	return slf
}

func (slf *stuConf) SetClogMaxConcurrentPusher(maxConcurrentPusher int) ice.Conf {
	slf.clogMaxConcurrentPusher = maxConcurrentPusher
	return slf
}

func (slf *stuConf) SetInternalBaseUrls(urls []string) ice.Conf {
	slf.internalBaseUrls = urls
	return slf
}

func (slf *stuConf) SetTxLockEngine(engine string, address string, dvalTimeout time.Duration, dvalTryFor *time.Duration) ice.Conf {
	slf.txLockEngineName = engine
	slf.txLockEngineAddress = address
	slf.txLockDvalTimeout = dvalTimeout
	slf.txLockDvalTryFor = dvalTryFor
	return slf
}

func (slf *stuConf) Commit() {
	mainConfCommit()
}
