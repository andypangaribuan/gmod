/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package json

func (*stuJson) Marshal(obj interface{}) ([]byte, error) {
	return api.Marshal(obj)
}

func (*stuJson) UnMarshal(data []byte, out interface{}) error {
	return api.Unmarshal(data, &out)
}

func (*stuJson) Encode(obj interface{}) (string, error) {
	return api.MarshalToString(obj)
}

func (*stuJson) Decode(jsonStr string, out interface{}) error {
	return api.UnmarshalFromString(jsonStr, &out)
}

func (*stuJson) MapToJson(maps map[string]interface{}) (string, error) {
	return api.MarshalToString(maps)
}
