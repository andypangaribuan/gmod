/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package crypto

import (
	"bytes"
	"crypto/aes"
	"fmt"
)

// ecc mode decryption
func ecbDecrypt(value, key []byte) ([]byte, error) {
	if !isValidKey(key) {
		return nil, fmt.Errorf("the length of the secret key is wrong, the current incoming length is% d", len(key))
	}
	if len(value) < 1 {
		return nil, fmt.Errorf("source data length cannot be 0")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(value)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("the source data length must be an integer multiple of %d, the current length is %d", block.BlockSize(), len(value))
	}

	var dst []byte
	tmpData := make([]byte, block.BlockSize())

	for index := 0; index < len(value); index += block.BlockSize() {
		block.Decrypt(tmpData, value[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}

	dst, err = pkcs5UnPadding(dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

// ecc mode encryption
func ecbEncrypt(src, key []byte) ([]byte, error) {
	if !isValidKey(key) {
		return nil, fmt.Errorf("the length of the secret key is wrong, the current incoming length is% d", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src) < 1 {
		return nil, fmt.Errorf("source data length cannot be 0")
	}

	src = pkcs5Padding(src, block.BlockSize())
	if len(src)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("the source data length must be an integer multiple of %d, the current length is %d", block.BlockSize(), len(src))
	}

	var dst []byte
	tmpData := make([]byte, block.BlockSize())

	for index := 0; index < len(src); index += block.BlockSize() {
		block.Encrypt(tmpData, src[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}

	return dst, nil
}

// pkcs5 filling
func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// remove pkcs5 filling
func pkcs5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unPadding := int(origData[length-1])

	if length < unPadding {
		return nil, fmt.Errorf("invalid un-padding length")
	}

	return origData[:(length - unPadding)], nil
}

func isValidKey(key []byte) bool {
	k := len(key)

	switch k {
	case 16, 24, 32:
		return true
	default:
		return false
	}
}
