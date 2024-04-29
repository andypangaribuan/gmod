/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"fmt"

	"github.com/andypangaribuan/gmod/fm"
)

func (slf *stuFuseRMainContext) regulator() *stuFuseRRegulator {
	return &stuFuseRRegulator{
		mcx:          slf,
		currentIndex: -1,
	}
}

func (slf *stuFuseRMainContext) severity() string {
	if slf.val.unrouted {
		return "unrouted"
	}

	severity := "unknown"

	switch {
	case slf.responseMeta.Code >= 100 && slf.responseMeta.Code <= 199:
		severity = "server"

	case slf.responseMeta.Code >= 200 && slf.responseMeta.Code <= 299:
		severity = "success"

	case slf.responseMeta.Code >= 300 && slf.responseMeta.Code <= 399:
		severity = "server"

	case slf.responseMeta.Code >= 400 && slf.responseMeta.Code <= 499:
		severity = "warning"

	case slf.responseMeta.Code >= 500 && slf.responseMeta.Code <= 599:
		severity = "error"
	}

	return severity
}

func (slf *stuFuseRMainContext) getUserId() *string {
	return slf.idCast(slf.userId)
}

func (slf *stuFuseRMainContext) getPartnerId() *string {
	return slf.idCast(slf.partnerId)
}

func (slf *stuFuseRMainContext) idCast(id any) *string {
	if id == nil {
		return nil
	}

	switch val := id.(type) {
	case string:
		return &val
	case *string:
		return val

	case int:
		return fm.Ptr(fmt.Sprint(val))
	case *int:
		if val != nil {
			return fm.Ptr(fmt.Sprint(*val))
		}

	case int32:
		return fm.Ptr(fmt.Sprint(val))
	case *int32:
		if val != nil {
			return fm.Ptr(fmt.Sprint(*val))
		}

	case int64:
		return fm.Ptr(fmt.Sprint(val))
	case *int64:
		if val != nil {
			return fm.Ptr(fmt.Sprint(*val))
		}
	}

	return nil
}
