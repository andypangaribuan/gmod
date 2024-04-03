/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package conf

func (slf *stuConf) SetZxEnvName(name string) {
	slf.zxEnvName = name
}

func (slf *stuConf) SetTimeZone(timezone string) {
	slf.timezone = timezone
}
