/**
 * Golang equivalent to php `strcmp`
 *
 * Binary safe string comparison
 * @see https://www.php.net/manual/en/function.strcmp.php
 *
 * @param string string1
 * @param string string2
 * @return int
 */
package functions

func Strcmp(s1 string, s2 string) int {
	s1_len := len(s1)
	s2_len := len(s2)
	len_ := s1_len
	if len_ > s2_len {
		len_ = s2_len
	}
	for i := 0; i < len_; i++ {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
	}
	if s1_len > s2_len {
		return 1
	} else if s1_len < s2_len {
		return -1
	} else {
		return 0
	}
}
