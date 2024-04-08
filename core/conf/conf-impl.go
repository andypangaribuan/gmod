/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package conf

func (slf *stuConf) SetZxEnvName(name string) {
	slf.zxEnvName = name
}

func (slf *stuConf) SetTimeZone(timezone string) {
	slf.timezone = timezone
}
