/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

type Response struct {
	Meta ResponseMeta `json:"meta"`
	Data any          `json:"data,omitempty"`
}

type ResponseMeta struct {
	Code       int    `json:"code"`
	SubCode    string `json:"sub_code,omitempty"`
	Message    string `json:"message,omitempty"`
	AppMessage string `json:"app_message,omitempty"`
}

type ResponseOpt struct {
	SubCode    string
	Message    string
	AppMessage string
}
