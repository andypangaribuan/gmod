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
	"mime/multipart"
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/mol"
	"github.com/pkg/errors"
)

func (slf *stuFuseRContext) Clog() clog.Instance {
	return slf.mcx.clog
}

func (slf *stuFuseRContext) Auth(obj ...any) any {
	if len(obj) > 0 {
		slf.mcx.authObj = obj[0]
	}

	return slf.mcx.authObj
}

func (slf *stuFuseRContext) UserId(id ...any) any {
	if len(id) > 0 {
		slf.mcx.userId = id[0]
		slf.pushUserIdToClog()
	}

	return slf.mcx.userId
}

func (slf *stuFuseRContext) PartnerId(id ...any) any {
	if len(id) > 0 {
		slf.mcx.partnerId = id[0]
		slf.pushPartnerIdToClog()
	}

	return slf.mcx.partnerId
}

func (slf *stuFuseRContext) SetFiles(files map[string]string) {
	if len(files) > 0 {
		slf.mcx.files = &files
	}
}

func (slf *stuFuseRContext) ReqHeader() *map[string]string {
	return slf.header
}

func (slf *stuFuseRContext) ReqParam() *map[string]string {
	return slf.param
}

func (slf *stuFuseRContext) ReqQuery() *map[string]string {
	return slf.queries
}

func (slf *stuFuseRContext) ReqForm() *map[string][]string {
	return slf.form
}

func (slf *stuFuseRContext) ReqFile() *map[string][]*multipart.FileHeader {
	return slf.file
}

func (slf *stuFuseRContext) GetHeader(key string, dval ...string) *string {
	if slf.header != nil && len(*slf.header) > 0 {
		val, ok := (*slf.header)[strings.ToLower(key)]
		if ok {
			return &val
		}
	}

	if len(dval) > 0 {
		return &dval[0]
	}

	return nil
}

func (slf *stuFuseRContext) GetClientIP() string {
	if slf.mcx.clientIP == "" {
		slf.mcx.clientIP = cip.getClientIP(slf.mcx.fcx)
	}

	return slf.mcx.clientIP
}

func (slf *stuFuseRContext) RouteMethod() string {
	return strings.ToLower(slf.mcx.fcx.Route().Method)
}

func (slf *stuFuseRContext) RoutePath() string {
	return slf.mcx.fcx.Route().Path
}

func (slf *stuFuseRContext) ReqParser(header any, body any) error {
	if header != nil {
		if slf.header == nil || len(*slf.header) == 0 {
			return errors.New("parser: have no header")
		}

		data, err := gm.Json.Marshal(slf.header)
		if err != nil {
			return err
		}

		err = gm.Json.Unmarshal(data, header)
		if err != nil {
			return err
		}

		switch h := header.(type) {
		case *mol.RequestHeader:
			if h.RFTimeRaw != "" {
				tm, err := time.Parse(time.RFC3339, h.RFTimeRaw)
				if err != nil {
					rfc3339 := "2006-01-02 15:04:05Z07:00"
					tm, err = time.Parse(rfc3339, h.RFTimeRaw)
				}

				if err == nil {
					h.RFTime = &tm
				}
			}

		default:
			v, err := gm.Util.ReflectionGet(header, "RequestHeader")
			if err == nil && v != nil {
				if h, ok := v.(mol.RequestHeader); ok && h.RFTimeRaw != "" {
					tm, err := time.Parse(time.RFC3339, h.RFTimeRaw)
					if err != nil {
						rfc3339 := "2006-01-02 15:04:05Z07:00"
						tm, err = time.Parse(rfc3339, h.RFTimeRaw)
					}

					if err == nil {
						h.RFTime = &tm

						_ = gm.Util.ReflectionSet(header, map[string]interface{}{"RequestHeader": h})
					}
				}
			}
		}
	}

	if body != nil {
		if slf.bodyParser == nil {
			return errors.New("parser: have no body")
		}

		err := slf.bodyParser(body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (slf *stuFuseRContext) ReqParserPQF(param any, query any, form any) error {
	if form != nil {
		if slf.bodyParser != nil {
			err := slf.bodyParser(form)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (slf *stuFuseRContext) LastResponse() (val any, meta ResponseMeta) {
	return slf.mcx.responseVal, slf.mcx.responseMeta
}

func (slf *stuFuseRContext) R200OK(val any, opt ...ResponseOpt) any {
	return slf.setResponse(200, val, opt...)
}

func (slf *stuFuseRContext) R201Created(val any, opt ...ResponseOpt) any {
	return slf.setResponse(201, val, opt...)
}

func (slf *stuFuseRContext) R202Accepted(val any, opt ...ResponseOpt) any {
	return slf.setResponse(202, val, opt...)
}

func (slf *stuFuseRContext) R204NoContent(val any, opt ...ResponseOpt) any {
	return slf.setResponse(204, val, opt...)
}

func (slf *stuFuseRContext) R301MovedPermanently(val any, opt ...ResponseOpt) any {
	return slf.setResponse(301, val, opt...)
}

func (slf *stuFuseRContext) R307TemporaryRedirect(val any, opt ...ResponseOpt) any {
	return slf.setResponse(307, val, opt...)
}

func (slf *stuFuseRContext) R308PermanentRedirect(val any, opt ...ResponseOpt) any {
	return slf.setResponse(308, val, opt...)
}

func (slf *stuFuseRContext) R400BadRequest(val any, opt ...ResponseOpt) any {
	return slf.setResponse(400, val, opt...)
}

func (slf *stuFuseRContext) R401Unauthorized(val any, opt ...ResponseOpt) any {
	return slf.setResponse(401, val, opt...)
}

func (slf *stuFuseRContext) R403Forbidden(val any, opt ...ResponseOpt) any {
	return slf.setResponse(403, val, opt...)
}

func (slf *stuFuseRContext) R404NotFound(val any, opt ...ResponseOpt) any {
	return slf.setResponse(404, val, opt...)
}

func (slf *stuFuseRContext) R406NotAcceptable(val any, opt ...ResponseOpt) any {
	return slf.setResponse(406, val, opt...)
}

func (slf *stuFuseRContext) R412PreconditionFailed(val any, opt ...ResponseOpt) any {
	return slf.setResponse(412, val, opt...)
}

func (slf *stuFuseRContext) R418Teapot(val any, opt ...ResponseOpt) any {
	return slf.setResponse(418, val, opt...)
}

func (slf *stuFuseRContext) R428PreconditionRequired(val any, opt ...ResponseOpt) any {
	return slf.setResponse(428, val, opt...)
}

func (slf *stuFuseRContext) R500InternalServerError(err error, opt ...ResponseOpt) any {
	errMessage := fm.TernaryR(err == nil, "", func() string { return err.Error() })
	stackTrace := fmt.Sprintf("%+v", err)
	if errMessage == stackTrace || len(strings.Split(stackTrace, "\n")) < 3 {
		stackTrace = fmt.Sprintf("%+v", errors.WithStack(err))
	}

	return slf.setErrResponse(500, errMessage, stackTrace, opt...)
}

func (slf *stuFuseRContext) R503ServiceUnavailable(val any, opt ...ResponseOpt) any {
	return slf.setResponse(503, val, opt...)
}
