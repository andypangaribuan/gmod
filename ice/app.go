/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

type AppEnv interface {
	IsProd() bool // production
	IsRC() bool   // release candidate
	IsBox() bool  // sandbox
	IsStg() bool  // staging
	IsDev() bool  // development

	GetName() string
	IsRP() bool // release candidate + production
}
