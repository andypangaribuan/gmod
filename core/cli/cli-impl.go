/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cli

import "github.com/andypangaribuan/gmod/gm"

func (slf *stuCli) FindSingleArg(command *string, key string) bool {
	index := cmdFivSingleArg(command, key)
	if index >= 0 {
		cmdRemoveAtIndex(command, index)
		return true
	}

	return false
}

func (slf *stuCli) GetFirstArg(command *string) string {
	arg := cmdGevFirstArg(command)
	if arg != "" {
		cmdRemoveAtIndex(command, 0)
	}

	return arg
}

func (slf *stuCli) GetFirstArgAsDate(command *string, format ...string) string {
	arg := cmdGevFirstArg(command)
	if arg != "" {
		layout := "yyyy-MM-dd"
		if len(format) > 0 && format[0] != "" {
			layout = format[0]
		}

		_, err := gm.Conv.Time.ToTime(arg, layout)
		if err == nil {
			cmdRemoveAtIndex(command, 0)
			return arg
		}
	}

	return ""
}
