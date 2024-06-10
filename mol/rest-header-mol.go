/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package mol

import "time"

type RequestHeader struct {
	Authorization string `json:"Authorization"`
	UID           string `json:"X-UID"`
	Language      string `json:"X-Language"`
	Version       string `json:"X-Version"`
	HashKey       string `json:"X-HashKey"`
	UserHashId    string `json:"X-UserHashId"`
	RFTimeRaw     string `json:"X-RFTime"` // supported format: "2006-01-02T15:04:05+07:00" or "2006-01-02 15:04:05+07:00"
	RFTime        *time.Time
}
