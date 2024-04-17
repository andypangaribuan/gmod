/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package clog

type Instance interface {
	DbqV1(mol DbqV1) error
	ServicePieceV1(mol ServicePieceV1) error
	ServiceV1(mol ServiceV1) error
}
