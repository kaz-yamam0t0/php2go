/**
 * Golang equivalent to php `md5`
 *
 * Calculate the md5 hash of a string
 * @see https://www.php.net/manual/en/function.md5.php
 *
 * @param string s
 * @param bool binary
 * @return string
 */
package functions

import (
	_md5 "crypto/md5"
	"encoding/hex"
)

func Md5(s string, args ...bool) string {
	binary_flg := false
	if len(args) > 0 {
		binary_flg = args[0]
	}
	_ = binary_flg

	hash := _md5.New()
	defer hash.Reset()
	hash.Write([]byte(s))
	if binary_flg {
		return string(hash.Sum(nil))
	}
	return hex.EncodeToString(hash.Sum(nil))
}
