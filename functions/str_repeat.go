/**
 * Golang equivalent to php `str_repeat`
 *
 * Repeat a string
 * @see https://www.php.net/manual/en/function.str-repeat.php
 *
 * @param string s
 * @param int times
 * @return string
 */
package functions

import "strings"

func StrRepeat(s string, times int) string {
	return strings.Repeat(s, times)
}
