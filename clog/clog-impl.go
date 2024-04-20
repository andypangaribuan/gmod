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

func (slf *stuInstance) DbqV1(mol *DbqV1, async ...bool) error {
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
		ExecFunction: mol.ExecFunc,
		ErrorMessage: fm.PbwString(mol.ErrorMessage),
		StackTrace:   fm.PbwString(mol.StackTrace),
		Host1:        mol.Host1,
		Host2:        fm.PbwString(mol.Host2),
		Duration1:    int32(mol.Duration1),
		Duration2:    fm.PbwInt32(mol.Duration2),
		StartedAt:    gm.Conv.Time.ToStrFull(mol.StartedAt),
		FinishedAt:   gm.Conv.Time.ToStrFull(mol.FinishedAt),
	}

	return grpcCall(*fm.GetFirst(async, true), client.DbqV1, req)
}

func (slf *stuInstance) ServicePieceV1(mol *ServicePieceV1, async ...bool) error {
	req := &sclog.RequestServicePieceV1{
		Uid:              slf.uid,
		SvcName:          svcName,
		SvcVersion:       svcVersion,
		SvcParentName:    fm.PbwString(mol.SvcParentName),
		SvcParentVersion: fm.PbwString(mol.SvcParentVersion),
		Endpoint:         mol.Endpoint,
		Url:              mol.Url,
		ReqVersion:       fm.PbwString(mol.ReqVersion),
		ReqHeader:        fm.PbwString(mol.ReqHeader),
		ReqParam:         fm.PbwString(mol.ReqParam),
		ReqQuery:         fm.PbwString(mol.ReqQuery),
		ReqForm:          fm.PbwString(mol.ReqForm),
		ReqBody:          fm.PbwString(mol.ReqBody),
		ClientIp:         mol.ClientIp,
		StartedAt:        gm.Conv.Time.ToStrFull(mol.StartedAt),
	}

	return grpcCall(*fm.GetFirst(async, true), client.ServicePieceV1, req)
}

func (slf *stuInstance) ServiceV1(mol *ServiceV1, async ...bool) error {
	req := &sclog.RequestServiceV1{
		Uid:              slf.uid,
		UserId:           fm.PbwString(mol.UserId),
		PartnerId:        fm.PbwString(mol.PartnerId),
		SvcName:          svcName,
		SvcVersion:       svcVersion,
		SvcParentName:    fm.PbwString(mol.SvcParentName),
		SvcParentVersion: fm.PbwString(mol.SvcParentVersion),
		Endpoint:         mol.Endpoint,
		Url:              mol.Url,
		Severity:         mol.Severity,
		ExecPath:         mol.ExecPath,
		ExecFunction:     mol.ExecFunc,
		ReqVersion:       fm.PbwString(mol.ReqVersion),
		ReqHeader:        fm.PbwString(mol.ReqHeader),
		ReqParam:         fm.PbwString(mol.ReqParam),
		ReqQuery:         fm.PbwString(mol.ReqQuery),
		ReqForm:          fm.PbwString(mol.ReqForm),
		ReqFiles:         fm.PbwString(mol.ReqFiles),
		ReqBody:          fm.PbwString(mol.ReqBody),
		ResData:          fm.PbwString(mol.ResData),
		ResCode:          int32(mol.ResCode),
		ErrMessage:       fm.PbwString(mol.ErrMessage),
		StackTrace:       fm.PbwString(mol.StackTrace),
		ClientIp:         mol.ClientIp,
		StartedAt:        gm.Conv.Time.ToStrFull(mol.StartedAt),
		FinishedAt:       gm.Conv.Time.ToStrFull(mol.FinishedAt),
	}

	return grpcCall(*fm.GetFirst(async, true), client.ServiceV1, req)
}
