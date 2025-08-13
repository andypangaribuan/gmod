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

import (
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/ice"
	"slices"
)

func (slf *stuQueryBuilder) And(condition string, args ...any) ice.DbQueryBuilder {
	slf.add("AND", condition, args...)

	return slf
}

func (slf *stuQueryBuilder) AndNotNill(condition string, arg any, validator ...func() (ok bool, args []any)) ice.DbQueryBuilder {
	if !fm.IsNil(arg) {
		ok, args := slf.validator(validator, arg)
		if ok {
			slf.add("AND", condition, args...)
		}
	}

	return slf
}

func (slf *stuQueryBuilder) AndNotNill2(condition string, arg1 any, arg2 any, validator ...func() (ok bool, args []any)) ice.DbQueryBuilder {
	if !fm.IsNil(arg1) && !fm.IsNil(arg2) {
		ok, args := slf.validator(validator, arg1, arg2)
		if ok {
			slf.add("AND", condition, args...)
		}
	}

	return slf
}

func (slf *stuQueryBuilder) AndNotNil(condition string, args ...any) ice.DbQueryBuilder {
	if slices.ContainsFunc(args, fm.IsNil) {
		return slf
	}

	ls, validator := slf.getArgsValidator1(args)
	if validator == nil {
		slf.add("AND", condition, ls...)
	} else {
		ok, ls := (*validator)()
		if ok {
			slf.add("AND", condition, ls...)
		}
	}

	return slf
}

func (slf *stuQueryBuilder) Build() (where string, args []any) {
	where = ""
	args = make([]any, 0)

	for i := range len(slf.conditions) {
		if where != "" {
			where += " " + slf.adapter[i]
		}

		if where != "" {
			where += " "
		}

		where += slf.conditions[i]
		args = append(args, slf.args[i]...)
	}

	return where, args
}
