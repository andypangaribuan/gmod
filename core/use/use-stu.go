/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package use

import "cloud.google.com/go/storage"

type stuUse struct{}

type stuUseGcs struct {
	bucket *storage.BucketHandle
}