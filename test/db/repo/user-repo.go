/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package repo

import (
	"github.com/andypangaribuan/gmod/ice"
	"github.com/andypangaribuan/gmod/test/db/entity"
)

var User *stuRepo[entity.User]

func init() {
	add(func(dbi ice.DbInstance) {
		User = new(dbi, "users", `
					created_at, updated_at, deleted_at, uid,
					name, address, height, gold_amount`,
			func(e *entity.User) []any {
				return []any{
					e.CreatedAt, e.UpdatedAt, e.DeletedAt, e.Uid,
					e.Name, e.Address, e.Height, e.GoldAmount,
				}
			})
	})
}
