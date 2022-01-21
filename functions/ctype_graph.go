/**
 * Golang equivalent to php `ctype_graph`
 *
 * Check for any printable character(s) except space
 * @see https://www.php.net/manual/en/function.ctype-graph.php
 *
 * @param interface{} text
 * @return bool
 */
package functions

func CtypeGraph(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if c < 0x21 || 0x7E < c {
			return false
		}
	}
	return true
}
