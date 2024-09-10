/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import "time"

type NoteV1 struct {
	ExecPath string
	ExecFunc string
	Key      *string
	SubKey   *string
	Data     string
}

type DbqV1 struct {
	SqlQuery     string
	SqlArgs      *string
	Severity     string
	ExecPath     string
	ExecFunc     string
	ErrorMessage *string
	StackTrace   *string
	Host1        string
	Host2        *string
	Duration1    int
	Duration2    *int
	StartedAt    time.Time
	FinishedAt   time.Time
}

type HttpCallV1 struct {
	Url        string
	Severity   string
	ReqHeader  *string
	ReqParam   *string
	ReqQuery   *string
	ReqForm    *string
	ReqFiles   *string
	ReqBody    *string
	ResData    *string
	ResCode    int
	ErrMessage *string
	StackTrace *string
	StartedAt  time.Time
	FinishedAt time.Time
}

type ServicePieceV1 struct {
	SvcParentName    *string
	SvcParentVersion *string
	Endpoint         string
	Url              string
	ReqVersion       *string
	ReqSource        *string
	ReqHeader        *string
	ReqParam         *string
	ReqQuery         *string
	ReqForm          *string
	ReqBody          *string
	ClientIp         string
	StartedAt        time.Time
}

type ServiceV1 struct {
	SvcParentName    *string
	SvcParentVersion *string
	Endpoint         string
	Url              string
	Severity         string
	ExecPath         string
	ExecFunc         string
	ReqVersion       *string
	ReqSource        *string
	ReqHeader        *string
	ReqParam         *string
	ReqQuery         *string
	ReqForm          *string
	ReqFiles         *string
	ReqBody          *string
	ResData          *string
	ResCode          int
	ResSubCode       string
	ErrMessage       *string
	StackTrace       *string
	ClientIp         string
	StartedAt        time.Time
	FinishedAt       time.Time
}

type GrpcV1 struct {
	Destination string
	Severity    string
	ExecPath    string
	ExecFunc    string
	ReqHeader   *string
	Data        *string
}
