/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fct

// UnmarshalJSON implements the json.Unmarshaler interface.
func (slf *FCT) UnmarshalJSON(decimalBytes []byte) error {
	err := slf.deci.UnmarshalJSON(decimalBytes)
	if err != nil {
		return err
	}

	slf.set(slf.deci)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (slf FCT) MarshalJSON() ([]byte, error) {
	return []byte(slf.deci.String()), nil
}
