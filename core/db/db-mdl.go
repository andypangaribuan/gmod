/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

type RepoOpt struct {
	WithDeletedAtIsNull *bool
	RWFetchWhenNull     *bool
}

type FetchOpt struct {
	WithDeletedAtIsNull *bool
	EndQuery            *string
}

type Update struct {
	Set          string
	Where        string
	SetArgs      []interface{}
	WhereArgs    []interface{}
	WithUpdateAt *bool
}
