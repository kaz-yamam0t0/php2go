/**
 * Golang equivalent to php `strrchr`
 *
 * Find the last occurrence of a character in a string
 * @see https://www.php.net/manual/en/function.strrchr.php
 *
 * @param string haystack
 * @param string needle
 * @return string
 */
package functions

import "errors"

func Strrchr(haystack string, args interface{}) (string, error) {
	var needle byte
	switch a := args.(type) {
	case byte:
		needle = a
	case string:
		needle = a[0]
	default:
		return "", errors.New("needle not found")
	}
	for i := len(haystack) - 1; i >= 0; i-- {
		if haystack[i] == needle {
			return haystack[i:], nil
		}
	}
	return "", errors.New("needle not found")
}
