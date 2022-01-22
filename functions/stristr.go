package functions

import (
	"errors"
)

// Case-insensitive strstr
// 
// Actually the sypnosis of Stristr() is like this:
// 
//  Stristr(s string, needle string[, before_needle bool = 0]) 
func Stristr(s string, needle string, args ...bool) (string, error) {
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
		for i := 0; i < s_len; i++ {
			if !cmpLower(s[i], needle[i]) {
				return "", errors.New("needle not included")
			}
		}
		return s[:], nil
	}
	if s_len < n_len {
		return "", errors.New("needle not included")
	}

	for pos := 0; pos < s_len-n_len+1; pos++ {
		if cmpLower(s[pos], needle[0]) {
			flg := true
			for c := 0; c < n_len; c++ {
				if !cmpLower(s[pos+c], needle[c]) {
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
