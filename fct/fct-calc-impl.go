/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fct

// supported operator: +, -, *, /, %
func Calc2(v1 FCT, operator1 string, v2 FCT) (FCT, error) {
	return calc(v1, operator1, v2)
}

// supported operator: +, -, *, /, %
func AnyCalc2(v1 any, operator1 string, v2 any) (FCT, error) {
	return calc(v1, operator1, v2)
}

// supported operator: +, -, *, /, %
func UnsafeCalc2(v1 FCT, operator1 string, v2 FCT) FCT {
	return unsafeCalc(v1, operator1, v2)
}

// supported operator: +, -, *, /, %
func UnsafeAnyCalc2(v1 any, operator1 string, v2 any) FCT {
	return unsafeCalc(v1, operator1, v2)
}

// supported operator: +, -, *, /, %
func Calc3(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3)
}

// supported operator: +, -, *, /, %
func AnyCalc3(v1 any, operator1 string, v2 any, operator2 string, v3 any) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3)
}

func UnsafeCalc3(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3)
}

func UnsafeAnyCalc3(v1 any, operator1 string, v2 any, operator2 string, v3 any) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3)
}

// supported operator: +, -, *, /, %
func Calc4(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT, operator3 string, v4 FCT) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3, operator3, v4)
}

// supported operator: +, -, *, /, %
func AnyCalc4(v1 any, operator1 string, v2 any, operator2 string, v3 any, operator3 string, v4 any) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3, operator3, v4)
}

// supported operator: +, -, *, /, %
func UnsafeCalc4(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT, operator3 string, v4 FCT) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3, operator3, v4)
}

// supported operator: +, -, *, /, %
func UnsafeAnyCalc4(v1 any, operator1 string, v2 any, operator2 string, v3 any, operator3 string, v4 any) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3, operator3, v4)
}

// supported operator: +, -, *, /, %
func Calc5(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT, operator3 string, v4 FCT, operator4 string, v5 FCT) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5)
}

// supported operator: +, -, *, /, %
func AnyCalc5(v1 any, operator1 string, v2 any, operator2 string, v3 any, operator3 string, v4 any, operator4 string, v5 any) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5)
}

// supported operator: +, -, *, /, %
func UnsafeCalc5(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT, operator3 string, v4 FCT, operator4 string, v5 FCT) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5)
}

// supported operator: +, -, *, /, %
func UnsafeAnyCalc5(v1 any, operator1 string, v2 any, operator2 string, v3 any, operator3 string, v4 any, operator4 string, v5 any) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5)
}

// supported operator: +, -, *, /, %
func Calc6(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT, operator3 string, v4 FCT, operator4 string, v5 FCT, operator5 string, v6 FCT) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5, operator5, v6)
}

// supported operator: +, -, *, /, %
func AnyCalc6(v1 any, operator1 string, v2 any, operator2 string, v3 any, operator3 string, v4 any, operator4 string, v5 any, operator5 string, v6 any) (FCT, error) {
	return calc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5, operator5, v6)
}

// supported operator: +, -, *, /, %
func UnsafeCalc6(v1 FCT, operator1 string, v2 FCT, operator2 string, v3 FCT, operator3 string, v4 FCT, operator4 string, v5 FCT, operator5 string, v6 FCT) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5, operator5, v6)
}

// supported operator: +, -, *, /, %
func UnsafeAnyCalc6(v1 any, operator1 string, v2 any, operator2 string, v3 any, operator3 string, v4 any, operator4 string, v5 any, operator5 string, v6 any) FCT {
	return unsafeCalc(v1, operator1, v2, operator2, v3, operator3, v4, operator4, v5, operator5, v6)
}
