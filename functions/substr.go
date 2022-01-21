/**
 * php `substr` with Golang
 *
 * Return part of a string
 * @see https://www.php.net/manual/en/function.substr.php
 *
 * @param string s
 * @param int offset
 * @param int length
 * @return string
 */
package functions

func Substr(s string, offset int, args ...int) string {
	var length int = 0
	if len(args) > 0 {
		length = args[0]
	}
	if offset < 0 {
		offset = len(s) + offset
	}
	if offset < 0 || len(s) < offset {
		return ""
	}

	if length != 0 {
		var end int = 0
		if length < 0 {
			end = len(s) + length
		} else {
			end = offset + length
		}
		if end < offset || len(s) < end {
			return ""
		}
		return s[offset:end]
	}
	return s[offset:]
}
