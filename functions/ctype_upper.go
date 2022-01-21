/**
 * Golang equivalent to php `ctype_upper`
 *
 * Check for uppercase character(s)
 * @see https://www.php.net/manual/en/function.ctype-upper.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeUpper(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !('A' <= c && c <= 'Z') {
			return false
		}
	}
	return true
}
