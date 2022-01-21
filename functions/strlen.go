/**
 * Golang equivalent to php `strlen`
 *
 * Get string length
 * @see https://www.php.net/manual/en/function.strlen.php
 *
 * @param string s
 * @return int
 */
package functions

func Strlen(s string) int {
	return len(s)
}
