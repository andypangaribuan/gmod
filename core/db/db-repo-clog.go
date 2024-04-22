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
)

func pushClogReport(cin clog.Instance, report *stuReport, err error) {
	if cin == nil || report == nil {
		return
	}

	var (
		execPathFuncSkipLevel = 2
		sqlArgs               *string
		errMessage            *string
		stackTrace            *string
		host1                 string
		host2                 *string
		duration1             int
		duration2             *int
	)

	execPath, execFunc := gm.Util.GetExecPathFunc(execPathFuncSkipLevel)

	if len(report.execReport.Args) > 0 {
		jons, err := gm.Json.Encode(report.execReport.Args)
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
		if idx == 0 {
			trace = strings.Replace(trace, msg, "", 1)
			trace = strings.TrimSpace(trace)
		}

		errMessage = &msg
		stackTrace = &trace
	}

	for i, h := range report.execReport.Hosts {
		if i == 0 {
			host1 = h.Host
			duration1 = int(h.DurationMs)
		} else {
			host2 = &h.Host
			duration2 = fm.Ptr(int(h.DurationMs))
		}
	}

	mol := &clog.DbqV1{
		SqlQuery:     report.execReport.Query,
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
		StartedAt:    report.execReport.StartedAt,
		FinishedAt:   report.execReport.FinishedAt,
	}

	cin.DbqV1(mol)
}
