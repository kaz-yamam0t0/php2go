package functions

// Determine if a string contains a given substring
func StrContains(s string, needle string) bool {
	s_len := len(s)
	n_len := len(needle)
	if n_len == 0 {
		return true
	}
	if s_len == n_len {
		return s == needle
	}
	if s_len < n_len {
		return false
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
				return true
			}
		}
	}
	return false
}
