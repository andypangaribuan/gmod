/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package entity

import (
	"time"

	"github.com/andypangaribuan/gmod/fct"
)

type User struct {
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
	Uid        string     `db:"uid"`
	Name       string     `db:"name"`
	Address    *string    `db:"address"`
	Height     *int       `db:"height"`
	GoldAmount *fct.FCT   `db:"gold_amount"`
}
