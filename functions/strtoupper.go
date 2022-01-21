/**
 * Golang equivalent to php `strtoupper`
 *
 * Make a string uppercase
 * @see https://www.php.net/manual/en/function.strtoupper.php
 *
 * @param string s
 * @return string
 */
package functions

func Strtoupper(s string) string {
	len_ := len(s)
	dst := make([]byte, len_)
	for i := 0; i < len_; i++ {
		c := s[i]

		if 'a' <= c && c <= 'z' {
			dst[i] = c - 'a' + 'A'
		} else {
			dst[i] = s[i]
		}
	}
	return string(dst)
}
