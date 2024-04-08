/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fc

// UnmarshalJSON implements the json.Unmarshaler interface.
func (slf *FCT) UnmarshalJSON(decimalBytes []byte) error {
	err := slf.vd.UnmarshalJSON(decimalBytes)
	if err != nil {
		return err
	}

	slf.set(slf.vd)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (slf FCT) MarshalJSON() ([]byte, error) {
	return []byte(slf.vd.String()), nil
}
