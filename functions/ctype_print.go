/**
 * Golang equivalent to php `ctype_print`
 *
 * Check for printable character(s)
 * @see https://www.php.net/manual/en/function.ctype-print.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypePrint(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if c < 0x20 || 0x7E < c {
			return false
		}
	}
	return true
}
