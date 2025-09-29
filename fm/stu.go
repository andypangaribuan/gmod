/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fm

import (
	"context"

	"github.com/andypangaribuan/gmod/clog"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type stuGrpcSender[RQ any, RS any] struct {
	ctx            context.Context
	destination    string
	logc           clog.Instance
	header         map[string]string
	req            *RQ
	getErrResponse func(code *wrapperspb.StringValue, message *wrapperspb.StringValue) *RS
}
