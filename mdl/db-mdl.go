/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package mdl

import "time"

type DbConnection struct {
	AppName             string
	Host                string
	Port                int
	Name                string
	Scheme              string
	Username            string
	Password            string
	UnsafeCompatibility *bool
	PrintUnsafeError    *bool
	AutoRebind          *bool
	PrintSql            *bool
	MaxLifeTimeIns      int
	MaxIdleTimeIns      int
	MaxIdle             int
	MaxOpen             int
}

type DbExecReport struct {
	Query      string
	Args       []any
	StartedAt  time.Time
	FinishedAt time.Time
	DurationMs int64
	Hosts      []*DbExecReportHost
}

type DbExecReportHost struct {
	Host       string
	StartedAt  time.Time
	FinishedAt time.Time
	DurationMs int64
}
