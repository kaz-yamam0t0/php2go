/**
 * Golang equivalent to php `ucwords`
 *
 * Uppercase the first character of each word in a string
 * @see https://www.php.net/manual/en/function.ucwords.php
 *
 * @param string s
 * @param string separators
 * @return string
 */
package functions

func Ucwords(s string, args ...string) string {
	var separators string
	if len(args) > 0 {
		separators = args[0]
	} else {
		separators = " \t\r\n\f\v"
	}

	len_ := len(s)
	dst := make([]byte, len_)
	for i := 0; i < len_; i++ {
		c := s[i]

		if 'a' <= c && c <= 'z' {
			if i == 0 {
				dst[i] = c - 'a' + 'A'
				continue
			}
			is_sep := false
			for j := len(separators) - 1; j >= 0; j-- {
				if s[i-1] == separators[j] {
					is_sep = true
					break
				}
			}
			if is_sep {
				dst[i] = c - 'a' + 'A'
				continue
			}
		}
		dst[i] = s[i]
	}
	return string(dst)
}
