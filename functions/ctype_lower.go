/**
 * Golang equivalent to php `ctype_lower`
 *
 * Check for lowercase character(s)
 * @see https://www.php.net/manual/en/function.ctype-lower.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeLower(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !('a' <= c && c <= 'z') {
			return false
		}
	}
	return true
}
