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

func (slf *stuFuseRMainContext) regulator() FuseRRegulator {
	return &stuFuseRRegulator{
		mcx:          slf,
		currentIndex: -1,
	}
}

func (slf *stuFuseRMainContext) severity() string {
	if slf.val.endpoint == "unrouted" {
		return slf.val.endpoint
	}

	severity := "unknown"

	switch {
	case slf.responseCode >= 100 && slf.responseCode <= 199:
		severity = "server"

	case slf.responseCode >= 200 && slf.responseCode <= 299:
		severity = "success"

	case slf.responseCode >= 300 && slf.responseCode <= 399:
		severity = "server"

	case slf.responseCode >= 400 && slf.responseCode <= 499:
		severity = "warning"

	case slf.responseCode >= 500 && slf.responseCode <= 599:
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
