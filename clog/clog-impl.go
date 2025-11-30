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
	"time"

	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

func (slf *stuInstance) NoteV1(mol *NoteV1, async ...bool) error {
	if client == nil {
		return nil
	}

	timenow := mrf1[time.Time]("mrf-util-timenow")
	if mol.ExecPath == "" && mol.ExecFunc == "" {
		mol.ExecPath, mol.ExecFunc = mrf2[string, string]("mrf-util-get-exec-path-func", 3)
	}

	req := &sclog.RequestNoteV1{
		Uid:          slf.uid,
		UserId:       pbwString(slf.userId),
		PartnerId:    pbwString(slf.partnerId),
		SvcName:      svcName,
		SvcVersion:   svcVersion,
		ExecPath:     mol.ExecPath,
		ExecFunction: mol.ExecFunc,
		Key:          pbwString(mol.Key),
		SubKey:       pbwString(mol.SubKey),
		Data:         mol.Data,
		OccurredAt:   timeToStrFull(timenow),
	}

	return grpcCall(*getFirst(async, true), client.NoteV1, "NoteV1", req)
}

func (slf *stuInstance) DbqV1(mol *DbqV1, async ...bool) error {
	if client == nil {
		return nil
	}

	req := &sclog.RequestDbqV1{
		Uid:          slf.uid,
		UserId:       pbwString(slf.userId),
		PartnerId:    pbwString(slf.partnerId),
		SvcName:      svcName,
		SvcVersion:   svcVersion,
		SqlQuery:     mol.SqlQuery,
		SqlArgs:      pbwString(mol.SqlArgs),
		Severity:     mol.Severity,
		ExecPath:     mol.ExecPath,
		ExecFunction: mol.ExecFunc,
		ErrorMessage: pbwString(mol.ErrorMessage),
		StackTrace:   pbwString(mol.StackTrace),
		DbName:       pbwString(mol.DbName),
		SchemaName:   pbwString(mol.SchemaName),
		Host1:        mol.Host1,
		Host2:        pbwString(mol.Host2),
		Duration1:    int32(mol.Duration1),
		Duration2:    pbwInt32(mol.Duration2),
		StartedAt:    timeToStrFull(mol.StartedAt),
		FinishedAt:   timeToStrFull(mol.FinishedAt),
	}

	return grpcCall(*getFirst(async, true), client.DbqV1, "DbqV1", req)
}

func (slf *stuInstance) HttpCallV1(mol *HttpCallV1, async ...bool) error {
	if client == nil {
		return nil
	}

	req := &sclog.RequestHttpCallV1{
		Uid:        slf.uid,
		UserId:     pbwString(slf.userId),
		PartnerId:  pbwString(slf.partnerId),
		SvcName:    svcName,
		SvcVersion: svcVersion,
		Url:        mol.Url,
		Severity:   mol.Severity,
		ReqHeader:  pbwString(mol.ReqHeader),
		ReqParam:   pbwString(mol.ReqParam),
		ReqQuery:   pbwString(mol.ReqQuery),
		ReqForm:    pbwString(mol.ReqForm),
		ReqFiles:   pbwString(mol.ReqFiles),
		ReqBody:    pbwString(mol.ReqBody),
		ResData:    pbwString(mol.ResData),
		ResCode:    int32(mol.ResCode),
		ErrMessage: pbwString(mol.ErrMessage),
		StackTrace: pbwString(mol.StackTrace),
		StartedAt:  timeToStrFull(mol.StartedAt),
		FinishedAt: timeToStrFull(mol.FinishedAt),
	}

	return grpcCall(*getFirst(async, true), client.HttpCallV1, "HttpCallV1", req)
}

func (slf *stuInstance) ServicePieceV1(mol *ServicePieceV1, async ...bool) error {
	if client == nil {
		return nil
	}

	req := &sclog.RequestServicePieceV1{
		Uid:              slf.uid,
		SvcName:          svcName,
		SvcVersion:       svcVersion,
		SvcParentName:    pbwString(mol.SvcParentName),
		SvcParentVersion: pbwString(mol.SvcParentVersion),
		Endpoint:         mol.Endpoint,
		Url:              mol.Url,
		ReqVersion:       pbwString(mol.ReqVersion),
		ReqSource:        pbwString(mol.ReqSource),
		ReqHeader:        pbwString(mol.ReqHeader),
		ReqParam:         pbwString(mol.ReqParam),
		ReqQuery:         pbwString(mol.ReqQuery),
		ReqForm:          pbwString(mol.ReqForm),
		ReqBody:          pbwString(mol.ReqBody),
		ClientIp:         mol.ClientIp,
		StartedAt:        timeToStrFull(mol.StartedAt),
	}

	return grpcCall(*getFirst(async, true), client.ServicePieceV1, "ServicePieceV1", req)
}

