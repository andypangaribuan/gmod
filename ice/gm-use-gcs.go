/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import (
	"io"
	"mime/multipart"
)

type UseGcs interface {
	Init(credential UtilEnvBase64, bucketName string) error
	Read(dirPath string, callback func(directory string, name string) (canNext bool)) error
	ReadLines(filePath string) (lines []string, err error)
	Write(filePath string, reader io.Reader) error
	WriteData(filePath string, data []byte) error
	WriteReqFile(parts *map[string][]*multipart.FileHeader, key string, dirPath string, overrideFileName string) (fileName string, fileExt string, err error)
}
