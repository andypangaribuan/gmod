/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package model

type DbConnection struct {
	AppName             string
	Host                string
	Port                int
	Name                string
	Scheme              string
	Username            string
	Password            string
	UnsafeCompatibility bool
	AutoRebind          bool
	PrintSql            bool
	MaxLifeTimeIns      int
	MaxIdleTimeIns      int
	MaxIdle             int
	MaxOpen             int
}
