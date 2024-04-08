/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import "testing"

type Test interface {
	Start(t *testing.T, fn func(t *testing.T))
	Printf(t *testing.T, format string, args ...any)
}
