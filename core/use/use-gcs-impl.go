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
	"io"
	"mime/multipart"

	"github.com/andypangaribuan/gmod/ice"
)

func (slf *stuUseGcs) Init(credential ice.UtilEnvBase64, bucketName string) error {
	return slf.init(credential, bucketName)
}

func (slf *stuUseGcs) Read(dirPath string, callback func(directory string, name string) (canNext bool)) error {
	return slf.read(dirPath, callback)
}

func (slf *stuUseGcs) ReadLines(filePath string) (lines []string, err error) {
	return slf.readLines(filePath)
}

func (slf *stuUseGcs) Write(filePath string, reader io.Reader) error {
	return slf.write(filePath, reader)
}

func (slf *stuUseGcs) WriteData(filePath string, data []byte) error {
	return slf.write(filePath, bytes.NewReader(data))
}

func (slf *stuUseGcs) WriteReqFile(parts *map[string][]*multipart.FileHeader, key string, dirPath string, overrideFileName string) (fileName string, fileExt string, err error) {
	return slf.writeReqFile(parts, key, dirPath, overrideFileName)
}

func (slf *stuUseGcs) Delete(filepath string) error {
	return slf.delete(filepath)
}
