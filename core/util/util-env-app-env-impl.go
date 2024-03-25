/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

func (slf *srUtilEnvAppEnv) IsProd() bool {
	return slf.val == "prod" || slf.val == "production"
}

func (slf *srUtilEnvAppEnv) IsRC() bool {
	return slf.val == "rc" || slf.val == "release-candidate"
}

func (slf *srUtilEnvAppEnv) IsBox() bool {
	return slf.val == "box" || slf.val == "sbox" || slf.val == "sandbox"
}

func (slf *srUtilEnvAppEnv) IsStg() bool {
	return slf.val == "stg" || slf.val == "staging"
}

func (slf *srUtilEnvAppEnv) IsDev() bool {
	return slf.val == "dev" || slf.val == "development"
}

func (slf *srUtilEnvAppEnv) GetName() string {
	switch {
	case slf.IsProd():
		return "prod"
	case slf.IsRC():
		return "rc"
	case slf.IsBox():
		return "sbox"
	case slf.IsStg():
		return "stg"
	case slf.IsDev():
		return "dev"
	}

	return slf.val
}

func (slf *srUtilEnvAppEnv) IsRP() bool {
	return slf.IsRC() || slf.IsProd()
}
