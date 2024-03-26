/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package model

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
	Args       []interface{}
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
