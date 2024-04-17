/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

import (
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

func (slf *stuInstance) DbqV1(mol DbqV1) error {
	req := &sclog.RequestDbqV1{
		Uid:          slf.uid,
		UserId:       fm.PbwString(mol.UserId),
		PartnerId:    fm.PbwString(mol.PartnerId),
		SvcName:      svcName,
		SvcVersion:   svcVersion,
		SqlQuery:     mol.SqlQuery,
		SqlArgs:      fm.PbwString(mol.SqlArgs),
		Severity:     mol.Severity,
		ExecPath:     mol.ExecPath,
		ExecFunction: mol.ExecFunction,
		ErrorMessage: fm.PbwString(mol.ErrorMessage),
		StackTrace:   fm.PbwString(mol.StackTrace),
		Host1:        mol.Host1,
		Host2:        fm.PbwString(mol.Host2),
		Duration1:    int32(mol.Duration1),
		Duration2:    fm.PbwInt32(mol.Duration2),
		StartedAt:    gm.Conv.Time.ToStrFull(mol.StartedAt),
		FinishedAt:   gm.Conv.Time.ToStrFull(mol.FinishedAt),
	}

	_, err := fm.GrpcCall(client.DbqV1, req)
	return err
}
