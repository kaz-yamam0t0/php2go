/**
 * Golang equivalent to php `ctype_alnum`
 *
 * Check for alphanumeric character(s)
 * @see https://www.php.net/manual/en/function.ctype-alnum.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeAlnum(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !(('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')) {
			return false
		}
	}
	return true
}