func (slf *stuInstance) ServiceV1(mol *ServiceV1, async ...bool) error {
	if client == nil {
		return nil
	}

	req := &sclog.RequestServiceV1{
		Uid:              slf.uid,
		UserId:           pbwString(slf.userId),
		PartnerId:        pbwString(slf.partnerId),
		SvcName:          svcName,
		SvcVersion:       svcVersion,
		SvcParentName:    pbwString(mol.SvcParentName),
		SvcParentVersion: pbwString(mol.SvcParentVersion),
		Endpoint:         mol.Endpoint,
		Url:              mol.Url,
		Severity:         mol.Severity,
		ExecPath:         mol.ExecPath,
		ExecFunction:     mol.ExecFunc,
		ReqVersion:       pbwString(mol.ReqVersion),
		ReqSource:        pbwString(mol.ReqSource),
		ReqHeader:        pbwString(mol.ReqHeader),
		ReqParam:         pbwString(mol.ReqParam),
		ReqQuery:         pbwString(mol.ReqQuery),
		ReqForm:          pbwString(mol.ReqForm),
		ReqFiles:         pbwString(mol.ReqFiles),
		ReqBody:          pbwString(mol.ReqBody),
		ResData:          pbwString(mol.ResData),
		ResCode:          int32(mol.ResCode),
		ResSubCode:       mol.ResSubCode,
		ErrMessage:       pbwString(mol.ErrMessage),
		StackTrace:       pbwString(mol.StackTrace),
		ClientIp:         mol.ClientIp,
		StartedAt:        timeToStrFull(mol.StartedAt),
		FinishedAt:       timeToStrFull(mol.FinishedAt),
	}

	return grpcCall(*getFirst(async, true), client.ServiceV1, "ServiceV1", req)
}

func (slf *stuInstance) GrpcV1(mol *GrpcV1, async ...bool) error {
	if client == nil {
		return nil
	}

	req := &sclog.RequestGrpcV1{
		Uid:              slf.uid,
		UserId:           pbwString(slf.userId),
		PartnerId:        pbwString(slf.partnerId),
		SvcName:          svcName,
		SvcVersion:       svcVersion,
		SvcParentName:    pbwString(slf.svcParentName),
		SvcParentVersion: pbwString(slf.svcParentVersion),
		Destination:      mol.Destination,
		Severity:         mol.Severity,
		ExecPath:         mol.ExecPath,
		ExecFunction:     mol.ExecFunc,
		ReqHeader:        pbwString(mol.ReqHeader),
		Data:             pbwString(mol.Data),
		ErrMessage:       pbwString(mol.ErrMessage),
		StackTrace:       pbwString(mol.StackTrace),
		StartedAt:        timeToStrFull(mol.StartedAt),
		FinishedAt:       timeToStrFull(mol.FinishedAt),
	}

	return grpcCall(*getFirst(async, true), client.GrpcV1, "GrpcV1", req)
}

func (slf *stuInstance) DistLockV1(mol *DistLockV1, async ...bool) error {
	if client == nil {
		return nil
	}

	req := &sclog.RequestDistLockV1{
		Uid:        slf.uid,
		UserId:     pbwString(slf.userId),
		PartnerId:  pbwString(slf.partnerId),
		SvcName:    svcName,
		SvcVersion: svcVersion,
		Engine:     mol.Engine,
		Address:    mol.Address,
		Key:        mol.Key,
		ErrWhen:    pbwString(mol.ErrWhen),
		ErrMessage: pbwString(mol.ErrMessage),
		StackTrace: pbwString(mol.StackTrace),
		ObtainAt:   timeToStrFull(mol.ObtainAt),
		StartedAt:  timeToStrFull(mol.StartedAt),
		FinishedAt: timeToStrFull(mol.FinishedAt),
	}

	return grpcCall(*getFirst(async, true), client.DistLockV1, "DistLockV1", req)
}
