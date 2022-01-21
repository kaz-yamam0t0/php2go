/**
 * Golang equivalent to php `strncasecmp`
 *
 * Binary safe case-insensitive string comparison of the first n characters
 * @see https://www.php.net/manual/en/function.strncasecmp.php
 *
 * @param string string1
 * @param string string2
 * @param int length
 * @return int
 */
package functions

func lower_(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}
func cmp_(a byte, b byte) bool {
	a = lower_(a)
	b = lower_(b)
	return (a - b) == 0
}
func Strncasecmp(s1 string, s2 string, length int) int {
	s1_len := len(s1)
	s2_len := len(s2)
	len_ := s1_len
	if len_ > s2_len {
		len_ = s2_len
	}
	if len_ > length {
		len_ = length
	}

	for i := 0; i < len_; i++ {
		if !cmp_(s1[i], s2[i]) {
			return int(lower_(s1[i])) - int(lower_(s2[i]))
		}
	}
	if len_ == length {
		return 0
	}

	if s1_len > s2_len {
		return 1
	} else if s1_len < s2_len {
		return -1
	} else {
		return 0
	}
}
