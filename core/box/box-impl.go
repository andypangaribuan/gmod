/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package box

import "github.com/andypangaribuan/gmod/fm"

func (slf *stuBox) Set(key, subKey string, val any) {
	slf.mx.Lock()
	defer slf.mx.Unlock()

	sbox, ok := slf.box[key]
	if !ok {
		slf.box[key] = map[string]any{subKey: val}
	} else {
		sbox[subKey] = val
		slf.box[key] = sbox
	}
}

func (slf *stuBox) GetValue(key, subKey string) any {
	slf.mx.RLock()
	defer slf.mx.RUnlock()

	sbox, ok := slf.box[key]
	if !ok {
		return nil
	}

	val, ok := sbox[subKey]
	if !ok {
		return nil
	}

	return val
}

func (slf *stuBox) GetSub(key string) map[string]any {
	slf.mx.RLock()
	defer slf.mx.RUnlock()

	sbox, ok := slf.box[key]
	return fm.Ternary(ok, sbox, nil)
}
