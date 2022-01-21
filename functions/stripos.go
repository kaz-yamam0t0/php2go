/**
 * Golang equivalent to php `stripos`
 *
 * Find the position of the first occurrence of a case-insensitive substring in a string
 * @see https://www.php.net/manual/en/function.stripos.php
 *
 * @param string haystack
 * @param string needle
 * @param int offset
 * @return interface{}
 */
package functions

func Stripos(s string, needle string, args ...int) int {
	offset := 0
	if len(args) > 0 {
		offset = args[0]
	}
	s_len := len(s)
	n_len := len(needle)
	if n_len == 0 {
		return 0
	}
	if s_len == n_len {
		for i := 0; i < s_len; i++ {
			if !cmpLower(s[i], needle[i]) {
				return -1
			}
		}
		return 0
	}
	if s_len-offset < n_len {
		return -1
	}

	for pos := offset; pos < s_len-n_len+1; pos++ {
		if cmpLower(s[pos], needle[0]) {
			flg := true
			for c := 0; c < n_len; c++ {
				if !cmpLower(s[pos+c], needle[c]) {
					flg = false
					break
				}
			}
			if flg {
				return pos
			}
		}
	}
	return -1
}
