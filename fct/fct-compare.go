/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fct

import "github.com/pkg/errors"

// supported operator: ==, !=, <, <=, >, >=
func Compare(left FCT, operator string, right FCT) (bool, error) {
	if !ifHaveIn(operator, "==", "!=", "<", "<=", ">=", ">") {
		return false, errors.New("fct compare: invalid operation")
	}

	switch operator {
	case "==":
		return left.deci.Equal(right.deci), nil

	case "!=":
		return !left.deci.Equal(right.deci), nil

	case "<":
		return left.deci.LessThan(right.deci), nil

	case "<=":
		return left.deci.LessThanOrEqual(right.deci), nil

	case ">":
		return left.deci.GreaterThan(right.deci), nil

	case ">=":
		return left.deci.GreaterThanOrEqual(right.deci), nil
	}

	return false, errors.New("fct compare: unhandled operation")
}

// supported operator: ==, !=, <, <=, >, >=
func AnyCompare(left any, operator string, right any) (bool, error) {
	leftValue, err := New(left)
	if err != nil {
		return false, errors.WithMessage(err, "fct compare: invalid left value")
	}

	rightValue, err := New(right)
	if err != nil {
		return false, errors.WithMessage(err, "fct compare: invalid right value")
	}

	return Compare(*leftValue, operator, *rightValue)
}
