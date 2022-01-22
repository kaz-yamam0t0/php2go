package functions

// Check for uppercase character(s)
func CtypeUpper(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !('A' <= c && c <= 'Z') {
			return false
		}
	}
	return true
}
