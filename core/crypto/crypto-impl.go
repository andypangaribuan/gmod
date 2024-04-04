/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func (*stuCrypto) AesEcbEncrypt(key, value string) (string, error) {
	val, err := ecbEncrypt([]byte(value), []byte(key))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(val), nil
}

func (*stuCrypto) AesEcbRawEncrypt(key, value string) ([]byte, error) {
	return ecbEncrypt([]byte(value), []byte(key))
}

func (*stuCrypto) AesEcbDecrypt(key, value string) (string, error) {
	data, err := hex.DecodeString(value)
	if err != nil {
		return "", err
	}

	val, err := ecbDecrypt(data, []byte(key))
	if err != nil {
		return "", err
	}

	return string(val), nil
}

func (*stuCrypto) AesEcbRawDecrypt(key string, value []byte) ([]byte, error) {
	return ecbDecrypt(value, []byte(key))
}

func (*stuCrypto) Md5(val string) string {
	hash := md5.Sum([]byte(val))
	return hex.EncodeToString(hash[:])
}

func (*stuCrypto) Sha256(val string) string {
	hash := sha256.Sum256([]byte(val))
	return hex.EncodeToString(hash[:])
}

func (*stuCrypto) Sha512(val string) string {
	hash := sha512.Sum512([]byte(val))
	return hex.EncodeToString(hash[:])
}
