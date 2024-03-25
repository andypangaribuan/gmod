/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package conf

func init() {
	sr := &srConf{
		zxEnvName: "ZX_ENV",
	}

	iceConf = sr
}
