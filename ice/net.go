/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package ice

type Net interface {
	IsPortUsed(port int, host ...string) bool
}
