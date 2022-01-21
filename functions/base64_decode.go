/**
 * Golang equivalent to php `base64_decode`
 *
 * Decodes data encoded with MIME base64
 * @see https://www.php.net/manual/en/function.base64-decode.php
 *
 * @param string s
 * @param bool strict
 * @return (interface{}, error)
 */
package functions

import (
	b64 "encoding/base64"
)

func Base64Decode(s string) (string, error) {
	b, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil

}
