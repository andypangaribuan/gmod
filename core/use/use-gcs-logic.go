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
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"strings"

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

func (slf *stuUseGcs) writeReqFile(parts *map[string][]*multipart.FileHeader, key string, dirPath string, overrideFileName string) (fileName string, fileExt string, err error) {
	if parts == nil {
		err = errors.New("no file")
		return
	}

	fh, ok := (*parts)[key]
	if !ok || len(fh) == 0 || fh[0] == nil {
		err = errors.New("empty file")
		return
	}

	var (
		part     = fh[0]
		name     = part.Filename
		fullPath string
		ls       = strings.Split(name, ".")
	)

	if len(ls) > 0 {
		fileExt = strings.TrimSpace(ls[len(ls)-1])
		fileExt = strings.ToLower(fileExt)
		name = strings.TrimSpace(name[:len(name)-len(fileExt)-1])
	}

	if overrideFileName != "" {
		name = overrideFileName
	}

	fileName = name
	fullPath = dirPath

	if fullPath[len(fullPath)-1:] != "/" {
		fullPath += "/"
	}

	fullPath += fileName

	if fullPath[len(fullPath)-1:] != "." {
		fullPath += "."
	}

	fullPath += fileExt

	file, er := part.Open()
	defer func() {
		_ = file.Close()
	}()

	if er != nil {
		err = errors.WithStack(er)
		return
	}

	buf := bytes.NewBuffer(nil)
	_, er = io.Copy(buf, file)
	if er != nil {
		err = errors.WithStack(er)
		return
	}

	err = slf.write(fullPath, buf)
	return
}
