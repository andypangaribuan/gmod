/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

type Crypto interface {
	AesEcbEncrypt(key, value string) (string, error)
	AesEcbRawEncrypt(key, value string) ([]byte, error)
	AesEcbDecrypt(key, value string) (string, error)
	AesEcbRawDecrypt(key string, value []byte) ([]byte, error)
	Md5(val string) string
	Sha256(val string) string
	Sha512(val string) string
}
