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

import "strings"

func cmdLs(command *string) []string {
	if command == nil {
		return []string{}
	}

	return strings.Fields(*command)
}

func cmdRemoveAtIndex(command *string, index int) {
	if command == nil {
		return
	}

	ls := cmdLs(command)
	if len(ls) >= index+1 {
		*command = strings.Join(sliceRemove(ls, index), " ")
	}
}

func cmdFivSingleArg(command *string, key string) (index int) {
	index = -1
	ls := cmdLs(command)

	for i := range ls {
		if key == ls[i] {
			index = i
			break
		}
	}

	return index
}

func cmdGevFirstArg(command *string) string {
	if command == nil {
		return ""
	}

	ls := cmdLs(command)

	if len(ls) > 0 {
		v := ls[0]
		if len(v) > 0 && v[0:1] != "-" && v[0:1] != "+" {
			return v
		}
	}

	return ""
}
