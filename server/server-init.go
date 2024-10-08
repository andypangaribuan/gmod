/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import "sync"

func init() {
	serverImpl = new(stuServer)
	cronMX = make(map[string]*sync.Mutex, 0)
	cronIsStartUp = make(map[string]bool, 0)
}
