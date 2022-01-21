/**
 * Golang equivalent to php `implode`
 *
 * Join array elements with a string
 * @see https://www.php.net/manual/en/function.implode.php
 *
 * @param string separator
 * @param interface{} array
 * @return string
 */
package functions

import "strings"

func Implode(separator string, a []string) string {
	return strings.Join(a, separator)
}
