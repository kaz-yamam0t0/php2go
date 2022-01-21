/**
 * Golang equivalent to php `strtolower`
 *
 * Make a string lowercase
 * @see https://www.php.net/manual/en/function.strtolower.php
 *
 * @param string s
 * @return string
 */
package functions

func Strtolower(s string) string {
	len_ := len(s)
	dst := make([]byte, len_)
	for i := 0; i < len_; i++ {
		c := s[i]

		if 'A' <= c && c <= 'Z' {
			dst[i] = c - 'A' + 'a'
		} else {
			dst[i] = s[i]
		}
	}
	return string(dst)
}
