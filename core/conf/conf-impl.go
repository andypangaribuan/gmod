/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package conf

import "github.com/andypangaribuan/gmod/ice"

func (slf *stuConf) SetZxEnvName(name string) ice.Conf {
	slf.zxEnvName = name
	return slf
}

func (slf *stuConf) SetTimezone(timezone string) ice.Conf {
	slf.timezone = timezone
	return slf
}

func (slf *stuConf) SetCLogAddress(address string, svcName string, svcVersion string) ice.Conf {
	slf.clogAddress = address
	slf.svcName = svcName
	slf.svcVersion = svcVersion
	return slf
}

func (slf *stuConf) SetInternalBaseUrls(urls []string) ice.Conf {
	slf.internalBaseUrls = urls
	return slf
}

func (slf *stuConf) Commit() {
	mainConfCommit()
}
