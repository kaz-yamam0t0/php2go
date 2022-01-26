package functions

import "errors"

// Find the first occurrence of a string
//
// Actually the sypnosis of Strstr() is like this:
//
//  Strstr(haystack string, needle string[, before_needle bool = false])
func Strstr(s string, needle string, args ...bool) (string, error) {
	before_needle := false
	if len(args) > 0 {
		before_needle = args[0]
	}
	s_len := len(s)
	n_len := len(needle)
	if n_len == 0 {
		return s[:], nil
	}
	if s_len == n_len {
		if s == needle {
			return s[:], nil
		} else {
			return "", errors.New("needle not included")
		}
	}
	if s_len < n_len {
		return "", errors.New("needle not included")
	}

	for pos := 0; pos < s_len-n_len+1; pos++ {
		if s[pos] == needle[0] {
			flg := true
			for c := 0; c < n_len; c++ {
				if s[pos+c] != needle[c] {
					flg = false
					break
				}
			}
			if flg {
				if before_needle {
					return s[0:pos], nil
				} else {
					return s[pos:], nil
				}
			}
		}
	}
	return "", errors.New("needle not included")
}
