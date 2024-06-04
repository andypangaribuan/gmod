/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import "strings"

func TrimSpace(val *string) *string {
	if val == nil {
		return nil
	}

	v := strings.TrimSpace(*val)
	return &v
}