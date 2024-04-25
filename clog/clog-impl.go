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
	"github.com/andypangaribuan/gmod/grpc/service/sclog"
)

func (slf *stuInstance) DbqV1(mol *DbqV1, async ...bool) error {
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
		Host1:        mol.Host1,
		Host2:        pbwString(mol.Host2),
		Duration1:    int32(mol.Duration1),
		Duration2:    pbwInt32(mol.Duration2),
		StartedAt:    timeToStrFull(mol.StartedAt),
		FinishedAt:   timeToStrFull(mol.FinishedAt),
	}

	return grpcCall(*getFirst(async, true), client.DbqV1, req)
}

func (slf *stuInstance) HttpCallV1(mol *HttpCallV1, async ...bool) error {
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

	return grpcCall(*getFirst(async, true), client.HttpCallV1, req)
}

func (slf *stuInstance) ServicePieceV1(mol *ServicePieceV1, async ...bool) error {
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

	return grpcCall(*getFirst(async, true), client.ServicePieceV1, req)
}

func (slf *stuInstance) ServiceV1(mol *ServiceV1, async ...bool) error {
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
		ErrMessage:       pbwString(mol.ErrMessage),
		StackTrace:       pbwString(mol.StackTrace),
		ClientIp:         mol.ClientIp,
		StartedAt:        timeToStrFull(mol.StartedAt),
		FinishedAt:       timeToStrFull(mol.FinishedAt),
	}

	return grpcCall(*getFirst(async, true), client.ServiceV1, req)
}
