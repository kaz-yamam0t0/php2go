/**
 * Golang equivalent to php `ctype_cntrl`
 *
 * Check for control character(s)
 * @see https://www.php.net/manual/en/function.ctype-cntrl.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeCntrl(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if c >= 0x20 {
			return false
		}
	}
	return true
}
