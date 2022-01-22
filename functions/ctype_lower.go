package functions

// Check for lowercase character(s)
func CtypeLower(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if !('a' <= c && c <= 'z') {
			return false
		}
	}
	return true
}
