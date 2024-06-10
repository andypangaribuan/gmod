/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package util

func (slf *stuEnvBase64) Key() string {
	return slf.key
}

func (slf *stuEnvBase64) Data() []byte {
	return slf.data
}
