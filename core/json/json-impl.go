/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package json

func (*stuJson) Marshal(obj any) ([]byte, error) {
	return api.Marshal(obj)
}

func (*stuJson) Unmarshal(data []byte, out any) error {
	return api.Unmarshal(data, &out)
}

func (*stuJson) Encode(obj any) (string, error) {
	return api.MarshalToString(obj)
}

func (*stuJson) Decode(jsonStr string, out any) error {
	return api.UnmarshalFromString(jsonStr, &out)
}

func (*stuJson) MapToJson(maps map[string]any) (string, error) {
	return api.MarshalToString(maps)
}
