/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

type Json interface {
	Marshal(obj any) ([]byte, error)
	UnMarshal(data []byte, out any) error
	Encode(obj any) (string, error)
	Decode(jsonStr string, out any) error
	MapToJson(maps map[string]any) (string, error)
}
