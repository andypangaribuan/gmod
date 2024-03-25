/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package conf

func (slf *srConf) SetZxEnvName(name string) {
	slf.zxEnvName = name
}

func (slf *srConf) SetTimeZone(zone string) {
	slf.timeZone = zone
}
