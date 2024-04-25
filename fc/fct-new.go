/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package fc

import (
	"log"
)

func SafeNew(val any) (*FCT, error) {
	v, err := safeNew(val)
	if err != nil {
		return nil, err
	}

	var f FCT
	f.set(*v)

	return &f, nil
}

func UNew(val any) FCT {
	v, err := SafeNew(val)
	if err != nil {
		log.Panicf("error when fct unew\n%+v\n", err)
	}

	return *v
}
