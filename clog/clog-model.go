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

type DbqV1 struct {
	UserId       *string
	PartnerId    *string
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

type ServicePieceV1 struct {
	SvcParentName    *string
	SvcParentVersion *string
	Endpoint         string
	Url              string
	ReqVersion       *string
	ReqHeader        *string
	ReqParam         *string
	ReqQuery         *string
	ReqForm          *string
	ReqBody          *string
	ClientIp         string
	StartedAt        time.Time
}

type ServiceV1 struct {
	UserId           *string
	PartnerId        *string
	SvcParentName    *string
	SvcParentVersion *string
	Endpoint         string
	Url              string
	Severity         string
	ExecPath         string
	ExecFunc         string
	ReqVersion       *string
	ReqHeader        *string
	ReqParam         *string
	ReqQuery         *string
	ReqForm          *string
	ReqFiles         *string
	ReqBody          *string
	ResData          *string
	ResCode          int
	ErrMessage       *string
	StackTrace       *string
	ClientIp         string
	StartedAt        time.Time
	FinishedAt       time.Time
}
