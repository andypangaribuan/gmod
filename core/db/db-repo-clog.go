/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import (
	"fmt"
	"strings"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/pkg/errors"
)

func pushClogReport(cin clog.Instance, report *stuReport, err error, execPathFuncSkipLevel int) {
	if cin == nil || report == nil {
		return
	}

	var (
		timenow    = gm.Util.Timenow()
		query      = report.query
		sqlArgs    *string
		errMessage *string
		stackTrace *string
		host1      string
		host2      *string
		duration1  int
		duration2  *int
		startedAt  = timenow
		finishedAt = timenow
	)

	if report.execReport != nil {
		query = report.execReport.Query
		startedAt = report.execReport.StartedAt
		finishedAt = report.execReport.FinishedAt
	}

	execPath, execFunc := gm.Util.GetExecPathFunc(execPathFuncSkipLevel)

	if report.execReport != nil && len(report.execReport.Args) > 0 {
		jons, err := gm.Json.Encode(report.execReport.Args)
		if err == nil {
			sqlArgs = &jons
		} else {
			sqlArgs = fm.Ptr(fmt.Sprintf("%v", sqlArgs))
		}
	} else if len(report.args) > 0 {
		jons, err := gm.Json.Encode(report.args)
		if err == nil {
			sqlArgs = &jons
		} else {
			sqlArgs = fm.Ptr(fmt.Sprintf("%v", sqlArgs))
		}
	}

	if err != nil {
		var (
			msg   = err.Error()
			trace = fmt.Sprintf("%+v", err)
			idx   = strings.Index(trace, msg)
		)

		if msg == trace || len(strings.Split(trace, "\n")) < 3 {
			trace = fmt.Sprintf("%+v", errors.WithStack(err))
		}

		if idx == 0 {
			trace = strings.Replace(trace, msg, "", 1)
			trace = strings.TrimSpace(trace)
		}

		errMessage = &msg
		stackTrace = &trace
	}

	if report.execReport != nil && report.execReport.Hosts != nil {
		for i, h := range report.execReport.Hosts {
			if i == 0 {
				host1 = h.Host
				duration1 = int(h.DurationMs)
			} else {
				host2 = &h.Host
				duration2 = fm.Ptr(int(h.DurationMs))
			}
		}
	}

	mol := &clog.DbqV1{
		SqlQuery:     query,
		SqlArgs:      sqlArgs,
		Severity:     fm.Ternary(err == nil, "success", "error"),
		ExecPath:     execPath,
		ExecFunc:     execFunc,
		ErrorMessage: errMessage,
		StackTrace:   stackTrace,
		Host1:        host1,
		Host2:        host2,
		Duration1:    duration1,
		Duration2:    duration2,
		StartedAt:    startedAt,
		FinishedAt:   finishedAt,
	}

	_ = cin.DbqV1(mol)
}
