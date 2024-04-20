/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

func (slf *stuFuseRMainContext) regulator() FuseRRegulator {
	return &stuFuseRRegulator{
		mcx:          slf,
		currentIndex: -1,
	}
}
