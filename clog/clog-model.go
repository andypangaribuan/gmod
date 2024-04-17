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
	ExecFunction string
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

}