/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package use

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

func (slf *stuUseGcs) Init(credential ice.UtilEnvBase64, bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credential.Data()))
	if err != nil {
		return errors.WithStack(err)
	}

	slf.bucket = client.Bucket(bucketName)
	return nil
}
