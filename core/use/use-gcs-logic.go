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
	"io"

	"cloud.google.com/go/storage"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

func (slf *stuUseGcs) init(credential ice.UtilEnvBase64, bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credential.Data()))
	if err != nil {
		return errors.WithStack(err)
	}

	slf.bucket = client.Bucket(bucketName)
	return nil
}

func (slf *stuUseGcs) write(filePath string, reader io.Reader) error {
	if slf.bucket == nil {
		return errors.New("do init first, before you use this")
	}

	ctx := context.Background()
	writer := slf.bucket.Object(filePath).NewWriter(ctx)

	_, err := io.Copy(writer, reader)
	if err != nil {
		_ = writer.Close()
		return errors.WithStack(err)
	}

	err = writer.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
