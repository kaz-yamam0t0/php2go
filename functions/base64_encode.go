/**
 * Golang equivalent to php `base64_encode`
 *
 * Encodes data with MIME base64
 * @see https://www.php.net/manual/en/function.base64-encode.php
 *
 * @param string s
 * @return string
 */
package functions

import (
	b64 "encoding/base64"
)

func Base64Encode(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}
