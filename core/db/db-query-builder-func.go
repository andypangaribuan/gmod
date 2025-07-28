/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

func (slf *stuQueryBuilder) add(adapter string, condition string, args ...any) {
	slf.adapter = append(slf.adapter, adapter)
	slf.conditions = append(slf.conditions, condition)
	slf.args = append(slf.args, args)
}

func (slf *stuQueryBuilder) validator(validator []func() (ok bool, args []any), arg ...any) (ok bool, args []any) {
	ok = true
	args = make([]any, 0)

	if len(validator) > 0 {
		ok, args = validator[0]()
	} else if len(arg) > 0 {
		args = append(args, arg...)
	}

	return
}

func (slf *stuQueryBuilder) getArgsValidator1(args []any) ([]any, *queryBuilderFuncValidator1) {
	var (
		ls = make([]any, 0)
		fn *queryBuilderFuncValidator1
	)

	for _, arg := range args {
		switch v := arg.(type) {
		case queryBuilderFuncValidator1:
			fn = &v

		default:
			ls = append(ls, arg)
		}
	}

	return ls, fn
}
