/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package lock

import (
	"fmt"
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
)

func pushClogReport(logc clog.Instance, key string, startedAt time.Time, err error, when string) {
	if logc == nil {
		return
	}

	var (
		finishedAt = time.Now()
		errWhen    *string
		errMessage *string
		stackTrace *string
	)

	if err != nil {
		errWhen = &when
		errMessage = fm.Ptr(err.Error())
		stackTrace = fm.Ptr(fmt.Sprintf("%+v", err))
	}

	mol := &clog.DistLockV1{
		Engine:     txLockEngineName,
		Address:    txLockEngineAddress,
		Key:        key,
		ErrWhen:    errWhen,
		ErrMessage: errMessage,
		StackTrace: stackTrace,
		StartedAt:  startedAt,
		FinishedAt: finishedAt,
	}

	_ = logc.DistLockV1(mol)
}
