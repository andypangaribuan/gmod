/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import "github.com/andypangaribuan/gmod/clog"

func (slf *stuXDB) Select(logc clog.Instance, query string, args ...any) ([]*map[string]any, error) {
	return slf.override(logc, slf.fetches(nil, query, args)).fetches()
}
