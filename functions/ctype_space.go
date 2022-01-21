/**
 * Golang equivalent to php `ctype_space`
 *
 * Check for whitespace character(s)
 * @see https://www.php.net/manual/en/function.ctype-space.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeSpace(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if c != 0x0A && c != 0x0D && c != 0x20 && c != 0x09 && c != 0x0b && c != 0x0c {
			return false
		}
	}
	return true
}
