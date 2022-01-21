/**
 * Golang equivalent to php `count_chars`
 *
 * Return information about characters used in a string
 * @see https://www.php.net/manual/en/function.count-chars.php
 *
 * @param string s
 * @param int mode (ignored)
 * @return interface{}
 */
package functions

func CountChars(s string, args ...int) [256]int {
	/*	mode := 0
		if len(args) > 0 {
			mode = args[0]
		} */
	var dst [256]int
	s_len := len(s)
	for i := 0; i < s_len; i++ {
		dst[s[i]]++
	}
	return dst
}
