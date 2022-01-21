/**
 * Golang equivalent to php `ctype_alpha`
 *
 * Check for alphabetic character(s)
 * @see https://www.php.net/manual/en/function.ctype-alpha.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeAlpha(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !(('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')) {
			return false
		}
	}
	return true
}
