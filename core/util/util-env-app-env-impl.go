/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package util

func (slf *stuUtilEnvAppEnv) IsProd() bool {
	return slf.val == "prod" || slf.val == "production"
}

func (slf *stuUtilEnvAppEnv) IsRC() bool {
	return slf.val == "rc" || slf.val == "release-candidate"
}

func (slf *stuUtilEnvAppEnv) IsBox() bool {
	return slf.val == "box" || slf.val == "sbox" || slf.val == "sandbox"
}

func (slf *stuUtilEnvAppEnv) IsStg() bool {
	return slf.val == "stg" || slf.val == "staging"
}

func (slf *stuUtilEnvAppEnv) IsDev() bool {
	return slf.val == "dev" || slf.val == "development"
}

func (slf *stuUtilEnvAppEnv) GetName() string {
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

func (slf *stuUtilEnvAppEnv) IsRP() bool {
	return slf.IsRC() || slf.IsProd()
}
