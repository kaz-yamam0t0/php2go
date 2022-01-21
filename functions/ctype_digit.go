/**
 * Golang equivalent to php `ctype_digit`
 *
 * Check for numeric character(s)
 * @see https://www.php.net/manual/en/function.ctype-digit.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeDigit(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !('0' <= c && c <= '9') {
			return false
		}
	}
	return true
}
