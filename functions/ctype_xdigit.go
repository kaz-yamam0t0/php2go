/**
 * Golang equivalent to php `ctype_xdigit`
 *
 * Check for character(s) representing a hexadecimal digit
 * @see https://www.php.net/manual/en/function.ctype-xdigit.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeXdigit(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !(('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')) {
			return false
		}
	}
	return true
}
