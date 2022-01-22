package functions

// Check for numeric character(s)
func CtypeDigit(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !('0' <= c && c <= '9') {
			return false
		}
	}
	return true
}
